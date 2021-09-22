package vo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type gatewayMethod struct {
	packagePath string
}

func NewGatewayMethod(structName, gatewayRootFolderName string,packagePath string) (map[string]int, error) {

	gm := gatewayMethod{packagePath: packagePath}

	existingFunc := map[string]int{}

	err := gm.readStruct(structName, gatewayRootFolderName, existingFunc)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// existingFunc map[string]int dibuag dari parameter
func (obj *gatewayMethod) readStruct(structName, folderPath string, existingFunc map[string]int) error {

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, folderPath, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {

		// read file by file
		for _, file := range pkg.Files {

			importPaths := map[string]string{}

			for _, decl := range file.Decls {

				switch gd := decl.(type) {

				case *ast.GenDecl:
					err := obj.generalDecl(structName, gd, importPaths, existingFunc)
					if err != nil {
						return err
					}

				case *ast.FuncDecl:
					//ast.Print(fset, gd)
					if !obj.findAndCollectImplMethod(gd, structName, existingFunc) {
						continue
					}
				}

			}

		}

	}

	return nil
}

func (obj *gatewayMethod) generalDecl(structName string, gd *ast.GenDecl, importPaths map[string]string, existingFunc map[string]int) error {
	for _, spec := range gd.Specs {

		// handle import
		is, ok := spec.(*ast.ImportSpec)
		if ok {
			handleImports(is, importPaths)
		}

		// it is type declaration
		ts, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		// the struct name must have a 'Gateway' suffix
		if ts.Name.String() != structName {
			continue
		}

		// gateway must be a struct type
		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			continue
		}

		// if struct list empty then nothing to do
		if st.Fields.List == nil {
			break
		}

		for _, fieldList := range st.Fields.List {

			switch ty := fieldList.Type.(type) {
			case *ast.SelectorExpr: // struct is extend another struct

				expression := ty.X.(*ast.Ident).String()
				pathWithGomod := importPaths[expression]
				pathOnly := strings.TrimPrefix(pathWithGomod, obj.packagePath+"/")
				structName := ty.Sel.String()
				err := obj.readStruct(structName, pathOnly, existingFunc)
				if err != nil {
					return err
				}

			}

		}

	}
	return nil
}

func (obj *gatewayMethod) findAndCollectImplMethod(fd *ast.FuncDecl, structName string, existingFunc map[string]int) bool {
	if fd.Recv == nil {
		return false
	}

	// read all the function that have receiver with gateway name
	if fd.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).String() != structName {
		return false
	}

	// collect all existing function that have been there in the file
	existingFunc[fd.Name.String()] = 1

	return true
}
