package files

import (
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/spf13/afero"
)

// ContextKey are the key use for the context parameter inside the Context
type ContextKey string

const (
	// RootKey is for the root directory
	RootKey string = "r"
)

// Ctx contains the information about the user regarding the files like right and root
type Ctx map[string]interface{}

// GetRoot get the root folder from the context
func (r Ctx) GetRoot() string {
	return r[RootKey].(string)
}

// FileInfo describes a file.
type FileInfo struct {
	Fs        afero.Fs    `json:"-"`
	Path      string      `json:"path"`
	Name      string      `json:"name"`
	Size      int64       `json:"size"`
	Extension string      `json:"extension"`
	ModTime   time.Time   `json:"modified"`
	Mode      os.FileMode `json:"mode"`
	IsDir     bool        `json:"isDir"`
	Type      string      `json:"type"`
}

// ListingDirectory list the contains of a directory
func ListingDirectory(ctx Ctx, folder string) ([]*FileInfo, error) {
	fullPath := path.Join(ctx.GetRoot(), folder)
	var files []*FileInfo
	f, err := os.Open(fullPath)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	for _, i := range list {

		item := &FileInfo{
			Name:      i.Name(),
			Size:      i.Size(),
			Extension: filepath.Ext(i.Name()),
			ModTime:   i.ModTime(),
			IsDir:     i.IsDir(),
			Path:      folder,
		}
		files = append(files, item)

	}
	return files, nil

}
