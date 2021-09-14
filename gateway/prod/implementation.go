package prod

import (
	"bufio"
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/mirzaakhena/gogen2/infrastructure/log"
	"github.com/mirzaakhena/gogen2/infrastructure/templates"
	"golang.org/x/tools/imports"
)

type prodGateway struct {
}

// NewProdGateway ...
func NewProdGateway() (*prodGateway, error) {
	return &prodGateway{}, nil
}

func (r *prodGateway) CreateFolderIfNotExist(ctx context.Context, folderPath string) error {
	if r.IsFileExist(ctx, folderPath) {
		return nil
	}
	if err := os.MkdirAll(folderPath, 0755); err != nil {
		return err
	}
	return nil
}

func (r *prodGateway) WriteFileIfNotExist(ctx context.Context, templateFile, outputFilePath string, obj interface{}) error {
	if r.IsFileExist(ctx, outputFilePath) {
		return nil
	}
	return r.WriteFile(ctx, templateFile, outputFilePath, obj)
}

func (r *prodGateway) WriteFile(ctx context.Context, templateFile, outputFilePath string, obj interface{}) error {
	var buffer bytes.Buffer

	scanner := bufio.NewScanner(bytes.NewReader([]byte(templateFile)))

	for scanner.Scan() {
		row := scanner.Text()
		buffer.WriteString(row)
		buffer.WriteString("\n")
	}

	tpl := template.Must(template.New("something").Funcs(FuncMap).Parse(buffer.String()))

	fileOut, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}

	if err := tpl.Execute(fileOut, obj); err != nil {
		return err
	}

	return nil
}

func (r *prodGateway) IsFileExist(ctx context.Context, filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func (r *prodGateway) Reformat(ctx context.Context, goFilename string) error {

	// reformat the import
	newBytes, err := imports.Process(goFilename, nil, nil)
	if err != nil {
		return err
	}

	// rewrite it
	if err := ioutil.WriteFile(goFilename, newBytes, 0644); err != nil {
		return err
	}

	return nil
}

func (r *prodGateway) GetPackagePath(ctx context.Context) string {

	var gomodPath string

	file, err := os.Open("go.mod")
	if err != nil {
		log.Error(ctx, "go.mod is not found. Please create it with command `go mod init your/path/project`\n")
		os.Exit(1)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			return
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		if strings.HasPrefix(row, "module") {
			moduleRow := strings.Split(row, " ")
			if len(moduleRow) > 1 {
				gomodPath = moduleRow[1]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Error(ctx, err.Error())
		os.Exit(1)
	}

	return strings.Trim(gomodPath, "\"")

}

func (r *prodGateway) GetInportTemplateFile(ctx context.Context) string {
	return templates.InportFile
}

func (r *prodGateway) GetOutportTemplateFile(ctx context.Context) string {
	return templates.OutportFile
}

func (r *prodGateway) GetInteractorTemplateFile(ctx context.Context) string {
	return templates.InteractorFile
}

func (r *prodGateway) GetTestTemplateFile(ctx context.Context) string {
	return templates.TestFile
}