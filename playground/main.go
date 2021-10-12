package main

import (
  "fmt"
  "strings"
)

func main() {

  s := "default/infrastructure/log/infra_log.go"

  i := strings.LastIndex(s, "/")
  fmt.Printf("%s\n",s[i+1:])

}
