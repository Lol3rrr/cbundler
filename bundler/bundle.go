package bundler

import (
  "strings"

  "cbundler/files"
)

func Bundle(name string, files []files.File) (string) {
  capsName := strings.ToUpper(name)

  result := ""
  result += "#ifndef " + capsName + "_H \n"
  result += "#define " + capsName + "_H \n"

  headerIncludes := getAllIncludes(files, true)
  for _, tmpInclude := range headerIncludes {
    result += "#include <" + tmpInclude + "> \n"
  }

  for _, tmp := range files {
    if tmp.IsHeader {
      result += tmp.Content
    }
  }

  result += "#ifndef " + capsName + "_IMPLEMENTATION" + "\n"
  result += "#define " + capsName + "_IMPLEMENTATION" + "\n"

  for _, tmp := range files {
    if !tmp.IsHeader {
      result += tmp.Content
    }
  }

  result += "#endif" + "\n"
  result += "#endif" + "\n"

  result = strings.ReplaceAll(result, "\n\n", "\n")

  return result
}
