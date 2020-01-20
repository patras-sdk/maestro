// Code generated for package crd by go-bindata DO NOT EDIT. (@generated)
// sources:
// config/crds/kudo.dev_instances.yaml
// config/crds/kudo.dev_operators.yaml
// config/crds/kudo.dev_operatorversions.yaml
package crd

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configCrdsKudoDev_instancesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5a\x5b\x6f\x1b\xb9\x15\x7e\xf7\xaf\x38\xf0\x3e\x38\xc6\x5a\xa3\xc6\xed\x43\x21\x14\x05\x52\xdb\x29\xd4\xa6\xb6\x10\x3b\x29\x16\xd9\x14\xa0\xc8\xa3\x19\xd6\x1c\x72\xc2\x8b\x64\x75\xb3\xff\xbd\x38\xe4\xcc\xe8\x36\xa3\x4b\xd0\x6c\xf9\x62\x8b\x22\xcf\xe5\x3b\xf7\x19\x9d\x0d\x06\x83\x33\x56\xc9\x8f\x68\x9d\x34\x7a\x04\xac\x92\xf8\xe2\x51\xd3\x27\x97\x3d\xff\xd1\x65\xd2\x0c\xe7\xaf\xa7\xe8\xd9\xeb\xb3\x67\xa9\xc5\x08\x6e\x82\xf3\xa6\x7c\x8f\xce\x04\xcb\xf1\x16\x67\x52\x4b\x2f\x8d\x3e\x2b\xd1\x33\xc1\x3c\x1b\x9d\x01\x30\xad\x8d\x67\xb4\xed\xe8\x23\x00\x37\xda\x5b\xa3\x14\xda\x41\x8e\x3a\x7b\x0e\x53\x9c\x06\xa9\x04\xda\xc8\xa1\xe1\x3f\xff\x5d\x76\x9d\xfd\xe1\x0c\x80\x5b\x8c\xd7\x9f\x64\x89\xce\xb3\xb2\x1a\x81\x0e\x4a\x9d\x01\x68\x56\xe2\x08\xa4\x76\x9e\x69\x8e\x2e\x7b\x0e\xc2\x64\x02\xe7\x67\xae\x42\x4e\xcc\x72\x6b\x42\x35\x82\x76\x3f\x5d\xa9\xe5\x48\x3a\x8c\xeb\xdb\x71\x4b\x49\xe7\xff\xbe\xb1\xfd\x4e\x3a\x1f\xbf\xaa\x54\xb0\x4c\xad\x71\x8b\xbb\x4e\xea\x3c\x28\x66\x57\xfb\x67\x00\x8e\x9b\x0a\x47\x70\x4f\xac\x2a\xc6\x51\xd0\x5e\x98\xda\x1a\xa7\x9a\xbd\xf3\xcc\x07\x37\x82\x5f\x7e\x3d\x03\x98\x33\x25\x45\xd4\x32\x7d\x69\x2a\xd4\x6f\x26\xe3\x8f\xbf\x7f\xe4\x05\x96\x2c\x6d\x02\x08\x74\xdc\xca\x2a\x9e\x6b\x45\x04\xe9\xc0\x17\x08\xe9\x28\xcc\x8c\x8d\x1f\x5b\x41\xe1\xcd\x64\x9c\xd5\x04\x2a\x6b\x2a\xb4\x5e\x36\x42\xd0\x5a\x33\x7a\xbb\xb7\xc5\xea\x82\x64\x49\x67\x40\x90\x99\x31\xb1\xac\x8d\x85\x02\x5c\x62\x6e\x66\xe0\x0b\xe9\xc0\x62\x65\xd1\xa1\x4e\x86\x5f\x23\x0b\x74\x84\x69\x30\xd3\x7f\x23\xf7\x19\x3c\xa2\x25\x22\xe0\x0a\x13\x94\x20\xdf\x98\xa3\xf5\x60\x91\x9b\x5c\xcb\xff\xb4\x94\x1d\x78\x13\x59\x2a\xe6\xb1\x36\x49\xb3\xa4\xf6\x68\x35\x53\x84\x62\xc0\x2b\x60\x5a\x40\xc9\x96\x60\x91\x78\x40\xd0\x6b\xd4\xe2\x11\x97\xc1\x3f\x8c\x25\x88\x66\x66\x04\x85\xf7\x95\x1b\x0d\x87\xb9\xf4\x8d\x9b\x73\x53\x96\x41\x4b\xbf\x1c\x46\x67\x95\xd3\xe0\x8d\x75\x43\x81\x73\x54\x43\x27\xf3\x01\xb3\xbc\x90\x1e\xb9\x0f\x16\x87\xac\x92\x83\x28\xb8\x8e\x5e\x9e\x95\xe2\x87\xd6\xd6\x17\x6b\x92\xfa\x25\xb9\x85\xf3\x56\xea\xbc\xdd\x8e\x5e\xd8\x8b\x3b\x39\x23\xd9\x97\xd5\xd7\x92\xfc\x2b\x78\x69\x8b\x50\x79\x7f\xf7\xf8\x04\x0d\xd3\x68\x82\x4d\xcc\x23\xda\xab\x6b\x6e\x05\x3c\x01\x25\xf5\x0c\x6d\x32\xdc\xcc\x9a\x32\x52\x44\x2d\x2a\x23\xb5\x8f\x1f\xb8\x92\xa8\x37\x41\x77\x61\x5a\x4a\x4f\x96\xfe\x12\xd0\x79\xb2\x4f\x06\x37\x31\xd8\x61\x8a\x10\x2a\xc1\x3c\x8a\x0c\xc6\x1a\x6e\x58\x89\xea\x86\x39\xfc\xee\xb0\x13\xc2\x6e\x40\x90\x1e\x06\x7e\x3d\x47\x6d\x1e\x4c\x68\xb5\xdb\x4d\x32\xe9\xb4\x50\x13\x84\x8f\x15\xf2\x8d\xd0\x10\xe8\xa4\x25\xf7\xf5\xcc\x23\x39\x7d\x73\x32\x5b\x23\xd5\x15\x8e\x75\xf8\x5b\xe6\x8d\xed\x88\xcb\x1d\x09\x1e\x36\xcf\x46\x71\xe5\x4c\x22\x39\x8d\xc5\x19\x5a\xa4\x1c\xe1\x0d\xf9\x50\xfa\x8a\x6f\xdf\xd9\x22\xdf\xf8\x4b\xb6\xb5\xdf\x27\x2d\xf4\x26\x91\x4e\x81\xdf\x4c\xc6\x4d\xe2\x48\xf9\x02\x1b\x39\x77\x38\xf6\x1a\xaf\x59\x33\x89\x4a\x4c\x98\x2f\x0e\x72\xbd\x18\xcf\x12\x9b\x18\x46\x11\x8e\x4a\x22\xc7\x8d\x7c\x14\x93\x26\x32\x91\x36\x3b\x48\x02\x90\xb7\x59\xac\xcf\x5f\xa5\xa0\xa9\x63\x73\x95\xc3\x3c\x93\x1a\x58\xca\xea\xf0\xb7\xc7\x87\xfb\xe1\x5f\x4d\x92\xb5\x93\x26\xe3\x1c\x9d\x4b\xae\x52\xa2\xf6\x57\xe0\x02\x2f\x80\xb9\xc6\x8b\x1e\xe9\x9b\xac\x64\x5a\xce\xd0\xf9\xac\xe6\x80\xd6\x7d\xba\xfe\xdc\x85\x19\xc0\x5b\x63\x01\x5f\x58\x59\x29\xbc\x02\x99\x50\x6e\xb3\x40\xe3\x14\xd2\x25\x20\x5a\x7a\xb0\x90\xbe\x90\xdd\x8a\x33\xa8\x8c\xa8\x15\x5e\x44\x45\x3d\x7b\x46\x30\xb5\xa2\x01\x41\xc9\x67\x1c\xc1\x39\x79\xd9\x9a\x88\xbf\x50\xc9\xfd\xf5\xbc\x93\xe6\xab\x45\x81\x16\xe1\x9c\x8e\x9c\x27\xc1\xda\x44\x4f\x7b\x8d\x7f\xac\x04\xf4\x05\xf3\xe0\xad\xcc\x73\xb4\xd8\x8d\x66\xcc\x5e\x94\x15\x2e\xc1\x58\xd2\x5d\x9b\x35\x02\x91\x2c\xd9\xac\x0e\x13\xb1\x23\xf0\xa7\xeb\xcf\x3d\xd2\x6e\xe2\x04\x52\x0b\x7c\x81\x6b\x90\x3a\xa1\x52\x19\x71\x99\xc1\x53\xf4\x88\xa5\xf6\xec\x85\xf8\xf0\xc2\x38\xd4\x60\xb4\x5a\x76\x4b\x6b\xa0\x60\x73\x04\x67\x4a\x84\x05\x2a\x35\x48\x59\x44\xc0\x82\x2d\x49\xff\xc6\x5c\xe4\x61\x0c\x2a\x66\xfd\x66\x09\xed\xa4\xfa\xf4\x70\xfb\x30\x4a\x52\x91\x0b\xe5\x9a\x44\xa1\xd4\x3c\x93\x54\x28\xa9\x42\xa6\x74\x4f\x3e\x19\xe1\x08\xc9\x39\xbc\x01\x5e\x30\x9d\x63\x27\xd9\xa8\x29\xc2\x2c\x50\x02\xce\x2e\x4e\x8d\xd6\xed\x5a\xd7\xac\x8e\x9a\xb7\x9d\x18\xfe\x4f\x95\xe3\x28\xb5\x62\x1b\x7a\x50\xad\xfb\x35\x7f\xde\xab\x16\x35\xc4\x56\xa3\xc7\xa8\x99\x30\xdc\x91\x52\x1c\x2b\xef\x86\x66\x8e\x76\x2e\x71\x31\x5c\x18\xfb\x2c\x75\x3e\x20\x47\x1c\x24\x4f\x70\xc3\xd8\xdc\x0e\x7f\x88\x7f\xbe\x49\x8b\xd8\xae\x1e\xa7\x4a\x3c\xfa\x5b\xe8\x43\x7c\xdc\xf0\x64\x75\x9a\x66\xe8\xd8\xaa\x74\xf1\xd8\x14\xc7\xad\x9b\x14\x12\x8b\x42\xf2\xa2\xe9\x6c\x57\xd9\xb3\x33\x46\x4a\x26\x52\xca\x65\x7a\xf9\xdd\xdd\x96\x80\x0c\x96\xe4\x59\x0e\xea\xb9\x6a\xc0\xb4\xa0\xff\x9d\x74\x9e\xf6\x4f\x46\x2e\xc8\x23\x82\xf4\xc3\xf8\xf6\xb7\x71\xe6\x20\x4f\x8e\xc8\xce\x2e\x8e\x56\xc5\x2c\x2b\xd1\xa3\xdd\x69\x60\x98\x10\x71\x72\x65\x6a\xb2\xa7\xc9\xf9\x26\x9e\x8a\xe9\xbb\x17\xe4\xc1\x1f\x6a\xe4\x2e\x9e\x62\x31\x64\x16\xc1\x2f\x0c\xa5\x7f\x6a\xe1\xe8\x3e\x60\x43\x00\x38\xd3\xd4\x5e\xb7\x15\x70\x04\xf0\xfa\x72\x47\x50\xa9\x85\xb4\xc8\xbd\x5a\x82\x2f\xac\x09\x79\x51\x37\xe4\xb1\x74\x00\x37\xd6\xa2\xab\x8c\x16\x54\x54\x5a\x54\x9a\xf4\xbe\xde\xd3\x66\x93\x16\xb3\x1d\x2e\x25\xab\x00\xae\x2f\x61\x87\x97\x43\x1f\x27\x93\xda\x41\x36\xe9\xad\xe3\xd1\xdd\xcf\xc1\xc6\x99\x7a\xaf\xe9\xad\x23\x1e\xb1\x88\x53\x11\xeb\xd1\x64\x57\xd6\xd8\x37\x8c\x1f\xe1\xe6\xc3\xfb\xf7\x77\xf7\x4f\xef\x7e\xaa\x31\xa5\x11\xe5\x21\x36\xc8\x6b\x23\xf3\xda\x23\x0a\x78\x35\xbe\xb9\xa4\x02\x29\x8c\xde\xad\x89\xb1\x0d\x88\xfd\x46\x23\xed\xd5\x7a\x5d\x5d\x48\xa5\xc8\x5a\x5c\x21\xb3\xc4\xe9\x9f\x85\x54\xd8\x5a\x07\x5e\x75\x98\x6e\x72\xe7\xa2\x91\x99\x72\x86\xee\xe2\x0b\xcd\x6c\x0e\xc5\x0a\xe6\x57\xd7\x97\x1b\x6c\x72\x39\x47\x07\xc1\x41\x49\xf1\x57\xcb\x7e\xb5\x43\x19\xb3\x3c\x03\x06\x93\xbb\xc6\x89\x66\xc6\x72\x14\x69\x7c\xb6\xa1\xf2\xa9\xc3\x48\xf9\x84\xf8\xd8\xa0\x75\x84\x95\x20\x7f\xa5\x8d\xdf\xf5\x33\xea\x2f\xa9\x63\x45\x01\x4b\xf4\x97\x3b\x2d\xc1\xbe\x91\x81\xc8\xde\xf7\x94\xce\x3d\x91\x46\x03\x27\xf5\xc4\xdb\xd7\x06\x2d\xc1\xe3\x42\xb3\x7b\xd6\x4b\x0f\x65\x0e\x4e\x7b\xf1\xd8\xc6\xbc\x67\xa6\x8e\x26\xea\x8e\x81\xef\x88\x79\x8f\xe5\xb9\xc5\x9c\x06\xe6\xc7\x1d\x01\x76\x84\x78\xb3\x75\x98\xbc\xb3\x49\xa5\x75\x6b\xd8\x3a\xb2\x6b\x04\xb5\x72\xde\xd1\x2e\xb7\xd3\x7e\x34\x71\x3a\x7c\xca\xd0\xc7\xbd\x9c\xe3\xe4\xdb\xec\xd8\x05\x76\xa7\xbe\x6d\x26\xa8\xd5\xe5\xac\xa2\x72\x98\x60\x6f\xd1\x8e\x95\xc8\x28\x65\xc2\xa9\x53\xe4\xde\xd4\xdd\x6d\x8f\xe3\xca\xc5\x86\x16\xe7\x93\x96\x1a\xac\x3f\x20\x8b\x83\x63\xda\x8e\xb9\x39\x5a\xe2\x67\x0d\x4f\x05\xba\xae\x26\x9c\xca\x43\x9a\x2b\xa3\xea\x2e\xf5\xf2\x96\x69\x17\x25\x72\x74\xb7\x7b\x7d\x85\x7b\x9c\xa3\xed\xa0\xd9\xa4\x42\xf8\xda\x73\x75\x45\x62\xff\x9a\x77\x10\x8f\x77\xee\xac\x35\x36\x7e\xfa\xd3\x20\xae\x3f\xc7\xed\x09\xa6\xc4\xbd\x4e\xfb\x5f\x7d\xbc\x3b\x68\xff\xbc\x5f\xac\x79\x8f\xd8\x3f\x26\x19\x06\xcd\xdf\xc1\x8f\xc7\xd3\xde\xbd\xdb\xc3\xe4\x08\xb9\xe7\xbd\x72\x7f\x85\xb7\xcc\x33\x05\x18\x71\x6b\x49\xc7\x7f\x6e\x0c\x65\x5d\xdf\xe5\x1c\x44\xf7\xeb\xee\xf0\xba\x2f\x86\x01\x14\x73\xfe\xad\xd4\xd2\x15\x28\xde\x87\xce\x4e\x19\xa8\x56\x94\xcc\x8f\x80\x9a\x88\x81\x97\x3b\x49\x36\x2d\x1d\x94\x62\x53\x85\x23\xf0\x36\x74\x1f\xd9\x9b\x11\x00\x4a\x74\x8e\xe5\x9d\xb9\xe4\xe0\xdd\xbe\x39\xec\xe0\xc5\xaa\x60\xae\x1b\x1b\x00\xe9\xb1\xec\xf9\x6a\x2b\xc2\x27\x44\xe5\x98\x08\xa7\x73\x3d\x04\xf7\x5b\x2a\xad\xbd\x10\x1d\xa5\x6f\x5a\xfd\x70\x9d\x40\xa4\x3f\x87\x37\xeb\x7f\x9f\xcb\x4f\x14\x10\xab\xbd\xf2\xed\x35\x70\x87\x0a\x8f\x1e\xab\x23\xac\x4c\x7c\xf7\x12\x3d\xc6\xd4\x69\x1d\x61\xf0\xb4\x8e\x02\x24\xad\x43\xc6\x3f\x99\xe0\x61\x47\x48\xeb\x74\x77\x38\x48\x12\x8e\x75\x98\x13\x95\xea\x6d\x0c\xba\x8e\x31\x6b\x59\xf7\x13\xbe\x23\x08\xed\x27\xb1\x0f\xda\xef\x11\x5d\x07\x00\xea\x79\x3e\xb0\x25\xcb\x87\xf1\x6d\x7a\x71\x45\xd4\xd2\xf8\x55\x18\x25\x1c\x04\x2d\xbf\x04\x84\xf1\x6d\xfd\x2e\xee\x0a\xa4\xe6\x2a\x88\x7e\x6b\x7c\xf8\x30\xbe\x75\x19\xc0\x5f\x90\xb3\xe0\x10\x16\x48\xd3\xd8\x85\x87\x87\xfb\x77\x3f\x01\xed\xc4\x13\xf5\x4c\x44\x4c\x69\x7c\x92\xf1\x9d\x61\x0f\xc9\xa4\x5c\xa4\x99\xa6\x9e\x28\x65\x8b\x17\x8d\x44\xda\xc7\xd6\xaa\x40\x55\x39\x28\xd9\x33\x82\x0b\x36\x69\xd2\x27\xe7\xf8\x36\xb5\x63\xf1\xd1\x05\x08\x13\x9f\xb8\xe6\xe8\x69\x26\x9b\xa9\xf8\x3e\xec\x74\xc0\xf7\x38\xcf\x06\xe0\x4e\x49\x8e\xf5\x93\x79\x9a\x1b\x75\x1c\xc5\xe3\x23\x85\x69\xf0\x04\x1a\x4f\x6f\xe7\x08\xb0\x74\x78\x9a\x00\xdd\x7d\xfd\x63\x05\x5a\x72\x15\x87\x5f\x42\x7a\xd6\xa5\x61\xc9\x4a\x15\x1f\x68\x1b\xed\xa4\xc0\xf8\x6e\x4b\xe6\x5a\xce\x24\x67\xda\xc3\x22\x8e\xb6\x91\x9d\xf4\x17\xbb\x03\xf8\xee\xf8\x78\xfc\x5c\xb6\xb5\xb5\xfa\x89\x40\xfd\x6b\x84\x76\x2b\x06\xc9\xa0\xfe\x5d\xc0\xea\x5b\x80\x34\x9b\xad\xb5\x25\xce\x1b\x4b\x39\x35\xed\xac\x22\x8c\x71\x8e\x95\x47\x71\xbf\xfd\x3b\x81\xf3\xd4\x56\x35\x3f\x03\x88\x1f\xb9\xd1\x69\x0e\x70\x23\xf8\xf4\xf9\x2c\x51\x45\xf1\xb1\x11\x86\x36\xff\x1b\x00\x00\xff\xff\x06\x63\xd2\x78\x58\x21\x00\x00")

func configCrdsKudoDev_instancesYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_instancesYaml,
		"config/crds/kudo.dev_instances.yaml",
	)
}

func configCrdsKudoDev_instancesYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_instancesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_instances.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configCrdsKudoDev_operatorsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4d\x8f\x1b\x37\x0c\xbd\xcf\xaf\x20\xd2\x43\x2e\xf5\xb8\x69\x7b\x28\x7c\x0b\xb6\x39\x2c\xda\xa4\x41\x36\xd8\x4b\xd1\x03\x67\x44\xdb\xec\xce\x48\x2a\x49\x19\xdd\xfc\xfa\x42\xd2\xf8\x63\xbc\xf6\x6e\x5b\x20\x73\xd3\x93\xc8\x47\x3d\x7e\x8c\x9a\xc5\x62\xd1\x60\xe4\x7b\x12\xe5\xe0\x57\x80\x91\xe9\x6f\x23\x9f\x57\xda\x3e\xfc\xa4\x2d\x87\xe5\xee\x4d\x47\x86\x6f\x9a\x07\xf6\x6e\x05\x37\x49\x2d\x8c\x9f\x48\x43\x92\x9e\x7e\xa6\x35\x7b\x36\x0e\xbe\x19\xc9\xd0\xa1\xe1\xaa\x01\x40\xef\x83\x61\x86\x35\x2f\x01\xfa\xe0\x4d\xc2\x30\x90\x2c\x36\xe4\xdb\x87\xd4\x51\x97\x78\x70\x24\x85\x61\xcf\xbf\xfb\xae\xfd\xbe\xfd\xb1\x01\xe8\x85\x8a\xf9\x67\x1e\x49\x0d\xc7\xb8\x02\x9f\x86\xa1\x01\xf0\x38\xd2\x0a\x42\x24\x41\x0b\xa2\xed\x43\x72\xa1\x75\xb4\x6b\x34\x52\x9f\xc9\x36\x12\x52\x5c\xc1\x01\xaf\x26\x53\x1c\xf5\x0e\xbf\x4d\xd6\x05\x1a\x58\xed\x97\x19\xfc\x2b\xab\x95\xad\x38\x24\xc1\xe1\x84\xad\xa0\xca\x7e\x93\x06\x94\x23\xde\x00\x68\x1f\x22\xad\xe0\x43\xa6\x8a\xd8\x93\x6b\x00\x76\x38\xb0\x2b\xd7\xa8\xe4\x21\x92\x7f\xfb\xf1\xf6\xfe\x87\xbb\x7e\x4b\x23\x56\x10\xc0\x91\xf6\xc2\xb1\x9c\x3b\xc4\x00\xac\x60\x5b\x82\x7a\x14\xd6\x41\xca\x72\xcf\x08\x6f\x3f\xde\x4e\xe6\x51\x32\x68\xbc\xbf\x62\xfe\x4e\x72\x7a\xc0\xce\x88\x5e\xe7\x48\xea\x19\x70\x39\x8b\x54\x09\xa7\x5c\x90\x03\xad\xd4\x61\x0d\xb6\x65\x05\xa1\x28\xa4\xe4\x6b\x5e\x4f\xdc\x42\x3e\x82\x1e\x42\xf7\x27\xf5\xd6\xc2\x1d\x49\x76\x02\xba\x0d\x69\x70\x39\xf5\x3b\x12\x03\xa1\x3e\x6c\x3c\x7f\x39\x78\x56\xb0\x50\x28\x07\x34\x9a\x14\xdf\x7f\xec\x8d\xc4\xe3\x90\x35\x4c\xf4\x2d\xa0\x77\x30\xe2\x23\x08\x65\x0e\x48\xfe\xc4\x5b\x39\xa2\x2d\xbc\x0f\x42\xc0\x7e\x1d\x56\xb0\x35\x8b\xba\x5a\x2e\x37\x6c\xfb\x2a\xee\xc3\x38\x26\xcf\xf6\xb8\x2c\xb5\xc8\x5d\xca\x09\x5d\x3a\xda\xd1\xb0\x54\xde\x2c\x50\xfa\x2d\x1b\xf5\x96\x84\x96\x18\x79\x51\x02\xf7\xa5\x88\xdb\xd1\x7d\x23\x53\xc9\xeb\xeb\x93\x48\xed\x31\x67\x5d\x4d\xd8\x6f\x0e\x70\x29\xb2\xab\xba\xe7\x5a\xcb\xd9\xc5\xc9\xac\xc6\x7f\x94\x37\x43\x59\x95\x4f\xef\xee\x3e\xc3\x9e\xb4\xa4\x60\xae\x79\x51\xfb\x68\xa6\x47\xe1\xb3\x50\xec\xd7\x24\x35\x71\x6b\x09\x63\xf1\x48\xde\xc5\xc0\xde\xca\xa2\x1f\x98\xfc\x5c\x74\x4d\xdd\xc8\x96\x33\xfd\x57\x22\xb5\x9c\x9f\x16\x6e\x4a\x2f\x43\x47\x90\xa2\x43\x23\xd7\xc2\xad\x87\x1b\x1c\x69\xb8\x41\xa5\xaf\x2e\x7b\x56\x58\x17\x59\xd2\x97\x85\x3f\x1d\x41\xf3\x83\x55\xad\x03\xbc\x9f\x15\x17\x33\xb4\x6f\xc1\xbb\x48\xfd\xac\x35\x1c\x29\x4b\x2e\x5f\x43\xa3\x5c\xf4\xb3\x39\x72\xbd\x1b\xcf\x19\x66\x1b\x57\xae\x52\xea\x28\x75\x24\x9e\x8c\xf4\x42\x33\xbf\x60\xe9\xc2\x7f\xb5\x19\x91\xbd\x21\x7b\x12\x3d\xb7\x61\xa3\xf1\x09\x78\xa6\xda\xfb\x83\xf9\x84\x77\xa4\x79\x2a\x1c\x06\xda\xd1\x7f\xfb\xc4\xd3\x35\xd5\xea\x47\x23\xf2\x70\x69\xe3\x2c\x84\x77\xf9\x5c\x69\x2d\x0f\xa1\x60\x38\x54\x63\x40\xe7\x84\xb4\x4c\x9c\x5c\x87\xd8\x97\x26\xb8\xe8\xb2\xfe\x2f\xdc\xb3\xf1\x3e\x2b\xe4\xd1\xc9\xbf\x88\x39\xff\x30\xea\x34\x48\x4a\x52\xac\x20\x08\x04\xd9\xa0\xe7\x2f\x65\xd4\x16\xf0\x7f\xc4\x70\xb1\xf2\x4f\xb7\x50\x04\x1f\x67\x3b\x49\x9e\xe8\x7c\x85\xe2\x72\x5b\x19\x5a\xd2\x97\x1b\xab\x1c\x9b\xb5\x56\xe8\x34\x0f\xaf\xe7\x7b\xeb\x02\xe7\x19\x74\x7c\x48\x4c\x6f\x96\x03\x54\xa2\x5a\x4c\xaf\x87\xe3\x2e\x40\xe5\x5d\x81\x49\xaa\xf5\xa0\x16\x04\x37\x34\x21\xc7\x2b\x61\xdf\x53\x34\x72\x1f\xce\x5f\x13\xaf\x5e\xcd\x1e\x0b\x65\xd9\x07\xef\xb8\xbe\x7f\xe0\xf7\x3f\x9a\xea\x95\xdc\xfd\x3e\x98\x0c\xfe\x13\x00\x00\xff\xff\x08\x2a\xd5\x2a\x7e\x09\x00\x00")

func configCrdsKudoDev_operatorsYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_operatorsYaml,
		"config/crds/kudo.dev_operators.yaml",
	)
}

func configCrdsKudoDev_operatorsYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_operatorsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_operators.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configCrdsKudoDev_operatorversionsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\xdd\x6f\xe3\x36\x12\x7f\xcf\x5f\x31\x48\x1f\xf6\x0a\xc4\xf6\x75\xdb\x87\x83\xdf\x16\xbb\xed\x21\xd7\x36\x09\x36\xd9\x1e\x0e\x6d\x0f\x4b\x4b\x63\x8b\x0d\x45\xaa\x24\x65\xaf\x6f\xd1\xff\xfd\x30\xfc\x90\x64\x5b\x94\x15\x67\xd3\xdb\x03\xa2\x97\xc4\x12\xc5\xf9\xe0\x7c\xfc\x66\x48\x9d\x4d\x26\x93\x33\x56\xf1\x9f\x50\x1b\xae\xe4\x1c\x58\xc5\xf1\x83\x45\x49\xbf\xcc\xf4\xfe\x6f\x66\xca\xd5\x6c\xfd\xd5\x02\x2d\xfb\xea\xec\x9e\xcb\x7c\x0e\xaf\x6b\x63\x55\xf9\x16\x8d\xaa\x75\x86\x6f\x70\xc9\x25\xb7\x5c\xc9\xb3\x12\x2d\xcb\x99\x65\xf3\x33\x00\x26\xa5\xb2\x8c\x6e\x1b\xfa\x09\x90\x29\x69\xb5\x12\x02\xf5\x64\x85\x72\x7a\x5f\x2f\x70\x51\x73\x91\xa3\x76\x14\x22\xfd\xf5\x5f\xa7\x2f\xa7\xdf\x9c\x01\x64\x1a\xdd\xeb\x77\xbc\x44\x63\x59\x59\xcd\x41\xd6\x42\x9c\x01\x48\x56\xe2\x1c\x54\x85\x9a\x59\xa5\xc3\x9b\x66\x7a\x5f\xe7\x6a\x9a\xe3\xfa\xcc\x54\x98\x11\xcd\x95\x56\x75\x35\x87\xe6\xbe\x7f\x33\xb0\xe3\x45\xb9\x0e\x93\x04\xf1\xdd\x13\xc1\x8d\xfd\xbe\xef\xe9\x0f\xdc\x58\x37\xa2\x12\xb5\x66\xe2\x90\x05\xf7\xd0\x70\xb9\xaa\x05\xd3\x07\x8f\xcf\x00\x4c\xa6\x2a\x9c\xc3\x15\xb1\x51\xb1\x0c\xf3\x33\x80\x35\x13\x3c\x77\x92\x7a\xc6\x54\x85\xf2\xd5\xcd\xe5\x4f\x5f\xdf\x66\x05\x96\xcc\xdf\x04\xc8\xd1\x64\x9a\x57\x6e\xdc\x3e\x63\xc0\x0d\xd8\x02\xc1\xbf\x01\x4b\xa5\xdd\xcf\x7d\xf6\xe0\xd5\xcd\xe5\x34\x4c\x57\x69\x7a\x6a\x79\x54\x07\x5d\x1d\x33\x68\xee\xed\x11\x7e\x41\x9c\x05\xa2\x39\x2d\x3c\x7a\xca\x81\x04\xe6\x60\x3c\x0f\x6a\x09\xb6\xe0\x06\x34\x56\x1a\x0d\x4a\x6f\x0a\x9d\x69\x81\x86\x30\x09\x6a\xf1\x1b\x66\x76\x0a\xb7\xe8\xf8\x04\x53\xa8\x5a\xe4\x64\x2d\x6b\xd4\x16\x34\x66\x6a\x25\xf9\x7f\x9a\x99\x0d\x58\xe5\x48\x0a\x66\x31\xac\x47\xbc\xb8\xb4\xa8\x25\x13\xa4\xd3\x1a\x2f\x80\xc9\x1c\x4a\xb6\x05\x8d\x44\x03\x6a\xd9\x99\xcd\x0d\x31\x53\xf8\x51\x69\x04\x2e\x97\x6a\x0e\x85\xb5\x95\x99\xcf\x66\x2b\x6e\xa3\xe1\x67\xaa\x2c\x6b\xc9\xed\x76\xe6\xcc\x97\x2f\x6a\xab\xb4\x99\xe5\xb8\x46\x31\x33\x7c\x35\x61\x3a\x2b\xb8\xc5\xcc\xd6\x1a\x67\xac\xe2\x13\xc7\xb8\x74\x76\x3f\x2d\xf3\x2f\x74\xf0\x12\xf3\xa2\xc3\xa9\xdd\x92\x15\x18\xab\xb9\x5c\x35\xb7\x9d\x41\x26\xf5\x4e\x06\x49\xcb\xcc\xc2\x6b\x9e\xff\x56\xbd\x74\x8b\xb4\xf2\xf6\xdb\xdb\x3b\x88\x44\xdd\x12\xec\xea\xdc\x69\xbb\x7d\xcd\xb4\x8a\x27\x45\x71\xb9\x44\xed\x17\x6e\xa9\x55\xe9\x66\x44\x99\x57\x8a\x4b\xeb\x7e\x64\x82\xa3\xdc\x55\xba\xa9\x17\x25\xb7\xb4\xd2\xbf\xd7\x68\x2c\xad\xcf\x14\x5e\x3b\xf7\x87\x05\x42\x5d\xe5\xcc\x62\x3e\x85\x4b\x09\xaf\x59\x89\xe2\x35\x33\xf8\xe4\x6a\x27\x0d\x9b\x09\xa9\xf4\xb8\xe2\xbb\x51\x6b\x77\xa0\xd7\x56\x73\x3b\xc6\x95\xde\x15\xda\x73\xc9\xdb\x0a\xb3\x1d\x0f\xc9\xd1\x70\x4d\x56\x6c\x99\x45\xb2\xfd\xbd\x17\xa6\x9d\x89\xfb\x9c\xd3\x3b\x68\xd5\xe3\xa0\x49\xc1\xc0\x47\x5d\x89\x19\xb1\x78\xeb\x1e\xee\xbf\xb8\x23\xc3\xeb\xbd\xc1\x8d\x00\x0c\x2c\x96\x15\x79\x5c\x1e\xed\xcf\x16\xcc\x42\xc6\x24\x2c\x70\x6f\x4a\x80\xda\x60\x4e\x6e\x1a\x88\xd3\xbf\x4c\x02\x97\xc6\x32\x99\xa1\x8f\x0d\xd8\x28\x60\x3a\x56\x96\x1c\x2b\x94\x39\xca\xec\x40\x31\x7b\x72\xbc\xe9\x0c\x04\xe6\x02\xba\x8b\x36\x42\xec\xcc\x11\x19\x51\x09\x46\xb8\xc5\xf2\x80\x50\x62\xd9\x1b\x92\x14\x6d\x96\xa8\x51\x66\x8e\xb6\xd7\x60\x9e\xa4\x91\x5e\xec\xb8\xe4\x7d\x31\x39\xc1\xcc\xab\x9b\xcb\x18\x89\xa3\x6c\x81\x19\x7b\x48\x77\x50\xd5\xfe\x5a\x72\x14\xf9\x0d\xb3\xc5\x08\xda\x2f\x2e\x97\x9e\x98\xb7\x0e\x05\x0c\x2a\x8e\x7e\xb5\x9b\x30\xef\x6c\x00\x59\x0e\x6a\xd9\x3b\x23\xc1\x06\x20\x37\xd6\x18\xde\xb8\xf0\xd1\x28\x18\x5d\x9b\x1c\x2c\xe3\x12\x98\x4f\x9e\xf0\x8f\xdb\xeb\xab\xd9\xdf\x55\x62\x4a\x27\x05\xb0\x2c\x43\x63\xbc\xfb\x95\x28\xed\x05\x98\x3a\x2b\x80\x99\xe8\x99\xb7\xf4\x64\x5a\x32\xc9\x97\x68\xec\x34\xd0\x40\x6d\x7e\x7e\xf9\x6b\xbf\xf6\x00\xbe\x53\x1a\xf0\x03\x2b\x2b\x81\x17\xc0\x83\x35\xc5\x10\x1b\xac\xc0\x25\x67\x52\x47\x33\x23\x6c\xb8\x2d\xb8\x4c\x69\x00\x2a\x95\x07\xb1\x37\x4e\x5c\xcb\xee\x11\x54\x10\xb7\x46\x10\xfc\x1e\xe7\x70\x4e\xe1\xa8\xc3\xe6\x47\x02\x37\x7f\x9c\x27\x66\xfd\xcb\xa6\x40\x8d\x70\x4e\x83\xce\x3d\x73\x4d\x26\xa5\x7b\xd1\x5e\x5a\x26\x9d\x83\x5b\xcd\x57\x2b\xd4\x0e\xa8\xf4\x5d\x2e\x41\x50\xe0\xfd\x12\x94\x26\x0d\x48\xd5\x99\xc2\x4d\x4c\xab\x57\x61\xc6\x97\x1c\xf3\x03\xa6\x7f\x7e\xf9\x6b\x92\xe3\x5d\x7d\x01\x97\x39\x7e\x80\x97\xc0\xa5\xd7\x4d\xa5\xf2\x2f\xa7\x70\xe7\xac\x63\x2b\x2d\xfb\x40\x94\xb2\x42\x19\x4c\x69\x56\x49\xb1\x25\x99\x0b\xb6\x46\x30\xaa\x44\xd8\xa0\x10\x93\xe8\xa2\x1b\xb6\x25\x2d\xc4\x85\x23\x7b\x63\x50\x31\x6d\x07\xad\x35\xe2\x97\xbb\xeb\x37\xd7\x73\xcf\x19\x19\xd4\xca\x81\x32\xca\x81\x4b\x4e\x88\x84\xa0\x88\xcf\xab\xce\x1a\xf7\xd2\x72\x7b\x99\xda\x9b\x0f\x45\xcf\x82\xc9\x15\x7a\x79\x11\x96\x35\xe5\xba\xe9\x8b\x53\xfc\x78\x1f\x5a\xb4\x57\x0f\xc8\xd8\x0f\x1c\xff\xa3\x54\x3d\x5a\x38\x57\x0d\x8c\x10\xee\xaa\x63\xe5\x83\xc2\x51\x65\xa2\x25\x5a\x74\xf2\xe5\x2a\x33\x24\x5a\x86\x95\x35\x33\xb5\x46\xbd\xe6\xb8\x99\x6d\x94\xbe\xe7\x72\x35\x21\xd3\x9c\x78\x1b\x30\x33\x57\x5e\xcc\xbe\x70\x7f\x4e\x96\xc5\x15\x06\x63\x05\x72\x83\xff\x0c\xa9\x88\x8e\x99\x9d\x24\x54\x13\x09\xaf\xc6\xad\x94\x5b\xa8\x18\x32\xcc\x41\x84\x6a\x32\xf8\x76\x0a\x6f\xe3\xd4\xa9\xf8\xb4\xe6\x2e\x05\xb3\x5a\x58\x43\x61\x67\xc9\x57\x27\x65\xc2\x88\xa8\xc7\xe7\xe2\x17\xb7\x5e\x82\x6c\xff\x5d\x72\xed\x4d\xc1\xb3\x22\x16\x48\x41\x84\x84\x04\x9c\x90\x79\xee\xd3\x0b\x93\xdb\x27\x77\x47\x32\x8a\x5a\x13\x47\xdb\x49\x28\xd9\x27\x4c\xe6\xf4\xbf\xe1\xc6\xd2\xfd\x93\xac\xa0\xe6\xa3\x42\xd0\xbb\xcb\x37\x7f\x8e\x93\xd6\xfc\xc4\x78\xb3\x1e\x6d\x02\xe7\x71\xc5\x33\x56\x91\xba\x4d\x10\xeb\xf7\x9a\x6b\x87\x41\x8c\x2b\xd5\x37\x94\x68\x77\x5a\x08\x87\x57\x50\x09\x5b\xa8\x75\x83\x30\x98\x46\x82\xb4\x6a\x43\xd5\xd5\x2f\x12\xbe\xf5\x40\x64\x0e\xff\xfe\x7a\xfa\xd5\xf4\x9b\xfe\xbc\x3a\x28\x5c\x60\xad\x67\xa1\x26\xbb\x5e\xdc\xf3\x7c\xdd\xe9\xa0\x1c\x12\xdc\xab\xa4\xba\x8f\x98\xd6\x6c\xbb\x5b\xa4\x06\xac\x3c\x88\xf1\xaf\xdd\x8c\x4d\x00\x88\xa8\xc3\x00\x4a\x55\xaf\x0a\x67\x2e\xba\x74\x5d\x07\xf2\x38\x81\x16\xb6\xaa\x3e\x60\x8f\x4b\x8a\x34\x96\xd0\x4b\xa9\x72\xbe\xdc\xb6\xa6\x47\xd5\x5a\xc8\xee\x7b\xaf\x0d\x41\xf6\x61\xc0\xfe\x18\xb8\x3e\xb8\x74\x83\x50\xfd\x51\x40\x1d\x58\x3f\x9e\x3a\x19\xa6\x7b\x5e\x7b\xe7\x7c\x02\x90\xfe\xe9\x21\xfa\x53\x00\xf4\x27\x81\xe7\x4f\x06\xce\x1f\x01\xcd\x1d\x08\xef\xe7\xf6\x24\x60\xde\x81\xe0\xbd\xb3\x3e\x14\x96\x1f\x02\xf0\xde\x69\x8f\x83\xf2\x41\x6f\x4d\x01\xf2\xcf\x1f\x8e\x0f\x8a\x95\x82\xe2\x9f\x1d\x10\x3f\x2a\x45\x12\x84\x7f\x96\x10\xfc\x48\x52\x3f\x0a\x5d\x1f\x07\x5c\x53\xc5\xec\xff\x03\x6c\x1d\xd4\x5c\x02\xb2\x7e\x66\x80\x75\x40\x84\x24\xf6\xaa\x98\x66\x25\x5a\xd4\x07\x00\x66\x4c\xcf\xf3\x26\xbe\xbd\x0b\x6c\xd7\x4c\x73\xb6\xe0\x82\xdb\x6d\x08\xcc\x7d\xbb\x6b\xbb\xd7\x02\x29\x9a\xfb\xce\xb0\xe5\xae\xbf\x4c\x80\xa1\x6d\x16\x3f\xb4\x5f\x1a\x8a\xbd\x11\xe8\xfc\x8d\x1f\xe9\x37\x55\xc2\x6b\x21\x7f\xfb\x54\xd9\x28\x89\x86\x54\x5a\xad\x79\x9e\xac\x33\x17\x1e\x37\xa6\xb9\x86\xe3\x85\x45\x97\xbb\x31\xec\x37\x3f\xda\x65\x60\x20\x94\x5c\xa1\xee\x0e\xa5\xb5\x28\xd4\x66\xa0\x81\xd7\x0a\xba\xe1\x42\xb8\x4d\x1b\x83\xf9\x69\x32\x70\x53\x09\xb6\x1d\x59\xe9\xbf\x69\x47\x87\xad\x04\xbf\x75\xb0\xd8\xc2\xbb\x4b\x73\x12\x03\x23\xbb\x41\xe7\x57\x01\xfd\x90\xfc\xdd\x1d\x8d\x00\x5d\x23\x27\x21\xcf\xc7\xdd\x8f\x64\x87\x59\xa0\x2b\xe5\xba\x40\xf3\xbd\xdf\xa6\x7e\x7d\xfd\xee\xea\xee\x3d\xcd\x22\xa1\x36\x71\x9b\xce\xfb\x8a\x20\x8b\x49\xb6\x81\x09\x8d\x05\x28\xf9\x8b\xf4\x9b\x4f\x2e\x9c\x57\x82\x67\xcc\xcc\x01\x3e\x7e\x84\xa9\xf3\x45\x33\x75\x54\xe0\x8f\x04\xba\x3c\xda\xdc\x48\x95\x7d\x07\x7a\x7b\x1b\x86\x76\xfa\x33\x01\x53\xef\x78\x4b\x9c\x11\x6c\xaa\x29\xbf\xc0\xc6\xa5\x68\xb9\x99\x10\x8d\xf3\x98\x0b\x70\x55\x31\xda\x02\x75\xc7\x37\xc9\x42\x4c\xbd\x5c\xf2\x61\xff\x5a\x28\x25\xb0\xb7\x66\x09\x68\x79\x84\x98\x77\x7e\x24\xf0\x9c\x52\x4c\xd3\x86\xaa\x04\x93\xde\x4c\x56\x68\x0d\xe0\x07\xcc\x6a\x0a\x59\x9b\x22\xd9\x73\xf6\x78\xb8\x0d\x98\x0e\x52\x9a\x68\x57\x97\xcd\x96\x58\xe8\x22\x77\x82\xd2\x7b\xbf\x73\xfa\x3e\xd5\x17\x5a\x12\x08\x26\x86\x1c\x04\x77\x5c\x39\x48\x8f\x1f\xb8\xb1\xa4\x43\x52\xdf\x86\x1b\x04\x6e\x5f\x18\x78\x9f\x63\x25\xd4\xf6\xfd\x09\x5e\x75\x42\x09\x4f\x8c\x1d\xc4\x67\x96\xe7\xee\x9c\x08\x13\x37\x03\x31\x7c\x37\xd5\x90\x80\xad\xad\x31\x30\xa8\xc3\xce\xdd\x4d\xc1\x8c\x5b\x17\x12\x1c\xfd\x86\xe3\x82\x2a\x24\xf2\x40\xdb\x17\xbf\x86\x33\x47\xe5\xe6\x1b\x61\x1c\x81\x70\xc9\x2a\x62\xc8\xbd\xe6\x35\xef\x4a\x48\xf7\x74\xb0\x24\x49\xa4\xd8\x14\xa5\x1d\xf1\xe3\x56\xa6\xb1\x58\x05\xd9\x63\x95\xfd\x7d\x03\x30\x12\x53\xc7\xdd\xff\x44\x60\x3d\xa6\x1f\x7f\xa5\xe3\xab\xbf\x8e\x84\x1a\x7f\x39\xee\x87\x66\xd9\xd1\xc2\xad\x93\x35\xa8\x9b\x5e\xed\x68\x3b\xea\xa3\xd9\x9e\x1e\x98\x14\x3a\x2a\x8a\xaa\x00\x63\x15\xc5\x29\xd6\x9e\xac\x48\x69\x07\x8e\x2d\x5d\x82\xf5\xce\x26\xba\x89\xc8\xda\xa0\xe3\xda\x37\xba\x06\xda\x7e\xf1\x72\x0b\xad\xb2\xac\xee\xd9\x3f\xee\x5e\x63\x56\xd0\x5f\xc7\xd6\x31\xd0\x1d\xb3\x9a\x61\x28\x33\xf7\x47\xa9\x8e\xd2\xe0\x83\x49\xa7\xc3\x50\xff\xb8\xde\x48\xf6\xd0\xe9\x8c\xd5\xcc\xe2\x6a\x3b\xda\x8c\xaf\x75\x8e\xbe\x3b\xd6\xf8\x73\xa1\x36\x1e\x80\xd4\x0b\xa7\x97\xd8\x40\x19\x5e\x63\xc1\xe4\xcc\x47\x9d\x16\xac\xb8\x83\x75\x39\xa8\x3a\x11\x73\xba\x72\x0d\xea\xf4\xa8\x86\x64\x2d\x04\x21\x97\x39\x58\x5d\xf7\x03\xa2\x61\xf5\x0d\x2b\xee\x54\x95\x75\xd4\x92\x90\x6c\xbc\xb2\x4e\x4d\x86\x07\x99\xab\x4d\x12\x94\xc6\xda\xa8\x45\x3f\xf7\x49\x0f\xea\x35\x49\xb4\xd7\xeb\x76\xf8\xf8\xa1\x73\xf8\xc5\x8d\x06\xb6\x66\x5c\x04\xf0\xe9\x75\x37\x70\x14\x09\x46\xd6\x84\x77\xcc\xdc\xfb\x52\x6a\x25\xd4\x82\x89\x0b\xa8\x94\xd8\x96\x4a\x57\x05\xcf\x80\x53\x4e\x2e\xe3\x29\xc0\xc8\x4e\x55\x2f\x04\xcf\x7a\xdb\x81\x2d\x8f\x8e\xe7\x07\xa6\xf2\xf4\x7e\xf7\xc9\xd5\xc3\x91\x17\xf7\x8f\x86\x0d\x68\xc9\x9d\x0c\xc3\x72\x81\xb9\xf1\x5a\x50\xc6\xf0\x28\xa9\x9b\xc8\x84\xde\xa9\xdb\xdc\x49\x05\x83\xda\xb7\xac\xd7\x8a\xe7\xb0\xd1\xdc\x1d\x00\xcc\xdc\xc1\x5c\xa8\xe5\xac\x64\xda\x14\x4c\x08\xd7\x46\xa6\xe4\xe1\x1b\xd5\xee\x04\x44\xc5\x74\xd2\x49\x32\xd4\x0e\x4c\xb8\x76\xa8\x09\x7b\xad\x34\xb5\x0a\x85\x10\xf1\xf8\x3d\x97\x39\xb1\x88\x90\xab\x8d\x34\x3c\xc7\x78\xdc\x33\x55\xcb\x54\x95\x56\x2c\x2b\x80\x9b\x0b\xcf\x8e\x93\x9f\xb0\xbf\x6b\x37\x3a\x68\x2f\x95\xf5\x0d\xe0\x40\x3b\xc0\xda\xa4\x3b\x93\x37\xfd\x66\x94\xf7\x2b\x43\x19\x9c\x47\x31\x17\x98\xa9\x12\x81\x95\x0b\xbe\xaa\x55\x6d\x9a\x13\xb1\xa1\x94\x48\xe1\x1f\x52\x8c\x9e\xc2\x3f\x11\x4a\xbe\x2a\x2c\x68\x5c\x73\xc3\xad\x77\x92\x56\x88\x6e\xef\x37\x84\x95\x21\xf4\x1f\xb9\x91\xc0\x8d\xa9\x13\xb5\xcb\xf1\xcc\x9d\x2b\x39\x90\xb1\x8f\xd5\x3e\x74\x2d\x99\x65\xe2\x71\x53\x54\xbc\x1a\x60\xe2\x68\x72\xdf\x8d\x90\xbc\xc2\x70\x4c\x92\xee\x2e\x42\x74\x67\xbe\x9c\x5e\xa1\xa4\xa0\xe4\xca\xc3\xc1\x7c\xc8\xdc\x44\x31\xfe\x44\x3c\x27\xf3\xb6\x41\x39\x04\x0a\xc7\x22\x26\xe2\xe9\x18\x6c\x19\x0d\x59\xee\x71\x10\x36\x3c\x6c\xae\x64\xac\x7b\xf0\x64\xa3\xd0\xd1\x88\xfc\x0f\xa3\x20\x54\xa5\x06\xf8\x1e\xc1\x71\x73\xb2\xfa\x11\xf6\x38\x4a\x31\x9f\x4c\xe2\x0d\x93\xf6\x5b\x9d\xe8\x3f\xc0\x28\x0f\x1c\x5c\xa2\x13\xea\xf4\x58\x37\x9d\x58\xab\x0f\xe8\x6f\x37\xef\x45\x32\x1e\x22\xc4\xaa\xad\x73\x52\xd6\x2a\xf8\xd7\xab\x1f\x7f\x68\x19\x02\xa1\xb2\xde\x82\x6e\xaf\x25\x47\xc1\x5d\xe4\xa8\x9d\xcb\xd3\x0d\xdd\x71\xfc\x70\x8a\x9d\x20\x44\xff\x29\xe3\x1e\x65\xd5\xd5\x4a\xb3\x9c\x16\xfc\x3b\xad\xca\x41\x6c\xf5\x6e\x67\xa8\x13\xcb\xe7\xf4\x3d\x40\x65\xda\xd3\xd2\x7e\xf6\x43\x2b\x72\x7b\xbd\x9f\x06\x8a\x7d\xa2\x93\x11\x27\x9e\x8d\x78\x3e\xd0\xbc\x27\xef\xf3\x81\xe6\xe7\x03\xcd\x9e\xe3\xe7\x03\xcd\xcf\x07\x9a\x47\x08\xf7\x7c\xa0\xf9\xb3\x3f\xd0\xfc\x7c\x18\xf8\xf9\x30\xf0\x69\x80\x3b\x71\x7a\x38\x41\xa6\xff\x33\x44\xcb\x6c\x6d\x46\x7f\x88\xe8\x46\xef\x7c\x8a\xa8\x16\x06\xf5\x7a\xe4\xb7\x88\x3d\x2c\xec\xdd\x6a\x3f\xdd\x0e\x5f\x89\x37\xb7\x1c\x93\x93\xf0\xbd\x76\xfb\x14\xc0\xd3\xef\x14\x54\x54\xb6\xb3\x55\x2c\xb1\x5a\x09\x09\xe4\x54\x16\xf3\xab\xfd\x0f\xb7\xcf\x7d\x96\x8d\x5f\x62\xbb\x9f\x99\x92\xbe\x68\x31\x73\xf8\xf9\xd7\x33\x08\xcd\x80\x08\xc2\xdd\xcd\xff\x06\x00\x00\xff\xff\x9b\xd8\x9c\xba\xf0\x3e\x00\x00")

func configCrdsKudoDev_operatorversionsYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_operatorversionsYaml,
		"config/crds/kudo.dev_operatorversions.yaml",
	)
}

func configCrdsKudoDev_operatorversionsYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_operatorversionsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_operatorversions.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"config/crds/kudo.dev_instances.yaml":        configCrdsKudoDev_instancesYaml,
	"config/crds/kudo.dev_operators.yaml":        configCrdsKudoDev_operatorsYaml,
	"config/crds/kudo.dev_operatorversions.yaml": configCrdsKudoDev_operatorversionsYaml,
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
	"config": &bintree{nil, map[string]*bintree{
		"crds": &bintree{nil, map[string]*bintree{
			"kudo.dev_instances.yaml":        &bintree{configCrdsKudoDev_instancesYaml, map[string]*bintree{}},
			"kudo.dev_operators.yaml":        &bintree{configCrdsKudoDev_operatorsYaml, map[string]*bintree{}},
			"kudo.dev_operatorversions.yaml": &bintree{configCrdsKudoDev_operatorversionsYaml, map[string]*bintree{}},
		}},
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
