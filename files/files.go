package files

type File struct {
  IsHeader bool
  Path string
  Includes []string
  Content string
}
