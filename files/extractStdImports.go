package files

import (
  "strings"
)

func extractStdImports(src string) (string, []string) {
  result := make([]string, 0)

  lineStart := strings.Index(src, "#include ")
  for lineStart != -1 {
    start := strings.Index(src, "<")
    if start < 0 {
      break
    }

    end := strings.Index(src[start + 1:], ">")
    if end < 0 {
      break
    }
    end += start + 1

    include := src[start + 1:end]
    result = append(result, include)

    src = strings.Replace(src, "#include <" + include + ">", "", 1)
    lineStart = strings.Index(src, "#include ")
  }

  return src, result
}
