package contract

import "io"

type IUpload interface {
	SetFile(io.Reader)
}

type Upload struct {
	File io.Reader
}

func (u *Upload) SetFile(r io.Reader) {
	u.File = r
}
