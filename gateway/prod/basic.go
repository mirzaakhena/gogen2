package prod

import (
	"bufio"
	"bytes"
	"context"
	"github.com/mirzaakhena/gogen2/infrastructure/log"
	"golang.org/x/tools/imports"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type basicUtilityGateway struct {
}

// CreateFolderIfNotExist ...
func (r *basicUtilityGateway) CreateFolderIfNotExist(ctx context.Context, folderPath string) (bool, error) {
	if r.IsFileExist(ctx, folderPath) {
		return true, nil
	}
	if err := os.MkdirAll(folderPath, 0755); err != nil {
		return false, err
	}
	return false, nil
}

// WriteFileIfNotExist ...
func (r *basicUtilityGateway) WriteFileIfNotExist(ctx context.Context, templateFile, outputFilePath string, obj interface{}) (bool, error) {
	if r.IsFileExist(ctx, outputFilePath) {
		return true, nil
	}
	return false, r.WriteFile(ctx, templateFile, outputFilePath, obj)
}

// WriteFile ...
func (r *basicUtilityGateway) WriteFile(ctx context.Context, templateFile, outputFilePath string, obj interface{}) error {
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

// PrintTemplate ...
func (r *basicUtilityGateway) PrintTemplate(ctx context.Context, templateString string, x interface{}) (string, error) {

	tpl, err := template.New("something").Funcs(FuncMap).Parse(templateString)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	if err := tpl.Execute(&buffer, x); err != nil {
		return "", err
	}

	return buffer.String(), nil

}

// IsFileExist ...
func (r *basicUtilityGateway) IsFileExist(ctx context.Context, filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// Reformat ...
func (r *basicUtilityGateway) Reformat(ctx context.Context, goFilename string, bytes []byte) error {

	// reformat the import
	newBytes, err := imports.Process(goFilename, bytes, nil)
	if err != nil {
		return err
	}

	// rewrite it
	if err := ioutil.WriteFile(goFilename, newBytes, 0644); err != nil {
		return err
	}

	return nil
}

// GetPackagePath ...
func (r *basicUtilityGateway) GetPackagePath(ctx context.Context) string {

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
