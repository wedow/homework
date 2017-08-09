// Code generated by go-bindata.
// sources:
// 1_initialize_schema.down.sql
// 1_initialize_schema.up.sql
// 2_add_username.down.sql
// 2_add_username.up.sql
// 3_add_parent_id.down.sql
// 3_add_parent_id.up.sql
// 4_add_location_details.down.sql
// 4_add_location_details.up.sql
// bindata.go
// DO NOT EDIT!

package migrations

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

var __1_initialize_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x29\xca\x2f\x50\x28\x49\x4c\xca\x49\x55\x28\xc8\x2f\x2e\x29\xb6\xe6\x02\x04\x00\x00\xff\xff\xf6\x8d\x0d\x67\x12\x00\x00\x00")

func _1_initialize_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_initialize_schemaDownSql,
		"1_initialize_schema.down.sql",
	)
}

func _1_initialize_schemaDownSql() (*asset, error) {
	bytes, err := _1_initialize_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_initialize_schema.down.sql", size: 18, mode: os.FileMode(420), modTime: time.Unix(1501615514, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_initialize_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xcc\xcd\x0d\xc2\x30\x10\x44\xe1\xf3\xba\x8a\x39\x26\x12\x1d\x50\x0c\xda\xc4\x03\xb2\xf0\x9f\xec\x41\x90\xee\x51\x72\x7c\xef\xf0\xed\x83\x2e\x42\xbe\x65\xa2\xb7\xa9\x89\x25\x58\x8a\x30\xb3\x2d\xbd\x26\x47\xf2\x8c\x3e\x52\xf1\x71\xe0\xcd\xe3\x16\x6c\x6f\x55\xac\x82\x89\x3f\x9d\x7d\x19\xf1\xe1\xe7\x4a\x85\x53\x5e\x3a\x22\x9f\xfe\xc9\x42\x6d\xdf\x65\x0d\xeb\x3d\xfc\x03\x00\x00\xff\xff\x7b\xca\x9f\x3d\x6b\x00\x00\x00")

func _1_initialize_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_initialize_schemaUpSql,
		"1_initialize_schema.up.sql",
	)
}

func _1_initialize_schemaUpSql() (*asset, error) {
	bytes, err := _1_initialize_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_initialize_schema.up.sql", size: 107, mode: os.FileMode(420), modTime: time.Unix(1501782884, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2_add_usernameDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xc8\x2f\x2e\x29\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x2d\x4e\x2d\xca\x4b\xcc\x4d\xb5\xe6\x02\x04\x00\x00\xff\xff\x57\x16\xee\xd5\x28\x00\x00\x00")

func _2_add_usernameDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__2_add_usernameDownSql,
		"2_add_username.down.sql",
	)
}

func _2_add_usernameDownSql() (*asset, error) {
	bytes, err := _2_add_usernameDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "2_add_username.down.sql", size: 40, mode: os.FileMode(420), modTime: time.Unix(1502226941, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2_add_usernameUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xc8\x2f\x2e\x29\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x2d\x4e\x2d\xca\x4b\xcc\x4d\x55\x08\x71\x8d\x08\x51\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\x51\x70\x71\x75\x73\x0c\xf5\x09\x51\x50\x57\xb7\xe6\x02\x04\x00\x00\xff\xff\x79\x09\xa6\xb0\x40\x00\x00\x00")

func _2_add_usernameUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__2_add_usernameUpSql,
		"2_add_username.up.sql",
	)
}

func _2_add_usernameUpSql() (*asset, error) {
	bytes, err := _2_add_usernameUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "2_add_username.up.sql", size: 64, mode: os.FileMode(420), modTime: time.Unix(1502226944, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __3_add_parent_idDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xc8\x2f\x2e\x29\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x48\x2c\x4a\xcd\x2b\x89\xcf\x4c\xb1\xe6\x02\x04\x00\x00\xff\xff\x86\x6f\x4f\x48\x29\x00\x00\x00")

func _3_add_parent_idDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__3_add_parent_idDownSql,
		"3_add_parent_id.down.sql",
	)
}

func _3_add_parent_idDownSql() (*asset, error) {
	bytes, err := _3_add_parent_idDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "3_add_parent_id.down.sql", size: 41, mode: os.FileMode(420), modTime: time.Unix(1502228419, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __3_add_parent_idUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xc8\x2f\x2e\x29\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x48\x2c\x4a\xcd\x2b\x89\xcf\x4c\x51\x70\xf2\x74\xf7\xf4\x0b\xb1\xe6\x02\x04\x00\x00\xff\xff\x7b\xe9\x36\xe8\x2f\x00\x00\x00")

func _3_add_parent_idUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__3_add_parent_idUpSql,
		"3_add_parent_id.up.sql",
	)
}

func _3_add_parent_idUpSql() (*asset, error) {
	bytes, err := _3_add_parent_idUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "3_add_parent_id.up.sql", size: 47, mode: os.FileMode(420), modTime: time.Unix(1502228997, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __4_add_location_detailsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xc8\x2f\x2e\x29\x56\xe0\xe2\x74\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x48\xce\x2c\xa9\xd4\x41\x15\xca\x49\x2c\x41\x17\xc9\xcf\x4b\x47\x13\x2a\x49\xcd\x2d\xb0\xe6\x02\x04\x00\x00\xff\xff\x22\xa4\x9a\x4b\x5e\x00\x00\x00")

func _4_add_location_detailsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__4_add_location_detailsDownSql,
		"4_add_location_details.down.sql",
	)
}

func _4_add_location_detailsDownSql() (*asset, error) {
	bytes, err := _4_add_location_detailsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "4_add_location_details.down.sql", size: 94, mode: os.FileMode(420), modTime: time.Unix(1502239135, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __4_add_location_detailsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xc8\x2f\x2e\x29\x56\xe0\xe2\x74\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x48\xce\x2c\xa9\x54\x08\x71\x8d\x08\xd1\x41\x11\xce\x49\x2c\x51\x70\x71\x75\xf6\xf4\x75\xf4\x41\x93\xc8\xcf\x4b\xc7\x2e\x53\x92\x9a\x5b\x00\x93\xb1\xe6\x02\x04\x00\x00\xff\xff\xd5\xd2\x2f\x3a\x77\x00\x00\x00")

func _4_add_location_detailsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__4_add_location_detailsUpSql,
		"4_add_location_details.up.sql",
	)
}

func _4_add_location_detailsUpSql() (*asset, error) {
	bytes, err := _4_add_location_detailsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "4_add_location_details.up.sql", size: 119, mode: os.FileMode(420), modTime: time.Unix(1502239110, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x4d\x6f\xdb\x38\x10\x3d\x8b\xbf\x82\x35\xd0\x42\x5a\x78\x6d\x49\xd6\x97\x0d\xe4\xd2\xa6\x0b\xf4\xb0\x2d\xb0\x6d\x4f\x3b\x0b\x83\x92\x48\x97\x58\x5b\x72\x25\xb9\x99\x24\xc8\x7f\x5f\x0c\x29\x27\x76\xd6\x1f\x6d\x9a\x1e\x14\xf3\x6b\xde\xbc\xa1\xde\x23\x95\xf1\x98\xbf\xa9\x4b\xc9\x17\xb2\x92\x8d\xe8\x64\xc9\xf3\x6b\xbe\xa8\x7f\xcf\x75\x55\x8a\x4e\x8c\xd8\x78\xcc\xdb\x7a\xd3\x14\xb2\x9d\x51\x3b\x98\xeb\x4a\x77\x5a\x2c\xf5\x8d\x9c\xb7\xc5\x17\xb9\x12\xa3\xb2\xbe\xaa\x46\xed\xd7\xe5\xb1\xf9\xcd\x7a\x3b\x1b\xce\x45\x59\xce\x37\xad\x6c\x2a\xb1\x92\x7b\x81\x8f\xa6\x1e\x62\x26\x66\x62\x2d\x1a\x59\x75\x73\x5d\xee\x05\x3d\x9e\x7b\x88\x8a\xcc\xcc\xb2\x2e\x44\xa7\xeb\x6a\x5e\xca\x4e\xe8\x65\xbb\x17\x7c\x64\xc9\x03\xc6\x76\x0f\x16\x35\xf5\x2e\x3f\xf0\xf7\x1f\x3e\xf1\xb7\x97\xef\x3e\xbd\x60\x6c\x2d\x8a\x7f\xc5\x42\xf2\x95\x5e\x34\x26\xbe\x65\x4c\xaf\xd6\x75\xd3\x71\x97\x39\x83\xfc\xba\x93\xed\x80\x39\x83\xa2\x5e\xad\x1b\xd9\xb6\xe3\xc5\x8d\x5e\xd3\x80\x5a\x75\xf4\xa3\x6b\xfb\x77\xac\xeb\x4d\xa7\x97\xd4\xa9\x4d\xc0\x5a\x74\x5f\xc6\x4a\x2f\x25\x35\x68\xa0\xed\x1a\x5d\x2d\xcc\x5c\xa7\x57\x72\xc0\x3c\xc6\xd4\xa6\x2a\xb6\xf4\xfe\x92\xa2\x74\xa9\xc1\xff\xfe\x87\xd2\x0e\x39\x6d\x20\xb7\x61\x1e\x77\xb7\xa3\xb2\x69\xea\xc6\xe3\xb7\xcc\x59\xdc\x98\x1e\x9f\x5d\x70\x62\x35\x7a\x2f\xaf\x08\x44\x36\xae\xa1\x4d\xfd\xd7\x1b\xa5\x64\x63\x60\x3d\x8f\x39\x5a\x99\x80\x17\x17\xbc\xd2\x4b\x82\x70\x1a\xd9\x6d\x9a\x8a\xba\x43\xae\x56\xdd\xe8\x2d\xa1\x2b\x77\x40\x40\xfc\xe5\xd7\x19\x7f\xf9\x6d\x60\x99\x98\x5c\x1e\x73\xee\x18\x73\xbe\x89\x86\xe7\x1b\xc5\x6d\x1e\x9b\x84\x39\x73\x4b\xe7\x82\xeb\x7a\xf4\xa6\x5e\x5f\xbb\xaf\xf2\x8d\x1a\xf2\xc5\x8d\xc7\x9c\x62\xf9\x76\xcb\x74\xf4\x66\x59\xb7\xd2\xf5\xd8\x73\xf1\x21\x18\x8b\x7f\x04\x48\x36\x8d\xe5\xdd\x0f\xe6\x1b\x35\x7a\x4d\xd4\x5d\x6f\x48\x2b\xd8\x1d\x63\xdd\xf5\x5a\x72\xd1\xb6\xb2\xa3\x2d\xdf\x14\x1d\xa1\x98\xfa\xfa\xf7\xc1\x1c\x5d\xa9\x9a\xf3\xba\x1d\xfd\xa1\x97\xf2\x5d\xa5\xea\xfb\xb8\xfe\x15\x6e\xc7\x77\x10\xcc\x3b\xe4\xbc\x7f\x8d\xcc\x69\xf5\x8d\xe9\xeb\xaa\x4b\x22\xe6\xac\xc8\xb3\xfc\x1e\xf4\xcf\xba\x94\x66\xf0\x93\x5e\x49\x4e\x32\x19\x51\x8b\xf2\x18\xa9\xb8\x4a\x3f\xce\xe5\xf1\xf7\x62\x25\x5d\xaf\xcf\x40\x39\xfb\x2a\x95\x1e\x51\x76\x76\x77\x22\xf6\xa3\xbe\xa1\x58\xc3\x66\x3f\x94\x88\x9e\x0c\x25\xae\xae\xb7\xcb\x7c\x1f\x80\x4a\x3b\x07\x40\xc5\xb9\xde\x43\xa1\xff\x43\xe8\xab\x3f\x0e\xf2\xae\xbd\xd4\x8d\xeb\xf1\xbc\xae\x97\xbb\xd1\x62\xd9\x9e\xa9\xfc\xba\xb5\x85\xcb\x46\x89\x42\xde\xde\xed\x44\xf7\x92\x20\x95\xcf\xe7\x07\x4e\xc2\xcb\xfa\xaa\xfa\xf8\x75\xc9\x2f\x7a\x69\xb8\x03\xc0\x40\x01\x66\x39\xa0\x9f\x01\xfa\xfe\xe1\x47\x29\xc0\x48\x00\x86\x53\xc0\x82\x7e\x15\x60\xec\x03\x86\x19\x60\x34\x05\x8c\x0a\x3b\x4e\xed\x38\xb6\xe3\x45\x66\xd7\x85\xd2\xc6\xe5\x09\xa0\x4c\x00\xfd\x10\xd0\x8f\xf6\xb1\xcd\x93\x00\x66\x25\xa0\x5f\x02\x26\x29\x60\x10\xee\x72\x18\x6c\x0f\x9d\x13\x65\xf5\xce\x38\x74\xe2\x6c\xfd\xb3\x73\x62\x31\xc7\x39\xb5\x47\x43\xe6\x38\x83\x53\x97\xcd\x60\xc8\x1c\xef\x5e\xe0\x27\x90\x88\xd1\x6f\xc6\xa0\xbb\x8c\x8c\x43\xef\x8f\xc1\xf3\x55\x9d\x3b\x73\xee\x8f\x0a\x63\xf6\xd9\xc5\x63\xe1\xdc\x92\xa5\x66\xfc\x4c\x49\x9c\xdc\x33\xe3\x41\x36\xe4\x64\x83\xd9\xae\x4b\xdc\x28\xf4\x3d\x33\x4e\xe2\x9e\x59\xf1\x7f\xae\x34\xba\x41\xec\x07\x49\x10\xc7\x41\x34\xe4\xbe\x77\xc7\x1c\x41\x04\x5e\x99\x9a\x6f\x4d\xa1\x33\xde\xd7\x4b\xec\x66\xe6\xef\xdd\xfd\x5b\x11\xc3\xb3\xc2\xfd\xbc\x7e\xaa\x6c\x03\x92\x26\x3d\xbd\xb4\x8a\x10\x70\xe2\x03\x06\x3e\x60\x14\x01\xca\x00\x50\x4d\x00\x73\x01\x98\x09\xc0\xc9\x14\x30\x4c\xac\xfc\x82\xd2\xca\xdc\x2f\x00\x4b\x01\x58\x90\x6c\x69\x6d\x08\xa8\x7c\xc0\xa9\x02\x94\x05\x60\x14\x00\x4e\x7d\x40\x29\x01\xe3\x00\x30\x0d\x01\xd3\x02\x50\x2a\xbb\x4e\x96\x80\xd9\xc4\x5a\x21\x0a\x01\x73\x09\x98\xc4\x80\x82\xda\x29\xa0\x98\x02\x66\x94\x37\x06\x8c\xb3\x9e\x87\x0f\x98\x53\x4c\x09\x98\x97\x96\x53\x94\x02\xaa\x10\x30\x2b\x00\x27\x94\x8b\x78\x50\xbe\x00\x50\xfa\xb6\x46\x39\x01\x0c\x12\xc0\xa4\x00\x4c\x94\xb5\xa3\x28\x00\xb3\xd0\xe6\x98\x28\xc0\x69\x09\x98\x96\x80\xc1\xd4\xc6\xd3\x1e\xc8\xd4\x5a\x3c\x23\x0e\x13\xc0\x58\x02\x4e\xc8\xea\xa1\xad\x53\x49\xc0\x62\x6a\xf9\x27\x25\x60\xa9\x6c\x0d\xb4\xa7\x32\x07\x9c\x94\x80\xaa\xb0\xfb\xf3\xd8\xda\x69\x6e\x8f\x07\xc2\xa1\x75\x49\xfe\xdd\xd6\x36\x2f\xfe\x79\x8c\x6d\xa0\x8e\xda\xda\x7e\x75\x9d\x37\xb5\x41\x79\xaa\xa5\x77\xab\xf9\x95\x86\xde\x16\xb3\xb5\xb3\x9f\xfe\xb0\x9f\xd3\x2c\xcc\xb2\xe7\xf1\xf3\xfe\x97\xf5\xcf\xdc\x41\x64\x2c\x45\x26\x9c\x5a\xd1\x93\xf8\xc8\x04\x14\x47\xfd\x34\xb2\x06\xa1\xf6\xb1\x7b\x28\x4e\x00\x53\xdf\x62\xd0\x5a\xbf\xbf\xcb\x68\x8c\xee\x20\x95\xf6\x73\xb1\x7d\xc8\x0c\x84\x43\x46\x8c\xa4\xfd\x35\x77\x5d\x6e\x0f\x96\x88\x0c\x1a\x9f\xbe\xdb\xe2\xd4\x9a\x92\x0e\x88\xb2\xe7\x75\xc8\x00\x07\x77\xea\x09\xea\x3f\x88\x63\xa4\x7f\xe4\x1f\xa0\x47\xba\x3f\x18\xff\x5d\xa2\x3f\x55\xc1\x73\x29\xfe\x68\x0d\xbd\xdc\x23\xff\x47\xd5\x1e\x86\x61\x32\x8d\x82\x5f\xa0\xf6\xa7\x5f\x5c\xcf\xa9\x75\x5a\x6b\xd6\x05\x3f\xa7\x73\xca\xb3\xcd\x6f\xbe\xd5\x32\x8b\x49\x97\x1c\x5d\xa2\xc6\x4b\x7d\xdf\xcf\x7b\xec\x9d\xbc\x86\x77\x0c\x98\x4e\xec\xa5\xaa\x62\xf8\x2f\x00\x00\xff\xff\xd2\x93\xf1\xef\x00\x10\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 12288, mode: os.FileMode(420), modTime: time.Unix(1502239160, 0)}
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
	"1_initialize_schema.down.sql": _1_initialize_schemaDownSql,
	"1_initialize_schema.up.sql": _1_initialize_schemaUpSql,
	"2_add_username.down.sql": _2_add_usernameDownSql,
	"2_add_username.up.sql": _2_add_usernameUpSql,
	"3_add_parent_id.down.sql": _3_add_parent_idDownSql,
	"3_add_parent_id.up.sql": _3_add_parent_idUpSql,
	"4_add_location_details.down.sql": _4_add_location_detailsDownSql,
	"4_add_location_details.up.sql": _4_add_location_detailsUpSql,
	"bindata.go": bindataGo,
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
	"1_initialize_schema.down.sql": &bintree{_1_initialize_schemaDownSql, map[string]*bintree{}},
	"1_initialize_schema.up.sql": &bintree{_1_initialize_schemaUpSql, map[string]*bintree{}},
	"2_add_username.down.sql": &bintree{_2_add_usernameDownSql, map[string]*bintree{}},
	"2_add_username.up.sql": &bintree{_2_add_usernameUpSql, map[string]*bintree{}},
	"3_add_parent_id.down.sql": &bintree{_3_add_parent_idDownSql, map[string]*bintree{}},
	"3_add_parent_id.up.sql": &bintree{_3_add_parent_idUpSql, map[string]*bintree{}},
	"4_add_location_details.down.sql": &bintree{_4_add_location_detailsDownSql, map[string]*bintree{}},
	"4_add_location_details.up.sql": &bintree{_4_add_location_detailsUpSql, map[string]*bintree{}},
	"bindata.go": &bintree{bindataGo, map[string]*bintree{}},
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

