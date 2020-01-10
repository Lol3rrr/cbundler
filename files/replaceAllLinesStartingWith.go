package files

import (
  "strings"
)

func replaceAllLinesStartingWith(src, key, replacement string) string {
  contains := strings.Index(src, key)
  for contains != -1 {
    src = replaceLineStartingWith(src, key, replacement)

    contains = strings.Index(src, key)
  }

  return src
}
