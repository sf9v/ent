// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package internal

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x51\x5f\x6b\xdb\x3e\x14\x7d\xb6\x3e\xc5\xf9\x99\xfe\xa8\xdd\xa5\x4a\xdb\xb7\x0d\xf2\x50\xda\x0c\x32\xb6\x76\x90\xc2\x1e\xba\x52\x14\xfb\x3a\x11\x75\x24\xef\x4a\x29\x0b\x42\xdf\x7d\x48\x4e\xc2\xf6\x64\x4b\xe7\xdc\xf3\x47\x37\x84\xe9\x85\xb8\xb3\xc3\x9e\xf5\x7a\xe3\x71\x73\x75\xfd\xf1\x72\x60\x72\x64\x3c\x3e\xab\x86\x56\xd6\xbe\x61\x61\x1a\x89\xdb\xbe\x47\x26\x39\x24\x9c\xdf\xa9\x95\xe2\x69\xa3\x1d\x9c\xdd\x71\x43\x68\x6c\x4b\xd0\x0e\xbd\x6e\xc8\x38\x6a\xb1\x33\x2d\x31\xfc\x86\x70\x3b\xa8\x66\x43\xb8\x91\x57\x47\x14\x9d\xdd\x99\x56\x68\x93\xf1\xaf\x8b\xbb\xf9\xc3\x72\x8e\x4e\xf7\x84\xc3\x1d\x5b\xeb\xd1\x6a\xa6\xc6\x5b\xde\xc3\x76\xf0\x7f\x99\x79\x26\x92\xe2\x62\x1a\xa3\x10\x21\xa0\xa5\x4e\x1b\x42\xb9\x55\xda\x94\x88\x51\x4c\xa7\xb8\x4b\x79\xd6\x64\x88\x95\xa7\x16\xab\x3d\xce\xc9\xf8\xe6\x74\x75\x2e\x71\xff\x88\x87\xc7\x27\xcc\xef\x17\x4f\x52\x0c\xaa\x79\x53\x6b\x42\xd2\x10\x42\x6f\x07\xcb\x1e\x95\x28\x4a\xeb\x4a\x51\x94\xab\xbd\xa7\xf4\x13\x02\x3c\x6d\x87\x5e\x79\x42\x39\xb2\x5c\xb6\xcc\xd0\xc0\xda\xf8\x0e\xe5\xff\xbf\x4a\xc8\xef\x07\xc5\x18\x45\x9d\x63\x9e\xad\x94\x23\x7c\x9a\x21\x7f\x8f\x78\x9a\x7d\x57\x0c\xd7\x6c\x68\xab\x1c\x66\x78\x7e\x21\xe3\xe5\xc2\x78\xe2\x4e\x35\x14\xb2\x34\x2b\xb3\x26\x9c\xbd\x4e\x70\x66\xd4\x36\xcb\xc8\x07\xb5\x25\x97\xf4\x8b\x22\x84\xcb\x83\x7e\x8c\x32\x1d\x4e\x51\x5c\x88\xe5\x61\x26\xc6\x49\xd6\x22\xd3\xe2\x32\x46\x11\x85\xe8\x76\xa6\xc9\x9d\xab\x1a\x41\x14\x29\x48\xaf\x0d\x39\x3c\xbf\x3c\xbf\xa4\xd2\xa2\xe8\x2c\xe3\x75\x72\xc8\x97\x7c\xc7\x28\xc7\xbc\x41\x14\xc5\x6a\x02\x62\x4e\xd8\x37\xc5\x6e\xa3\xfa\x65\x06\xab\x91\x53\x8b\xa2\xd0\x5d\x66\xfc\x37\x83\xd1\x7d\x9e\x29\x3a\xa5\xfb\x8a\x98\x13\x9c\x2a\x8c\xbe\x33\xa8\x61\x20\xd3\x56\xf9\x38\xc1\xaa\x16\x09\xb5\x4e\x2e\x7d\x6b\x77\x5e\xfe\x60\xed\xa9\xca\xfb\x90\x5f\xac\x36\x47\xe2\x18\xb7\x2a\x7f\x9a\xb2\xae\xeb\x53\xb7\xa3\x4b\xb2\xb7\x9c\x4b\x8e\x5a\xc4\x3c\x6a\x2d\x3d\x6b\xb3\x4e\x1c\x39\x4f\x9c\xaa\xfe\x90\x45\x32\x71\xfe\x5b\xfb\xea\x3a\xcb\xfd\xb3\xfa\xb1\xd9\xb8\xf9\xc3\x8b\xc6\x28\xfe\x04\x00\x00\xff\xff\x95\x06\x0f\xa4\x50\x03\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 848, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x5a\x5f\x6f\xdc\x36\x12\x7f\x5e\x7d\x8a\x89\x81\x1a\x52\xb0\xd5\xf6\x8a\xa2\xb8\xdb\xdc\x1e\x50\xb4\x29\xea\xeb\xd5\x0d\x9a\xa4\x2f\x41\xe0\xca\x12\xe5\x65\x2c\x91\x5b\x92\xeb\xd8\x75\xf3\xdd\x0f\x33\x43\x4a\x94\x56\xbb\x76\x12\xdb\x2f\x11\x87\xf3\xf7\xc7\xe1\x70\xc8\xcd\x62\x01\xdf\xeb\xcd\x8d\x91\x17\x6b\x07\x5f\x7f\xf5\x8f\x7f\x7d\xb9\x31\xc2\x0a\xe5\xe0\xc7\xa2\x14\xe7\x5a\x5f\xc2\x89\x2a\x73\xf8\xae\x69\x80\x98\x2c\xe0\xbc\xb9\x12\x55\x9e\x2c\x16\xf0\x6a\x2d\x2d\x58\xbd\x35\xa5\x80\x52\x57\x02\xa4\x85\x46\x96\x42\x59\x51\xc1\x56\x55\xc2\x80\x5b\x0b\xf8\x6e\x53\x94\x6b\x01\x5f\xe7\x5f\x85\x59\xa8\xf5\x56\x55\xa8\x42\x2a\x62\xf9\xdf\xc9\xf7\xcf\x4f\x5f\x3e\x87\x5a\x36\x22\xd0\x8c\xd6\x0e\x2a\x69\x44\xe9\xb4\xb9\x01\x5d\x83\x8b\xec\x39\x23\x44\x9e\x24\x9b\xa2\xbc\x2c\x2e\x04\x34\xba\xa8\x92\x44\xb6\x1b\x6d\x1c\xa4\xc9\xec\x48\xa8\x52\x57\x52\x5d\x2c\xde\x59\xad\x8e\x92\xd9\x51\xdd\x3a\xfc\xc7\x88\xba\x11\xa5\x3b\x4a\x92\xd9\xd1\x85\x74\xeb\xed\x79\x5e\xea\x76\x51\xfb\x80\xa5\x2a\xb7\xe7\x85\xd3\x66\x21\x14\xf1\xdf\xc5\xb3\xb0\xe5\x5a\xb4\xc5\x42\x54\x17\xe2\x63\xf8\x6b\x29\x9a\xea\x63\x04\xa4\xaa\xc4\xf5\x51\x92\x25\x08\xdb\x4b\xa2\x81\x11\x7e\xc1\x2c\x14\x0a\x84\x72\xb9\x9f\x70\xeb\xc2\xc1\xfb\xc2\x12\x2e\xa2\x82\xda\xe8\x16\x0a\x28\x75\xbb\x69\x24\x2e\x8e\x15\x06\x3c\x76\x79\xe2\x6e\x36\x22\xa8\xb4\xce\x6c\x4b\x07\xb7\xc9\xec\xb4\x68\x05\x00\x20\x45\xaa\x0b\xa0\xbf\x3f\x10\xcd\xe5\x91\x2a\x5a\x31\xd7\xad\x74\xa2\xdd\xb8\x9b\xa3\x3f\x92\xd9\xf7\x5a\xd5\xf2\x02\xc8\x87\xf0\xed\x99\x4b\x1a\x0e\xd9\x9f\x57\x17\xc2\x02\xc0\x9b\xb7\x4f\xf1\x33\xd6\x8d\x40\xda\x21\xf7\x8f\x88\x95\x25\x6e\xfa\x8c\xb8\x09\xc6\x11\xfb\x09\x22\x25\x2c\xb2\xd3\x67\xc4\x2e\x79\x6a\xc8\xff\x93\xd6\x97\xde\x99\x17\xda\x4a\x27\xb5\x0a\xfc\x6b\x9c\x1a\x72\xbf\xd0\x8d\x2c\x6f\x00\xce\xb5\x6e\x00\x06\xb0\x6c\x68\x6a\xc0\xfe\x81\x96\xab\x53\x5b\x09\x5b\x1a\x79\x2e\x2c\x14\x40\xae\xc3\x26\x4c\xf9\xac\xe7\xd5\xf6\x6b\xd2\xc9\xf5\xab\xd2\x45\x04\x20\x95\x03\x58\x2c\x80\x31\xa1\xd0\x82\x16\xd6\xdd\x48\xeb\xf2\x64\xf6\x8b\xbc\x16\xd5\x89\x42\x11\x72\x7a\xb1\x80\x13\x55\xc9\xb2\x70\xc2\x82\xac\x23\x01\xcc\x98\x16\xb9\xbf\x94\x8a\x05\xa5\x3a\xf1\x7a\xd9\x16\x91\x86\xb6\x5a\x22\xb1\x2d\x0e\x97\x1d\xda\x4d\x4e\xa6\x7f\x42\x6e\xb2\xe0\x6e\x6a\xf2\x5f\x9c\xa0\xd1\xdf\xde\x5c\x3d\x51\xb5\xee\xd9\x9e\x52\xe8\xf9\xab\x9b\x8d\x88\x27\xbc\x34\x9a\x1f\x4a\xbf\x2a\x62\x53\x87\x6d\xbb\x62\x94\xf7\x2f\xe5\x5f\x91\xe3\x4f\xa5\x72\xdf\x7e\xb3\x4f\xd8\xca\xbf\x46\xa6\x9f\xab\x6d\x6b\x3b\xb6\x37\x6f\x27\x8d\x87\x6d\x84\xbc\x43\xf1\xd7\x4a\xfe\xb9\xed\xcc\xc7\xf9\xbb\x2b\xbe\x25\xde\xa1\xfc\xa9\x6c\x9a\xe2\xbc\x11\xf7\x91\x57\x9e\x77\xa8\xe1\xd7\x0d\x66\x73\xd1\xdc\x47\x83\xf6\xbc\x43\x0d\x3f\x88\xba\xd8\x36\xee\x5e\x31\x54\xcc\x3b\xa9\xe0\xf7\xa2\x41\x28\xa4\x72\xc2\x60\xd9\xbd\xfd\xb0\x5f\xc1\xd9\x15\x32\x8f\xb0\xdc\x54\x85\x13\xc1\x9b\x3b\xb0\x24\xde\xb3\x49\x77\x4e\xda\x76\xeb\x3a\x50\x0f\xeb\x91\x81\x77\xa8\xe2\xf7\xa2\x91\x15\x9e\x16\x94\x1a\xb4\x51\xf7\xaa\xb8\xea\x78\x47\x79\xe9\xb4\x29\x2e\xc4\xcf\xe2\x06\xee\x4c\x6a\xcb\xbc\x67\x97\xe2\x66\x5c\x19\x7d\xb5\xa2\xbf\xa7\xc3\xe1\x48\x49\x28\x7b\x23\x37\x84\x42\xf2\xd5\x7d\xd0\xb0\x81\x77\xa4\x82\xea\x27\xee\x66\xe4\x6d\x8b\xcd\x1b\x8e\xe6\xed\xc4\x31\xc6\xb5\xf6\x6c\x77\x8f\x7f\xa7\x94\x76\x05\xfa\x67\x87\x4a\xe2\x7c\xf1\x4a\x8a\x9e\x77\xa2\xf2\xd3\xe9\xb6\x5b\x09\x89\xfc\x09\x85\x90\xe4\x26\xeb\xe0\xce\x9a\xed\x2d\x7f\x01\x9b\x03\x42\x07\xab\xde\x5e\xa1\x71\xb1\xfb\x4d\xd4\x53\x3d\x44\x2c\x63\x44\x7d\xb6\xeb\xe1\x6f\xa2\x0e\x7c\x7d\x7b\xb0\x23\xb8\xb7\xb8\xed\xe4\xcd\x81\x9a\x76\xa2\xae\x84\xb1\xe2\x90\x98\x64\x96\xb1\x8b\x7f\x6e\xa5\x11\xd5\x01\x39\xe3\x59\xf6\xee\xb4\xa7\xd8\xea\xe4\x11\xe1\x8e\x0d\xc6\x19\xc5\xe7\xf1\x6e\x4a\x31\xfd\x13\x72\x8a\x05\xfb\xa4\x9a\x42\xf2\x00\x82\xa1\x95\x8b\x4f\xa4\xbb\x5b\xb9\x09\xee\xa9\x56\x2e\x02\xa7\x4b\xa1\x7b\xa1\x74\x2a\xde\x53\xe6\x94\x46\x50\x9b\x53\xa8\x80\x08\x3a\xc5\xb0\xd0\x17\x77\x64\x1b\xa7\x4d\x9e\xd4\x5b\x55\x06\xc9\x54\x54\x7e\x81\x7e\xe8\x38\x32\x9f\x8e\xb7\xc9\x4c\x09\x58\xae\xe0\x18\x87\xb7\xc9\x0c\xb7\xc7\x32\x2c\xbf\xa8\xf2\x57\xc5\xc5\x1c\xa9\x37\x1b\xb1\x8c\xa8\xb8\xab\x92\x19\xed\xda\x88\x8c\x43\x24\x33\xee\xcb\x40\xe6\x21\x4e\xf8\x1c\x5d\xfa\x09\x3f\xc4\x99\x90\x85\x4b\x9e\x09\x43\x9e\xaa\x3b\x3b\x34\x55\x07\x3b\x3d\xa6\x4b\x9c\xe9\x87\xf3\x64\xf6\x21\x99\xc9\x1a\x8c\xa8\x31\x3a\x16\x7b\x46\xc3\x27\x2b\x50\xb2\xc1\xc8\x67\x4a\x20\x19\x56\x1d\x52\x46\xd4\x19\xd3\xa3\xe5\x5a\x01\xf3\x45\x34\x52\x6f\x84\xdb\x1a\x05\x4a\xf4\x0b\xc5\xdd\xdd\xee\x4a\x71\x4f\x4a\x4b\xc5\x9f\x53\x6b\x45\xc2\x69\x5d\x85\x3e\x2e\x5e\xad\x94\xaf\x0b\x73\x10\xc6\xe0\xf8\x96\xa2\x13\xc6\x60\x74\x75\x95\x3f\x37\x26\xcd\x9e\x11\x21\x8a\x2f\x78\x28\x9b\x39\xd4\xad\x43\x2e\x6d\xea\x94\xf3\x13\xbe\xf8\x73\x09\x5f\x5c\x1d\xcd\x51\x9e\x10\x45\xf1\x8c\x42\xb3\x84\xda\x31\xd9\xbc\x1d\xad\x33\x40\x27\x40\x2b\x5a\xeb\xe1\x0c\x52\xe6\xa3\x44\xe2\x19\x9f\x4b\xd4\xfb\x2d\xe3\x09\xa2\x8c\x13\x87\xa7\xfa\xdc\x09\x3d\xdb\xb2\xf7\x21\x74\x66\xc9\xac\xeb\xc7\xfa\xd9\x40\xc1\x59\xdf\xdd\x2c\x7b\xbd\xa1\xdf\x61\xb4\xc8\x76\xdc\x07\x2d\xc9\xf6\xa0\x33\xea\x39\xbb\x4e\x67\xd9\xc5\xdc\xf5\x33\xa3\xac\xe4\xe9\x41\x62\x46\x5d\x0e\xcd\x37\x42\xa5\x75\x95\xf7\xd4\x8c\x94\x84\x9e\xa0\xb3\xd1\x51\x68\xba\xeb\x0d\x3a\x1b\x1d\x05\xe7\xa3\x63\x7f\x49\xe7\xfe\xa5\x48\xa7\x0f\xff\x8c\xf7\x4a\xad\x0d\x9c\xcd\xa1\x70\xb8\xf0\xa6\x50\x58\x59\xaa\x3c\xee\x1e\x30\x0f\x6c\x1d\x93\xde\x14\x8e\xf2\x20\xcd\xde\xc2\x0a\x0a\x17\x36\x9d\xad\x29\x09\x60\x75\x77\x26\xb6\xd2\x5a\xac\x85\x54\xbe\x25\x0a\xa1\x23\x21\x3f\x8f\xe6\xa8\x0b\x4d\x64\x9d\x6e\xbc\x76\x2c\x57\x40\xf7\x0d\xc4\x0d\xef\x21\xd9\x33\xa6\x3f\x59\xc1\x57\xc1\x4f\xba\x9f\xac\xe0\x18\x27\x48\x18\x0f\x1c\xbe\x21\xfa\xf6\x15\xa8\x19\x86\xb2\x50\x70\x2e\x80\x5e\x59\x44\x05\x4e\x13\xcf\x85\x50\xc2\x14\xb4\x3f\x51\xf2\x47\x6d\x40\x5c\x17\xed\xa6\x11\x73\x50\xda\xe1\xa5\x77\xab\x4a\x6a\x0a\x1b\x79\x29\xc0\xc9\x56\xe4\xa7\xfa\x7d\x4e\x5e\x9e\xcd\xc3\xde\xc4\x0a\x9f\xff\x52\x18\xbb\x2e\x9a\xb4\xcf\x3b\xbf\x57\x23\x84\x6c\x9d\x0f\x1a\xfa\x55\x94\xa5\x71\xb9\xb1\xf5\x1c\x65\xfa\x9a\xc3\x87\xde\x6e\xcd\xe1\x9b\x2d\xd5\x1c\xfe\x9c\xaa\x39\x24\x9c\xca\xea\x1a\x6f\x70\x95\xb8\x1e\x1e\x11\xac\xfa\xb6\xb3\x7d\x4c\x04\xf4\x96\x8e\x4a\xbf\x9d\x64\x75\x4d\x5d\x20\xed\x60\x3e\x15\x97\xdd\x04\x8f\xc7\x7b\x1b\x67\xfa\x9d\x1d\x6f\x18\x9c\x19\xd5\x71\x8e\xd4\x63\xe8\xdf\x76\x78\xb5\x68\xa5\xa2\xb7\xa2\x2e\xad\xf1\x4b\x43\x01\xff\x7d\xf9\xeb\x29\x0a\x53\x2f\xe1\x17\xba\x12\xbc\xd0\xc4\x82\x0a\xbc\xb0\x3e\x7f\x27\x4a\xe7\xff\xf1\x08\x0d\x8c\xa6\x36\xd8\xc6\x16\xc5\x5b\xca\x20\x3d\x87\x37\x6f\xcf\x6f\x1c\xd7\xcf\xa8\x40\x5b\xaa\xa1\x2c\x8b\x98\xf1\x63\xd2\x32\xbc\x8b\xf0\x30\xcd\xe2\x73\x54\x2a\x7e\x25\x4c\xfd\xdb\x1e\x1d\xb4\xbf\xd6\xde\x72\x96\xf9\xed\xd6\x1d\x6f\x3e\xc9\x6c\x8e\x6b\x4e\x0f\x1a\x81\xf5\xde\x67\x81\x0f\xaa\x3b\x0c\xec\xf8\x2c\x18\x9b\xe1\x15\x7d\x78\x3b\xdc\x62\x75\xb6\x8a\x5a\x50\x52\x05\x43\x9d\x23\x0f\x61\xcb\x57\x3b\xd1\x17\x3b\xb2\xce\x1b\x91\x93\x19\x2b\xda\x66\x23\x54\x95\x7a\xc2\xbc\xef\xa7\xa2\x5d\x92\x66\x99\x87\xc9\xbf\xc7\xc5\x01\xf8\xe7\xbb\xc7\x0c\x01\xb7\x6e\x17\x84\xf7\xc1\x87\x11\x1e\x0f\xa3\x40\x4e\x82\x93\xf1\xd6\x9f\x8c\x66\xb4\xe8\xf4\xb0\xf8\xf8\xb9\xc5\x2f\x92\x0f\x6f\xc7\x0b\x0e\x8a\xb1\xcd\x7c\x65\x79\xad\xda\x41\x6d\xe1\x02\x61\xf9\x18\x90\x57\x42\xc1\xf9\xb6\xae\x85\x01\x2a\x29\xbe\xba\x86\xc7\x4d\x2a\x13\x23\x0d\xe9\xf9\xb6\xf6\x35\x01\x3b\x37\x26\xce\xf7\x55\x86\x01\x0c\xe4\x61\xa7\x0e\x15\xcd\xc1\x1e\x06\x42\x18\x13\x27\x44\xdd\xa7\x83\xf5\xd5\x97\x44\xa2\x76\x31\xf7\x07\xa0\x9d\x68\x19\x77\x55\xa3\xee\xe8\xf8\x89\x4f\x9f\xae\xea\xd0\x97\xf5\xef\xa7\x4e\x7b\x74\xfc\xdd\x24\x2e\x97\x1e\xb0\xd4\x82\x87\x25\x83\x71\xe9\x1a\xd7\x57\x82\x0d\x7d\x23\xed\x83\xfd\x35\xa8\x78\x07\x76\x57\x0c\x91\x9c\x43\x1b\x6d\x19\x76\x99\x2e\x03\x78\xc1\xa7\xce\x62\xba\x06\xb7\xd7\x5d\xfd\x4d\x66\x33\x7f\xc5\x8b\xbd\xf1\x85\xb1\xbd\xce\x7a\xb8\x27\x90\x1d\xb6\x3f\x68\xbd\xcb\x5b\x15\x65\x2d\xfa\x4b\x0e\xbf\x1b\xac\x69\xdd\xaf\xe8\x0c\x5b\x01\x6f\xbf\xbf\x3e\x0c\x77\x33\xb2\x4d\xb8\xf2\xb1\xbe\x90\x33\xd8\xa2\x74\x2f\x61\x2b\x38\x0e\xdf\xac\x91\xca\x89\xef\x08\xde\xcd\x89\xe4\x5f\xeb\x89\xe8\x0c\x9f\xf5\xb3\xe8\x29\x7e\x09\x72\xde\x2b\x0f\xc9\x1a\x95\x2b\xdf\x3c\x80\xad\x03\x20\xfb\x0e\x89\x87\x06\x7d\xdf\xe1\xf0\x49\xa7\x03\x69\x3d\x74\x3e\x3c\x82\xf7\x7b\xcf\x85\xcf\x39\x18\xc8\x00\xff\x90\x14\x87\xc1\x87\xc3\x83\xe7\x7d\xef\x3f\x99\x0c\xde\xf3\x6f\x5c\x91\xef\x3f\xb1\x43\x0f\x98\x8f\xd9\xb8\xea\x0d\x4b\x9e\x4f\x54\xae\x79\x7c\x57\xf9\x84\x9a\x37\xe8\xa3\xf6\x16\xbd\xfd\x75\xe6\xa3\xcb\xde\x74\x15\xb9\x5f\x11\xd9\xbf\xac\xdd\x19\xb1\xb7\x3c\x04\x6c\x89\xe7\xae\x5d\xbe\x83\xf9\x24\x76\x71\x3b\xb2\x17\xba\x7d\x89\xfa\x91\xc0\x4d\xa5\xe1\x7d\xb3\xb0\x4b\x42\x4e\xac\x2e\x01\xeb\xa2\xe1\x67\xaf\x0f\xf7\x0e\x79\xd0\x1a\xed\x8d\xd9\xff\x6e\x1b\x07\x3d\xec\xa9\xee\x11\xb5\xcd\xfd\x0f\xc3\x2b\x60\x75\x9e\x77\xda\xcd\x1a\xf8\x69\x2a\x83\xbe\xab\xe8\xfd\x91\x35\x3c\xe9\x2e\xb6\xf0\xf7\xdf\x38\x3a\x51\xb5\xce\x4f\xb7\xad\x30\xb2\x4c\xb3\x51\x3f\x43\x1e\xa8\x39\xe8\x4b\x6e\x55\xe2\x3b\x71\x9e\xd6\x8d\x2e\xdc\xb7\xdf\x70\x14\x4f\xf4\x65\x2c\x1c\xd7\x97\xad\x12\xd7\x1b\x51\x3a\x51\x8d\x2e\xfb\xf4\xce\xd0\x3d\x31\x2c\xf9\x8d\x21\x7e\x62\xb0\xef\xa5\x2b\xd7\xe0\xd8\x3a\xb9\x8a\xe7\xff\x33\xb4\x54\x16\x56\x80\x83\xff\xac\x20\xfe\xa1\xd5\xfd\x13\x8e\x8f\xc1\xc1\xbf\x47\xe4\x6f\xbf\x59\x62\x25\x1b\xdf\xea\xf9\xe1\x42\x65\xd3\xea\x5e\xcb\x69\x7d\xaf\xe5\x5e\x85\xdb\x5e\xe3\x54\xc1\xea\x2b\x06\xbc\x37\xc5\xc6\xc6\x3f\xcd\x7b\x7a\xa1\x2a\xee\x83\x02\xa1\x15\x6e\xad\x2b\x78\x2f\xdd\x1a\x8c\x28\xf5\x15\x37\xbf\x42\xd9\xad\x11\xa0\x34\x6c\x0a\x25\x4b\x0b\x52\x81\xef\x54\xa5\xba\xf0\x65\x2e\xaa\x50\x75\x15\xfd\x2a\x09\x9e\x98\xc1\x9b\xb7\xfd\x2f\xe8\x1f\x32\x48\x7d\x31\x8a\xc8\xe3\x9b\x74\x25\xb0\xfd\x46\xf5\x3e\x5f\x64\x0d\x57\xb4\x2f\xd9\x39\xec\x63\xaf\x06\xc5\x89\x1e\x57\x06\x29\xf1\xc5\xab\x10\x1d\x3b\xdf\xbd\x7d\xce\xe1\x8a\x5a\x9c\x3a\x14\x26\xca\x42\xaa\xff\xd8\xe9\x85\xec\xaa\xf2\x10\xc0\x7c\x84\x2e\x37\x04\x3b\xe0\x32\xf9\x73\xa1\x8c\xef\xc0\x31\x9a\x4c\x0f\x60\xd2\x5b\x3e\x62\xc9\x9d\x4a\x4f\x7c\x0c\x24\x07\xf1\x0d\xc0\x64\x20\x85\x6f\x90\x26\x71\x8c\x85\x77\xa1\x0c\x9d\xc9\x0e\x98\x61\xe2\x73\xe1\x1c\xde\xc8\x63\x40\xc3\x4c\x80\x94\xdf\xbe\x10\x53\xd9\xfd\x27\x9c\x8e\xfe\x88\xb0\x86\x48\x27\x80\x95\x5d\xdf\x76\x08\xda\x2e\x90\x31\xb8\x7c\x53\xdb\x81\x96\xc9\x9f\x0b\xec\xa1\x1b\x5c\xca\xed\x1e\xe3\xf7\x4b\x7f\x8b\x7b\x14\xfc\x38\x9c\x09\xf4\xd8\x89\xc3\xd8\x71\x14\x3b\xc8\xf1\x61\xbf\x83\x1c\x93\x3f\x17\xb9\x41\x2f\x13\x25\x24\xd3\x43\x3a\xe2\x88\xb2\x91\x9b\x90\x9e\xf8\x88\x50\x72\x7c\x13\x50\xae\x7d\xf3\x73\x08\x4a\xef\xfe\x18\x4a\xdf\x5a\xec\x60\xe9\xe9\x9f\x0b\xe6\xc1\x2e\x29\xf5\xed\x0c\x92\x5f\x44\x8d\xd2\xa3\x80\xe7\x03\x9a\x40\x6f\x13\xba\xab\x43\xf0\xf9\x40\x7a\xfc\x28\xc4\xee\x6d\xc2\x41\xfc\x3a\x91\x0d\x46\x74\x6d\xd0\x06\x5c\xfe\xb3\x54\x55\x9a\xc1\x6a\xd5\xcd\xbf\x70\xd4\x96\xcd\x1c\xac\xc0\xe5\xcf\x1b\xd1\xa6\x83\xbe\xc1\x25\x1f\x92\xff\x07\x00\x00\xff\xff\xf4\xc4\xa9\x64\x4c\x2b\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 11084, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
