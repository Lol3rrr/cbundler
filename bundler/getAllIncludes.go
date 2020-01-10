package bundler

import (
  "cbundler/files"
  "cbundler/utils"
)

func getAllIncludes(files []files.File, isHeader bool) ([]string) {
  result := make([]string, 0)

  for _, tmp := range files {
    if tmp.IsHeader != isHeader {
      continue
    }

    for _, tmpInclude := range tmp.Includes {
      if !utils.Contains(tmpInclude, result) {
        result = append(result, tmpInclude)
      }
    }
  }

  return result
}
