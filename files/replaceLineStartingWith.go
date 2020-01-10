package files

import (
  "strings"
)

func replaceLineStartingWith(src, key, replacement string) string {
  start := strings.Index(src, key)
  if start < 0 {
    return src
  }

  lineEnd := strings.Index(src[start:], "\n")
  if lineEnd < 0 {
    return src
  }
  lineEnd += start

  src = strings.Replace(src, src[start:lineEnd], replacement, 1)

  return src
}
