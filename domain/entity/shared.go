package entity

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// GetPackageName ...
func GetPackageName(rootFolderName string) string {
	i := strings.LastIndex(rootFolderName, "/")
	return rootFolderName[i+1:]
}

//// IsExist ...
//func IsExist(rootFolderName, typeName string, isWantedType func(expr ast.Expr) bool) (bool, error) {
//
//	fset := token.NewFileSet()
//
//	pkgs, err := parser.ParseDir(fset, rootFolderName, nil, parser.ParseComments)
//	if err != nil {
//		return false, err
//	}
//
//	// in every package
//	for _, pkg := range pkgs {
//
//		// in every files
//		for _, file := range pkg.Files {
//
//			// in every declaration like type, func, const
//			for _, decl := range file.Decls {
//
//				// focus only to type
//				gen, ok := decl.(*ast.GenDecl)
//				if !ok || gen.Tok != token.TYPE {
//					continue
//				}
//
//				for _, specs := range gen.Specs {
//
//					ts, ok := specs.(*ast.TypeSpec)
//					if !ok {
//						continue
//					}
//
//					// focus only to some struct for example
//					//if _, ok = ts.Type.(*ast.StructType); !ok {
//					//  continue
//					//}
//
//					if !isWantedType(ts.Type) {
//						continue
//					}
//
//					// entity already exist, abort the command
//					if ts.Name.String() != typeName {
//						continue
//					}
//
//					return true, nil
//				}
//			}
//
//		}
//	}
//
//	return false, nil
//}

func IsExist(fset *token.FileSet, rootFolderName string, foundExactType func(file *ast.File, ts *ast.TypeSpec) bool) bool {

	pkgs, err := parser.ParseDir(fset, rootFolderName, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// in every package
	for _, pkg := range pkgs {

		// in every files
		for _, file := range pkg.Files {

			// in every declaration like type, func, const
			for _, decl := range file.Decls {

				// focus only to type
				gen, ok := decl.(*ast.GenDecl)
				if !ok || gen.Tok != token.TYPE {
					continue
				}

				for _, specs := range gen.Specs {

					ts, ok := specs.(*ast.TypeSpec)
					if !ok {
						continue
					}

					// focus only to some struct for example
					//if _, ok = ts.Type.(*ast.StructType); !ok {
					//  continue
					//}

					if foundExactType(file, ts) {
						return true
					}

					// entity already exist, abort the command
					//if ts.Name.String() != typeName {
					//  continue
					//}
					//
					//fmt.Printf("File : %v\n", fset.Position(file.Pos()).Filename)

				}
			}
		}
	}
	return false
}
