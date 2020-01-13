package files

import (
  "errors"
  "strings"
  "io/ioutil"
)

func loadFile(path string) (File, error) {
  rawData, err := ioutil.ReadFile(path)

  header := true
  index := strings.Index(path, ".h")
  if index < 0 {
    header = false
  }

  cIndex := strings.Index(path, ".c")
  if index < 0 && cIndex < 0 {
    return File{}, errors.New("Not a fitting File")
  }

  data := string(rawData)
  includes := make([]string, 0)

  data = removeCustomImports(data)
  data, includes = extractStdImports(data)


  tmp := File{
    IsHeader: header,
    Path: path,
    Includes: includes,
    Content: data,
  }

  return tmp, err
}
