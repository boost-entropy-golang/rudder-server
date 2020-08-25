// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package migrator

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// MigrationAssets contains SQL migration scripts and templates
var MigrationAssets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 8, 21, 18, 35, 58, 746206690, time.UTC),
		},
		"/jobsdb": &vfsgen۰DirInfo{
			name:    "jobsdb",
			modTime: time.Date(2020, 8, 21, 18, 35, 58, 745819855, time.UTC),
		},
		"/jobsdb/000001_create_tables.down.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "000001_create_tables.down.tmpl",
			modTime:          time.Date(2020, 8, 21, 18, 35, 58, 745530408, time.UTC),
			uncompressedSize: 265,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x84\x8e\xb1\x8a\x83\x40\x10\x86\xfb\x7d\x8a\xbf\xb8\x56\x5f\xc0\xea\x0e\x3d\x10\x0e\x4e\xa2\x45\x52\x2d\x2b\x6e\x44\x91\x55\x76\x47\x50\x86\x79\xf7\x10\x13\x93\x90\x26\xe5\xcc\x7c\xff\xfc\x5f\x14\x21\xf5\xe3\x04\xe3\x56\xd8\xa5\x0b\xd4\xb9\x16\x8d\x21\x13\x2c\x05\xc5\xec\x8d\x6b\x2d\xe2\xf4\xbe\x11\x51\x00\x90\x1e\xfe\x0b\x54\xdf\x3f\x7f\x19\x98\xbf\xe2\xc2\xdb\x73\xb7\x88\xe8\x7e\xac\x83\x66\x8e\x45\x92\x4f\x9c\x0e\x64\x68\x7e\xd0\xcc\xd6\x35\x22\x4a\xed\x42\xb4\x4e\x16\x3b\x67\xf5\x75\x54\xb7\x77\xa7\x22\x7b\x3b\x24\xcf\x58\x3f\xce\xde\x99\x01\x64\xea\x61\x0f\x6c\xfd\xf9\x2f\xb2\x63\x5e\x56\x25\x98\x5f\x45\x36\x3c\x51\xea\x12\x00\x00\xff\xff\xa5\xa4\xe3\xe3\x09\x01\x00\x00"),
		},
		"/jobsdb/000001_create_tables.up.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "000001_create_tables.up.tmpl",
			modTime:          time.Date(2020, 8, 21, 18, 35, 58, 745646260, time.UTC),
			uncompressedSize: 1358,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8c\x54\x51\x6f\xe2\x38\x10\x7e\xe7\x57\xcc\x03\x12\xad\x04\x9c\x74\xf7\x76\x7d\x0a\xe0\xdb\x4d\x0f\x02\x82\x74\xaf\x7d\x8a\x26\xf6\x00\xee\x1a\x3b\x67\x3b\x2d\x51\x94\xff\x7e\x72\x48\xba\x57\xb6\x2b\x35\x8f\x33\xdf\xf7\xcd\xcc\x97\x19\x4f\x26\x10\x6b\xe9\x25\x2a\x78\x21\xeb\xa4\xd1\x60\xf6\x70\x6f\x72\xb7\x98\x81\xc7\x5c\x91\x1b\x43\x8e\x8e\x04\x18\x0d\x9e\x4e\x85\x42\x4f\x20\xd0\x23\x14\xd6\xbc\x48\x41\x02\xf2\x0a\x08\xf9\xb1\xa7\x49\xed\x3c\x6a\x4e\x83\xc9\x04\xe6\x47\xe2\xdf\xe1\xd9\xe4\x4e\xe4\xbf\x39\xf2\x65\x31\x3d\x18\xd8\x1b\x0b\xa8\x14\xe0\x0b\x4a\x15\x8a\xbc\x57\x9e\x0e\x02\xf5\xde\x94\x56\xa3\x72\x97\x36\x06\xf3\x2d\x8b\x52\x06\x69\x34\x5b\x32\x88\xff\x82\x64\x9d\x02\x7b\x8c\x77\xe9\x0e\xea\x7a\xba\xb1\xb4\x97\xe7\xa6\xc9\x9e\x2f\x2c\xb8\x19\x00\x00\x48\x01\xb3\xf8\xcb\x8e\x6d\xe3\x68\x09\x9b\x6d\xbc\x8a\xb6\x4f\xf0\x37\x7b\x1a\xb7\x59\x53\x90\x45\x1f\x66\xfe\x16\x6d\xe7\x5f\xa3\xed\xcd\x1f\xbf\xdf\xb6\xc2\xc9\xc3\x72\x79\xc1\x08\xa3\x09\x66\xeb\xf5\x92\x45\xc9\x15\x2b\x2b\xb0\x52\x06\x05\xdc\xef\xd6\xc9\xec\x8a\xe7\x3c\x5a\x9f\x79\x79\x22\x48\xe3\x15\xdb\xa5\xd1\x6a\x73\x05\x21\x2d\xae\x00\xb7\x77\xdd\xe4\x79\xe0\x7b\x02\x5f\x15\x34\x86\xd2\x5d\x4c\x0e\x36\x76\x3f\x65\xb0\x58\xc3\x70\x08\x33\xf6\x25\x4e\x5a\xb1\xde\x9e\xa7\x0d\x0b\xb8\xac\xe5\x67\x81\xdf\xa6\xc3\x17\xed\x80\x25\x0f\xab\x9b\xb7\x40\xff\x8d\x5e\x51\x7a\xa9\x0f\xa3\xf1\xcf\x29\x3a\x13\x2f\x7f\x95\x74\x25\xe7\x44\x82\xc4\x47\xc9\x4e\x34\xb3\xe4\x6d\xf5\x11\x60\x8f\x52\x7d\x4c\xc5\xdc\x58\x4f\x62\x74\x7b\xd7\xe6\xd8\xe3\x9c\x6d\xd2\x78\x9d\xbc\x21\xff\xf9\xca\x12\x10\x65\xa1\x24\x0f\x63\x9a\xfc\x99\xb8\x87\x34\x44\x75\xa9\xd4\xdd\x80\x25\x0b\x18\x0e\x2f\x76\x2e\xd0\xa3\x23\xdf\x39\x07\x68\x09\xb4\xf1\xc0\x2d\xa1\x27\x01\x42\x5a\xe2\x5e\x55\xc1\xe1\x93\x3c\x74\x1b\xe1\xb8\x95\x85\x77\x63\xf0\x47\x6a\x29\x3d\xfc\x45\x62\xb7\xe8\x23\xd7\x05\x17\x3b\xd8\x97\x9a\xb7\xbc\xb0\xfc\x84\x62\x1a\x0a\xa7\x47\x82\x7f\x4b\xb2\xd5\xe5\x0f\x4a\xfd\x03\x7f\x2a\x9d\x07\x54\xaf\x58\xf5\x22\x6d\x21\xf1\xbe\xd5\xd2\x49\x7d\x68\x13\xe1\x36\x9c\x07\xc7\x8f\x74\xc2\xfe\x54\xdb\x22\xdf\xba\xb3\x2d\x0b\x11\x30\xed\x6d\xd1\x59\xba\xe0\xfd\xb5\x1e\x47\x0d\x39\x81\xa0\xbd\xd4\x24\xc6\xbd\xfe\xbb\xab\xfe\x4e\x15\x8c\x3a\xcb\xdc\x08\xbc\x01\xe9\xc3\xc2\xd3\xcf\xaa\x52\x0b\xc9\xc9\x75\xb3\x4a\x07\xd2\x01\x6a\xa0\x33\x9e\x0a\x45\xe1\x21\xc1\x37\xec\xb5\xb3\x7d\xf1\x1f\xf1\xbe\x0d\xf7\xe7\x60\x32\x09\x92\x75\x6d\x51\x1f\x08\xa6\x7d\x37\x4d\x13\xc2\xed\x2a\x2f\x53\xb6\xed\xde\x82\xba\x1e\xfe\xff\xfc\x73\x97\xd5\xf5\xb4\x69\x20\x5a\x2c\x60\xbe\x5e\x3e\xac\x12\xd0\xf4\x9a\x71\xa3\xca\x93\x86\x94\x3d\xa6\x77\x9f\x91\x69\x2f\xa8\xfc\xa4\x58\x5d\x93\x16\x4d\xf3\x5f\x00\x00\x00\xff\xff\x7d\x20\xde\x2b\x4e\x05\x00\x00"),
		},
		"/jobsdb/000002_alter_dataset_tables.down.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "000002_alter_dataset_tables.down.tmpl",
			modTime:          time.Date(2020, 8, 21, 18, 35, 58, 745757531, time.UTC),
			uncompressedSize: 802,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xa4\x91\x41\x6a\xf3\x30\x10\x85\xf7\x39\xc5\x2c\x0c\xf9\x7f\x68\x72\x01\xaf\x9c\x58\x6d\x03\x8e\x6c\x52\x99\xb6\x2b\x21\x5b\x93\xa0\x60\xec\x20\xc9\x34\x41\xe8\xee\xc5\x76\x1a\x48\x71\x42\x69\x67\x61\xf0\xcc\x9b\x6f\xc4\x7b\xce\x69\x51\xef\x10\xe6\xb1\xb0\xc2\xa0\x35\xde\x4f\x00\x00\xa2\x84\x91\x0d\xb0\x68\x91\x10\x70\x2e\x98\x67\x1a\xb7\xea\xe8\x3d\xdf\x37\x85\xe1\xce\xcd\xbd\x87\x78\x93\x66\xb0\x4c\x93\x7c\x4d\xa1\x35\xa8\xb9\x92\xe1\x8f\x97\x07\xcd\x79\xbb\xd4\x28\x2c\x4a\x2e\x2c\x18\xb4\x10\x93\xc7\x28\x4f\x18\xd0\x3c\x49\x7e\x49\xc4\xe3\x41\x69\x1c\x07\xf6\xc4\xd9\x0c\x34\x0e\x77\x61\xdf\x14\xdc\x58\x61\x91\xdb\xd3\x01\xa1\xfb\xf4\x9a\x38\x85\x20\x80\x05\x79\x5a\xd1\xfe\xbf\xab\xe5\x86\x44\x8c\x00\x7b\xcf\xc8\xb7\xbd\x8b\xa4\x7f\xf0\x0b\x10\x9a\xaf\xff\x5d\x35\xbf\x6a\xfa\x21\x94\x55\xf5\x6e\xfa\x30\x3e\xc6\x23\x96\xed\x3d\x81\x69\xcb\x12\x51\xa2\xbc\x25\x38\x1f\xe0\x1a\xad\x3e\xdd\x12\x6d\x85\xaa\x6e\x23\x44\xd1\x68\x8b\x72\xfa\x3f\xbc\x9a\x93\xb7\x25\xc9\xd8\x2a\xa5\x57\xdd\xd7\x67\x42\x41\xb6\x87\x4a\x95\x9d\x1d\x4d\xb1\xc7\xd2\x02\xeb\xba\x75\x5b\x55\x03\x82\xd0\x18\x82\xe0\xec\xff\xbd\x44\x7b\x53\xdb\xd1\x5c\x2f\x96\x8f\x25\x10\xfe\x95\xac\xe4\x80\x5d\x51\x16\x4e\x9c\xc3\x5a\x7a\xff\x19\x00\x00\xff\xff\xea\x0e\x28\x43\x22\x03\x00\x00"),
		},
		"/jobsdb/000002_alter_dataset_tables.up.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "000002_alter_dataset_tables.up.tmpl",
			modTime:          time.Date(2020, 8, 21, 18, 35, 58, 745873428, time.UTC),
			uncompressedSize: 718,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xa4\x91\x51\x6b\x9c\x40\x14\x85\xdf\xfd\x15\xe7\x41\x48\x0a\xd9\x85\x42\xe9\x8b\x4f\xee\x3a\xc9\x0a\x66\x14\x77\xb6\xd9\x3e\xc9\xb8\xde\x46\x83\xa8\xcc\x8c\x69\xc2\x30\xff\xbd\x8c\xdb\x6e\x28\x14\x52\xda\x37\xef\xf5\xdc\xef\xde\x73\xc6\x5a\x25\x87\x47\xc2\x3a\x91\x46\x6a\x32\xda\xb9\x00\x00\xe2\x4c\xb0\x12\x22\xde\x64\x0c\xd6\x86\xeb\x42\xd1\xb7\xee\xc5\xb9\xea\x69\xac\x75\x65\xed\xda\xb9\x9f\x9a\x6d\x9e\x1d\xee\x39\x4e\x8a\xa4\xa1\xa6\x92\x06\x9a\x0c\x12\x76\x1b\x1f\x32\x01\x9e\x3f\x5c\x7f\x88\xfe\x0d\x49\x2f\x53\xa7\xe8\x7f\x89\x49\xf2\x8b\x97\xde\x82\xe7\x02\xec\x98\xee\xc5\x1e\xb3\x26\x55\x75\x0d\x04\x3b\x8a\xa5\xcf\x0f\x59\x76\xd9\x72\xb5\xfa\x78\x15\x05\xcb\x96\xd5\x0a\xdb\x71\x78\x26\x65\xf0\x34\xd6\x95\x36\xd2\x10\xcc\x88\x67\xa9\x4e\xad\x54\x37\x68\xd4\x38\xc1\xbc\x4e\xf4\xf6\xbf\xf2\xe5\xbb\x37\x2e\xda\xf9\x8f\xde\xdf\x36\x89\xaf\x05\xc3\x97\xb8\xdc\xee\xe2\xf2\xfa\xf3\xa7\xbf\xb0\xfe\x0e\xd6\x7b\xf6\xcc\x4d\x7a\x97\x72\x11\x05\x17\x97\xf7\xdd\x63\x6b\xa0\x4d\xd7\xf7\xa8\xc9\x07\xd4\xa0\x7e\xc5\x68\x5a\x52\x7e\x52\x37\x35\xba\x41\x1b\x39\x9c\xe8\x06\xbd\xd4\x06\xe3\x40\x98\xa7\xc6\xbf\x3c\xbe\xfb\xb9\x73\x16\x2d\x2d\x79\xac\x17\x72\x92\x23\x0c\xb1\x61\x77\x29\x5f\xea\xa5\x57\xe6\xc5\xf9\x8a\xdf\x23\x8b\x2e\x0a\x76\xdc\xb2\x42\xa4\x39\xc7\xc3\x8e\x71\xe4\x62\xc7\xca\x3d\x84\xff\x1e\xe6\xbe\x3f\x0b\x19\x4f\x10\x86\x51\x60\x2d\x0d\x8d\x73\xc1\x8f\x00\x00\x00\xff\xff\x3e\x89\xb1\x86\xce\x02\x00\x00"),
		},
		"/node": &vfsgen۰DirInfo{
			name:    "node",
			modTime: time.Date(2020, 8, 24, 8, 57, 49, 415538602, time.UTC),
		},
		"/node/000001_create_event_schema.down.sql": &vfsgen۰CompressedFileInfo{
			name:             "000001_create_event_schema.down.sql",
			modTime:          time.Date(2020, 8, 21, 18, 35, 58, 746029747, time.UTC),
			uncompressedSize: 279,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xd2\xd5\xd5\xe5\xd2\xd5\xd5\x55\x70\x2d\x4b\xcd\x2b\x51\xf0\xcd\x4f\x49\xcd\x29\x06\x09\x70\x71\xb9\x04\xf9\x07\x28\x84\x38\x3a\xf9\xb8\x2a\x78\xba\x29\xb8\x46\x78\x06\x87\x04\x2b\xa4\x82\x94\xc5\xe7\x82\x95\x59\x43\xd5\x78\xfa\xb9\xb8\x46\x60\x57\x13\x5f\x5e\x94\x59\x92\x1a\x9f\x9d\x5a\x19\x9f\x99\x97\x92\x5a\x61\xcd\xc5\x05\xb3\x50\x21\x38\x39\x23\x35\x37\x51\x21\x2c\xb5\xa8\x38\x33\x3f\x0f\x9f\xa5\xc5\x60\x95\xf1\x65\x50\x95\xc4\xd8\x9b\x99\x02\xb3\x90\x18\xa5\x50\x0b\x32\x12\x8b\x33\xa0\xda\x00\x01\x00\x00\xff\xff\xcc\x7f\x3b\xaa\x17\x01\x00\x00"),
		},
		"/node/000001_create_event_schema.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000001_create_event_schema.up.sql",
			modTime:          time.Date(2020, 8, 21, 18, 35, 58, 746148807, time.UTC),
			uncompressedSize: 945,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x91\x4d\x6f\xe2\x30\x10\x86\xcf\xc9\xaf\x98\x1b\x44\x22\x97\x5d\x69\x0f\x70\x32\x60\x76\xbd\xcd\x07\x4d\x1c\x0a\x27\x2b\xc5\x83\x70\x0b\xa1\x4a\x0c\x2d\xaa\xfa\xdf\x2b\x87\xcf\x40\x11\x70\xcd\x3c\xf3\x4e\xfc\xbc\xae\xeb\xda\xae\xeb\x02\x5d\x61\xa6\xc1\x5f\x48\x9c\x15\xe6\x83\x6d\x77\x22\x4a\x38\x05\x4e\xda\x1e\x05\xd6\x83\x20\xe4\x40\x87\x2c\xe6\x31\xa0\x81\xc5\xbc\x84\xa1\x6e\x5b\x96\x92\x10\xd3\x88\x11\x0f\xfa\x11\xf3\x49\x34\x82\x07\x3a\x6a\xd8\x96\xb5\x5c\x2a\x09\x03\x12\x75\xfe\x91\xa8\xfe\xfb\x8f\x53\xa6\x04\x89\xe7\x99\xe1\x7b\xae\x34\x8a\x57\x5c\x1f\x88\x5f\x55\x62\x73\x48\xaf\xdf\x10\x38\x1d\xf2\x1f\x66\xe5\x4f\x08\x25\x31\xd3\x6a\xa2\x30\xaf\x72\xd0\xa5\x3d\x92\x78\x1c\x6a\x35\xb3\x32\xce\x31\xd5\x28\x45\xaa\x81\x33\x9f\xc6\x9c\xf8\xfd\x73\x36\x08\x9f\xea\x8e\xd3\xda\x1b\x60\x41\x97\x0e\x2f\x1b\x10\xfb\x67\x08\x95\x49\xfc\x80\x30\x38\x11\xb4\x07\x4c\xe8\x4e\x38\xc4\xe3\x29\xce\x53\x18\x60\x5e\xa8\x45\x76\x5d\x7a\x51\xf2\x62\xb5\xe5\x77\xde\xdb\xec\xef\xbd\xea\xe1\xcc\xde\xc5\x8a\xb6\x47\xa7\x69\x31\xbd\x58\xd2\x86\x81\xff\x71\x18\xb4\x2b\x83\x39\xea\x54\xa6\xfa\x74\x74\x28\xe5\xf3\xab\xd6\x6c\xbe\x14\x8b\xec\xd9\xe0\x13\x95\x17\x5a\x14\x88\xd9\xd5\x76\x0c\x3e\x4b\x6f\xa5\xef\xe8\x52\xc9\x43\x89\x67\xc2\xab\xa0\xd3\xda\x85\x26\x01\x7b\x4c\x6e\xca\x3e\xd2\x79\xf3\x9d\x06\x1c\x6d\x39\xad\xef\x00\x00\x00\xff\xff\x0d\x70\xf8\x26\xb1\x03\x00\x00"),
		},
		"/node/000002_create_col_counters_event_schema.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000002_create_col_counters_event_schema.up.sql",
			modTime:          time.Date(2020, 8, 24, 10, 49, 48, 329042306, time.UTC),
			uncompressedSize: 309,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8c\x90\x4f\x4b\x03\x31\x10\x47\xcf\xcd\xa7\xf8\xdd\xaa\x87\x80\xe7\xf6\x94\xed\xa6\x12\x89\x59\xec\x26\xe0\x6d\x89\xdd\x80\x91\x36\xd1\xfc\xe9\x45\xfc\xee\x92\x05\x3d\x89\xf4\x38\xcc\xf0\xde\x63\x28\xa5\x84\x52\x0a\x36\xcf\xd8\xc5\x53\x3d\x07\x94\x88\x5c\x62\x72\x38\xc6\x1a\x8a\x4b\x19\xb3\x2d\xb6\x5d\x11\xc2\xa4\xe6\x07\x68\xd6\x49\x8e\x7c\x7c\x75\x67\x3b\x5d\x5c\xca\x3e\x86\x0c\xb2\x5a\xb1\xbe\xc7\x6e\x90\xe6\x51\x41\xec\xa1\x06\x0d\xfe\x2c\x46\x3d\xe2\x3d\xf9\x8b\x2d\x6e\x6a\x24\x3c\x8c\x83\xea\x96\xad\x32\x52\xa2\xe7\x7b\x66\xa4\xc6\xfa\xf3\x6b\xbd\xd9\xbc\xe5\x18\x5e\xb6\xff\x89\x7e\x3d\x6a\xd4\x07\x26\x94\x46\x0d\xfe\xa3\xba\xa9\x56\x3f\xc3\x28\xf1\x64\x38\x6e\xda\x70\x7b\x1d\xe7\x8f\xde\x12\x8b\x3d\x4d\xcb\x03\xd0\x89\xfb\x26\xf9\xc9\xbc\xdb\x92\xef\x00\x00\x00\xff\xff\x22\x3a\x04\x17\x35\x01\x00\x00"),
		},
		"/node/000002_del_col_counters_event_schema.down.sql": &vfsgen۰CompressedFileInfo{
			name:             "000002_del_col_counters_event_schema.down.sql",
			modTime:          time.Date(2020, 8, 24, 10, 49, 59, 165304150, time.UTC),
			uncompressedSize: 230,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x9c\x8f\x41\x0a\x83\x30\x10\x45\xf7\x9e\xe2\x5f\x20\x27\x70\x95\x5a\x0b\x82\xd5\xa2\x29\x74\x17\x82\x09\x34\xa0\x99\x36\x99\xf1\xfc\xc5\x9e\xa0\x74\xff\xff\x7b\x3c\xa5\x54\xa5\x94\x82\xf6\x1e\x0d\xad\xb2\x25\x30\xa1\x30\xe5\x80\x85\x24\x71\xc8\x05\xde\xb1\x3b\x56\x55\xa5\x7b\xd3\x4e\x30\xfa\xd4\xb7\x28\xcb\x33\x6c\xce\xee\x21\x97\x48\xa9\xe0\x3c\x8d\x37\x34\x63\x7f\xbf\x0e\xe8\x2e\x68\x1f\xdd\x6c\x66\xbc\x72\xdc\x1d\x07\x7b\x30\xea\x5f\xfe\xc3\x6c\x26\xdd\x0d\x06\x92\xe2\x5b\x82\x15\x89\xbe\xfe\x43\xcc\xc4\x6e\xb5\xdf\x86\xfa\x13\x00\x00\xff\xff\x73\xa7\x01\x8c\xe6\x00\x00\x00"),
		},
		"/warehouse": &vfsgen۰DirInfo{
			name:    "warehouse",
			modTime: time.Date(2020, 8, 21, 18, 35, 58, 746258607, time.UTC),
		},
		"/warehouse/000001_create_tables.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000001_create_tables.up.sql",
			modTime:          time.Date(2020, 8, 21, 18, 35, 58, 746310726, time.UTC),
			uncompressedSize: 4017,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xd4\x57\x4f\x73\x9b\x3e\x10\x3d\xff\xf8\x14\x7b\xf0\x21\x9e\x89\x6f\xbf\xe9\x85\x13\xb6\xd5\x84\x16\x83\x07\x94\xd6\x39\x69\x54\x50\x12\xcd\x10\xf0\x80\xdc\xa6\xdf\xbe\x03\x22\x20\x30\xc8\x4a\xe2\x64\xda\x9b\xc7\xbb\xfb\xd8\x3f\x6f\xf5\x24\x6b\xb1\xb0\x16\x0b\xf8\xf5\x40\x4a\x41\xef\x79\x76\x4f\xee\x78\xca\xca\xea\x6f\x6b\x15\x22\x07\x23\xc0\xce\xd2\x43\xe0\x7e\x06\x3f\xc0\x80\x76\x6e\x84\xa3\x23\x7f\xb8\xb0\x00\x00\x78\x02\x4b\xf7\x2a\x42\xa1\xeb\x78\xb0\x0d\xdd\x8d\x13\xde\xc2\x57\x74\x7b\x59\x5b\xd3\x3c\xa6\x82\xe7\x19\x60\xb4\xc3\x35\x9a\x7f\xe3\x79\xd2\x56\xe6\x87\x22\x66\x84\x27\xf0\xcd\x09\x57\xd7\x4e\x78\xf1\xe9\xff\xf9\xc0\x27\x61\xa5\xe0\x59\x0d\xa1\x77\x2c\xe3\x07\xf6\x48\xe1\x4b\x14\xf8\xcb\x81\x89\x15\x45\x5e\xd4\x09\x34\xae\x82\x8a\x43\xa9\x62\xc9\xff\xef\x78\x51\x0a\xc2\x7e\xb2\x4c\x10\x2a\x00\xbb\x1b\x14\x61\x67\xb3\x6d\x2a\xa1\x1a\xa3\xc8\x05\x4d\xa5\xb5\xac\xda\xe1\xfa\xcd\xb7\xe2\x82\x51\xc1\x92\x5e\xc8\x20\xbd\xc3\x3e\x99\x76\x99\xdb\x96\xe5\x78\x18\x85\xcd\x48\x8e\x86\x56\x21\x38\xeb\x35\xac\x02\xef\x66\xe3\x0f\x46\xa6\xaf\x68\x32\x4c\x5b\xea\x64\xd4\x48\x0f\xec\x96\x50\xae\xbf\x46\xbb\x13\x84\x22\x3c\x21\x3c\x4b\xd8\x13\x04\xfe\x08\xdb\x5a\xba\x5c\x0e\x58\x71\xaa\x47\x20\x8d\x4d\xd2\xcd\xf4\xf1\xed\x16\xa9\x14\xb0\x2d\x6b\x1d\xc0\x6c\x06\x4b\x74\xe5\xfa\x75\xa5\xeb\x30\xd8\x4a\x3f\x05\xb1\x0a\x67\x44\xfc\xde\x33\xbb\x76\x42\xbb\x15\xda\x62\x37\xf0\xe1\xfb\x35\xf2\x21\xc0\xd7\x28\x8c\x00\x57\xbf\xb3\x43\x9a\xda\x16\xf2\xd7\x30\x9b\xd9\x56\xb7\x75\x69\x4e\x13\xe3\x95\xeb\x9c\x8d\xf6\x4d\x2d\x9c\x48\xd7\x96\x8b\x1f\xb6\x8b\xaa\x63\xd5\x29\x8d\xab\xa0\x3f\x52\x46\x32\xfa\xc8\xc6\xb2\x7a\xdd\x5a\x8d\xf0\x41\xe9\x62\x8f\x0c\xea\xe7\xab\x41\x57\x39\xe8\xa3\xcf\xc9\xfe\x0e\x97\x34\xfd\xef\xb7\x98\x74\xe9\xf5\x16\x43\xe5\xc4\xe4\x56\x5c\x2a\xc5\xcd\x55\xfa\x1d\xf6\x55\xbc\x09\xf7\x1a\x4f\x33\xe2\x19\x10\xa8\x4a\xa5\xdc\xd3\x98\x7d\x30\xc9\x4a\x41\x0b\x41\x74\xab\xc1\xb2\x44\x6b\x97\x08\x6d\xdf\xc7\xc2\x27\x8d\xc7\x72\xf3\x52\xe9\xaa\x2d\x6f\xd6\x28\x69\x7c\x62\xf1\x98\x7e\xf1\x47\x9e\xdd\x97\xea\x97\xde\x41\xb9\x9e\x99\xf7\xf7\x68\x96\xae\x27\xd3\xab\xae\x36\xeb\xe4\x92\x37\x45\x13\x49\x83\xde\x1e\xb7\xfb\x25\x6d\x73\x73\xac\xd1\xd3\x62\x1c\xfa\x05\xaa\xf9\x1c\x74\x0e\xbd\x94\x58\xe7\x90\x4b\x79\x8c\x19\x9e\x5a\xff\x0d\x03\x8c\x0e\xaf\x2e\xdf\x9e\x9f\x56\xaa\x0c\x97\x7b\x78\xf9\xec\x18\x27\x78\x85\xa4\xb9\x47\x76\x41\xef\xb0\x8b\xfd\x26\xbd\x50\xd6\xf4\x58\x2a\x7b\x9a\xf2\x15\x6d\xd5\x33\xbc\x07\x45\xd4\xb9\x4c\x89\xe1\x60\xda\x6a\xc8\x50\x02\xcd\xd3\x7e\x25\xe9\x55\xc4\x73\x50\x5f\x2a\x83\xd1\xcb\x4c\x7a\xbe\x81\xed\xff\x80\x8c\x9b\x3f\xf1\x4e\xde\x0f\xeb\xb1\xb5\x1c\x3c\xea\xa3\xee\x84\x3d\xfd\xa8\x69\x30\x06\xc1\x6d\xdb\xfa\x8f\x9c\xe7\xc1\x0d\x6f\x6f\xad\xfb\xdc\xfe\x13\x00\x00\xff\xff\x89\x94\x7c\x27\xb1\x0f\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/jobsdb"].(os.FileInfo),
		fs["/node"].(os.FileInfo),
		fs["/warehouse"].(os.FileInfo),
	}
	fs["/jobsdb"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/jobsdb/000001_create_tables.down.tmpl"].(os.FileInfo),
		fs["/jobsdb/000001_create_tables.up.tmpl"].(os.FileInfo),
		fs["/jobsdb/000002_alter_dataset_tables.down.tmpl"].(os.FileInfo),
		fs["/jobsdb/000002_alter_dataset_tables.up.tmpl"].(os.FileInfo),
	}
	fs["/node"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/node/000001_create_event_schema.down.sql"].(os.FileInfo),
		fs["/node/000001_create_event_schema.up.sql"].(os.FileInfo),
		fs["/node/000002_create_col_counters_event_schema.up.sql"].(os.FileInfo),
		fs["/node/000002_del_col_counters_event_schema.down.sql"].(os.FileInfo),
	}
	fs["/warehouse"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/warehouse/000001_create_tables.up.sql"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
