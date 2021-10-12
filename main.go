package main

import "github.com/mirzaakhena/gogen2/domain/service"

func main() {
  //application.Run(registry.NewGogen2()())

  fileRenamer := map[string]string{
    "username": "mirza",
    "usecase":  "createorder",
    "gateway":  "prod",
  }

  service.CreateEverythingExactly("default/infrastructure/log", fileRenamer)

}
