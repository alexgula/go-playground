package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _esc_localFS struct{}

var _esc_local _esc_localFS

type _esc_staticFS struct{}

var _esc_static _esc_staticFS

type _esc_file struct {
	compressed string
	size       int64
	local      string
	isDir      bool

	data []byte
	once sync.Once
	name string
}

func (_esc_localFS) Open(name string) (http.File, error) {
	f, present := _esc_data[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_esc_staticFS) Open(name string) (http.File, error) {
	f, present := _esc_data[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		gr, err = gzip.NewReader(bytes.NewBufferString(f.compressed))
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (f *_esc_file) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_esc_file
	}
	return &httpFile{
		Reader:    bytes.NewReader(f.data),
		_esc_file: f,
	}, nil
}

func (f *_esc_file) Close() error {
	return nil
}

func (f *_esc_file) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_esc_file) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_esc_file) Name() string {
	return f.name
}

func (f *_esc_file) Size() int64 {
	return f.size
}

func (f *_esc_file) Mode() os.FileMode {
	return 0
}

func (f *_esc_file) ModTime() time.Time {
	return time.Time{}
}

func (f *_esc_file) IsDir() bool {
	return f.isDir
}

func (f *_esc_file) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _esc_local
	}
	return _esc_static
}

var _esc_data = map[string]*_esc_file{

	"/favicon.ico": {
		local: "favicon.ico",
		size:  4286,
		compressed: "\x1f\x8b\b\x00\x00\tn\x88\x00\xff\xec\x931J\xc4@\x14\x86_\x18dA\xc1\x8dI\xb0\x13R\x89\xa57P\x10/\xe1%\xbc\x82\xe0!\xc4^M\xd8\x13\xd8\x19\xdcVl-4\x9a\xceB\x844B\x8a\xb0\xe3?\x1b\x04\x97M¼$\x934\xf9\x97\x8f\x85\to>\xde\x1b\x1e\x91\x85\x9f\xef\xab\u007f\x9fB\x9bh\x97\x88\x0e\x00\x8e蘊\xf3" +
			"e\xf0\xcd\xd9*\x183\xa6\xabġ\xb7o\xda\xf1\x1ez\xcf@rx\xbd\xb3\x8f\xfeկ|\xd3q\xbe\x04\xb6\xcfu\x96\xe3\xe6\\\u007f7\xdej\x86t\xd7\xf9\xfbꡬ&\x0e\xbc\xaf.\xdcM\xfd]\xb9\x9b\xf8\xdf\x02gO\xcb?sNM\xf8q\xf6\xd3W\xef\x15\xfe^\xdeݔ\xff\xe9\x9a6\xda\xecN\x1b\xff\xf7\x8d\xb3\xcdq\x9b\x9c" +
			"\u007f\xd3ā{2\xa4_Ŕ_\xf7\x1e\x13o\x80\xb9^q\xee\xd0\xf1ǡs\xae\u05f7{\xcf\xed\xe1cf\x1fj\xce\xe1\xb3\xea\x0e\x9d]\xac\x9f\xd9\xce\x19w\x9f\xda\xee_\xf9\xfc\x86\xf5\x17\xb3po\x87\xf4\xff\x05o\xba\xc9t<\xd6͒\xeb\x1f3\xc6D\xe4Z\x1e\xa4\x8ch\xbe \x129\xc8@j\xe5\"\x11`\x92\x8bH1\xcd" +
			"@\n\"p\x01P\x02\x92\xa9\xbcL&R\xa4b\x81\x9alY\xab\ue44aȚ\xaf\xbb\xa4\xfc\r\x00\x00\xff\xffR\x8e\x90z\xbe\x10\x00\x00",
	},

	"/": {
		isDir: true,
		local: "/",
	},
}
