package main

import (
  "go/ast"
  "go/parser"
  "go/token"
)

func main() {

  fset := token.NewFileSet()
  pkgs, err := parser.ParseDir(fset, "pkg", nil, parser.ParseComments)
  if err != nil {
    panic(err)
  }

  ast.Print(fset, pkgs)

  //// in every package
  //for _, pkg := range pkgs {
  //
  //  // in every files
  //  for _, file := range pkg.Files {
  //
  //    // in every declaration like type, func, const
  //    for _, decl := range file.Decls {
  //
  //      // focus only to type
  //      gen, ok := decl.(*ast.GenDecl)
  //      if !ok || gen.Tok != token.FUNC {
  //        continue
  //      }
  //
  //
  //    }
  //  }
  //}


}
