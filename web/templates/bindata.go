// Code generated by go-bindata.
// sources:
// web/templates/common/base.go.html
// web/templates/common/head.go.html
// web/templates/common/menu.go.html
// web/templates/monitor/details.go.html
// web/templates/monitor/index.go.html
// web/templates/monitor/menu.go.html
// web/templates/query/index.go.html
// web/templates/index/index.go.html
// DO NOT EDIT!

package templates

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _webTemplatesCommonBaseGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\x4a\x2c\x4e\x55\xaa\xad\xe5\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x01\xaa\xc9\x48\x4d\x4c\x51\x52\xd0\x03\xa9\x82\x48\xda\x24\xe5\xa7\x54\xa2\xaa\xc9\x4d\xcd\x2b\x85\xa8\x41\x16\x4d\xce\xcf\x2b\x49\xcd\x2b\x81\x6a\xd6\x87\x68\xb3\xd1\x87\x59\x91\x9a\x97\x02\x14\x07\x04\x00\x00\xff\xff\xfd\x8f\xc0\x67\x8d\x00\x00\x00")

func webTemplatesCommonBaseGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesCommonBaseGoHtml,
		"web/templates/common/base.go.html",
	)
}

func webTemplatesCommonBaseGoHtml() (*asset, error) {
	bytes, err := webTemplatesCommonBaseGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/common/base.go.html", size: 141, mode: os.FileMode(436), modTime: time.Unix(1512123257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webTemplatesCommonHeadGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x54\x4d\x73\xe2\x38\x10\xbd\xe7\x57\xb8\x7c\xe1\x90\xb5\x85\x81\x5d\x48\x0a\x67\x2b\x4b\x76\xf8\xc8\x17\x81\x04\xc2\x5c\xa6\x84\xd5\xb6\xe5\xc8\x92\x23\xc9\x80\x87\xe1\xbf\x8f\x1c\x86\x84\xca\x64\x3e\x72\x18\x4e\xea\x47\xf7\xeb\xf7\x9e\x55\x5a\xaf\x09\x84\x94\x83\x65\xc7\x80\x89\xbd\xd9\x1c\xb4\xcb\xc3\xc9\x81\x65\x7e\xed\x14\x34\xb6\x82\x18\x4b\x05\xda\xb7\x73\x1d\x3a\x2d\x7b\xff\xaf\x58\xeb\xcc\x81\xc7\x9c\x2e\x7c\x7b\xe5\xe4\xd8\x09\x44\x9a\x61\x4d\xe7\x0c\x6c\x2b\x10\x5c\x03\x37\x73\x14\x7c\x20\x11\xec\x26\x35\xd5\x0c\x4e\xd6\x6b\x37\xc3\x11\x7c\x7a\xaa\x36\x9b\x36\xda\xc2\x7b\xe4\x1c\xa7\xe0\xdb\x04\x54\x20\x69\xa6\xa9\xe0\x7b\x94\xf6\xf7\x8d\x0b\x0a\xcb\x4c\x48\xbd\xd7\xb5\xa4\x44\xc7\x3e\x81\x05\x0d\xc0\x79\x2a\xfe\xb2\x28\xa7\x9a\x62\xe6\xa8\x00\x33\xf0\xbd\x1d\x11\xa3\xfc\xc1\x8a\x25\x84\x7e\xa5\x34\xa5\x8e\x11\x0a\x0d\x8d\x72\x23\x21\x22\x06\x38\xa3\xca\x35\xe6\x50\xa0\xd4\xbf\x21\x4e\x29\x2b\xfc\xeb\x0c\xf8\xe1\x18\x73\x75\xd8\x11\x9c\x00\x57\x40\x8e\xeb\xd5\xea\x97\x67\xbc\x62\x49\x60\x7e\x45\xe9\x82\x81\x8a\x01\x74\xc5\xd2\x45\x06\x7e\x45\xc3\x4a\x97\x4c\x95\xfd\xe5\x65\xaf\xfd\xd2\x6b\x6f\xd5\xd8\x3b\x35\x29\x5e\x05\x84\xbb\x73\x21\xb4\xd2\x12\x67\x65\x51\x0a\x7a\x06\x50\xdd\xad\xbb\xcd\x92\xf6\x05\x73\x53\x6a\xba\x94\xb2\x8d\x6d\x0d\x91\xa4\xba\x30\x3b\x62\x5c\x6f\x35\x9c\xff\x26\x33\x4a\xc7\xfd\x0f\x70\xee\x91\x6e\x3a\x18\x9d\x3e\x14\x41\xde\x3b\xed\x8d\xa2\x7a\xed\x3a\xbd\x0b\x96\xcb\xa6\xe0\xf5\xd1\x8c\x44\x8d\x09\x3e\x1c\xa6\xe3\x5b\xf5\x19\x9d\xff\xd3\x5a\xcc\xc9\xff\x49\xdc\xc8\x4d\xce\x52\x28\x25\x24\x8d\x28\xf7\x6d\xcc\x05\x2f\x52\x91\x2b\xfb\x4f\x9b\x72\x74\x0c\x29\xfc\xcc\x9a\xec\x15\xe2\xca\xa3\x23\x35\xb9\x9f\x34\xf8\x59\x75\x90\x6b\xc6\xbb\x58\xb1\xce\x20\xef\x34\xf3\x65\x42\xf2\xe9\xd1\x78\x22\x2f\x16\xa3\x99\x10\xc3\xac\x36\x9f\xce\xa2\x34\x1a\xdc\xf4\xef\x97\x0c\x8d\xb3\x5f\x59\xdb\xde\x48\x4b\xc9\xc0\xb7\x11\xc2\x09\x5e\xbd\xbe\x26\x25\x86\x18\x9d\x2b\x94\x3c\xe6\x20\x0b\xe4\xb9\x9e\xe7\x56\xbf\x55\x4f\xda\x13\x43\xd7\x46\x5b\xaa\x37\x78\xdf\x1b\x51\xf2\xfa\xb3\x27\x6f\x46\x73\x1b\xfc\xdd\xbf\xa1\xf3\x6a\xad\xf9\xb8\x28\x92\xf1\x65\xd8\x4b\xae\x2f\xf1\xc5\x43\x98\x4f\x27\xab\x8f\xab\xbb\x21\xef\x0c\x4e\x9b\xac\x96\x76\xa6\x57\xfd\xac\x7b\x94\x76\x3b\x67\xad\x65\xf7\xaa\x1f\x0c\xcf\x9a\xb7\x2b\xfc\xe3\x68\x7e\xc3\x8b\xd1\x9e\x98\x7c\x98\xc8\x49\xc8\xb0\x84\x57\x51\x31\x41\xb0\x8a\x8d\x70\xd4\x70\xbd\xa6\xeb\xed\x80\x77\xa4\x45\xea\x66\x81\x90\x91\x39\xb8\x8b\xc6\x1b\x93\x6d\xb4\x7d\xde\xd6\x6b\xe0\xc4\xbc\x76\x5f\x03\x00\x00\xff\xff\x25\x6d\x7e\xc6\x00\x05\x00\x00")

func webTemplatesCommonHeadGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesCommonHeadGoHtml,
		"web/templates/common/head.go.html",
	)
}

func webTemplatesCommonHeadGoHtml() (*asset, error) {
	bytes, err := webTemplatesCommonHeadGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/common/head.go.html", size: 1280, mode: os.FileMode(436), modTime: time.Unix(1512123257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webTemplatesCommonMenuGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\x41\x4e\xc4\x30\x0c\x45\xf7\x3d\x85\x65\x36\xb0\x60\x7a\x81\xb4\x0b\x38\x08\x72\x13\x8f\x5a\xd5\x38\x28\x6d\x90\x50\x34\x77\xc7\x30\x4d\x80\x61\xf5\x63\xfd\x67\x7f\xc7\xa5\x04\x3e\x2f\xca\x80\xaf\xac\x19\x2f\x97\xce\x29\xbd\x83\x17\xda\xb6\x01\xed\x39\x51\x82\xab\x3c\x1a\x49\x59\x76\x1c\x3b\x00\x17\x96\x46\xf9\xa8\x3b\xd9\x8c\xf4\xed\xfc\xf5\x8e\xd6\x99\x29\x34\xdf\x08\xba\xf1\xa7\x44\x1a\x10\xe6\xc4\xe7\x01\xef\x10\xa2\x3e\xcb\xe2\xd7\x01\xe7\x65\xdb\x63\xfa\x38\x4d\xe4\xd7\xfb\x07\x1c\x9f\x4c\x5d\x4f\x47\x50\x6f\x49\xff\x33\x7d\x14\xa1\xb7\x8d\xeb\xde\xb5\xfe\x89\xcf\xf2\x2b\xbf\x62\x26\x8d\x28\x65\x92\xe8\xd7\xeb\x55\x5e\xbe\x3e\xc8\xba\x23\x9c\xec\x3e\x15\x60\x0d\xad\x72\x7d\x96\x9b\x95\x8e\x87\xeb\x6d\xec\xd8\x95\x02\xc6\x83\x35\x7c\x06\x00\x00\xff\xff\x42\x5b\x78\xaf\x71\x01\x00\x00")

func webTemplatesCommonMenuGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesCommonMenuGoHtml,
		"web/templates/common/menu.go.html",
	)
}

func webTemplatesCommonMenuGoHtml() (*asset, error) {
	bytes, err := webTemplatesCommonMenuGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/common/menu.go.html", size: 369, mode: os.FileMode(436), modTime: time.Unix(1512123257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webTemplatesMonitorDetailsGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x58\x5f\x6f\xdb\x38\x12\x7f\xf7\xa7\x98\xf3\x01\x1b\x09\xb5\xe5\xa4\x79\xbb\x24\x06\xee\xb6\xb8\x6b\x0e\xbb\xd9\x62\x5b\xb4\x0f\x77\x87\x82\x96\x68\x9b\xa9\x24\xfa\x48\xca\x8e\xcf\xd0\x77\xbf\x19\x52\x7f\x48\xc5\x8e\xb3\x05\x5a\x14\xb1\xc4\x99\xf9\x71\xf8\xe3\xcc\x70\xa8\xc3\x21\xe3\x4b\x51\x72\x18\xa7\xb2\x34\xbc\x34\xe3\xba\x1e\xdd\x66\x62\x0b\x69\xce\xb4\xbe\x1b\x2b\xb9\x1b\xcf\x47\x00\xfe\x18\xa9\x32\x34\x52\x56\x32\x94\xe5\xd3\x22\x9b\x5e\xbd\x6d\x65\xeb\xab\xf9\xe1\x90\x18\x61\x72\x5e\xd7\xb7\x33\x7c\x7d\x66\xb3\x61\x25\xcf\xc1\xfe\x9d\xa2\x3b\xac\xca\x4d\x63\x7d\x5a\x6f\xcd\x59\x26\xca\x55\xa7\x47\x33\x5d\xcf\x3f\xb1\x45\xce\x41\x1b\x66\x84\x36\x22\xd5\x38\xdf\x75\x87\x34\x43\xa8\x93\xb0\xd3\x85\xcc\xf6\x3e\x9a\xb1\x50\x8d\x8a\x7b\xb1\x7f\xa7\xda\x28\xb1\xe1\x99\xa7\x4b\xda\xe4\x8f\x3f\x42\x63\x2a\x1c\xb0\x6a\x60\x99\xa0\x49\x15\x3e\x09\x59\x42\x59\x15\x0b\xa4\x12\x90\x69\x8d\xae\xdc\x8d\x91\xba\x0f\xad\xf4\x76\x66\xd6\x2f\xa0\x7c\xc4\x95\x56\x1a\xe4\x12\xd0\x01\xd8\x28\x99\x72\xad\xe5\x00\x8c\x94\xf8\x19\xa0\x07\xeb\x04\x01\x15\x88\xc0\x56\x5c\x43\xce\x56\x2b\x64\x18\x16\x7c\x2d\xca\x0c\xde\x7f\xf9\x35\x84\xfd\x6d\xb9\xd4\xdc\x4c\x7f\x61\xab\x33\xd8\xef\xc5\x6a\x0d\x3b\xf4\x41\x41\xc1\xd4\x37\x88\xa4\xb5\x6c\xbd\x2e\xf9\x93\x69\x67\xc5\xc9\x68\x4a\x0c\x30\x5d\x15\x3c\x8b\xc3\x19\xd1\x85\x53\x53\xf9\x6a\x5f\x94\x30\x7c\xfa\xfb\x0b\x8b\xc6\x20\x75\xda\xd7\xe3\xf9\x7d\xb9\xa9\xcc\x59\x45\x5a\x6f\x65\x8e\x6a\xe2\x88\x7a\xcd\xce\xcf\xd5\x69\x8f\xe6\x8b\xbd\xe1\xfa\x94\x30\xe3\x39\xdb\x9f\x12\x7e\x07\xea\xd0\x63\xd2\x18\x84\xef\xad\xa1\x7c\x00\x91\x79\x91\xfa\x59\xf0\xdd\x78\x60\x47\x5a\x5e\xd2\xcc\x6c\x8a\x0c\x32\xce\xe5\xbb\x4e\x31\x6d\x0c\x98\xfd\x06\x43\xc2\xe0\x9e\xcf\x1e\xd9\x96\xb9\x51\x44\x1d\x35\x36\x5b\xa6\x30\xf2\xb4\xe9\x32\x00\xee\x20\xbb\x4e\x72\x99\xb2\x3c\x8a\x6f\x7c\x35\xc5\xcb\x8c\xab\x4e\x51\xa3\xe6\xb2\x2a\x53\x7a\x8e\x3a\xa7\x75\x7c\x18\x75\xfe\x91\x55\xb5\xc9\x90\xb1\xce\xea\x83\xad\x29\x9e\x25\x4a\x19\xda\xf4\xab\x24\xa3\x0e\xee\x3e\x23\x7f\x50\xe5\x5f\x97\xff\xb9\x19\x28\x51\xd9\xd1\xad\xf8\xea\xa8\xb8\x22\xf9\x58\x55\x65\x49\xb5\xcb\xd7\x10\xcb\xe8\x4f\x16\x20\xb1\x35\x2c\xf9\x9d\xa7\x72\xcb\x15\xe6\xc0\x21\xd8\x3b\x0f\xc5\x69\x0c\x81\xea\x10\xd4\xc7\xc4\x42\x90\xe7\x2f\x20\x6a\x27\x3f\x09\x47\xab\x70\xb9\x8b\x49\x8f\x06\x3e\xf6\xfb\x5d\x01\xd3\x60\xc4\xd5\x07\x1c\xbc\xba\x19\x8d\x3c\x94\xd9\x0c\x44\x89\x64\xb2\x5c\xfc\x8f\x53\x05\x28\x60\xb1\x87\xa6\xee\x0f\x66\xdb\x51\x2a\x53\x26\xe3\x6c\x97\x43\x42\x05\x65\xee\x8b\xc2\xbf\x51\xfc\x1f\x93\x4a\x9b\xcb\xa7\x6c\x9d\xf4\xa4\xb1\x85\x7e\x47\x39\x89\xd2\xe9\xd5\x4d\xb8\x36\xa4\x40\x19\xd8\x09\xac\x1e\x0c\xd2\x35\x9d\x93\x19\xc8\x0d\xc7\x3c\xa5\x68\xc6\xff\x54\xf4\x6c\xd5\xa1\x0a\xe8\x18\x1b\xd0\xe3\x0a\xa3\x3b\x74\x84\x06\xd5\xc6\x02\x30\x2c\xc4\x3b\x7e\xa1\x38\x64\x42\x6f\xd0\x03\xaa\x96\xac\x2f\xfc\x21\x4e\xce\xcd\x05\x1e\x0e\x85\x30\x3d\xde\x03\x2b\x38\x0e\x16\x1c\x4f\xb1\x54\xc3\x52\xc9\xa2\xc5\xdf\x4f\x80\x69\xd2\xc4\xdd\x90\x50\x4a\x43\xee\x97\x2b\x1e\x82\xb2\x72\x5f\x48\x74\x80\x5c\x79\xac\xb4\x2b\xdd\x98\x55\xc8\x9a\x02\x5b\xa9\x90\xa4\xbc\x42\xc7\xdd\x4a\x71\x06\x6a\x2c\xb2\x63\x24\xfe\x4c\xf4\x20\x89\x5f\x13\x4b\x54\x13\xac\x96\x9b\xf8\xe6\x74\x1c\x77\xb9\x01\x3f\xfd\x04\xd8\x5c\xb8\x42\xf0\x09\x2b\x0b\xd4\x35\xdc\x61\x24\xf7\x27\xe1\x20\xd8\x83\x69\xfb\x97\x84\x48\x8a\xa8\x4f\x69\x39\xaa\xeb\xf8\x54\x1a\xb8\x1d\xda\x71\x62\x69\xcd\xb6\xdc\xad\x17\x99\x93\xb0\xe2\x8e\x6a\xb6\x5d\x59\x6a\x27\xcd\x1e\x18\xf6\x0d\x09\x31\xcf\xf7\x19\x51\x70\x23\x28\x42\x4a\x64\x14\x78\xb1\x31\x7b\xc8\xb1\x79\x99\xa0\x3a\xec\x64\x95\x67\x90\x2a\x4e\x81\xca\xe0\x81\x3d\x84\xa4\x78\x0b\xd0\x98\x4c\x51\x9c\x58\x57\xa2\x18\xe6\x70\x79\x6c\xe5\x6d\xd4\x7a\x86\xcb\x9c\x99\x5f\xd9\x26\x1a\x5b\xd9\x38\x4e\x0a\xce\x4a\x0f\x69\x06\x57\x97\xf6\x5f\x48\xc7\xe8\x58\xa8\xa5\x6b\x9e\x7e\x6b\x96\x65\x99\xa1\xd0\xe4\x5b\x21\xb1\xbe\x50\x55\xa4\x7a\x40\x71\xa3\x1d\x4d\x05\x51\x86\xbf\x62\x10\xb7\xa9\x54\x18\x92\xa6\xe1\x75\x10\x38\x74\x38\x7c\x6c\x2a\x6d\x70\x50\x24\x48\x7e\x64\xd6\x42\xf7\xa7\x44\x43\x53\x67\x32\xa0\xc4\x4e\xb5\xd9\xdb\x1d\x73\xc5\x1b\xfd\x61\x5b\x29\x32\x28\x64\x26\x96\x36\xbb\x70\x1b\x90\xad\x9c\xa5\x3c\xb0\x25\x57\xd2\x4a\xb5\x9e\xfc\xf3\xe3\x6f\x0f\x09\x9e\x12\x9a\x47\xf6\x91\x9a\xc4\x72\x85\x10\x2e\x70\xe3\x20\x98\xec\xcc\x88\xbf\xa4\xf9\x3a\x82\x9c\x07\x18\x98\x1a\xc3\xbb\xcc\xf4\xb3\xe9\x8c\x28\xf8\x3b\xb2\xba\x83\xa8\xe4\x3b\x78\x87\x51\x11\xb5\x3e\x24\x0f\x72\x17\x63\xb1\xed\x04\xdd\xa2\xad\xa4\xdd\xc7\xe4\x32\x74\xc4\x2f\xb0\x3d\xd6\xa0\x80\xf7\x50\xbe\x80\x20\x5b\x97\x02\xc6\xa1\x3b\x4c\xdd\x0e\x92\x42\x24\x17\x8f\x9f\xe9\x65\x02\x5a\xa5\xf6\x29\x86\xc3\xa0\x65\x51\xdc\x54\xaa\x7c\x36\x0c\xf0\xb3\xac\x4a\xf3\x17\x68\x31\x12\xfb\x4e\xc7\x4d\x03\xe5\x06\x26\xcf\xec\x6c\x09\xf7\xec\x5c\x49\xf7\xec\xec\xc0\xd0\xae\x0e\x29\x0a\x63\xdd\x6d\xc5\x4a\xc9\x0a\xef\x00\xae\x8a\x53\xfd\x2a\xb8\x5a\xf1\x2f\x58\xf4\x7b\x12\xad\x70\xe2\x91\xd7\x0c\x74\x9c\x0c\x62\xc2\x3f\xce\xda\x82\xe8\xcf\x13\xf7\xa9\x6a\x97\x8b\xa9\x8a\x4d\x72\x90\xa9\xfd\x76\x3c\x03\x6e\x4f\xb3\x33\xc8\x56\xed\x45\xe4\x53\x5c\xb8\xf6\xf8\x14\x19\x4e\xea\xb3\xd1\x8e\x9c\xa2\x23\x38\xa2\x07\x5e\x3b\xdb\xef\x26\x24\x3c\xdf\xcf\x61\x9f\xa7\xc4\x8f\x95\xfe\x39\x2c\x4e\xba\x29\x4e\x13\x97\xe7\x61\x8d\x6a\x02\xff\xe2\xd6\x64\xf3\x8b\x37\x5e\xb3\xf9\xe6\x02\x5b\xea\x6c\xfe\xef\xf2\xe2\x4d\xb0\x82\x46\xd3\xb5\x6e\x67\x94\xba\x8e\xed\x15\x60\x5e\x37\x77\x46\xbb\x2b\x1c\x89\x91\x7f\x17\x4f\x3c\x8b\xde\xc6\x67\x4c\xba\x00\xff\x83\x26\x76\x03\x3a\x9b\xcb\x57\xd9\xd8\xc3\x2c\x98\x07\x0a\xfd\x32\x4d\x5d\xbc\xbd\xde\x3d\x2f\x90\x8e\xfa\xd7\x87\x46\xed\xed\xf7\xe1\x80\x27\x24\xff\x2f\xf8\x7d\x8b\xd7\xb1\xd4\xc3\x96\xdb\x06\x26\x06\x6a\x7f\xa7\x49\xfe\x41\x43\x37\x1e\x22\xcf\x35\x7f\x8d\x61\x7f\x5b\x0a\xac\xcb\x0c\x8d\x47\xcf\x6c\x7f\xc1\x46\xc4\x66\x88\x91\x1f\x98\x50\xda\xe5\x08\xa6\x82\x54\x26\xea\x6e\x4c\x6c\xb2\x88\x0f\x4d\x08\xdb\x23\xf0\xbe\x34\x11\x5d\x90\xe8\x3c\xea\x06\x16\x34\x70\xe3\x77\x55\x78\x08\xba\xcb\xd8\x04\x78\x49\x5f\x07\xa8\x35\x50\xbc\xc0\xde\xce\x76\x0b\x81\x3f\x99\xbb\x07\x6a\x9e\x63\x67\x10\x8d\xff\x1c\xde\x4a\xe3\x46\xf0\xd7\x3c\x8f\xc6\x49\x27\x5b\xc8\x27\x14\x11\x56\xd4\xad\x67\xe2\x5d\xf5\xe2\x43\x9b\x7a\x19\xdd\xe7\xc0\xf7\x2e\x4b\xd6\xa6\xc8\xa3\x63\xd7\x45\x5f\xeb\xc7\xbb\x94\x58\x72\xb0\xf4\xb0\xcd\x06\x37\x2a\x1a\x1b\xec\x6a\x13\xfb\x61\x0a\xc3\x6d\x1c\x20\x4f\xb0\x9b\xc4\x93\xf5\x75\xae\x27\xfc\x09\x9b\xde\x38\x71\x94\x47\xf1\x68\x18\xaa\x3b\x51\x66\x72\x47\xa5\xeb\x9e\x5c\xc0\xd2\xd7\x6f\xba\xd7\x4a\x21\x07\x8f\x1a\x87\xc6\xd8\x3f\x2f\xb0\xc7\xfb\xba\x61\x66\x5d\xd7\x33\x5a\xe4\x2c\x68\xcf\x71\x10\xdf\x71\x3b\x75\x22\xb2\xa7\xba\x46\x7f\x87\x97\xf8\xce\xc1\x7a\x02\x6f\xb1\x61\xe9\x0b\x25\x75\x6b\x78\x3f\x75\x0d\x99\xbd\x3f\xe6\xfb\xd1\x8f\xf3\xe0\x76\xe6\xbe\x4d\x34\x9f\x2b\xdb\x8f\x87\xcd\x43\xf3\xd3\xa6\xce\xff\x03\x00\x00\xff\xff\x49\x42\x44\xfa\x47\x15\x00\x00")

func webTemplatesMonitorDetailsGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesMonitorDetailsGoHtml,
		"web/templates/monitor/details.go.html",
	)
}

func webTemplatesMonitorDetailsGoHtml() (*asset, error) {
	bytes, err := webTemplatesMonitorDetailsGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/monitor/details.go.html", size: 5447, mode: os.FileMode(436), modTime: time.Unix(1512123376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webTemplatesMonitorIndexGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x95\x51\x6f\xd3\x30\x10\xc7\xdf\xf7\x29\xac\xa8\xf0\x44\x13\x31\xed\x69\xa4\x91\x00\xa1\xaa\xa3\x5a\x2a\x48\x79\x45\x6e\x7c\x6d\x2c\x32\xdb\x72\x9c\x15\x14\xf9\xbb\x63\xa7\x49\x48\x9d\x96\x3c\x94\x4d\xec\xa5\xb5\x7d\xff\x3b\xff\x7c\xbe\xf8\xaa\x8a\xc0\x96\x32\x40\x5e\xca\x99\x02\xa6\x3c\xad\xaf\x42\x42\x1f\x51\x9a\xe3\xa2\x98\x79\x92\xef\xbd\xe8\x0a\xa1\xfe\x9a\x95\x62\xe3\x24\x6b\x8b\x6b\xcb\xa7\x0f\x64\xfa\xf6\xba\xb1\x19\x6b\x76\x1d\xad\x24\x4f\xa1\x28\xb8\x2c\xc2\xc0\x4c\x1b\x4b\x55\x4d\x36\xb8\x80\xef\x02\xab\x0c\xdd\xce\x90\xdf\xcd\x0c\x44\x2b\x91\x98\xed\x00\x4d\x28\x23\xf0\xf3\x0d\x9a\x88\x36\x52\xed\xd0\xcd\x8a\xce\x63\x0c\xc7\x55\x08\xcc\x20\x47\xf5\xef\xd4\xe4\x02\x97\xb9\x3a\xd2\x9e\x50\x4f\x33\xc0\x84\xb2\x9d\xa3\xb3\x47\xbd\x39\x16\x2a\xaa\x72\xf0\xa2\x10\xa3\x4c\xc2\x76\xe6\xf5\x4f\xac\x75\xd0\xe1\x07\xc6\x50\x9f\x50\x6b\x2f\x32\xe3\xce\xe0\xcf\x25\x16\x99\xf9\xe5\xa5\xd0\x3a\x0c\xf0\x60\xcb\x20\xbb\x71\x70\x03\xc3\x3b\x76\x82\x0d\x27\xbf\x06\xf8\x55\x85\xf6\xd4\x5c\xc5\x64\x67\x37\xb5\x09\x76\x41\x50\x2f\xcd\x4d\xe8\x32\x77\xc3\xf4\x6e\x0d\xc8\x0e\x16\xc4\xde\x9b\x1d\xd5\x11\xeb\xd8\xfe\x82\x89\x52\x7d\x55\x12\xf0\x43\x31\x88\x69\xa2\xe6\x34\x5a\xdc\xaf\xd6\xc9\xad\x2d\x12\xeb\xeb\x27\x5c\xd0\x54\x6b\xf4\x5a\x62\x29\xdf\x99\x65\x21\x29\x53\x5b\xe4\xbd\x4a\xbc\x43\x78\xff\x23\x27\x90\xda\x2c\x19\xef\x13\x4c\xc0\xc8\x89\xad\xc6\x59\xef\xb8\xd9\x28\xc1\x9b\x1c\xce\xa1\xde\xc5\x8b\xfb\xff\x02\x75\xc9\xf9\x8f\x52\xfc\x95\x75\x19\xc7\x9f\xd7\xab\xe7\x80\x6d\x6b\x89\x37\x44\x3d\xce\x79\xb7\x78\x86\x72\xfe\x25\x5e\xaf\x92\xf7\x1f\x96\x9f\x6a\xd2\x3f\x41\x46\x78\x7b\xc2\x8b\xa8\x73\xce\xc5\x71\x5e\xc5\xa1\x5a\xcf\x67\xf5\x90\x53\xeb\x38\xc2\x58\x4b\x9e\xa8\x00\xe2\x52\x8d\x7e\x58\xf1\x3a\x79\x86\x2f\x2b\x0c\x86\x2f\x83\x79\x60\x8c\xd4\x7d\x44\x06\x0f\x96\xb3\x70\x34\xed\x6f\xd6\x19\x9a\xc1\xc5\x6d\xea\x1b\x85\xfd\xc5\x1d\xea\xd1\x04\xa9\xb5\x76\xf0\x32\xfb\x92\x25\x77\x5b\x92\x5d\x6b\x4b\xe5\x49\xfb\xd0\x3f\x2b\x87\xe6\xaf\x95\xfc\x0e\x00\x00\xff\xff\x33\x4f\x5a\x19\xeb\x08\x00\x00")

func webTemplatesMonitorIndexGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesMonitorIndexGoHtml,
		"web/templates/monitor/index.go.html",
	)
}

func webTemplatesMonitorIndexGoHtml() (*asset, error) {
	bytes, err := webTemplatesMonitorIndexGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/monitor/index.go.html", size: 2283, mode: os.FileMode(436), modTime: time.Unix(1512123257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webTemplatesMonitorMenuGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x8e\xc1\xaa\xc2\x30\x10\x45\xf7\xef\x2b\x86\xd0\xc5\x13\xa4\xd9\x4b\xda\xad\x3f\xe0\xbe\xc4\x76\x6a\x03\x3a\x09\x49\xab\x42\x98\x7f\x77\xa2\xd0\x55\x11\x37\xc9\xc0\x3d\xdc\x7b\x72\x1e\x70\x74\x84\xa0\x6e\x48\x4b\xd7\x7b\x9a\x91\x66\xc5\xfc\x97\x73\xb4\x74\x41\xa8\x1c\x0d\xf8\xdc\x43\x15\xa2\xef\xe1\xd0\x40\x5d\x0e\x4c\xc9\xc7\x24\x18\x80\xb9\xba\xd6\x58\x98\x22\x8e\x8d\xca\xb9\xaa\xcf\x36\x61\x17\xec\x3c\x31\xeb\x95\xd5\x92\xbc\x9b\x98\x55\x2b\x77\x09\xea\x63\xb4\x61\x92\xd7\x2f\x81\x19\xfe\x57\x78\x67\xb4\x6d\x8d\x96\x62\xd1\x40\x1a\xbe\xea\xdc\x1d\x3e\x7e\x30\x29\xd8\xa6\xc4\xc9\x07\xd7\x97\xf9\x42\x6c\x2e\x7f\xfe\x57\x00\x00\x00\xff\xff\xfc\x37\x76\xe3\x2b\x01\x00\x00")

func webTemplatesMonitorMenuGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesMonitorMenuGoHtml,
		"web/templates/monitor/menu.go.html",
	)
}

func webTemplatesMonitorMenuGoHtml() (*asset, error) {
	bytes, err := webTemplatesMonitorMenuGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/monitor/menu.go.html", size: 299, mode: os.FileMode(436), modTime: time.Unix(1512123257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webTemplatesQueryIndexGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\xcb\x6e\xdb\x3a\x10\xdd\xe7\x2b\x26\x82\x71\xb3\xb8\x95\x85\x26\xbb\x46\x56\x37\xcd\xaa\x40\x81\xa2\x8b\x2e\x03\x5a\x1c\x4b\x44\x69\xd2\xe5\x23\x8e\x61\xf8\xdf\x3b\x12\x25\xc5\x96\x69\x25\x45\x51\xa0\x02\x2c\x7b\x34\x0f\x9e\x39\x3c\x1c\x6b\xbf\xe7\xb8\x12\x0a\x21\x29\xb5\x72\xa8\x5c\x72\x38\x5c\xe5\x5c\x3c\x41\x29\x99\xb5\x8b\xc4\xe8\x6d\x52\x5c\x01\x5d\xc7\x4f\x9b\x60\x46\x69\xa6\xf3\x8d\xfd\x2f\x59\x31\x6f\xa9\x65\xba\xe6\xe9\xfb\xdb\x51\x0c\x45\xd5\xb7\xc5\x57\x8f\x66\x07\x9f\x71\x67\xf3\x8c\xcc\x51\xc4\x7e\x2f\x56\x30\x47\x63\xb4\x21\xa4\xa3\xec\xa3\x35\x98\x44\xe3\xa0\xbd\xa7\x9c\xa9\x0a\x4d\x6f\x08\xbb\x16\xd6\x8a\xa5\xc4\x04\x8c\x96\xd8\xc5\x9e\x61\xa1\x7a\x4b\xef\x9c\x56\xe0\x76\x1b\x8a\x0a\x46\x32\x34\x21\xb5\xa5\x0a\x9c\x39\xd6\xd7\x1c\x2a\xe5\x76\xc3\x54\xf1\x9f\x13\x6b\xb4\xf7\x79\xd6\x5a\x79\x16\x0a\x44\x96\xb1\xce\x68\x55\x15\x0f\x4d\x53\xd7\x14\x1e\x4c\x6a\x35\xf4\x39\x7f\x88\x77\x9b\x51\xbb\xe7\xf4\xa0\xe2\x67\xa1\x81\xb4\x2d\x33\x4a\xa8\xea\xcd\xb4\x75\xf1\xff\x3c\x6f\xdf\x03\xce\x53\xe6\x3a\xf0\x7f\xce\x5d\xcf\x9e\xd5\xde\x94\x68\x23\xfe\x7c\xa5\xcd\x1a\xb4\xb2\x7e\xb9\x16\x6e\x91\x6c\x85\xe2\x7a\x3b\x97\xba\x64\x4e\x10\x0d\x0b\xb8\x21\x40\x4b\x66\xf1\x71\xc3\x5c\x7d\x38\x64\x64\x5a\x94\x58\x3a\xe4\x8f\xa1\x2e\x3d\xbc\x81\xff\xc1\x22\x33\x65\x3d\x7f\x62\xd2\xe3\x3d\x18\x74\xde\x28\x58\x31\x69\xf1\x3e\xc2\xf3\x78\xf3\x84\xda\x78\x97\x56\x46\xfb\x0d\x1c\xfd\x4e\x65\x75\x21\x79\xa2\x40\xba\x74\x6a\x22\xab\xcd\x9c\xda\x66\x4a\x07\xfa\xa4\x34\x5e\x98\x97\x0e\xb8\xd1\x1b\x62\x45\xa5\x4e\x57\x95\xec\x05\x10\x8c\x45\xd2\x7b\x93\x22\x46\x4d\x2b\x8b\x41\x3f\x8c\x68\x21\xa9\xbc\xa6\x8e\x13\xa4\x5e\xf6\xe9\x03\x8e\x35\x2a\xff\x4a\x83\xcd\xb5\xdf\xcf\x86\x9d\x83\x0f\x0b\x38\xde\xc7\x37\x24\x9b\x66\xf4\xc0\x8c\x14\x81\xcf\xef\x60\x16\x3a\x6a\xeb\x5c\xd6\xd3\x19\x7a\x29\x8a\x9c\x41\x6d\x70\xb5\x48\x8e\xf1\xb4\x4a\x9a\xf5\x2c\x35\xdc\x0d\x46\x9e\x31\xe2\x86\x12\xdf\x80\x31\xae\xfa\x13\x04\x99\x97\x13\x0a\x8a\x1d\xa6\x23\x77\xab\x2a\x10\x7c\x91\x04\x79\x27\x9d\x62\x1c\x3e\xbb\x41\x2f\xcd\x19\x4a\x9b\xbf\x15\x1a\x2d\x34\x5f\xf0\xa7\x17\x06\xf9\xdf\x97\x6d\x38\xb4\x97\x64\x9b\x14\xdf\x5a\xc8\xaf\xcb\x6c\x82\x84\x0b\xae\x3c\x6b\x7a\x3e\x7f\x4e\x1b\x42\xe7\x3d\x36\x67\xa6\xe7\xf4\x68\x2a\x7f\xd1\xd0\x49\x0c\x56\xda\x2b\x7e\x0d\x9f\x04\x87\x9d\xf6\x64\x9a\x0a\x1d\x38\x0d\xcc\x39\x56\xd6\xe0\x6a\x5c\x7f\xbc\x80\x32\x26\x8f\x51\xe8\xc8\x0c\xe3\xb2\x1d\x61\x47\x79\x79\x6d\xb2\xf8\xbb\x02\x1d\x63\x94\xd0\xde\x87\x69\x11\x2c\xeb\x4b\x82\x6f\x27\xde\x24\x42\x5c\x8d\x8c\x37\x04\x44\x38\xae\xef\x4e\x43\x9d\x70\x34\x7c\x9a\x29\xf3\x03\x77\xcd\x31\xa9\xef\x8a\xa9\xde\xe2\x0b\x2e\x35\xdf\xc5\x56\xdb\x18\x6c\x4a\x77\xbd\xe7\x59\x63\xff\x16\x73\x2f\x5c\x77\xae\xee\xab\xf7\xfc\x0a\x00\x00\xff\xff\x05\xda\xbe\x9a\xaf\x09\x00\x00")

func webTemplatesQueryIndexGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesQueryIndexGoHtml,
		"web/templates/query/index.go.html",
	)
}

func webTemplatesQueryIndexGoHtml() (*asset, error) {
	bytes, err := webTemplatesQueryIndexGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/query/index.go.html", size: 2479, mode: os.FileMode(436), modTime: time.Unix(1512123257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webTemplatesIndexIndexGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\x4d\x6e\xc3\x20\x10\x85\xf7\x3e\xc5\x08\x65\x59\xc7\x6a\x95\x55\x65\x7b\xd1\x03\x54\xbd\xc2\xc8\x4c\x02\x12\x1e\x22\xa0\x3f\xd2\x88\xbb\x97\x46\x34\x2d\x6e\x15\x2f\x0c\x33\xef\x83\x37\x0f\x11\x4d\x47\xcb\x04\x6a\xf1\x9c\x88\x93\xca\xb9\x1b\xb5\x7d\x83\xc5\x61\x8c\x93\x0a\xfe\x5d\xcd\x1d\xc0\xef\xde\x17\x8a\xe5\x50\xb8\x28\x5b\xcd\xf5\xab\xee\xef\x1f\xaa\x06\x20\x12\x90\x4f\x04\x3b\xcb\x9a\x3e\xee\x60\xb7\xf8\xf5\xec\xb9\x78\xc1\xe3\x04\xfb\x6b\x15\x8b\x33\xd4\xef\xf6\x8d\x5b\xe2\x8c\x4c\x0e\x2e\xff\xbe\xc4\xc1\x57\x97\x1a\xf6\x1f\xba\x37\x84\xda\xf2\x69\xc3\x15\xd2\x1c\x5a\x30\xd9\xe4\x48\xcd\x23\x82\x09\x74\x9c\x94\xc8\x4f\x80\xfd\x13\x46\x7a\xc1\x64\x72\x1e\xd4\xdc\x28\xcf\xb8\x52\xce\xe3\x80\x7f\x0c\x06\x73\x98\xbb\x76\xba\xa1\x8c\xd7\x84\x6b\x1b\x4d\x29\x42\xac\xeb\x53\x5d\x85\xba\xa9\xcb\x37\xf2\x19\x00\x00\xff\xff\xcd\x13\xfe\x16\xdd\x01\x00\x00")

func webTemplatesIndexIndexGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesIndexIndexGoHtml,
		"web/templates/index/index.go.html",
	)
}

func webTemplatesIndexIndexGoHtml() (*asset, error) {
	bytes, err := webTemplatesIndexIndexGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/index/index.go.html", size: 477, mode: os.FileMode(436), modTime: time.Unix(1512123257, 0)}
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
	"web/templates/common/base.go.html": webTemplatesCommonBaseGoHtml,
	"web/templates/common/head.go.html": webTemplatesCommonHeadGoHtml,
	"web/templates/common/menu.go.html": webTemplatesCommonMenuGoHtml,
	"web/templates/monitor/details.go.html": webTemplatesMonitorDetailsGoHtml,
	"web/templates/monitor/index.go.html": webTemplatesMonitorIndexGoHtml,
	"web/templates/monitor/menu.go.html": webTemplatesMonitorMenuGoHtml,
	"web/templates/query/index.go.html": webTemplatesQueryIndexGoHtml,
	"web/templates/index/index.go.html": webTemplatesIndexIndexGoHtml,
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
	"web": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"common": &bintree{nil, map[string]*bintree{
				"base.go.html": &bintree{webTemplatesCommonBaseGoHtml, map[string]*bintree{}},
				"head.go.html": &bintree{webTemplatesCommonHeadGoHtml, map[string]*bintree{}},
				"menu.go.html": &bintree{webTemplatesCommonMenuGoHtml, map[string]*bintree{}},
			}},
			"index": &bintree{nil, map[string]*bintree{
				"index.go.html": &bintree{webTemplatesIndexIndexGoHtml, map[string]*bintree{}},
			}},
			"monitor": &bintree{nil, map[string]*bintree{
				"details.go.html": &bintree{webTemplatesMonitorDetailsGoHtml, map[string]*bintree{}},
				"index.go.html": &bintree{webTemplatesMonitorIndexGoHtml, map[string]*bintree{}},
				"menu.go.html": &bintree{webTemplatesMonitorMenuGoHtml, map[string]*bintree{}},
			}},
			"query": &bintree{nil, map[string]*bintree{
				"index.go.html": &bintree{webTemplatesQueryIndexGoHtml, map[string]*bintree{}},
			}},
		}},
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
