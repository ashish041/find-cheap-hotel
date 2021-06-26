package logger

import (
	"io"
	"os"

	"github.com/maledog/logrot"
)

// WriteTo opens file for logging, and rotates and gzips it upon reaching
// maxSize and keeps up to maxFiles logs. If the file does not exist it
// is created with the given permission.
func WriteTo(file string, perm os.FileMode, maxSize int64, maxFiles int) (io.WriteCloser, error) {
	return logrot.Open(file, perm, maxSize, maxFiles)
}

// RotateOption specifies the log rotation parameters.
type RotateOption struct {
	Perm     os.FileMode
	MaxSize  int64
	MaxFiles int
}

// MustWriteTo opens file for logging, and rotates and gzips it upon reaching
// 200,000,000 bytes, and keeps up to 7 logs. If the file does not exist it is
// created with permission bits 0644. It panics on error. The RotateOption
// argument can be used to change the default rotation setting.
func MustWriteTo(file string, ro ...RotateOption) io.WriteCloser {
	o := RotateOption{
		Perm:     0644,
		MaxSize:  200 * 1000 * 1000,
		MaxFiles: 7,
	}

	for _, oo := range ro {
		if oo.Perm > 0 {
			o.Perm = oo.Perm
		}
		if oo.MaxSize > 0 {
			o.MaxSize = oo.MaxSize
		}
		if oo.MaxFiles > 0 {
			o.MaxFiles = oo.MaxFiles
		}
	}

	w, err := WriteTo(file, o.Perm, o.MaxSize, o.MaxFiles)
	if err != nil {
		panic(err)
	}
	return w
}
