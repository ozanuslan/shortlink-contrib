// Code generated by go-bindata.
// sources:
// link.graphqls
// mutation.graphqls
// query.graphqls
// schema.graphqls
// DO NOT EDIT!

package schema

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

var _linkGraphqls = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xcf\x6a\xc6\x20\x10\xc4\xef\xfb\x14\x53\x3c\x27\xbd\xe7\xda\xaf\x85\x0f\x4a\xbf\x43\xf3\x02\xf9\xb3\x18\x89\x55\xd1\xcd\x21\x94\xbc\x7b\xa9\x29\x54\x92\x78\x10\xfc\xcd\x3a\x33\xab\xf0\xe2\x47\x86\x66\xc7\xb1\x13\x1e\xd1\xaf\x08\xd1\x8b\x1f\x2a\xcd\xae\xd2\x5e\xf8\x2b\xd8\x4e\xb8\xc6\xed\x81\x8f\x47\x8b\xd7\xdb\xbd\xad\x49\x21\xf9\x25\x0e\xdc\x20\xcc\xfa\xd9\x1a\x37\xe7\xab\xce\x7f\x89\x14\xda\x35\x70\x22\x85\x6a\x3f\x44\xb2\x06\xc6\xbb\x71\x33\xbe\x09\x00\x96\x68\x1b\xe0\x53\xa2\x71\xfa\x09\x19\x4d\x5d\x9a\x8e\x6c\xe4\x34\x44\xd3\x73\xc9\x37\x22\xe3\xc2\x22\xd9\xee\xcd\x58\xe1\x78\xcf\xef\xc2\x79\x1f\x2e\xc4\x22\xe1\x5a\xfb\x4f\x3a\xeb\x5b\xd1\x3f\xfd\xc5\xfc\xee\xdb\x20\xa3\x63\xa7\x74\x2e\xb5\x4f\x5f\x39\xff\x04\x00\x00\xff\xff\xcc\xb6\x2b\x95\x83\x01\x00\x00")

func linkGraphqlsBytes() ([]byte, error) {
	return bindataRead(
		_linkGraphqls,
		"link.graphqls",
	)
}

func linkGraphqls() (*asset, error) {
	bytes, err := linkGraphqlsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "link.graphqls", size: 387, mode: os.FileMode(436), modTime: time.Unix(1574556559, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mutationGraphqls = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\xf0\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xa8\xe6\x52\x50\x50\x50\x28\x2d\x48\x49\x2c\x49\xf5\xc9\xcc\xcb\xd6\x28\x2d\xca\xb1\x52\x08\x2e\x29\xca\xcc\x4b\xd7\x51\xc8\x48\x2c\xce\x40\xf0\x52\x52\x8b\x93\x8b\x32\x93\x52\x61\x22\x9a\x56\x0a\x4e\xf9\xf9\x39\xa9\x89\x79\x60\x43\x92\x8b\x52\xc9\x33\x04\xa4\x45\x11\x6c\x44\x4a\x6a\x4e\x2a\xd4\x08\x64\x5d\x08\x8b\x14\xb9\x6a\xb9\x00\x01\x00\x00\xff\xff\xc9\x68\x79\xa5\xc1\x00\x00\x00")

func mutationGraphqlsBytes() ([]byte, error) {
	return bindataRead(
		_mutationGraphqls,
		"mutation.graphqls",
	)
}

func mutationGraphqls() (*asset, error) {
	bytes, err := mutationGraphqlsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mutation.graphqls", size: 193, mode: os.FileMode(436), modTime: time.Unix(1573254298, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _queryGraphqls = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xc1\x4a\xc4\x30\x14\x45\xf7\xf9\x8a\x2b\x6e\x66\x36\x7e\x40\x76\x22\x08\x03\x2e\x14\x05\x17\xc3\x2c\xaa\xbc\x26\x8f\x86\x97\x98\xbc\x22\x45\xfa\xef\xd2\xb4\x16\x03\xb3\xcb\x3d\xe7\xe6\x06\x72\x8b\x37\x4f\x78\x19\x29\x4f\xd0\x29\x11\x32\xa5\x4c\x85\x44\x0b\xba\x10\x10\x7b\xa8\x27\x90\x68\x9e\x90\x22\x2f\x9c\x45\x63\xa5\xf7\xcf\xa7\x3b\x53\x6f\xad\x03\x3f\x06\x00\x02\xcb\x70\xf0\x5d\xf1\x16\xaf\x9a\x59\xdc\xd1\xe2\x89\x65\xb8\xd9\x6d\x39\xf4\x1c\x94\xf2\xca\x1f\xeb\xf9\x24\x69\xd4\xa3\xc5\xb9\x56\x2f\x66\x36\x86\x17\xb4\x6d\xfc\x2b\x6d\xcf\x08\xfd\xed\xd7\x48\x5f\x4d\x0c\xad\x0d\xda\x44\xd7\x5a\xd7\xda\xcf\x28\xda\xb1\x94\x06\x4a\xd4\x87\x6b\xfc\x83\xf4\x9b\x48\x2c\xce\x2b\xbc\x6c\xd4\xb1\x94\x77\xd6\xfd\x13\xcc\x6c\x7e\x03\x00\x00\xff\xff\xc9\x0b\x3c\x11\x6c\x01\x00\x00")

func queryGraphqlsBytes() ([]byte, error) {
	return bindataRead(
		_queryGraphqls,
		"query.graphqls",
	)
}

func queryGraphqls() (*asset, error) {
	bytes, err := queryGraphqlsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "query.graphqls", size: 364, mode: os.FileMode(436), modTime: time.Unix(1574466313, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGraphqls = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x50\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x3a\x60\x91\xdc\xd2\x92\xc4\x92\xcc\xfc\x3c\x2b\x05\x5f\x28\x4b\x87\xab\x96\x0b\x10\x00\x00\xff\xff\x96\xb8\xd3\x7f\x35\x00\x00\x00")

func schemaGraphqlsBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphqls,
		"schema.graphqls",
	)
}

func schemaGraphqls() (*asset, error) {
	bytes, err := schemaGraphqlsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphqls", size: 53, mode: os.FileMode(436), modTime: time.Unix(1571565241, 0)}
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
	"link.graphqls":     linkGraphqls,
	"mutation.graphqls": mutationGraphqls,
	"query.graphqls":    queryGraphqls,
	"schema.graphqls":   schemaGraphqls,
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
	"link.graphqls":     &bintree{linkGraphqls, map[string]*bintree{}},
	"mutation.graphqls": &bintree{mutationGraphqls, map[string]*bintree{}},
	"query.graphqls":    &bintree{queryGraphqls, map[string]*bintree{}},
	"schema.graphqls":   &bintree{schemaGraphqls, map[string]*bintree{}},
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
