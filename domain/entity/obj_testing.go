package entity

import (
	"fmt"
	"github.com/mirzaakhena/gogen2/domain/vo"
)

// ObjTesting depend on (which) usecase that want to be tested
type ObjTesting struct {
	TestName   vo.Naming
	ObjUsecase ObjUsecase
}

type ObjDataTesting struct {
	PackagePath string
	UsecaseName string
	TestName    string
	Methods     vo.OutportMethods
}

func NewObjTesting(testName string, objUsecase ObjUsecase) (*ObjTesting, error) {

	var obj ObjTesting
	obj.TestName = vo.Naming(testName)
	obj.ObjUsecase = objUsecase

	return &obj, nil
}

func (o ObjTesting) GetData(PackagePath string, outportMethods vo.OutportMethods) *ObjDataTesting {
	return &ObjDataTesting{
		PackagePath: PackagePath,
		UsecaseName: o.ObjUsecase.UsecaseName.String(),
		TestName:    o.TestName.LowerCase(),
		Methods:     outportMethods,
	}
}

func GetTestFileName(o ObjTesting) string {
	return fmt.Sprintf("%s/testcase_%s_test.go", GetUsecaseRootFolderName(o.ObjUsecase), o.TestName.LowerCase())
}
