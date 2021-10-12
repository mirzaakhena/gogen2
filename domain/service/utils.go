package service

import (
  "fmt"
  "github.com/mirzaakhena/gogen2/gateway/prod"
  "github.com/mirzaakhena/gogen2/infrastructure/templates"
  "os"
  "strings"
  "text/template"
)

func CreateEverythingExactly(path string, fileRenamer map[string]string) error {

  ff := FileAndFolders{
    Folders: map[string]int{},
    Files:   make([]string, 0),
  }

  err := readFolders(path, &ff)
  if err != nil {
    return err
  }

  for folder := range ff.Folders {

    s := replaceVariable(folder, fileRenamer)

    err := os.MkdirAll(s, 0755)
    if err != nil {
      return err
    }
  }

  for _, fileRaw := range ff.Files {

    file := replaceVariable(fileRaw, fileRenamer)

    i := strings.LastIndex(file, "/")
    nameFileWithExtOnly := fmt.Sprintf("%s", file[i+1:])

    j := strings.LastIndex(nameFileWithExtOnly, "._")
    nameFileWithoutUnderscore := fmt.Sprintf("%s/%s%s", file[:i], nameFileWithExtOnly[:j+1], nameFileWithExtOnly[j+2:])

    fileOut, err := os.Create(nameFileWithoutUnderscore)
    if err != nil {
      return err
    }

    templateData, err := templates.AppTemplates.ReadFile(fileRaw)
    if err != nil {
      return err
    }

    tpl, err := template.
      New("something").
      Funcs(prod.FuncMap).
      Parse(string(templateData))

    if err != nil {
      return err
    }

    err = tpl.Execute(fileOut, struct{}{})
    if err != nil {
      return err
    }

  }

  //fmt.Printf("%v\n", ff)

  return nil
}

func replaceVariable(folder string, fileRenamer map[string]string) string {
  s := folder
  for k, v := range fileRenamer {
    s = strings.ReplaceAll(s, fmt.Sprintf("${%v}", k), v)
  }
  return s
}

type FileAndFolders struct {
  Folders map[string]int
  Files   []string
}

func readFolders(path string, ff *FileAndFolders) error {

  dirs, err := templates.AppTemplates.ReadDir(path)
  if err != nil {
    return err
  }

  for _, dir := range dirs {

    name := dir.Name()

    if dir.IsDir() {

      s := fmt.Sprintf("%s/%s", path, name)
      //fmt.Printf("ketemu folder %s\n", s)

      for k := range ff.Folders {
        //fmt.Printf("k=%v\n", k)
        if strings.Contains(s, k) {
          //fmt.Printf("remove %v from %v\n", k, ff.Folders)
          delete(ff.Folders, k)
        }
      }

      ff.Folders[s] = 1

      err = readFolders(fmt.Sprintf("%s/%s", path, name), ff)
      if err != nil {
        return err
      }

    } else {
      s := fmt.Sprintf("%s/%s", path, name)
      ff.Files = append(ff.Files, s)
    }

  }

  return nil

}
