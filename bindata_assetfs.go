// Code generated by go-bindata.
// sources:
// www/index.html
// DO NOT EDIT!

package main

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _wwwIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\x4d\x6f\xd4\x30\x10\x3d\x3b\xbf\x62\xc8\x25\x89\x54\x62\xf5\xba\x72\x7c\x61\x91\x8a\x54\xb4\x48\xf4\x82\x10\x07\x6f\x32\x51\x02\x8e\x1d\x9c\xf1\x42\x84\xfa\xdf\xd1\xd8\x94\x42\xf9\x38\x8d\x3c\x6f\xde\xbc\xe7\xa7\x51\xcf\x8e\xa7\x17\x77\xef\xde\xbc\x84\x89\x16\xab\x0b\xf5\x50\xd0\x0c\xba\x10\x8a\x66\xb2\xa8\x6f\xfd\x6e\x2c\xed\x27\x87\x70\xe3\x17\xfc\xe2\xc3\x27\x25\x33\x54\x28\x99\x67\xd5\xd9\x0f\x3b\x53\xa6\xeb\xbf\xcf\x4f\xd7\x8c\x8e\x3e\x2c\xba\x10\x42\x59\x73\x46\x0b\xa3\x0f\x5d\x49\xf8\x95\x9e\xcf\x6e\x8d\x54\xea\x57\x5c\x0e\x4a\x26\x38\x0d\x26\x00\x68\x5f\x31\x4f\x96\xe0\xcc\x82\xbf\xb1\xc0\x44\xf2\xa3\xef\xe3\xf6\x07\x63\x8b\xe7\x65\xfe\xc9\x79\x78\x5d\x8c\x8d\xd8\x95\x47\xef\xb0\x64\x57\xf2\x87\x2d\x35\xcc\x97\x47\x77\xfa\x14\xe9\xa9\x9b\x6d\x35\x4e\x2b\x99\x0a\x13\x33\x41\x6d\x7d\x98\xd7\x5f\x5d\xca\x8f\xe6\x62\x72\x97\x15\xc4\xe0\xfb\xb8\xa0\xa3\xf6\x73\xc4\xb0\xbf\x45\x8b\x3d\xf9\x50\x57\x2c\x5c\x35\xad\x77\xd9\x1a\x74\x30\x46\xd7\xd3\xec\x5d\x8d\x0d\x7c\x2b\x84\x10\xd8\xae\x01\x2f\xe8\xe8\x88\xa3\x89\x96\xea\xa6\xe0\xb6\x45\x02\xce\x1c\x3a\xf8\xdf\x72\xd0\xf0\x3e\x7d\xfe\x31\xaf\x0f\x55\xd3\xa6\x04\x78\xcd\x88\xd4\x4f\x75\x25\xb1\x9f\x7c\x75\x95\x04\xc5\x82\x34\xf9\xe1\x00\xd5\xea\x37\xaa\xae\x52\x8f\xa5\x0e\x49\x90\x9f\xf7\x4d\x4b\x13\xba\x3a\xe0\xb6\x42\xa7\x81\x6b\xcb\x0a\x75\xf3\x04\xc9\x1b\xff\xe5\x90\x73\xac\x9a\x76\x76\x0e\xc3\xcd\xdd\xeb\x5b\xe8\xd2\xaa\x2c\x51\x08\x71\xcf\x19\xe7\x18\xf9\xd4\xf2\x8d\x29\x99\xae\xf4\x7b\x00\x00\x00\xff\xff\xda\x8e\xbf\x44\xbc\x02\x00\x00")

func wwwIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwIndexHtml,
		"www/index.html",
	)
}

func wwwIndexHtml() (*asset, error) {
	bytes, err := wwwIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/index.html", size: 700, mode: os.FileMode(420), modTime: time.Unix(1501613665, 0)}
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
	"www/index.html": wwwIndexHtml,
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
	"www": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{wwwIndexHtml, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
