package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __000001_create_links_table_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\xc8\xc9\xcc\xcb\x2e\xb6\xe6\x02\x04\x00\x00\xff\xff\xc8\xd3\xf7\xad\x1c\x00\x00\x00")

func _000001_create_links_table_down_sql() ([]byte, error) {
	return bindata_read(
		__000001_create_links_table_down_sql,
		"000001_create_links_table.down.sql",
	)
}

var __000001_create_links_table_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x51\x41\xaf\x9b\x30\x18\xbb\xe7\x57\xf8\x56\x90\x8a\xb4\x4d\xea\x65\x68\x87\x8e\x66\x1b\x12\x0d\x1a\x05\x75\xb7\x2a\x94\x4c\x64\xa4\xa1\x4a\xc2\xd4\xfe\xfb\x29\x0b\xdd\xda\xa7\xa7\xbe\x9c\xc0\xb1\xfd\xf9\x73\x92\x04\x99\x11\xdc\x09\x70\x38\xde\x2a\x81\x9f\xa3\x81\x92\x7a\xb0\x24\xab\xe8\xba\xa6\xa8\xd7\x9f\x0b\x3a\x43\x11\x01\x00\xd9\x21\x1c\x2b\x8c\xe4\x6a\xfe\xd1\xa3\x83\x9e\x94\x22\xb8\x3f\xc7\x51\x5b\x67\xb8\xd4\x2e\x58\x1c\xce\xc3\x23\xe1\x6c\xe4\x89\x9b\x2b\x06\x71\x5d\xfe\xbd\x99\xcc\xec\xf8\x9b\x9b\x63\xcf\x4d\xf4\x61\xb5\x8a\xff\xb9\x07\x4e\xcf\x6d\xff\x16\xa7\x13\xf6\x68\x64\x2b\xe0\xc4\xc5\x05\xe8\x97\x1d\x35\x6e\x1f\x2d\x5e\x06\x8f\x53\x42\xb2\x72\xbb\xa5\xac\x46\xc9\xee\x17\x47\xbe\xc3\xa2\x90\x7a\x80\x92\xd6\x2d\x52\x42\xd6\x45\x4d\xab\x87\x6e\xbc\x55\xb9\x67\x1e\x2d\x61\xfb\xd1\x38\x8f\x7b\xcb\xd0\x63\xc3\xf2\xef\x0d\x45\xce\x36\xf4\xc7\xdc\x85\xec\x0e\x93\xd4\x9d\xb8\x04\x31\x9b\x87\x45\xb2\x8b\x9f\xea\xfc\xfa\xaf\x2b\xfd\x8d\xd7\x26\x09\x72\x96\x15\xcd\x86\x26\x81\x96\x24\xe0\x16\xe2\xc2\x4f\x67\x25\x3e\x62\x47\x0b\x9a\xd5\x90\xdd\xd2\xf7\xbd\x0c\x85\x7e\xa9\xca\xed\x6c\xb4\xff\x46\x2b\xea\x1f\xfa\x13\xde\xbf\x4b\x9f\x44\xf1\x7d\xfc\x9f\xdf\xec\x72\xf6\x15\xad\x33\x42\xcc\x59\x6e\x31\x10\x4d\x46\xc5\x29\xf9\x13\x00\x00\xff\xff\x65\x2d\x2b\x29\x71\x02\x00\x00")

func _000001_create_links_table_up_sql() ([]byte, error) {
	return bindata_read(
		__000001_create_links_table_up_sql,
		"000001_create_links_table.up.sql",
	)
}

var __000002_add_test_link_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _000002_add_test_link_down_sql() ([]byte, error) {
	return bindata_read(
		__000002_add_test_link_down_sql,
		"000002_add_test_link.down.sql",
	)
}

var __000002_add_test_link_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\xcf\x6f\x82\x30\x14\xc7\xef\xfc\x15\x2f\xbd\xa0\x09\x68\xa6\x37\x3c\x55\x6d\xb6\x66\x48\x17\xe8\xbc\x9a\x82\x55\x98\x48\x4d\x5b\xb6\xb9\x65\xff\xfb\xe2\xaf\x2c\x0e\xb2\x93\x19\x27\xf2\xf2\xfa\xf9\xbe\x4f\xde\xf3\x7d\xe0\x31\x8e\x12\x3c\xe1\x94\x45\xe0\xff\xd7\xe7\x8c\xc9\x3d\x8d\x1c\x00\x00\x9a\xb0\x10\x1f\xd3\x43\x32\x27\x21\xc4\x04\x4f\x61\xc2\x66\x33\xca\x39\x99\x8e\x1c\x27\xc1\x73\xf2\xc4\x68\xc4\xc1\xbe\x2f\x32\x2d\x85\x95\x8b\xa5\x5c\x89\xba\xb4\x8b\xb2\xa8\x36\x66\xe4\x38\x34\x4a\x48\xcc\x81\x46\x9c\xc1\xb1\xd6\xa9\x75\xe9\x41\x2e\x4c\xee\xc1\x52\x9a\x4c\x17\xa9\xf4\xe0\xc5\xa8\xaa\x7b\x4c\x9d\xe3\xf0\x99\x24\xd0\x71\x73\x6b\x77\x26\xe8\xf7\x53\x61\xc5\x87\xd2\x3d\x5d\xbb\x1e\xb8\xdb\xfd\x83\x30\xf9\xdd\xe1\x77\xb6\x87\x9d\xd4\x46\x55\xa2\x84\x37\x99\x9a\xc2\xca\x43\xf9\x13\xd5\xba\x44\x01\x6a\xbe\x47\x1e\xa0\x43\x30\x0a\xd0\x19\x83\x3c\x74\x99\x01\x05\xa8\x05\x88\xbe\xdc\xee\x4d\x2c\xd6\x85\xcd\xeb\xb4\x97\xa9\xed\x65\xa0\x1f\x9b\xc1\xd9\x46\x64\x99\xaa\x2d\xa8\x15\x9c\xba\xdb\x6c\x9a\x9c\x86\xd5\xa0\x61\xf5\x1b\x7c\x33\xab\xd7\x4d\xbb\xd1\xf0\xb2\x1f\xb1\x96\xa0\x2a\x38\xf5\xb5\xf9\x5c\x13\x1a\x2e\xc3\xe6\x86\xae\x90\x67\x13\xdf\x87\x98\x85\xe1\x18\x4f\x1e\x81\xb3\x3f\xce\xf1\x74\xbe\x23\xe7\x3b\x00\x00\xff\xff\x41\x53\xa8\x1d\x5f\x03\x00\x00")

func _000002_add_test_link_up_sql() ([]byte, error) {
	return bindata_read(
		__000002_add_test_link_up_sql,
		"000002_add_test_link.up.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"000001_create_links_table.down.sql": _000001_create_links_table_down_sql,
	"000001_create_links_table.up.sql":   _000001_create_links_table_up_sql,
	"000002_add_test_link.down.sql":      _000002_add_test_link_down_sql,
	"000002_add_test_link.up.sql":        _000002_add_test_link_up_sql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"000001_create_links_table.down.sql": &_bintree_t{_000001_create_links_table_down_sql, map[string]*_bintree_t{}},
	"000001_create_links_table.up.sql":   &_bintree_t{_000001_create_links_table_up_sql, map[string]*_bintree_t{}},
	"000002_add_test_link.down.sql":      &_bintree_t{_000002_add_test_link_down_sql, map[string]*_bintree_t{}},
	"000002_add_test_link.up.sql":        &_bintree_t{_000002_add_test_link_up_sql, map[string]*_bintree_t{}},
}}
