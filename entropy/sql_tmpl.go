package entropy

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

func digest_tmpl() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x0a, 0x76,
		0xf5, 0x71, 0x75, 0x0e, 0x51, 0xe0, 0x52, 0x50, 0xc8, 0x4c, 0xd1, 0x01,
		0x51, 0x3e, 0xfe, 0xe1, 0xae, 0x41, 0x1a, 0x1e, 0xae, 0x11, 0x1a, 0xbe,
		0x2e, 0xa6, 0x1a, 0xee, 0x41, 0xfe, 0xa1, 0x01, 0xf1, 0xce, 0xfe, 0x7e,
		0xce, 0x8e, 0x21, 0x1a, 0x65, 0xa9, 0x45, 0xc5, 0x99, 0xf9, 0x79, 0x3a,
		0xea, 0xea, 0x9a, 0x40, 0xa0, 0xe0, 0x18, 0xac, 0x90, 0x92, 0x99, 0x9e,
		0x5a, 0x5c, 0xc2, 0xe5, 0x16, 0xe4, 0xef, 0xab, 0x50, 0x16, 0x5f, 0x5d,
		0xad, 0xa0, 0x17, 0x90, 0x5f, 0x5c, 0x92, 0x96, 0x59, 0xa1, 0x50, 0x5b,
		0xcb, 0x05, 0xd6, 0xab, 0xe0, 0x14, 0x09, 0x34, 0x59, 0x81, 0xcb, 0x3f,
		0xc8, 0xc5, 0x35, 0x08, 0xca, 0x71, 0x71, 0x0d, 0x76, 0xb6, 0xe6, 0x02,
		0x04, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xd8, 0xbf, 0xab, 0x7a, 0x00, 0x00,
		0x00,
	},
		"digest.tmpl",
	)
}

func repo_tmpl() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x7c, 0x90,
		0xcd, 0x4e, 0x83, 0x40, 0x14, 0x85, 0xf7, 0x3c, 0xc5, 0x59, 0x74, 0x02,
		0x24, 0xad, 0x0f, 0x50, 0xe3, 0x02, 0xed, 0x10, 0x89, 0xfc, 0x05, 0xa6,
		0x11, 0x56, 0x84, 0x74, 0xa6, 0x66, 0x12, 0xf9, 0x09, 0x60, 0x53, 0xd3,
		0xf0, 0xee, 0xce, 0x80, 0x56, 0x8c, 0xc6, 0xdd, 0xe5, 0x72, 0xce, 0x9c,
		0xef, 0x9e, 0xcb, 0x05, 0xab, 0x16, 0xdb, 0x3b, 0xdc, 0xc4, 0x4d, 0x3f,
		0x1c, 0xe5, 0x19, 0xe3, 0x68, 0x18, 0x0f, 0x09, 0x75, 0x18, 0x05, 0x73,
		0xee, 0x7d, 0x0a, 0xcf, 0x45, 0x18, 0x31, 0xd0, 0xcc, 0x4b, 0x59, 0x8a,
		0x53, 0x31, 0x5b, 0xc6, 0x11, 0x16, 0x0c, 0x40, 0x72, 0x30, 0x9a, 0xb1,
		0xb5, 0x1a, 0x4f, 0xa2, 0xeb, 0x65, 0x53, 0x4f, 0xdf, 0x93, 0x25, 0xdc,
		0xfb, 0xbe, 0xfe, 0x51, 0x0c, 0xfd, 0x50, 0x56, 0x2d, 0x98, 0x17, 0xd0,
		0x94, 0x39, 0x41, 0x8c, 0x1d, 0x75, 0x9d, 0xbd, 0xcf, 0xac, 0x94, 0x25,
		0xae, 0xde, 0x5a, 0x26, 0xc9, 0x37, 0xa4, 0xda, 0x10, 0x0e, 0xf2, 0xb8,
		0x25, 0xc1, 0x96, 0x1c, 0xcd, 0x35, 0xcc, 0x30, 0x7a, 0x36, 0x6d, 0x5b,
		0x3f, 0xa1, 0x42, 0xbb, 0xb2, 0x7e, 0x11, 0x58, 0xd5, 0x65, 0x25, 0xd6,
		0x58, 0x71, 0xd1, 0x1f, 0x66, 0xee, 0xb2, 0x1b, 0xe4, 0xa0, 0x72, 0x7b,
		0x8d, 0x0e, 0x4c, 0xda, 0x49, 0xa5, 0x19, 0xd5, 0x7c, 0x68, 0x5e, 0xdf,
		0xaa, 0x9a, 0xbd, 0xb7, 0x62, 0x76, 0xa9, 0xed, 0x12, 0x4e, 0x29, 0x44,
		0xcd, 0x67, 0x6f, 0x9c, 0x78, 0x81, 0x93, 0xe4, 0x78, 0xa2, 0xb9, 0x25,
		0xb9, 0x6d, 0xd8, 0xb7, 0xd7, 0x2e, 0xbc, 0x70, 0x47, 0xb3, 0xaf, 0x4b,
		0xae, 0x1d, 0x14, 0x92, 0x9f, 0x11, 0x85, 0x8b, 0x56, 0xac, 0x4f, 0x89,
		0xb6, 0xfe, 0x82, 0x2e, 0xfe, 0x24, 0xfe, 0x91, 0xb0, 0x80, 0xff, 0x2f,
		0x65, 0x21, 0x53, 0x49, 0xdf, 0x37, 0x7c, 0x04, 0x00, 0x00, 0xff, 0xff,
		0xb4, 0x40, 0x87, 0x64, 0xd0, 0x01, 0x00, 0x00,
	},
		"repo.tmpl",
	)
}

func upsert_tmpl() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x94, 0x53,
		0x51, 0x6b, 0xf2, 0x40, 0x10, 0x7c, 0x4e, 0x7e, 0xc5, 0x3c, 0x18, 0x12,
		0x21, 0xfa, 0x03, 0xfc, 0x3e, 0x5b, 0x42, 0xba, 0x62, 0x4a, 0xbc, 0x93,
		0xe4, 0x40, 0xfa, 0x24, 0xa1, 0x89, 0x25, 0x50, 0x63, 0x31, 0xc1, 0x16,
		0xc4, 0xff, 0xde, 0xbb, 0x4b, 0x62, 0xaf, 0xea, 0x43, 0x3d, 0x08, 0xdc,
		0x2e, 0xb3, 0xbb, 0x73, 0xb3, 0x93, 0x88, 0xa5, 0x94, 0x08, 0xf0, 0x04,
		0x09, 0x2d, 0xe3, 0x20, 0x24, 0x44, 0x4c, 0x70, 0x1c, 0xd6, 0xc7, 0x23,
		0xc6, 0xcb, 0x5d, 0xdd, 0x6c, 0xca, 0x2f, 0x9c, 0x4e, 0xb0, 0x01, 0x4f,
		0x7e, 0x80, 0xcc, 0xef, 0xb3, 0xea, 0xad, 0xc0, 0xa0, 0xca, 0xb6, 0x85,
		0x8f, 0x41, 0x5e, 0xd4, 0xaf, 0x98, 0x4c, 0x25, 0x3a, 0xdb, 0x37, 0x65,
		0x53, 0xee, 0xaa, 0x5a, 0x16, 0x68, 0xac, 0x46, 0x6b, 0x9c, 0xcc, 0xf8,
		0x7d, 0x79, 0x51, 0xe5, 0x3d, 0xa0, 0xcc, 0xdb, 0xec, 0xa1, 0xd8, 0xd7,
		0xb2, 0xb0, 0x0d, 0xd6, 0x4d, 0xdd, 0x64, 0xdb, 0x0f, 0x79, 0x1f, 0xda,
		0x56, 0x4a, 0x31, 0x85, 0x02, 0x3f, 0xed, 0xee, 0x18, 0x0e, 0x54, 0xc5,
		0xe7, 0xf8, 0x8a, 0x82, 0x49, 0xc2, 0xea, 0x51, 0x8a, 0xc9, 0x39, 0x38,
		0xd3, 0xd1, 0x99, 0x30, 0x48, 0x09, 0xed, 0x55, 0x9d, 0xd5, 0x9c, 0x18,
		0x76, 0xef, 0x79, 0x8f, 0xc2, 0xff, 0x07, 0xb3, 0xc8, 0x40, 0x0a, 0x85,
		0x4c, 0x45, 0x32, 0x13, 0xd1, 0x82, 0x3c, 0xd7, 0x79, 0x19, 0x39, 0xdb,
		0x91, 0x93, 0xc3, 0x99, 0x4f, 0x9c, 0xc5, 0xc4, 0xd9, 0xb8, 0x3e, 0x5c,
		0xc6, 0x57, 0xee, 0xd0, 0xa8, 0xa1, 0x58, 0x0e, 0x0b, 0x79, 0x10, 0x53,
		0x1a, 0x92, 0xa7, 0xc6, 0x74, 0x72, 0xf8, 0x7f, 0xea, 0xd4, 0xb7, 0x22,
		0xf6, 0x84, 0x20, 0x45, 0x53, 0xdb, 0xd6, 0x2c, 0xe1, 0x0b, 0x78, 0x68,
		0x85, 0x3c, 0x2b, 0x73, 0xbf, 0x96, 0xea, 0x3c, 0xaa, 0xa6, 0x37, 0x04,
		0xbd, 0xde, 0xac, 0x81, 0xef, 0x97, 0x6c, 0xa4, 0x3a, 0xad, 0x6c, 0x6b,
		0xa8, 0x22, 0xa9, 0x9e, 0x6d, 0xc5, 0x34, 0x13, 0x78, 0xe6, 0x11, 0x93,
		0x3e, 0xd3, 0x4f, 0xb8, 0x20, 0x7c, 0x3f, 0xdd, 0x9b, 0x3c, 0xaf, 0x59,
		0x9a, 0xf4, 0x7e, 0xf9, 0xd0, 0xf4, 0xa2, 0x66, 0xa4, 0x95, 0xbc, 0xfc,
		0x33, 0xba, 0x37, 0xc8, 0x4d, 0x81, 0xb3, 0xce, 0x4a, 0x98, 0x6a, 0x83,
		0x94, 0xf9, 0x3f, 0xfb, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xba, 0x40,
		0x26, 0x60, 0x03, 0x00, 0x00,
	},
		"upsert.tmpl",
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
	"digest.tmpl": digest_tmpl,
	"repo.tmpl": repo_tmpl,
	"upsert.tmpl": upsert_tmpl,
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
	"digest.tmpl": &_bintree_t{digest_tmpl, map[string]*_bintree_t{
	}},
	"repo.tmpl": &_bintree_t{repo_tmpl, map[string]*_bintree_t{
	}},
	"upsert.tmpl": &_bintree_t{upsert_tmpl, map[string]*_bintree_t{
	}},
}}