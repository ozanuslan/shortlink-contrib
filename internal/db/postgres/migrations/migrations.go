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

var __000001_create_links_table_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x41\x6f\x9c\x30\x10\x85\xef\xfe\x15\x4f\xb9\x04\xa4\x45\x4a\xab\xe6\x52\xd4\x03\x65\x67\x5b\xab\xac\x69\xc1\x28\x9b\x13\x62\x83\xdb\x75\x43\x21\xb2\x4d\xb4\x3f\xbf\xf2\x9a\x55\xd3\x68\x0f\xad\x4f\x68\x78\xf3\xcd\xd3\x9b\xc9\x2b\xca\x24\x81\x76\x92\x44\xcd\x4b\x01\xbe\x81\x28\x25\x68\xc7\x6b\x59\xe3\x6a\x9e\x75\x9f\x4c\xd6\x3e\x5d\xa5\x8c\x25\x09\x72\xa3\x3a\xa7\xd0\xc1\x75\xfb\x41\xe1\xfb\x64\x30\xe8\xf1\xd1\xb2\x05\x24\xb3\x8f\x05\xbd\x82\x04\x41\xc4\x00\x40\xf7\x08\xaf\x69\xf8\xfa\x24\x12\x4d\x51\x60\x4d\x9b\xac\x29\x24\xfc\xb8\xf6\x87\x1a\x95\xe9\x9c\x6a\x9f\xdf\x45\xf1\x8a\xe1\xe5\xcb\x4b\x51\xcb\x2a\xe3\x42\x42\xf7\xed\x89\x8c\xaf\x15\xdf\x66\xd5\x3d\xbe\xd0\x7d\xa4\xfb\xa5\x63\x36\x43\xe8\x78\xee\xcc\xc3\xa1\x33\xd1\xdb\xdb\xdb\x18\xe3\xe4\x30\xce\xc3\x10\x34\x87\xce\x1e\xfe\xd6\xdc\xc4\x78\xa5\xe9\x95\x7d\x30\x7a\xaf\xe0\xd4\xd1\x85\xd2\x4f\x3b\x8d\x38\x7f\xec\xcf\xce\xce\x6d\x2c\x4e\x19\xcb\xcb\xed\x96\x84\x44\x29\x96\x44\x82\x53\x5e\xe3\xba\xd0\xe3\x23\x06\x6d\xdd\x75\xca\x58\x56\x48\xaa\x5e\x4a\x4e\x03\xca\x3b\xe1\xab\x25\xec\x61\x32\xce\xd7\x3d\x32\x04\xdc\x08\xfe\xad\x21\x70\xb1\xa6\xdd\xa5\x9c\x5b\xdd\xb7\xb3\x1e\x7b\x75\x0c\x28\xb1\x8c\xf6\xc9\xfc\x07\xc5\x47\x73\x99\xe3\xff\xc4\xe1\x18\xb8\xc8\x8b\x66\x4d\x49\x90\x25\x09\x3a\x0b\x75\xec\x7e\x3d\x0d\xea\x3d\x6a\x2a\x28\xf7\x5b\x5a\xf9\x5d\xac\x42\xd8\x9b\xaa\xdc\x2e\xa0\xbb\xcf\x54\x91\xbf\x87\x0f\x78\x73\x93\xfe\xb3\x31\x9f\xdc\x1f\x37\x4d\xcd\xc5\x27\xec\x9d\x51\x6a\x71\x76\x36\x85\x68\x36\x43\x9c\xb2\xdf\x01\x00\x00\xff\xff\x5a\x64\xea\x0f\xe1\x02\x00\x00")

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
