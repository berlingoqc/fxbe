package files

import (
	"io"
	"os"
)

const BufferSize = 512

// UploadFile create a new file with the contains of the io.Reader
func UploadFile(fullpath string, src io.Reader) error {
	dst, err := os.Create(fullpath)
	if err != nil {
		return err
	}
	buf := make([]byte, BufferSize)
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := dst.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}
