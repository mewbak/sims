// Code generated by go-bindata.
// sources:
// lines2out1.dat
// easy.dat
// hard.dat
// impossible.dat
// DO NOT EDIT!

package main

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

var _lines2out1Dat = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x96\x3f\x8f\xdb\x30\x0c\xc5\x67\xdf\xe7\x68\xb7\x0c\x22\xa9\x5b\x8c\xe2\xa6\x0e\xee\xd2\x6e\xb7\x14\x87\xa0\x43\xc7\x16\x1d\x92\xef\x5f\x38\xb2\x19\x3d\x8a\xb6\xfe\x40\xc3\x89\xa1\xf5\x3b\x3e\xf2\x45\xf1\x75\x99\xa7\x4f\xdf\x7f\xfd\xf9\x3d\x7d\xfe\xf6\xf7\xdf\xfd\xf6\x93\xe7\x70\x09\x1f\x5f\x78\x7e\xbd\xbc\xbe\xe5\x1f\xd2\x47\x1e\x31\x44\x02\x51\xcc\x22\xba\x04\x88\x08\x22\x86\x48\x20\xca\x29\x0c\x14\x06\x0a\x03\x85\x81\xc2\x40\x11\xa0\x08\x50\x04\x28\x02\x14\x01\x4a\x04\x4a\x04\x4a\x04\x4a\x04\x4a\x7c\x50\x7e\xdc\x6f\xa6\xcb\xfc\x06\x9f\x52\xfe\x50\xea\x5e\x16\x42\x96\x31\xcb\x98\x15\xcc\x0a\x66\x23\x66\x57\x19\x2f\xd7\xaf\xf3\xf4\x1e\xae\xef\x34\xad\x2b\x3c\x56\xfb\x2e\xe8\x7e\x5b\xca\xe3\x2c\xdb\xb3\x0b\xc8\x54\x9e\x0c\xb0\xf2\xda\xc8\xf0\xe2\x80\x5a\xa3\x75\x8d\x77\xde\x92\xf2\xfb\x0a\x85\x12\x7f\x67\x89\x4f\x1e\x99\xe7\xda\xc9\xe4\xf2\xf8\xb0\x82\xae\x9a\x95\x27\x0d\xda\x8e\xc9\x1a\x29\x2f\x76\x75\xcd\x92\x35\x97\x78\xb4\xfa\x6f\x6c\xb2\xe4\xf9\x8f\x56\xff\x8d\x79\xcf\xfc\x55\x5e\x1c\xfc\x76\x90\xe7\x3f\x42\xff\xd5\x3b\x78\xe8\x40\xe5\xa1\x86\x56\x72\x41\x54\x1e\xbb\x27\x6a\xe4\x42\xb5\xf2\xa4\x52\x8b\x4f\x2e\x9e\x54\x5e\x6c\xee\x1a\xba\x8f\x3c\xff\x71\xf2\xcb\xe8\xdd\x42\xd6\x2f\x9c\xfc\x32\x7a\xfb\x91\xf5\x0b\x5b\xbf\xb4\x6b\x0f\xde\xf7\x83\x93\x5f\x9e\xf9\x76\xb2\xd1\xae\x3c\x76\x4f\xd4\xc9\x86\xa8\x3c\xa9\xd4\x72\x3a\x55\xeb\x17\x4e\x7e\x69\xeb\xda\xe9\x0d\x9b\x78\xb2\xcf\x77\xd4\x31\x3a\xf1\x9d\x57\xcc\xb7\xb5\x56\x93\x51\x1e\xc1\xff\x19\x26\x2b\x8f\xdd\xe7\x5a\xc8\xa0\x5e\x79\x52\xa9\xe0\x98\x0c\x44\xe5\xc5\x66\x6d\xfe\xaf\x11\xe5\xf3\x8d\xde\x3c\xda\xd8\x86\xa9\x3c\xbc\xc3\x7b\xc8\xe4\xf2\xd8\x3d\xd1\x46\xce\x66\xac\x3c\xa9\xd4\x52\x7b\x3b\xd8\x26\xa2\xbc\xd8\xdc\xb5\x93\xb7\x83\x6d\x1e\xcb\xf6\x7e\x55\x2e\x7b\xa2\x24\x98\x3a\x95\xc7\x2e\xa5\x46\x2e\x66\xa2\x3c\xa9\x9c\x3d\x73\x5d\xb6\x57\x5e\x6c\x56\x79\xf4\xcb\x19\xd4\xcf\x0b\xe5\x7e\x09\x9d\x9d\xa4\x52\x2f\xe5\x7e\xf1\xb5\x9d\x4d\xb5\xd0\x4b\xb9\x5f\x6a\x53\x38\xbb\x19\x36\xbd\x8c\xf5\xf5\x2b\x27\xac\x8f\xb1\xbe\x7e\xe5\x84\xf5\xc9\x11\xaf\xa7\xd2\x47\x8d\x2f\xff\x03\x00\x00\xff\xff\xe2\x8f\x28\x15\x1c\x10\x00\x00")

func lines2out1DatBytes() ([]byte, error) {
	return bindataRead(
		_lines2out1Dat,
		"lines2out1.dat",
	)
}

func lines2out1Dat() (*asset, error) {
	bytes, err := lines2out1DatBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lines2out1.dat", size: 4124, mode: os.FileMode(420), modTime: time.Unix(1568628202, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _easyDat = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xf7\xb0\xe2\x54\xf1\x4b\xcc\x4d\xe5\x54\xf5\xcc\x2b\x28\x2d\x89\x36\xb2\x32\xd0\x31\x88\xb5\x31\xb2\x32\xd4\x31\xb1\x43\x16\x34\x8c\x45\xe6\x19\xa1\xf0\x8c\x63\x39\x55\xfd\x4b\x4b\xd0\xf4\x1b\xd9\xa1\x88\x1a\xc6\x72\xc5\xbb\x58\x71\xba\x96\xa5\xe6\x95\xc4\x1b\x70\x1a\x72\x1a\x80\xa1\x21\xa7\x01\x92\xb8\x21\x27\x4c\x06\x55\xdc\x08\x26\x06\x22\x91\xc4\x8d\x11\xa6\x70\x1a\x72\x01\x02\x00\x00\xff\xff\x7d\xc7\x16\x70\xce\x00\x00\x00")

func easyDatBytes() ([]byte, error) {
	return bindataRead(
		_easyDat,
		"easy.dat",
	)
}

func easyDat() (*asset, error) {
	bytes, err := easyDatBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "easy.dat", size: 206, mode: os.FileMode(420), modTime: time.Unix(1568628239, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _hardDat = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xf7\xb0\xe2\x54\xf1\x4b\xcc\x4d\xe5\x54\xf5\xcc\x2b\x28\x2d\x89\x36\xb2\x32\xd0\x31\x88\xb5\x31\xb2\x32\xd4\x31\xb1\x43\x16\x34\x8c\x45\xe6\x19\xa1\xf0\x8c\x63\x39\x55\xfd\x4b\x4b\xd0\xf4\x1b\xd9\xa1\x88\x1a\xc6\x72\xc5\xbb\x58\x71\xba\x96\xa5\xe6\x95\xc4\x1b\x70\x1a\x82\x21\x88\x36\x40\x12\x87\x88\x40\xe5\x90\xc4\x8d\x20\x2a\x21\x18\x49\xdc\x18\xaa\xde\x00\x2c\x0e\x08\x00\x00\xff\xff\xdd\x0b\x6c\xb8\xce\x00\x00\x00")

func hardDatBytes() ([]byte, error) {
	return bindataRead(
		_hardDat,
		"hard.dat",
	)
}

func hardDat() (*asset, error) {
	bytes, err := hardDatBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "hard.dat", size: 206, mode: os.FileMode(420), modTime: time.Unix(1568626501, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _impossibleDat = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xf7\xb0\xe2\x54\xf1\x4b\xcc\x4d\xe5\x54\xf5\xcc\x2b\x28\x2d\x89\x36\xb2\x32\xd0\x31\x88\xb5\x31\xb2\x32\xd4\x31\xb1\x43\x16\x34\x8c\x45\xe6\x19\xa1\xf0\x8c\x63\x39\x55\xfd\x4b\x4b\xd0\xf4\x1b\xd9\xa1\x88\x1a\xc6\x72\xc5\xbb\x58\x71\xba\x96\xa5\xe6\x95\xc4\x1b\x70\x1a\x72\xc2\x31\x92\x38\x42\x14\x55\xdc\x08\x22\x02\x86\x86\x48\xe2\xc6\x10\x11\x88\x1e\x2e\x40\x00\x00\x00\xff\xff\x49\x76\x16\xe1\xce\x00\x00\x00")

func impossibleDatBytes() ([]byte, error) {
	return bindataRead(
		_impossibleDat,
		"impossible.dat",
	)
}

func impossibleDat() (*asset, error) {
	bytes, err := impossibleDatBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "impossible.dat", size: 206, mode: os.FileMode(420), modTime: time.Unix(1568626501, 0)}
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
	"lines2out1.dat": lines2out1Dat,
	"easy.dat": easyDat,
	"hard.dat": hardDat,
	"impossible.dat": impossibleDat,
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
	"easy.dat": &bintree{easyDat, map[string]*bintree{}},
	"hard.dat": &bintree{hardDat, map[string]*bintree{}},
	"impossible.dat": &bintree{impossibleDat, map[string]*bintree{}},
	"lines2out1.dat": &bintree{lines2out1Dat, map[string]*bintree{}},
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

