package types

import "io"

type FileReadData struct {
	File io.Reader
	Name string
}
