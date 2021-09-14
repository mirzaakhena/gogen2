package main

import (
  "github.com/mirzaakhena/gogen2/application"
  "github.com/mirzaakhena/gogen2/application/registry"
)

func main() {
  application.Run(registry.NewGogen2()())
}
