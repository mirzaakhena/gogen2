package entity

import (
  "strings"
)

func GetPackageName(rootFolderName string) (string) {
  i := strings.LastIndex(rootFolderName, "/")
  return rootFolderName[i+1:]
}
