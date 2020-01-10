package utils

func Contains(key string, arr []string) bool {
  for _, tmp := range arr {
    if tmp == key {
      return true
    }
  }

  return false
}
