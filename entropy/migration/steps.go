package migration

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

func entropy_db_001_initial_schema_sql() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xc4, 0x90,
		0xc1, 0x4b, 0xc3, 0x30, 0x14, 0xc6, 0xef, 0xfd, 0x2b, 0xde, 0xb1, 0x81,
		0x1d, 0xa6, 0xb0, 0x93, 0xa7, 0x18, 0x5e, 0x67, 0x30, 0x7b, 0x99, 0xcf,
		0x54, 0xd8, 0x69, 0x14, 0x17, 0x24, 0xa0, 0xed, 0x4c, 0x5b, 0xc1, 0xff,
		0xde, 0x56, 0x37, 0xd7, 0xd9, 0xe1, 0x45, 0x70, 0xd7, 0xbc, 0xef, 0x0b,
		0xbf, 0xdf, 0xa7, 0x18, 0xa5, 0x43, 0x70, 0xf2, 0xda, 0x20, 0x44, 0xbf,
		0xad, 0xea, 0xd0, 0x54, 0x31, 0xf8, 0x1a, 0x52, 0x48, 0x00, 0xc2, 0x06,
		0x34, 0x39, 0x9c, 0x23, 0xc3, 0x92, 0xf5, 0x42, 0xf2, 0x0a, 0x6e, 0x71,
		0x05, 0x32, 0x77, 0x56, 0x53, 0xd7, 0x5d, 0x20, 0xb9, 0x49, 0x97, 0xab,
		0xab, 0x36, 0x3e, 0x7a, 0x78, 0x90, 0xac, 0x6e, 0x24, 0xa7, 0x17, 0xd3,
		0xa9, 0x00, 0xb2, 0x0e, 0x28, 0x37, 0xa6, 0xbf, 0x97, 0xc5, 0xcb, 0x2f,
		0xd7, 0x9c, 0xf4, 0x5d, 0x8e, 0x90, 0x7e, 0xfd, 0x32, 0xf9, 0x4c, 0x0b,
		0xb0, 0x04, 0xca, 0x52, 0x66, 0xb4, 0x72, 0xc0, 0xb8, 0x34, 0x52, 0x21,
		0x24, 0xe2, 0x2a, 0x49, 0x8e, 0x98, 0xdb, 0x32, 0xbc, 0xb6, 0x7e, 0xbd,
		0x2d, 0x62, 0x13, 0x9a, 0x50, 0x95, 0xeb, 0xbe, 0xbc, 0xa3, 0xff, 0xf6,
		0x79, 0xdf, 0x5b, 0x9c, 0x64, 0xe9, 0x1f, 0x07, 0x76, 0xe9, 0xa1, 0xb6,
		0x43, 0xe9, 0x03, 0x99, 0x65, 0xd4, 0x73, 0xfa, 0x11, 0x10, 0x1d, 0x59,
		0x86, 0x8c, 0xa4, 0xf0, 0xfe, 0x68, 0xbe, 0x34, 0x6c, 0xc4, 0x18, 0x36,
		0x16, 0xe5, 0xd3, 0x80, 0xf5, 0xdf, 0x30, 0xf7, 0x93, 0x0e, 0x60, 0x4f,
		0xef, 0x36, 0xee, 0x8c, 0x25, 0x6a, 0xdf, 0xfc, 0x45, 0xe1, 0xad, 0x78,
		0x6e, 0x0f, 0xaf, 0x97, 0xb3, 0xd9, 0x19, 0xc5, 0x3e, 0x02, 0x00, 0x00,
		0xff, 0xff, 0x79, 0x0f, 0xa9, 0xb4, 0xfe, 0x02, 0x00, 0x00,
	},
		"entropy/db/001_initial_schema.sql",
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
	"entropy/db/001_initial_schema.sql": entropy_db_001_initial_schema_sql,
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
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"entropy": &_bintree_t{nil, map[string]*_bintree_t{
		"db": &_bintree_t{nil, map[string]*_bintree_t{
			"001_initial_schema.sql": &_bintree_t{entropy_db_001_initial_schema_sql, map[string]*_bintree_t{
			}},
		}},
	}},
}}
