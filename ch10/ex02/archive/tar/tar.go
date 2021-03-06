package tar

import (
	"archive/tar"
	"io"
	"os"

	archives "github.com/negibokken/golang-training/ch10/ex02/archive"
)

type reader struct {
	tarReader *tar.Reader
	file      *os.File
	toWrite   string
}

func (r *reader) Read(b []byte) (written int, err error) {
	for len(b) > 0 {
		if len(r.toWrite) > 0 {
			n := copy(b, r.toWrite)
			written += n
			r.toWrite = r.toWrite[n:]
			b = b[n:]
		}
		n, err := r.tarReader.Read(b)
		written += n
		b = b[n:]
		switch err {
		case io.EOF:
			h, err := r.tarReader.Next()
			if err != nil {
				return written, err
			}
			if h.Typeflag == tar.TypeDir {
				continue
			}
			r.toWrite = h.Name + ":\n"
		case nil:
			continue
		default:
			return written, err
		}
	}
	return written, nil
}

func NewReader(f *os.File) (io.Reader, error) {
	return &reader{tar.NewReader(f), f, ""}, nil
}

func init() {
	archives.RegisterFormat("tar", "ustar", 257, NewReader)
}
