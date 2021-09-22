package entity

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/mirzaakhena/gogen2/application/apperror"
	"github.com/mirzaakhena/gogen2/domain/vo"
	"os"
)

// ObjError depend on (which) usecase that want to be tested
type ObjError struct {
	ErrorName vo.Naming
}

// ObjDataError is object that used in template
type ObjDataError struct {
	ErrorName string
}

// NewObjError Constructor
func NewObjError(errorName string) (*ObjError, error) {

	if errorName == "" {
		return nil, apperror.ErrorNameMustNotEmpty
	}

	var obj ObjError
	obj.ErrorName = vo.Naming(errorName)

	return &obj, nil
}

// TODO add function is Exist

// GetData ...
func (o ObjError) GetData() *ObjDataError {
	return &ObjDataError{
		ErrorName: o.ErrorName.String(),
	}
}

// GetErrorRootFolderName ...
func GetErrorRootFolderName() string {
	return fmt.Sprintf("application/apperror")
}

// GetErrorEnumFileName ...
func GetErrorEnumFileName() string {
	return fmt.Sprintf("%s/error_enum.go", GetErrorRootFolderName())
}

// GetErrorFuncFileName ...
func GetErrorFuncFileName() string {
	return fmt.Sprintf("%s/error_func.go", GetErrorRootFolderName())
}

// InjectCode ...
func (o ObjError) InjectCode(templateCode string) ([]byte, error) {

	// reopen the file
	file, err := os.Open(GetErrorEnumFileName())
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var buffer bytes.Buffer
	for scanner.Scan() {
		row := scanner.Text()

		buffer.WriteString(row)
		buffer.WriteString("\n")
	}

	if err := file.Close(); err != nil {
		return nil, err
	}

	// write the template in the end of file
	buffer.WriteString(templateCode)
	buffer.WriteString("\n")

	return buffer.Bytes(), nil
}
