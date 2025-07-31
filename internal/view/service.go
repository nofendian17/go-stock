package view

import (
	"embed"
	"io/fs"
)

type Service interface {
	GetFS() fs.FS
}

type service struct {
	files fs.FS
}

func New(web embed.FS) Service {
	subFS, _ := fs.Sub(web, "web/dist")
	return &service{files: subFS}
}

func (s *service) GetFS() fs.FS {
	return s.files
}
