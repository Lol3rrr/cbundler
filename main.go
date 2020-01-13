package main

import (
  "fmt"
  "flag"
  "strings"
  "io/ioutil"

  "cbundler/files"
  "cbundler/bundler"
)

func main() {
  dirNamePtr := flag.String("dir", "src", "The root directory of the source code")
  namePtr := flag.String("name", "shl", "The name of the output")
  blPtr := flag.String("bl", "", "A comma seperated blacklist of blacklist files")

  flag.Parse()

  fmt.Printf("Starting... \n")
  fmt.Printf("Loading Folder: '%s' \n", *dirNamePtr)

  blacklist := strings.Split(*blPtr, ",")
  fileList := files.LoadDirs(*dirNamePtr, blacklist)
  bundleFile := bundler.Bundle(*namePtr, fileList)

  ioutil.WriteFile(*namePtr + ".h", []byte(bundleFile), 0644)

  fmt.Printf("Done \n")
}
