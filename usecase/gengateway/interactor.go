package gengateway

import (
	"context"
	"github.com/mirzaakhena/gogen2/domain/entity"
	"github.com/mirzaakhena/gogen2/domain/service"
	"github.com/mirzaakhena/gogen2/domain/vo"
)

//go:generate mockery --name Outport -output mocks/

type genGatewayInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenGateway
func NewUsecase(outputPort Outport) Inport {
	return &genGatewayInteractor{
		outport: outputPort,
	}
}

// Execute the usecase GenGateway
func (r *genGatewayInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	objUsecase, err := entity.NewObjUsecase(req.UsecaseName)
	if err != nil {
		return nil, err
	}

	objGateway, err := entity.NewObjGateway(req.GatewayName, *objUsecase)
	if err != nil {
		return nil, err
	}

	err = service.ConstructApplication(ctx, r.outport)
	if err != nil {
		return nil, err
	}

	err = service.ConstructLog(ctx, r.outport)
	if err != nil {
		return nil, err
	}

	_, err = r.outport.CreateFolderIfNotExist(ctx, entity.GetGatewayRootFolderName(*objGateway))
	if err != nil {
		return nil, err
	}

	packagePath := r.outport.GetPackagePath(ctx)

	outportMethods, err := vo.NewOutportMethods(req.UsecaseName, packagePath)
	if err != nil {
		return nil, err
	}

	gatewayFile := entity.GetGatewayFileName(*objGateway)

	// file gateway impl is not exist, we create one
	if !r.outport.IsFileExist(ctx, gatewayFile) {

		gatewayTemplate := r.outport.GetGatewayTemplate(ctx)

		data := objGateway.GetData(packagePath, outportMethods)
		err := r.outport.WriteFile(ctx, gatewayTemplate, gatewayFile, data)
		if err != nil {
			return nil, err
		}

		err = r.outport.Reformat(ctx, gatewayFile, nil)
		if err != nil {
			return nil, err
		}

		return res, nil
	}

	// file gateway impl file is already exist, we want to inject non existing method

	existingFunc, err := vo.NewGatewayMethod(objGateway.GatewayName.CamelCase(), entity.GetGatewayRootFolderName(*objGateway), packagePath)
	if err != nil {
		return nil, err
	}

	// collect the only methods that has not added yet
	notExistingMethod := vo.OutportMethods{}
	for _, m := range outportMethods {
		if _, exist := existingFunc[m.MethodName]; !exist {
			notExistingMethod = append(notExistingMethod, m)
		}
	}

	gatewayCode := r.outport.GetGatewayMethodTemplate(ctx)

	// we will only inject the non existing method
	data := objGateway.GetData(packagePath, notExistingMethod)

	templateHasBeenInjected, err := r.outport.PrintTemplate(ctx, gatewayCode, data)
	if err != nil {
		return nil, err
	}

	bytes, err := objGateway.InjectToGateway(templateHasBeenInjected)
	if err != nil {
		return nil, err
	}

	// reformat outport.go
	err = r.outport.Reformat(ctx, entity.GetGatewayFileName(*objGateway), bytes)
	if err != nil {
		return nil, err
	}

	return res, nil
}
