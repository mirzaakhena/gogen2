package entity

import (
	"github.com/mirzaakhena/gogen2/application/apperror"
	"github.com/mirzaakhena/gogen2/domain/vo"
)

// ObjError depend on (which) usecase that want to be tested
type ObjError struct {
	ErrorName vo.Naming
}

type ObjDataError struct {
	ErrorName string
}

func NewObjError(errorName string) (*ObjError, error) {

	if errorName == "" {
		return nil, apperror.ErrorNameMustNotEmpty
	}

	var obj ObjError
	obj.ErrorName = vo.Naming(errorName)

	return &obj, nil
}

func (o ObjError) GetData() *ObjDataError {
	return &ObjDataError{
		ErrorName: o.ErrorName.String(),
	}
}