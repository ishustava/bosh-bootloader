// Code generated by go-bindata.
// sources:
// templates/bosh_director.tf
// templates/cf_dns.tf
// templates/cf_lb.tf
// templates/concourse_lb.tf
// templates/jumpbox.tf
// templates/vars.tf
// DO NOT EDIT!

package gcp

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

var _templatesBosh_directorTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x95\xd1\x72\xa2\x30\x14\x86\xaf\xe5\x29\xce\x64\xf6\xb2\xba\x96\x05\xf5\xa6\x4f\xb2\xe3\x64\x02\x1e\x59\xb6\x91\xc3\x84\xa0\x9d\xe9\xf8\xee\x3b\x21\x04\xb0\x2a\x1b\x6d\x9d\x4e\x2f\xa0\xf8\xff\x7f\x4e\xbe\x9c\x03\x54\xeb\xb2\xd6\xc0\x0a\xd4\x07\x52\xaf\xbc\x10\x3b\x64\xf0\x1e\x00\x00\xec\x85\xac\x11\x5e\x80\xfd\x78\xcf\x88\x32\x89\x3c\xa5\x5d\x59\x6b\xe4\xad\x7a\x96\x24\x72\xea\xee\x8d\xf3\xc8\x82\x63\x10\xb8\xcc\xaa\x4e\x6e\x8b\xed\x0d\x4d\xb2\xfd\xf7\x42\x70\x42\xd5\x1f\x4e\x25\x16\x5c\x8b\xcc\x33\x7b\x9b\x2b\x3c\x08\x29\x67\xc6\x3c\x35\xe6\x6b\xc1\x9b\x5c\x61\xaa\x49\x9d\x84\x4f\x7c\x93\x9d\xfb\x42\xfa\xdf\x7a\x57\x26\xf4\x76\x35\x77\x2f\xd4\x0c\x8b\x3d\xcf\x37\xc7\x69\xab\x3d\xf1\xe7\x85\x46\x55\x08\x79\xcf\xae\x9d\x77\x50\x96\xc2\x8a\x6a\x95\x22\xb0\xcb\xa7\xcb\x80\x0d\xce\xd7\xae\x65\xdc\x93\xc9\x79\xb9\x4e\x14\x00\x88\x5a\x13\x4f\x15\x8a\x93\x03\xad\xe0\x05\xb6\x42\x56\x38\xba\x72\xaf\x6f\x17\xb7\x0f\x86\x6b\x4f\xce\xd6\x6e\x35\x01\x40\x5e\xf2\x34\xdf\x28\xae\x44\x91\x35\x3c\x9e\xe7\xb3\xe6\xef\xe7\xf3\xc2\xfc\xde\x66\xb7\x19\x1e\x3d\x5d\xa1\xdc\x72\x99\x17\xaf\xff\x21\xe6\x30\x33\x60\xf8\x66\x41\xf7\x35\x9b\x13\x3a\x2b\xba\x93\xf5\x65\xdd\x3a\x69\x01\x80\x2d\xc7\xee\xd7\x10\xfe\xcd\xdc\x86\xe7\x6c\x6d\x04\x42\x4a\x3a\xb4\x5d\x52\x92\xd2\x56\x14\x86\xec\x09\xd8\x62\xb5\x58\x99\x6b\x18\xc7\x71\xcc\xd6\x56\xa3\x48\x53\x4a\xd2\xd4\xa2\xd3\xd2\x54\x77\x34\x39\x5a\xa8\x0c\xb5\x69\x3c\x9b\x70\xba\x99\x6e\xa4\xd8\xda\x17\x53\x6f\x19\xe7\xd4\xeb\xbe\x02\x94\x47\xfd\x9e\xd0\x56\x51\xf4\xab\xb9\xae\xa2\xe8\x0b\x21\xba\xb7\xc7\x8d\x20\x3b\x9b\x07\xcc\x4e\xfb\x68\xa0\x83\xbd\x7c\x84\x7a\x17\x20\xf7\x0a\xf3\x67\xe3\x1c\x53\x4d\xbe\x88\x2e\x5a\x1e\x48\x6a\xb0\xa9\xab\x9d\x17\x85\xb6\xf7\xc2\x38\x8c\xe7\xf6\x66\xb9\x5c\x7e\x47\xb3\xb5\x5f\x25\x03\xa7\x79\x30\x8a\xf2\x83\xf8\x81\x10\xdd\xc7\x72\x7c\x7a\x3f\xc3\xab\x3b\xa6\xa7\xf1\x99\xba\xb9\x35\x3d\xdb\xf1\x9b\x5a\x70\xc0\x2a\x4f\x77\x3d\x2c\x9f\x61\xbe\xa6\xa9\x37\x77\x0d\xfc\xbf\x00\x00\x00\xff\xff\x79\x21\x66\x2d\xa9\x0a\x00\x00")

func templatesBosh_directorTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesBosh_directorTf,
		"templates/bosh_director.tf",
	)
}

func templatesBosh_directorTf() (*asset, error) {
	bytes, err := templatesBosh_directorTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/bosh_director.tf", size: 2729, mode: os.FileMode(420), modTime: time.Unix(1507150150, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCf_dnsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x96\x41\x8f\x9b\x30\x10\x85\xef\xfc\x8a\x91\xd5\x53\xa5\xa0\x95\x7a\xce\xa1\x52\xcf\xbd\xf4\x58\x45\xc8\xb1\x27\x04\xc9\xd8\xd6\xcc\x40\x9a\xae\xf8\xef\x95\x09\x10\xd2\x2e\x6a\x38\x64\x95\xc3\x72\x01\xcc\xf3\xbc\x37\x9f\x11\xa6\xd5\x54\xe9\xbd\x43\x50\x7c\x66\xc1\xba\xb0\xa1\xd6\x95\x57\xf0\x9a\x01\xc8\x39\x22\x6c\x41\xb1\x50\xe5\x4b\x95\x75\x59\x46\xc8\xa1\x21\x83\xa0\xca\x10\x4a\x87\x85\xf5\x5c\xd4\xda\xeb\x12\x6d\xf1\x3b\x78\x54\xa0\xd0\xb7\xfd\xf0\xe5\x36\x15\xf2\xba\x46\x18\x8e\x2d\xa8\x4f\xaf\xad\xa6\x3c\xc9\x2a\xdb\x6d\x7a\x59\x06\x90\xa6\x8c\xc2\x49\x74\x93\xaa\xcb\x7b\x1d\xb2\xa1\x2a\x4a\x15\x7c\xd2\x7d\xfb\xfe\x03\x52\x09\x38\x04\x02\x39\x22\xdc\x54\x07\xf4\x6d\x45\xc1\xd7\xe8\xa5\x6f\x20\x34\x12\x1b\xf9\xab\xdd\x3e\x2e\x23\xb5\x48\x7c\x49\xdc\x6a\xd7\xe0\x25\xc6\x42\xa3\xf9\xbc\xcd\x3c\x05\x1f\x2b\x74\xcb\xa4\x08\x4d\x20\x5b\x30\x8a\x02\x75\xaa\x9c\x35\x9a\xec\xc6\x7a\xfe\x87\xd3\x16\xd4\xe7\xfc\x4e\xf3\x91\x5c\x77\xc1\x13\xd1\x5b\x2e\x7a\x3a\x3f\x47\x73\x13\xea\xd8\x08\x16\xa5\x0b\x7b\xed\x0a\x6d\x2d\x21\x73\x6e\x0e\x9b\xe1\x52\xed\xc6\x05\x9f\xfc\xbf\xa6\x72\x22\xee\xba\x72\x5f\x5e\x5e\xb2\x0c\x60\x9e\x64\x25\xa3\x4e\xa5\x02\x44\x56\x8b\xe6\x3e\xe0\x34\xf9\xbf\x11\xf3\xe1\xdc\xa9\xdd\x7d\x80\xcd\x61\xc3\x7c\xdc\x44\x0a\xbf\xce\x6f\x01\x66\x3e\x3e\x00\xf1\x2c\xf8\xd5\xfd\x69\xe8\xbe\x95\x6e\x35\x58\x31\x71\xe9\xa5\x15\x13\x1f\xcb\x34\x79\x53\x68\x04\xe9\x29\xa1\x5e\xe3\xad\xa6\x6a\x43\x8c\x0e\x69\x89\xec\xf0\xf8\xb1\x74\x4f\x4f\xf4\x21\xb8\x89\xb5\x9a\xa6\x0b\x65\x49\x58\x6a\x09\x8b\x44\x67\x92\x0f\xaa\x2b\xf7\xac\x13\x2f\x6f\x5b\x27\xfe\xc0\x79\xe7\x0e\x45\x68\x8f\xcd\x7e\x86\x71\x3b\x0d\x3e\x92\xe1\x60\xbb\x9b\xfd\xe4\x4d\xfc\xde\x1b\xdc\xd0\xed\x0d\xba\x3f\x01\x00\x00\xff\xff\xea\xd9\xbe\x05\x97\x0a\x00\x00")

func templatesCf_dnsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_dnsTf,
		"templates/cf_dns.tf",
	)
}

func templatesCf_dnsTf() (*asset, error) {
	bytes, err := templatesCf_dnsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_dns.tf", size: 2711, mode: os.FileMode(420), modTime: time.Unix(1507151140, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCf_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\x4b\x8f\xdb\x36\x17\x5d\x7f\xfe\x15\x84\xf0\x2d\x5a\x20\xf2\xd8\x1e\x37\x75\x17\x59\x15\xdd\xa6\x5d\x74\x17\x0c\x04\x8a\xba\xb2\x09\x73\x44\x96\xa4\xec\x18\x83\xf9\xef\x05\x49\xc9\xd6\x93\x7a\xd8\x41\x31\xc9\xc2\x1a\xf1\xde\x73\xa9\x73\x1f\x3a\x92\x4e\x58\x52\x1c\x33\x40\x81\x52\x2c\x22\x20\x35\x4d\x29\xc1\x1a\x02\xf4\xb6\x40\x48\x5f\x04\xa0\x2f\x28\x50\x5a\xd2\x6c\x1f\x2c\xde\x17\x8b\x5e\x8f\x48\x48\x7a\x32\xbf\x47\xb8\xf4\x7a\xf3\x5c\x8b\x5c\xa3\x40\xf2\x5c\x83\x8c\x62\x4c\x8e\x90\x25\x91\x02\x79\xa2\xa4\x08\x7a\xc2\x2c\xb7\x7e\xff\x7f\xdb\x73\xbe\x67\x10\x11\xfe\x2a\x72\x0d\x4d\xf3\xa5\x43\x09\x59\x1c\x16\x2b\x61\xb9\x92\xe1\x57\x78\xef\x8a\xc8\xe2\x88\x0a\x17\xc7\x17\x69\xcf\x78\x8c\x59\x84\x93\x44\x82\x52\x4b\x92\x86\xe5\x61\xf1\x5b\x07\x57\xea\x10\x09\xc9\xbf\x5f\xc6\xe2\x57\x80\x95\x3a\x84\xd6\xb7\x1b\x5a\x13\x11\x4d\xdb\x7b\x05\x5b\x13\x11\x3a\xe7\x6e\xf0\xb3\x9a\x01\x7a\xee\x21\x81\x48\x48\x0e\x79\x3c\x19\xd1\xb9\xd5\x31\x25\x28\x9e\x4b\x02\x28\x68\x78\xa5\x54\xc2\x19\x33\x16\xa0\xa0\x3c\x0c\x49\xea\xa2\x99\xa4\x23\xf7\xcf\x06\x3c\x61\xb9\x84\xec\x14\xd1\xe4\x3d\x24\x69\xc8\x05\x64\xc1\x02\xa1\x04\x04\x64\x89\x8a\x78\x86\xbe\xa0\x6f\xcd\x00\x19\xe8\x33\x97\xc7\x65\x1c\xb3\xb0\x38\x0e\x5e\x0c\xb8\x3b\xbe\x82\x0f\xbb\x95\x45\xb8\x40\x08\x33\xc6\xcf\x05\x23\x42\x72\xcd\x09\x67\x06\x46\x13\x11\xb8\x93\x5c\x6a\xe5\xb0\xbf\x05\xbb\x55\xf0\x09\x05\xdb\xed\xb3\x0d\xfc\x6e\x00\x1c\x1b\x91\xc4\xd9\x1e\x94\x35\x5a\x2d\xed\xff\xa7\x55\xf0\x62\x0c\x34\x96\x7b\xd0\x91\xc6\x7b\xb7\x7c\x77\xef\xbc\x78\xd3\x50\xef\x8f\x00\x05\xb7\x0e\xa9\xe4\xa2\x23\x0b\xfe\xec\x16\xb0\x29\x97\x67\x2c\x13\x9a\xed\x23\x99\x33\x70\xf0\x07\xad\x45\x78\x5b\x09\xdd\xca\x88\xbc\x1b\x47\xc3\x32\x15\xe5\x7e\x67\xb7\x7c\xc9\x33\xea\x2b\x83\x22\x0d\x26\xa4\x1b\x08\xcb\x72\xe7\x2c\x2e\xba\x5c\x01\x4b\x23\x46\xb3\xa3\xc5\x33\x89\x77\x69\x35\x78\xbb\xd5\x7d\xfc\xa8\xd9\x04\xa9\xff\x80\x21\x55\xa7\x48\x8d\xe3\xc8\xf4\x85\x97\xa4\x56\x0e\x2a\xf5\x53\x46\x68\xf1\xd2\x26\xc6\xda\x3b\x63\x3b\x34\x14\x91\x54\x68\x6a\xa7\x46\x20\x01\x33\x76\x41\x18\x31\x8e\x13\x14\x63\x86\x33\x02\x12\xc5\xb9\x46\x8c\x2a\x0d\x09\xc2\x0a\xe1\x0c\x19\x10\x74\x05\xc9\x25\x8b\x5e\xb1\xe8\xe5\xa6\x58\xaf\x11\x92\x4b\x16\x9a\x73\x55\x4a\x46\x5e\xbd\x6a\x5e\xbe\xf2\x5c\x7f\x3f\x09\xaa\x9b\x85\xd2\x61\x0a\x15\xaa\x9b\x8b\xbb\x09\x41\xa8\x21\x46\x7a\x86\x60\xc3\xca\xe0\x9a\x3f\xab\x58\xfe\xb9\xd7\x52\x49\x41\x01\x71\x23\x34\x12\x12\x52\xfa\xbd\xc5\x65\x47\x15\xe5\x0a\xa4\x61\xe4\x44\x13\x48\xcc\x25\xa0\x42\x43\xa1\x23\x5c\xd0\x93\x3d\x53\x89\x86\x04\xa6\xd2\x36\xc4\x4d\x69\xb9\x30\x29\x65\xf0\x93\x89\xe5\xd1\x64\x3f\xdb\x1d\x54\xe1\xbc\xae\xce\x9c\xd1\x14\xc8\x85\x30\x40\x6f\x8b\xff\x11\x09\x06\x2b\x86\x94\x4b\x88\x12\x50\x5a\x72\xb3\x01\x2d\x73\xb0\x37\x2a\x1f\x73\x45\x2a\x1b\xc5\x58\x24\xd3\x7f\xcf\x28\x26\xb8\xe5\x2f\xc5\x39\xd3\xe5\x4d\xec\x4e\x91\x38\xb6\xa5\x0e\x80\x99\x3e\x44\xe4\x00\xe4\xe8\xf6\x2f\xf2\x98\x51\x12\xba\x85\xb0\x58\xe8\xec\xa8\x9e\x91\xeb\x00\xec\x35\xd9\x39\x55\x0d\x61\xa8\x76\x43\xaf\x8d\xb4\x5b\xed\x56\x66\x55\xc2\x3f\x39\x28\x1d\x09\xac\x0f\x95\x38\x4f\x0e\x27\x18\xcc\x46\x2b\xe8\x63\xae\xab\x9c\xd6\x3d\x1b\x1f\xde\xf7\x48\xe9\x67\x6a\xc2\xb7\xc7\xce\x22\xaa\x3a\x7c\x0c\x19\xe8\x84\xe0\x6e\xe5\xd3\x81\xeb\xe7\xd5\x72\xb3\x5e\x5b\x2d\xb8\xd9\x18\xfb\xe7\x5f\x96\xeb\xdf\xdc\x89\xf5\x67\xeb\x5a\x15\x87\xe8\x81\xf2\xb0\xfd\xf8\x53\x44\x12\x9c\xb3\xa1\x87\xb9\x8a\x69\xfd\x31\xe8\xf6\xec\xd6\x5b\x0a\x35\xdd\x79\xf5\x1c\x98\x22\x37\xbb\x09\x65\xd6\x05\xde\x5f\x63\x57\xeb\x8f\xf3\xb0\xb1\xd9\x6c\x36\xb7\xfa\x1a\x7c\x8c\x18\xc8\x9a\xff\xee\x59\xab\x8e\x99\xa9\x33\x4d\x00\x4a\x51\x9e\x45\x38\x4d\x69\x46\xb5\xbd\x07\x7e\xfd\xf3\xeb\x1f\x03\x79\xed\x12\xcd\xfd\xe9\x1d\xda\x47\x4d\xe8\x4e\x2b\xf0\x5e\x75\x6b\x60\x6c\x3e\x9c\x16\xaf\x26\xef\xef\xdf\xff\x6a\x28\xf4\xc7\xbd\x57\x98\xdf\xb4\x95\xf7\x0b\x23\xba\xb6\xde\x59\x37\xdf\x51\xad\x55\x31\xff\x08\x6d\xb5\x5e\x6d\xb6\xe1\xf3\xe6\xd7\xcf\xbb\xf9\xcd\xd5\x62\xd7\xdf\x5d\xb5\xa1\xd8\xc9\xee\x10\xaf\x33\x14\x83\x27\x8b\x5e\x01\x54\x4f\x67\x9f\x66\xb8\x53\x31\xb4\xe6\xcd\x2c\x56\xbc\x13\xc7\x08\xb8\x0a\x29\x36\xb1\xb6\x1a\xda\xd9\x6d\x31\xd8\x99\xe3\x4f\x0b\x84\xfc\x79\xee\x1c\x64\xbe\x3c\x0c\xf3\x3f\x71\x94\x55\x36\xed\x9d\x65\x95\x26\x78\xc4\x44\x1b\xf1\x36\x73\xfe\x28\x3b\xab\xc9\xc2\xe3\x3c\xf0\xae\xcb\x18\x4c\xab\xcf\x51\x88\x93\xeb\x71\x64\x29\x76\x88\xff\x51\x73\xa7\xb3\x1e\xcf\xaa\x78\xad\x34\xaa\x1a\xaf\xd6\xd3\x6b\xf1\xac\xfc\x35\x68\x5f\x17\x3d\xa0\xf8\x9a\x6f\xbd\x67\xd1\x31\x89\x8d\x1f\x40\xc6\x6e\xf5\x43\xb8\x68\x7e\x01\x98\xdb\x85\xc5\x97\x80\xf6\xc7\x9b\x06\xb0\xb9\x83\x0e\x01\x97\x62\xe3\x8a\x5a\xf1\x9d\xa0\x53\x9c\x73\x3d\x6f\xed\x9c\x39\xab\xab\xe8\x2f\x75\xc6\x44\x91\x31\xf9\x01\x71\xb7\xdd\x16\xea\x62\xaa\xb8\xa8\x31\x3d\x5a\x56\xb4\xc8\xe8\x63\x62\xd2\xd8\x9b\x80\xfa\x40\xed\xdf\x95\x58\x5f\x72\x27\x77\x63\x41\xb1\xbf\x1f\x4d\x06\xef\xef\xc8\x8e\x2f\x68\xff\x06\x00\x00\xff\xff\x43\xa8\x44\x67\xd9\x1d\x00\x00")

func templatesCf_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_lbTf,
		"templates/cf_lb.tf",
	)
}

func templatesCf_lbTf() (*asset, error) {
	bytes, err := templatesCf_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_lb.tf", size: 7641, mode: os.FileMode(420), modTime: time.Unix(1507150150, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConcourse_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x93\x3f\x8f\xdc\x20\x10\xc5\xeb\xf0\x29\x46\xa3\x94\xc1\x85\x73\x6d\xaa\x28\xed\x25\x45\xba\xe8\x84\x58\x7b\xec\x45\xc7\x31\x08\xf0\x5a\xd1\xc9\xdf\x3d\xc2\x66\xbd\xbe\xdc\x9f\x5d\x69\xb5\x95\x47\xe3\xc7\x63\xf8\x3d\xe0\x21\xf9\x21\x01\x36\xec\x1a\x1e\x42\x24\x95\x74\xe8\x29\x29\xcf\x6c\x11\x9e\xc5\xa7\x83\xb6\x03\xc1\x37\xc0\xcf\xcf\x3d\x73\x6f\x49\x35\xfc\xe4\x87\xf4\x42\x59\x2d\xb5\x9c\x6b\xa7\x9f\x68\x42\x31\x09\xf1\xda\xdd\xee\x94\xf1\xd9\x17\x00\xe0\x7d\x6b\xdd\xb6\x81\x62\xac\xd6\x85\xf2\xd8\x29\xdf\xc5\x3f\x50\xe4\x21\x34\x04\xf8\xdf\xfa\xce\x04\x1a\xb5\xb5\x08\x78\x2c\xe5\xea\xb5\x6c\x9f\xa7\xcc\x43\xcc\xdb\x1f\x74\xa8\xc8\x1d\x94\x69\xa7\x93\x4e\xb2\x27\x87\x59\x4a\x69\xe4\xf0\xf8\xe6\xa4\xe5\x5f\xb5\xdb\x59\x79\xac\x0b\x00\x01\xa0\xad\xe5\xb1\x9c\xd6\x07\x4e\xdc\xb0\xcd\x36\xa9\xf1\xb8\x34\x39\xa4\xb8\x8c\xf1\x07\xef\xee\xbe\xe2\x17\xc0\xba\xae\x6b\x7c\x10\x00\x53\xb6\x28\x94\x93\xee\xe3\x2c\x3a\x1d\xe3\xe1\x43\x04\x05\x14\x6e\xe8\xcb\xb5\xb7\x02\x78\xff\xf4\x1f\x03\x7e\x71\x4b\x70\x93\xfe\x85\xde\x02\x20\x52\x8c\x86\x9d\xd2\x5d\x67\x9c\x49\x7f\xb3\xfe\xfe\xe7\xfd\x8f\x33\xc9\x72\x18\x75\x68\x8d\xeb\x55\x18\x2c\x21\x60\x8c\x7b\x79\xea\xca\xa5\xbb\x4d\xf8\x4c\xca\x31\xee\x71\xe5\xbc\x51\x5f\x78\xdb\x23\xd9\x4e\x59\xe3\x1e\xa7\xec\x92\xf3\x54\x41\xbb\x9e\x66\x97\x39\x4a\x01\x60\xbc\xda\xc6\xff\xfb\xfb\xaf\xd2\x2d\x89\xbc\xbd\xe5\xd5\xaf\xe0\x15\xab\x7d\x4a\x3e\x5e\x45\x6b\x76\xb8\x19\xaf\xfc\x02\x6e\x8c\xeb\x5f\x00\x00\x00\xff\xff\x5c\xc2\x73\xd4\xf2\x04\x00\x00")

func templatesConcourse_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesConcourse_lbTf,
		"templates/concourse_lb.tf",
	)
}

func templatesConcourse_lbTf() (*asset, error) {
	bytes, err := templatesConcourse_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/concourse_lb.tf", size: 1266, mode: os.FileMode(420), modTime: time.Unix(1507150150, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJumpboxTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xcf\xc1\x8a\x83\x30\x10\x06\xe0\xf3\xe6\x29\x86\xb0\x57\x15\x84\x5c\x84\x7d\x96\x90\x35\x83\xb5\x44\x13\x26\x33\x22\x88\xef\x5e\x4a\xad\x56\xe8\xa5\xbd\x86\xfc\xdf\xff\x0f\x61\x8e\x42\x2d\x82\xee\x62\xec\x02\xda\x36\x0e\x49\x18\xad\xf3\x9e\x30\x67\x0d\xfa\x2a\x43\xfa\x8f\x73\xd1\x27\x0d\x8b\x02\x18\xdd\x80\xf0\x07\xfa\x77\x99\x1c\x95\x38\x4e\xb6\xf7\x6b\xf1\xf2\x4b\xad\x4a\x45\xe1\x24\xbc\x87\xad\x50\x78\xa4\x01\x26\x17\x64\x03\xde\x77\x96\x87\x55\x6e\x4f\x6b\x53\xd7\x27\x17\x67\x46\x1a\x5d\xb0\xcf\x55\x5f\xba\x27\xd4\xf7\x84\x2d\x47\x3a\x8e\x5f\xd4\xcf\xce\x5e\x98\x53\x6e\xaa\xea\xb3\xd9\xc6\x18\x73\x2f\xb9\x05\x00\x00\xff\xff\xb2\xf6\x55\xa8\x69\x01\x00\x00")

func templatesJumpboxTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesJumpboxTf,
		"templates/jumpbox.tf",
	)
}

func templatesJumpboxTf() (*asset, error) {
	bytes, err := templatesJumpboxTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jumpbox.tf", size: 361, mode: os.FileMode(420), modTime: time.Unix(1507150150, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\x41\x0a\xc3\x20\x10\x05\xd0\x75\x3c\x85\x0c\x59\xb4\x9b\xde\xa0\x67\x29\x36\x99\xca\x14\x99\x91\x89\x08\xad\x78\xf7\x22\x16\xcc\xa6\xa4\x4b\xf9\x4f\x3e\xf3\xb3\x53\x72\xf7\x80\x16\xa2\xca\x13\x97\x74\xa3\x15\x6c\x31\x53\x7a\x45\xb4\x57\x0b\x5b\x52\x62\x0f\xa6\x1a\x33\xac\xa2\x27\xe1\x63\xf7\x16\xc6\x63\x85\x9c\xff\x6a\x5d\x14\x57\xe4\x44\x2e\x6c\xbf\x70\x54\xc9\xb4\xa2\x5a\xf0\x22\x3e\xf4\xf2\xdd\xbf\xc6\xe7\xf2\xa0\x80\x27\x98\x4b\x76\x7a\xd9\x85\x15\xce\x15\xcc\xf4\x1d\xa2\xd3\x46\xc6\x32\x2d\xee\xb7\x8f\xb4\xbf\x6b\xab\xff\x04\x00\x00\xff\xff\x39\xda\x2a\x22\x4d\x01\x00\x00")

func templatesVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVarsTf,
		"templates/vars.tf",
	)
}

func templatesVarsTf() (*asset, error) {
	bytes, err := templatesVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vars.tf", size: 333, mode: os.FileMode(420), modTime: time.Unix(1507150150, 0)}
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
	"templates/bosh_director.tf": templatesBosh_directorTf,
	"templates/cf_dns.tf": templatesCf_dnsTf,
	"templates/cf_lb.tf": templatesCf_lbTf,
	"templates/concourse_lb.tf": templatesConcourse_lbTf,
	"templates/jumpbox.tf": templatesJumpboxTf,
	"templates/vars.tf": templatesVarsTf,
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
	"templates": &bintree{nil, map[string]*bintree{
		"bosh_director.tf": &bintree{templatesBosh_directorTf, map[string]*bintree{}},
		"cf_dns.tf": &bintree{templatesCf_dnsTf, map[string]*bintree{}},
		"cf_lb.tf": &bintree{templatesCf_lbTf, map[string]*bintree{}},
		"concourse_lb.tf": &bintree{templatesConcourse_lbTf, map[string]*bintree{}},
		"jumpbox.tf": &bintree{templatesJumpboxTf, map[string]*bintree{}},
		"vars.tf": &bintree{templatesVarsTf, map[string]*bintree{}},
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

