package fileserver

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

type FileServer struct {
	NewPath string
	dir     string
	path    string
	mode    int8
	w       http.ResponseWriter
	r       *http.Request
}

type FileInfo struct {
	IsDir bool  `json:"isdir"`
	Size  int64 `json:"size"`
}

func NewFileServer(dir string, path string, mode int8) *FileServer {
	return &FileServer{
		dir:  dir,
		path: path,
		mode: mode,
	}
}

// it won't close r.body
func (s *FileServer) Serve(w http.ResponseWriter, r *http.Request) error {
	if !vaildpath(s.path) || !vaildpath(s.NewPath) {
		return errors.New("not allow char in url")
	}
	if !strings.HasPrefix(s.path, "/") {
		s.path = "/" + s.path
	}
	s.w = w
	s.r = r
	switch s.mode {
	case ModeRead:
		return s.read()
	case ModeWrite:
		return s.write()
	case ModeCreateFile:
		return s.createfile()
	case ModeCreateDir:
		return s.createdir()
	case ModeDelete:
		return s.delete()
	case ModeRename:
		return s.rename()
	default:
		return ErrUnknowMode
	}
	// return nil
}

func (s *FileServer) read() error {
	fullpath := s.dir + s.path
	info, err := os.Stat(fullpath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		dir, err := os.ReadDir(fullpath)
		if err != nil {
			return err
		}
		dirs := make(map[string]*FileInfo, len(dir))
		for _, v := range dir {
			f, err := v.Info()
			if err != nil {
				return err
			}
			dirs[v.Name()] = &FileInfo{
				IsDir: v.IsDir(),
				Size:  f.Size(),
			}
		}
		return json.NewEncoder(s.w).Encode(dirs)
	}
	file, err := os.Open(fullpath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(s.w, file)
	return err
}

func (s *FileServer) write() error {
	fullpath := s.dir + s.path
	info, err := os.Stat(fullpath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return ErrIsDir
	}
	file, err := os.OpenFile(fullpath, os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, s.r.Body)
	return err
}

func (s *FileServer) createfile() error {
	fullpath := s.dir + s.path
	if _, err := os.Stat(fullpath); !os.IsNotExist(err) {
		return err
	}
	file, err := os.Create(fullpath)
	if err != nil {
		return err
	}
	return file.Close()
}
func (s *FileServer) createdir() error {
	fullpath := s.dir + s.path
	if _, err := os.Stat(fullpath); !os.IsNotExist(err) {
		return err
	}
	return os.MkdirAll(fullpath, os.ModePerm)
}

func (s *FileServer) delete() error {
	fullpath := s.dir + s.path
	info, err := os.Stat(fullpath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return os.RemoveAll(fullpath)
	}
	return os.Remove(fullpath)
}

func (s *FileServer) rename() error {
	fullpath := s.dir + s.path
	_, err := os.Stat(fullpath)
	if err != nil {
		return err
	}
	os.Rename(fullpath, s.dir+s.NewPath)
	return nil
}

func vaildpath(v string) bool {
	if strings.ContainsAny(v, ":*?|<>") {
		return false
	}
	if !strings.Contains(v, "..") {
		return true
	}
	for _, ent := range strings.FieldsFunc(v, func(r rune) bool {
		return r == '/' || r == '\\'
	}) {
		if ent == ".." {
			return false
		}
	}
	return true
}
