package files

import (
  "os"
  "path/filepath"

  "cbundler/utils"
)

func LoadDirs(dir string, blacklist []string) ([]File) {
  result := make([]File, 0)

  filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
    if info == nil || info.IsDir() {
      return nil
    }

    data, err := loadFile(path)
    if err != nil {
      return nil
    }

    if utils.Contains(path, blacklist) {
      return nil
    }

    result = append(result, data)

    return nil
  })

  return result
}
