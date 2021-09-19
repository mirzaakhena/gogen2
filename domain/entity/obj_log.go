package entity

import (
	"fmt"
)

func GetLogRootFolderName() string {
	return fmt.Sprintf("infrastructure/log")
}

func GetLogInterfaceFileName() string {
	return fmt.Sprintf("%s/log.go", GetLogRootFolderName())
}

func GetLogImplementationFileName() string {
	return fmt.Sprintf("%s/log_default.go", GetLogRootFolderName())
}
