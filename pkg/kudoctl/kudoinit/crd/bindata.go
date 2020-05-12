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

var _configCrdsKudoDev_instancesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5a\x5b\x6f\xe3\xb8\x15\x7e\xf7\xaf\x38\xc8\x3e\x64\x82\x8d\xe5\x4e\x0a\x14\x85\x51\x14\x98\x66\x66\x0b\xb7\xdb\x24\x18\x27\x53\x2c\x66\xa7\x00\x4d\x1e\x5b\x6c\x28\x52\xc3\x8b\x1d\x77\x67\xff\x7b\x71\x48\x49\xbe\x48\xf2\x65\xbb\xb3\xe5\x4b\x22\x8a\x3c\xd7\xef\x5c\x48\x6b\x30\x1c\x0e\x07\xac\x94\x1f\xd0\x3a\x69\xf4\x18\x58\x29\xf1\xc5\xa3\xa6\x27\x97\x3d\xff\xd1\x65\xd2\x8c\x96\xaf\x67\xe8\xd9\xeb\xc1\xb3\xd4\x62\x0c\xb7\xc1\x79\x53\xbc\x47\x67\x82\xe5\xf8\x16\xe7\x52\x4b\x2f\x8d\x1e\x14\xe8\x99\x60\x9e\x8d\x07\x00\x4c\x6b\xe3\x19\x4d\x3b\x7a\x04\xe0\x46\x7b\x6b\x94\x42\x3b\x5c\xa0\xce\x9e\xc3\x0c\x67\x41\x2a\x81\x36\x72\xa8\xf9\x2f\x7f\x97\xdd\x64\x7f\x18\x00\x70\x8b\x71\xfb\xa3\x2c\xd0\x79\x56\x94\x63\xd0\x41\xa9\x01\x80\x66\x05\x8e\x41\x6a\xe7\x99\xe6\xe8\xb2\xe7\x20\x4c\x26\x70\x39\x70\x25\x72\x62\xb6\xb0\x26\x94\x63\x68\xe6\xd3\x96\x4a\x8e\xa4\xc3\xa4\xda\x1d\xa7\x94\x74\xfe\xef\x3b\xd3\xdf\x4b\xe7\xe3\xab\x52\x05\xcb\xd4\x16\xb7\x38\xeb\xa4\x5e\x04\xc5\xec\x66\x7e\x00\xe0\xb8\x29\x71\x0c\x77\xc4\xaa\x64\x1c\x05\xcd\x85\x99\xad\xec\x54\xb1\x77\x9e\xf9\xe0\xc6\xf0\xd3\xcf\x03\x80\x25\x53\x52\x44\x2d\xd3\x4b\x53\xa2\x7e\xf3\x30\xf9\xf0\xfb\x29\xcf\xb1\x60\x69\x12\x40\xa0\xe3\x56\x96\x71\x5d\x23\x22\x48\x07\x3e\x47\x48\x4b\x61\x6e\x6c\x7c\x6c\x04\x85\x37\x0f\x93\xac\x22\x50\x5a\x53\xa2\xf5\xb2\x16\x82\xc6\x96\xd3\x9b\xb9\x3d\x56\x97\x24\x4b\x5a\x03\x82\xdc\x8c\x89\x65\xe5\x2c\x14\xe0\x12\x73\x33\x07\x9f\x4b\x07\x16\x4b\x8b\x0e\x75\x72\xfc\x16\x59\xa0\x25\x4c\x83\x99\xfd\x1b\xb9\xcf\x60\x8a\x96\x88\x80\xcb\x4d\x50\x82\xb0\xb1\x44\xeb\xc1\x22\x37\x0b\x2d\xff\xd3\x50\x76\xe0\x4d\x64\xa9\x98\xc7\xca\x25\xf5\x90\xda\xa3\xd5\x4c\x91\x15\x03\x5e\x03\xd3\x02\x0a\xb6\x06\x8b\xc4\x03\x82\xde\xa2\x16\x97\xb8\x0c\xfe\x61\x2c\x99\x68\x6e\xc6\x90\x7b\x5f\xba\xf1\x68\xb4\x90\xbe\x86\x39\x37\x45\x11\xb4\xf4\xeb\x51\x04\xab\x9c\x05\x6f\xac\x1b\x09\x5c\xa2\x1a\x39\xb9\x18\x32\xcb\x73\xe9\x91\xfb\x60\x71\xc4\x4a\x39\x8c\x82\xeb\x88\xf2\xac\x10\xdf\x34\xbe\xbe\xdc\x92\xd4\xaf\x09\x16\xce\x5b\xa9\x17\xcd\x74\x44\x61\xaf\xdd\x09\x8c\xe4\x5f\x56\x6d\x4b\xf2\x6f\xcc\x4b\x53\x64\x95\xf7\xef\xa6\x8f\x50\x33\x8d\x2e\xd8\xb5\x79\xb4\xf6\x66\x9b\xdb\x18\x9e\x0c\x25\xf5\x1c\x6d\x72\xdc\xdc\x9a\x22\x52\x44\x2d\x4a\x23\xb5\x8f\x0f\x5c\x49\xd4\xbb\x46\x77\x61\x56\x48\x4f\x9e\xfe\x1c\xd0\x79\xf2\x4f\x06\xb7\x31\xd8\x61\x86\x10\x4a\xc1\x3c\x8a\x0c\x26\x1a\x6e\x59\x81\xea\x96\x39\xfc\xea\x66\x27\x0b\xbb\x21\x99\xf4\xb8\xe1\xb7\x73\xd4\xee\xc2\x64\xad\x66\xba\x4e\x26\x9d\x1e\xaa\x83\x70\x5a\x22\xdf\x09\x0d\x81\x4e\x5a\x82\xaf\x67\x1e\x09\xf4\xf5\xca\x6c\x8b\x54\x57\x38\x56\xe1\x6f\x99\x37\xb6\x23\x2e\x5b\x12\xdc\xef\xae\x8d\xe2\xca\xb9\x44\x02\x8d\xc5\x39\x5a\xa4\x1c\xe1\x0d\x61\x28\xbd\xe2\xfb\x7b\xf6\xc8\xd7\x78\xc9\xf6\xe6\xfb\xa4\x85\xde\x24\xd2\x29\xf0\x9b\x87\x49\x9d\x38\x52\xbe\xc0\x5a\xce\x16\xc7\x5e\xe7\xd5\x63\x2e\x51\x89\x07\xe6\xf3\xa3\x5c\x2f\x27\xf3\xc4\x26\x86\x51\x34\x47\x29\x91\xe3\x4e\x3e\x8a\x49\x13\x99\x48\x93\x1d\x24\x01\x08\x6d\x16\xab\xf5\xd7\x29\x68\xaa\xd8\xdc\xe4\x30\xcf\xa4\x06\x96\xb2\x3a\xfc\x6d\x7a\x7f\x37\xfa\xab\x49\xb2\x76\xd2\x64\x9c\xa3\x73\x09\x2a\x05\x6a\x7f\x0d\x2e\xf0\x1c\x98\xab\x51\x34\xa5\x37\x59\xc1\xb4\x9c\xa3\xf3\x59\xc5\x01\xad\xfb\x78\xf3\xa9\xcb\x66\x00\xdf\x19\x0b\xf8\xc2\x8a\x52\xe1\x35\xc8\x64\xe5\x26\x0b\xd4\xa0\x90\x2e\x19\xa2\xa1\x07\x2b\xe9\x73\xd9\xad\x38\x83\xd2\x88\x4a\xe1\x55\x54\xd4\xb3\x67\x04\x53\x29\x1a\x10\x94\x7c\xc6\x31\x5c\x10\xca\xb6\x44\xfc\x89\x4a\xee\xcf\x17\x9d\x34\x5f\xad\x72\xb4\x08\x17\xb4\xe4\x22\x09\xd6\x24\x7a\x9a\xab\xf1\xb1\x11\xd0\xe7\xcc\x83\xb7\x72\xb1\x40\x8b\xdd\xd6\x8c\xd9\x8b\xb2\xc2\x15\x18\x4b\xba\x6b\xb3\x45\x20\x92\x25\x9f\x55\x61\x22\x5a\x02\x7f\xbc\xf9\xd4\x23\xed\xae\x9d\x40\x6a\x81\x2f\x70\x03\x52\x27\xab\x94\x46\x5c\x65\xf0\x18\x11\xb1\xd6\x9e\xbd\x10\x1f\x9e\x1b\x87\x1a\x8c\x56\xeb\x6e\x69\x0d\xe4\x6c\x89\xe0\x4c\x81\xb0\x42\xa5\x86\x29\x8b\x08\x58\xb1\x35\xe9\x5f\xbb\x8b\x10\xc6\xa0\x64\xd6\xef\x96\xd0\x4e\xaa\x8f\xf7\x6f\xef\xc7\x49\x2a\x82\xd0\x42\x93\x28\x94\x9a\xe7\x92\x0a\x25\x55\xc8\x94\xee\x09\x93\xd1\x1c\x21\x81\xc3\x1b\xe0\x39\xd3\x0b\xec\x24\x1b\x35\x45\x98\x07\x4a\xc0\xd9\xe5\xb9\xd1\xba\x5f\xeb\xea\xd1\x51\xf3\xf6\x13\xc3\xff\xa9\x72\x9c\xa4\x56\x6c\x43\x8f\xaa\x75\xb7\x85\xe7\x83\x6a\x51\x43\x6c\x35\x7a\x8c\x9a\x09\xc3\x1d\x29\xc5\xb1\xf4\x6e\x64\x96\x68\x97\x12\x57\xa3\x95\xb1\xcf\x52\x2f\x86\x04\xc4\x61\x42\x82\x1b\xc5\xe6\x76\xf4\x4d\xfc\xf3\x8b\xb4\x88\xed\xea\x69\xaa\xc4\xa5\xbf\x85\x3e\xc4\xc7\x8d\xce\x56\xa7\x6e\x86\x4e\xad\x4a\x97\xd3\xba\x38\xee\xed\xa4\x90\x58\xe5\x92\xe7\x75\x67\xbb\xc9\x9e\x9d\x31\x52\x30\x91\x52\x2e\xd3\xeb\xaf\x0e\x5b\x32\x64\xb0\x24\xcf\x7a\x58\x9d\xab\x86\x4c\x0b\xfa\xdf\x49\xe7\x69\xfe\x6c\xcb\x05\x79\x42\x90\x3e\x4d\xde\xfe\x36\x60\x0e\xf2\xec\x88\xec\xec\xe2\x68\x94\xcc\xb2\x02\x3d\xda\x56\x03\xc3\x84\x88\x27\x57\xa6\x1e\x0e\x34\x39\xbf\x88\xa7\x62\xfa\xdd\x0b\xf2\xe0\x8f\x35\x72\x97\x8f\xb1\x18\x32\x8b\xe0\x57\x86\xd2\x3f\xb5\x70\xb4\x1f\xb0\x26\x00\x9c\x69\x6a\xaf\x9b\x0a\x38\x06\x78\x7d\xd5\x12\x54\x6a\x21\x2d\x72\xaf\xd6\xe0\x73\x6b\xc2\x22\xaf\x1a\xf2\x58\x3a\x80\x1b\x6b\xd1\x95\x46\x0b\x2a\x2a\x8d\x55\xea\xf4\xbe\xdd\xd3\x66\x0f\x8d\xcd\x5a\x5c\x0a\x56\x02\xdc\x5c\x41\x8b\x97\x43\x1f\x4f\x26\x15\x40\x76\xe9\x6d\xdb\x23\x3e\x51\x36\xe9\x6e\xec\xe0\x9f\xb9\x54\xd8\x68\x03\xaf\x5e\x5f\xd5\x9a\x3b\xc8\x59\x59\xa2\x76\x54\xea\xed\x1a\xbc\x2c\x10\x18\x04\x87\xb6\x2a\x60\x6d\x79\xd9\x46\xd5\x6b\x60\x1b\xb1\x5f\xdd\x5c\x6d\x0c\x9a\x0c\x1e\x03\xdd\xd1\x11\x49\x34\x07\x6a\x27\x7d\x48\xf7\x18\x2d\xca\xab\x1c\xf5\x16\xba\x40\x18\x74\xfa\xf2\xd2\x57\xa2\x00\x66\x8b\x8c\xd8\xa3\x95\x46\x48\x0e\x33\xc6\x9f\x43\x19\xfb\xaf\xde\x56\x86\xa2\xc3\x4a\x51\x9f\xf0\xf0\x45\xba\x68\xd4\x6a\xef\x5c\x2a\xcc\xe0\x4d\x83\x5b\xb5\xae\x7a\x33\x13\xad\x62\x8d\x29\xda\x46\x35\x96\x00\xc4\x51\xc5\x66\x82\xca\xec\x86\x49\xca\x23\x64\x0f\x1b\xb4\x8e\xc0\x50\x4c\xbb\xbd\x9a\xdf\xa2\x79\x67\x3c\x8e\x61\xc7\xab\x95\xf3\xea\xd3\x50\x34\x68\x6c\xbb\x88\x63\x0f\xf6\xda\x36\x8d\x9d\xde\x64\x0a\xb7\x4f\xef\xdf\xbf\xbb\x7b\xfc\xfe\x87\x2a\x0a\xe8\x50\x79\x1f\x8f\x34\x5b\x97\x1c\x5b\x97\x4a\xf0\x6a\x72\x7b\x45\xa6\x15\x46\xb7\x71\x15\x1b\xb7\x64\xcf\x4a\xda\xeb\xed\x4e\x68\x25\x95\xa2\xf8\xe2\x0a\x99\x25\x4e\xef\x18\xcf\xf7\x62\xb0\x45\x33\x67\x14\xa8\x41\xcb\xcf\x01\x81\x12\xa3\x33\xf5\x59\x20\xe2\x86\x54\x8f\x24\x66\x94\x2c\x87\x1b\xa8\x49\x9f\x18\x52\x03\xd8\x81\x56\x8d\x2b\x22\xb7\x9f\xfd\x0e\x1d\xc3\xca\x2a\x9e\xba\x12\xf8\xc1\xa4\x5f\xdd\x46\x1d\xcb\xfb\x8d\x8f\xa7\x71\x3d\x70\x56\x12\x2c\xd2\x91\xb7\x39\xea\xc6\xaa\x60\x94\x32\xe1\xfc\x13\xdd\x29\xd5\x87\x6c\x1c\x2f\x45\x88\x52\x02\x4a\x6e\x94\x70\xb5\x0f\x26\x6f\xab\x7b\x9e\x6b\x90\x9a\xab\x20\xba\x18\xd1\x78\x7a\x9a\xbc\x75\x19\xc0\x5f\x90\xb3\xe0\xa8\xff\x26\xd4\x5c\x7a\xb8\xbf\xfb\xfe\x07\x4a\x26\x69\x45\x05\x11\x62\xa9\x81\x29\x99\x6e\xa3\x92\x02\x71\x77\x1f\xfd\x4a\xc2\xc6\x4a\x52\x7b\xd4\x3e\xc6\x41\x8e\xaa\x74\x50\xd0\x11\xca\x05\x5b\x69\x41\xcc\xe2\xdb\x58\xfc\x3a\x49\x0a\x13\xfb\xf8\x05\x7a\xc2\xfc\x5c\xc5\x5b\x96\x5f\xa5\x3e\x76\x5f\x7e\xb4\x70\xd1\x7d\xfd\x91\xe0\xb0\x7d\x01\x62\x66\x55\xfe\x6c\xdd\x80\x9c\x70\x01\x42\x38\x9e\x76\x42\xf2\xb4\x4a\xbd\x23\xe4\xc5\x43\x43\x0d\xb6\xef\x26\xe3\x99\x3d\x4d\xc7\xb2\x18\x63\xf4\x47\x0d\x8f\x39\xba\xae\xf3\x0f\x55\xe6\x74\xa4\x8f\x2a\x25\x4f\x79\xcb\xb4\x8b\x12\x39\xda\xdb\x3d\xbe\xc0\x1d\x95\xa9\x0e\x9a\x75\x4e\x83\x2f\x3d\x5b\x37\x24\x0e\x8f\x65\x07\xf1\xb8\xe7\x9d\xb5\xc6\xc6\xa7\x3f\x0d\xe3\xf8\x73\x9c\x7e\xc0\x94\x81\xb7\x69\xff\xab\x8f\x77\x07\xed\x1f\x0f\x8b\xb5\xec\x11\xfb\xdb\x24\xc3\xb0\xfe\x3b\xfc\xf6\x74\xda\xed\xbd\x3d\x4c\x4e\x90\x7b\xd9\x2b\xf7\x17\xf8\x8e\x79\xa6\x00\xa3\xdd\x1a\xd2\xf1\x9f\x5b\x53\x94\x0a\x7d\x17\x38\x88\xee\x97\xf6\xbd\xc1\xa1\x64\x0d\xa0\x98\xf3\x4f\xe9\xa6\x74\xf3\xeb\x46\x67\xd8\xcf\x8d\x2d\x98\x1f\x03\xad\x1d\x52\xb3\xd3\x7d\x99\x70\x28\xb3\x02\x14\xe8\x1c\x5b\x74\x16\x87\xa3\x7b\xfb\x0e\xb9\x47\x37\x96\x39\x73\xdd\xda\x03\x48\x8f\x45\xcf\xab\xbd\x18\x7e\x20\x2a\xa7\xc4\x30\xad\xeb\x21\x78\xd8\x17\x69\x1c\x34\xd1\x49\xfa\xa6\xd1\x6f\xae\x33\x88\xf4\x17\xe5\x7a\xfc\xfa\xc5\xf9\x4c\x01\xb1\x3c\x28\xdf\x41\x07\x77\xa8\x30\xf5\x58\x9e\xe0\x65\xe2\x7b\x90\xe8\x29\xae\x4e\xe3\x04\x87\xa7\x71\x92\x41\xd2\x38\xe6\xfc\xb3\x09\x1e\x07\x42\x1a\xe7\xc3\xe1\x28\x49\x38\x15\x30\x67\x2a\xd5\x7b\x60\xee\x5a\xc6\xac\x65\xdd\xd7\xa7\x27\x10\x3a\x4c\xe2\x90\x69\xbf\x46\x74\x1d\x31\x50\x4f\xfb\xfb\xb5\x1a\xe0\xff\xad\x05\xee\x21\xb9\xdb\x18\x9f\xdb\x04\xf7\xc9\xb9\xd3\x1a\x9f\xdc\x06\x1f\x31\xf8\x01\xf0\xec\x18\xdc\x29\xc9\xb1\xfa\xd9\x63\x86\x80\x3a\xde\x73\xc4\xfb\x9a\x59\xf0\x64\x34\x9e\x7e\xfa\x24\x83\xa5\xc5\xb3\x64\xd0\xf6\xb9\xde\x0a\xb4\x04\x15\x87\x9f\x43\xba\x48\xd4\xb0\x66\x85\x8a\xbf\x16\x18\xed\xa4\x88\x87\x43\x27\x17\x5a\xce\x25\x67\xda\xc3\x2a\xde\x83\x44\x76\xd2\x5f\xb6\x0f\x06\xda\xec\x4b\x7f\x7a\x8f\xbf\x37\xb5\xf9\xfe\xa2\xfa\xd4\xa3\x99\x8a\x41\x32\xac\x3e\xba\xd8\xbc\x05\x48\x7d\xfe\x18\xbc\x0d\x49\x5d\xe7\x8d\xa5\x9c\x9a\x66\x36\x11\xc6\x38\xc7\xd2\xa3\xb8\xdb\xff\x08\xe3\x22\x35\x4e\xf5\x37\x16\xf1\x91\x1b\x9d\x3a\x7d\x37\x86\x8f\x9f\x06\x89\x2a\x8a\x0f\xb5\x30\x34\xf9\xdf\x00\x00\x00\xff\xff\x32\xc9\x61\x20\xb5\x22\x00\x00")

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

var _configCrdsKudoDev_operatorsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4d\x8f\x1b\x37\x0c\xbd\xcf\xaf\x20\xd2\x43\x2e\xf5\xb8\x69\x81\xa2\xf0\x2d\xd8\xe6\xb0\x68\x93\x06\xd9\x60\x2f\x45\x0f\x9c\x11\x6d\xb3\x3b\x23\xa9\x24\x65\x74\xf3\xeb\x0b\x49\xe3\x8f\xf1\xda\xbb\x6d\x81\xcc\x4d\x4f\x22\x1f\xf5\xf8\x31\x6a\x16\x8b\x45\x83\x91\xef\x49\x94\x83\x5f\x01\x46\xa6\xbf\x8d\x7c\x5e\x69\xfb\xf0\x93\xb6\x1c\x96\xbb\x37\x1d\x19\xbe\x69\x1e\xd8\xbb\x15\xdc\x24\xb5\x30\x7e\x22\x0d\x49\x7a\xfa\x99\xd6\xec\xd9\x38\xf8\x66\x24\x43\x87\x86\xab\x06\x00\xbd\x0f\x86\x19\xd6\xbc\x04\xe8\x83\x37\x09\xc3\x40\xb2\xd8\x90\x6f\x1f\x52\x47\x5d\xe2\xc1\x91\x14\x86\x3d\xff\xee\xbb\xf6\xfb\xf6\xc7\x06\xa0\x17\x2a\xe6\x9f\x79\x24\x35\x1c\xe3\x0a\x7c\x1a\x86\x06\xc0\xe3\x48\x2b\x08\x91\x04\x2d\x88\xb6\x0f\xc9\x85\xd6\xd1\xae\xd1\x48\x7d\x26\xdb\x48\x48\x71\x05\x07\xbc\x9a\x4c\x71\xd4\x3b\xfc\x36\x59\x17\x68\x60\xb5\x5f\x66\xf0\xaf\xac\x56\xb6\xe2\x90\x04\x87\x13\xb6\x82\x2a\xfb\x4d\x1a\x50\x8e\x78\x03\xa0\x7d\x88\xb4\x82\x0f\x99\x2a\x62\x4f\xae\x01\xd8\xe1\xc0\xae\x5c\xa3\x92\x87\x48\xfe\xed\xc7\xdb\xfb\x1f\xee\xfa\x2d\x8d\x58\x41\x00\x47\xda\x0b\xc7\x72\xee\x10\x03\xb0\x82\x6d\x09\xea\x51\x58\x07\x29\xcb\x3d\x23\xbc\xfd\x78\x3b\x99\x47\xc9\xa0\xf1\xfe\x8a\xf9\x3b\xc9\xe9\x01\x3b\x23\x7a\x9d\x23\xa9\x67\xc0\xe5\x2c\x52\x25\x9c\x72\x41\x0e\xb4\x52\x87\x35\xd8\x96\x15\x84\xa2\x90\x92\xaf\x79\x3d\x71\x0b\xf9\x08\x7a\x08\xdd\x9f\xd4\x5b\x0b\x77\x24\xd9\x09\xe8\x36\xa4\xc1\xe5\xd4\xef\x48\x0c\x84\xfa\xb0\xf1\xfc\xe5\xe0\x59\xc1\x42\xa1\x1c\xd0\x68\x52\x7c\xff\xb1\x37\x12\x8f\x43\xd6\x30\xd1\xb7\x80\xde\xc1\x88\x8f\x20\x94\x39\x20\xf9\x13\x6f\xe5\x88\xb6\xf0\x3e\x08\x01\xfb\x75\x58\xc1\xd6\x2c\xea\x6a\xb9\xdc\xb0\xed\xab\xb8\x0f\xe3\x98\x3c\xdb\xe3\xb2\xd4\x22\x77\x29\x27\x74\xe9\x68\x47\xc3\x52\x79\xb3\x40\xe9\xb7\x6c\xd4\x5b\x12\x5a\x62\xe4\x45\x09\xdc\x97\x22\x6e\x47\xf7\x8d\x4c\x25\xaf\xaf\x4f\x22\xb5\xc7\x9c\x75\x35\x61\xbf\x39\xc0\xa5\xc8\xae\xea\x9e\x6b\x2d\x67\x17\x27\xb3\x1a\xff\x51\xde\x0c\x65\x55\x3e\xbd\xbb\xfb\x0c\x7b\xd2\x92\x82\xb9\xe6\x45\xed\xa3\x99\x1e\x85\xcf\x42\xb1\x5f\x93\xd4\xc4\xad\x25\x8c\xc5\x23\x79\x17\x03\x7b\x2b\x8b\x7e\x60\xf2\x73\xd1\x35\x75\x23\x5b\xce\xf4\x5f\x89\xd4\x72\x7e\x5a\xb8\x29\xbd\x0c\x1d\x41\x8a\x0e\x8d\x5c\x0b\xb7\x1e\x6e\x70\xa4\xe1\x06\x95\xbe\xba\xec\x59\x61\x5d\x64\x49\x5f\x16\xfe\x74\x04\xcd\x0f\x56\xb5\x0e\xf0\x7e\x56\x5c\xcc\xd0\xbe\x05\xef\x22\xf5\xb3\xd6\x70\xa4\x2c\xb9\x7c\x0d\x8d\x72\xd1\xcf\xe6\xc8\xf5\x6e\x3c\x67\x98\x6d\x5c\xb9\x4a\xa9\xa3\xd4\x91\x78\x32\xd2\x0b\xcd\xfc\x82\xa5\x0b\xff\xd5\x66\x44\xf6\x86\xec\x49\xf4\xdc\x86\x8d\xc6\x27\xe0\x99\x6a\xef\x0f\xe6\x13\xde\x91\xe6\xa9\x70\x18\x68\x47\xff\xed\x13\x4f\xd7\x54\xab\x1f\x8d\xc8\xc3\xa5\x8d\xb3\x10\xde\xe5\x73\xa5\xb5\x3c\x84\x82\xe1\x50\x8d\x01\x9d\x13\xd2\x32\x71\x72\x1d\x62\x5f\x9a\xe0\xa2\xcb\xfa\xbf\x70\xcf\xc6\xfb\xac\x90\x47\x27\xff\x22\xe6\xfc\xc3\xa8\xd3\x20\x29\x49\xb1\x82\x20\x10\x64\x83\x9e\xbf\x94\x51\x5b\xc0\xff\x11\xc3\xc5\xca\x3f\xdd\x42\x11\x7c\x9c\xed\x24\x79\xa2\xf3\x15\x8a\xcb\x6d\x65\x68\x49\x5f\x6e\xac\x72\x6c\xd6\x5a\xa1\xd3\x3c\xbc\x9e\xef\xad\x0b\x9c\x67\xd0\xf1\x21\x31\xbd\x59\x0e\x50\x89\x6a\x31\xbd\x1e\x8e\xbb\x00\x95\x77\x05\x26\xa9\xd6\x83\x5a\x10\xdc\xd0\x84\x1c\xaf\x84\x7d\x4f\xd1\xc8\x7d\x38\x7f\x4d\xbc\x7a\x35\x7b\x2c\x94\x65\x1f\xbc\xe3\xfa\xfe\x81\xdf\xff\x68\xaa\x57\x72\xf7\xfb\x60\x32\xf8\x4f\x00\x00\x00\xff\xff\x7d\x5f\xed\xd8\x7e\x09\x00\x00")

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

var _configCrdsKudoDev_operatorversionsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x3b\x5d\x6f\xe3\x36\xb6\xef\xf9\x15\x07\xd3\x87\xb9\x05\x62\xfb\x76\x2e\x70\x71\xe1\xb7\xc1\x4c\x7b\x91\xed\x34\x09\x26\x99\x2e\x16\xdd\x02\xa1\xa5\x63\x8b\x0d\x45\x72\x49\xca\x1e\x6f\xd1\xff\xbe\x38\x87\x94\x64\x39\x92\xac\x24\x93\xee\xac\x5e\x12\x4b\xe4\xe1\xf9\xfe\xd2\xd1\xd9\x6c\x36\x3b\x13\x56\xfe\x8c\xce\x4b\xa3\x97\x20\xac\xc4\xcf\x01\x35\xfd\xf2\xf3\xfb\xff\xf3\x73\x69\x16\xdb\xef\x56\x18\xc4\x77\x67\xf7\x52\xe7\x4b\x78\x57\xf9\x60\xca\x8f\xe8\x4d\xe5\x32\x7c\x8f\x6b\xa9\x65\x90\x46\x9f\x95\x18\x44\x2e\x82\x58\x9e\x01\x08\xad\x4d\x10\x74\xdb\xd3\x4f\x80\xcc\xe8\xe0\x8c\x52\xe8\x66\x1b\xd4\xf3\xfb\x6a\x85\xab\x4a\xaa\x1c\x1d\x9f\x50\x9f\xbf\xfd\xef\xf9\x9b\xf9\xff\x9e\x01\x64\x0e\x79\xfb\xad\x2c\xd1\x07\x51\xda\x25\xe8\x4a\xa9\x33\x00\x2d\x4a\x5c\x82\xb1\xe8\x44\x30\x2e\xed\xf4\xf3\xfb\x2a\x37\xf3\x1c\xb7\x67\xde\x62\x46\x67\x6e\x9c\xa9\xec\x12\x9a\xfb\x71\x67\x42\x27\x92\x72\x95\x80\x24\xf2\xf9\x89\x92\x3e\xfc\xd8\xf7\xf4\x83\xf4\x81\x57\x58\x55\x39\xa1\x1e\xa2\xc0\x0f\xbd\xd4\x9b\x4a\x09\xf7\xe0\xf1\x19\x80\xcf\x8c\xc5\x25\x5c\x12\x1a\x56\x64\x98\x9f\x01\x6c\x85\x92\x39\x53\x1a\x11\x33\x16\xf5\xdb\xeb\x8b\x9f\xff\xe7\x26\x2b\xb0\x14\xf1\x26\x40\x8e\x3e\x73\xd2\xf2\xba\x63\xc4\x40\x7a\x08\x05\x42\xdc\x01\x6b\xe3\xf8\xe7\x31\x7a\xf0\xf6\xfa\x62\x9e\xc0\x59\x47\x4f\x83\xac\xd9\x41\xd7\x81\x1a\x34\xf7\x8e\x0e\x7e\x4d\x98\xa5\x43\x73\x12\x3c\xc6\x93\xd3\x11\x98\x83\x8f\x38\x98\x35\x84\x42\x7a\x70\x68\x1d\x7a\xd4\x51\x15\x0e\xc0\x02\x2d\x11\x1a\xcc\xea\x37\xcc\xc2\x1c\x6e\x90\xf1\x04\x5f\x98\x4a\xe5\xa4\x2d\x5b\x74\x01\x1c\x66\x66\xa3\xe5\x3f\x1b\xc8\x1e\x82\xe1\x23\x95\x08\x98\xe4\x51\x5f\x52\x07\x74\x5a\x28\xe2\x69\x85\xe7\x20\x74\x0e\xa5\xd8\x83\x43\x3a\x03\x2a\x7d\x00\x8d\x97\xf8\x39\xfc\x64\x1c\x82\xd4\x6b\xb3\x84\x22\x04\xeb\x97\x8b\xc5\x46\x86\x5a\xf1\x33\x53\x96\x95\x96\x61\xbf\x60\xf5\x95\xab\x2a\x18\xe7\x17\x39\x6e\x51\x2d\xbc\xdc\xcc\x84\xcb\x0a\x19\x30\x0b\x95\xc3\x85\xb0\x72\xc6\x88\x6b\xd6\xfb\x79\x99\x7f\xe3\x92\x95\xf8\xd7\x07\x98\x86\x3d\x69\x81\x0f\x4e\xea\x4d\x73\x9b\x15\x72\x90\xef\xa4\x90\x24\x66\x91\xb6\x45\xfc\x5b\xf6\xd2\x2d\xe2\xca\xc7\xef\x6f\x6e\xa1\x3e\x94\x45\xd0\xe5\x39\x73\xbb\xdd\xe6\x5b\xc6\x13\xa3\xa4\x5e\xa3\x8b\x82\x5b\x3b\x53\x32\x44\xd4\xb9\x35\x52\x07\xfe\x91\x29\x89\xba\xcb\x74\x5f\xad\x4a\x19\x48\xd2\xff\xa8\xd0\x07\x92\xcf\x1c\xde\xb1\xf9\xc3\x0a\xa1\xb2\xb9\x08\x98\xcf\xe1\x42\xc3\x3b\x51\xa2\x7a\x27\x3c\xbe\x38\xdb\x89\xc3\x7e\x46\x2c\x3d\xcd\xf8\x43\xaf\xd5\x5d\x18\xb9\xd5\xdc\xae\xfd\x4a\xaf\x84\x8e\x4c\xf2\xc6\x62\xd6\xb1\x90\x1c\xbd\x74\xa4\xc5\x41\x04\x24\xdd\x3f\xda\x30\x3f\x00\xdc\x67\x9c\xd1\x40\x6d\x8f\x81\x0e\x12\x06\xd1\xeb\x6a\xcc\x08\xc5\x1b\x7e\x78\xbc\xb1\x43\xc3\xbb\xa3\xc5\x0d\x01\x02\x02\x96\x96\x2c\x2e\xaf\xf5\x2f\x14\x22\x40\x26\x34\xac\xf0\x08\x24\x40\xe5\x31\x27\x33\x4d\x87\xd3\xbf\x42\x83\xd4\x3e\x08\x9d\x61\xf4\x0d\xd8\x30\x60\x3e\x95\x96\xda\x9f\x8d\xd2\x70\xc5\x32\xfb\x88\x6b\x74\x48\x87\x91\x02\x09\xa9\x3d\xa0\x36\xd5\xa6\x60\x9d\x73\x25\x7b\x23\xc2\x4b\x61\x80\xbd\xa9\x1e\x90\x20\x35\x49\x3b\x80\x71\x50\x9a\x5c\xae\xf7\x8c\xb2\x23\xb0\x24\xc5\xe4\xb5\x8e\xb6\x0d\xc9\x0d\x06\x9d\x6b\x2f\x09\x6f\xaf\x2f\x6a\x87\x5a\xf3\xca\x45\x7a\x1e\x9c\x38\xca\x2f\xba\xd6\x12\x55\x7e\x2d\x42\x71\xf2\xd4\xd7\x17\xeb\x44\x1f\x8b\xd7\x80\x00\x2b\x31\x8a\xab\xf1\xd3\x2c\x44\x14\x79\xbc\xd9\x03\x12\x80\xac\xd0\x61\x5a\x7f\x1e\x9d\x49\xd2\x99\xd6\xb7\x93\x48\x40\xc4\xd8\x07\x7f\xb9\xb9\xba\x5c\xfc\xbf\x89\xb8\xf6\xc2\x14\x59\x86\xde\x47\xdb\x29\x51\x87\x73\xf0\x55\x56\x80\xf0\xb5\x59\xdd\xd0\x93\x79\x29\xb4\x5c\xa3\x0f\xf3\x74\x02\x3a\xff\xcb\x9b\x5f\xfb\x78\x06\xf0\x83\x71\x80\x9f\x45\x69\x15\x9e\x83\x8c\x5c\x6e\xbc\x63\xad\x3c\xd2\x47\x46\x34\xf0\x60\x27\x43\x21\xfb\x09\x17\x60\x4d\x9e\x08\xde\x31\xa1\x41\xdc\x23\x98\x44\x68\x85\xa0\xe4\x3d\x2e\xe1\x15\x69\xd6\x01\x8a\xbf\x53\x56\xf2\xc7\xab\x5e\x98\xff\xb5\x2b\xd0\x21\xbc\xa2\x25\xaf\x22\x62\x4d\x00\xa4\x7b\xb5\x7e\xb4\x08\xb2\x5d\x06\x27\x37\x1b\x74\xd8\xcf\x4d\xf6\xea\xe4\x2d\xbf\x25\xf5\x96\x6b\xd0\xe6\x00\x00\x83\x25\x99\x59\xcc\xe4\x5a\x62\xfe\x00\xe1\x5f\xde\xfc\x3a\x80\x6d\x97\x4f\x20\x75\x8e\x9f\xe1\x0d\x48\x1d\xb9\x62\x4d\xfe\xed\x1c\x6e\x59\x23\xf6\x3a\x88\xcf\x74\x4e\x56\x18\x8f\x1a\x8c\x56\xfb\x7e\x6c\x0d\x14\x62\x8b\xe0\x4d\x89\xb0\x43\xa5\x66\xd1\x2b\xe5\xb0\x13\x7b\xa2\xbf\x16\x17\x69\x98\x00\x2b\x5c\xe8\xa6\x16\xbd\x50\x6f\xaf\xde\x5f\x2d\x23\x56\xa4\x42\x1b\xce\xa1\x28\x64\xad\x25\x25\x10\x94\x39\xc4\x30\x48\x3a\xc9\xec\xa8\xa2\x72\x90\x5b\x2b\x84\xde\x3c\xf4\x7a\xc0\x6e\x83\xb9\xbb\xae\x28\x30\xcd\x5f\x3f\xd6\x5a\x8f\x73\x80\xfa\xea\xc9\x05\x8e\x1d\xc3\xbf\x29\xa2\x4e\x22\x8b\x13\xf6\x93\x64\x5d\x1e\xe8\xf3\x28\x59\x54\x3a\x38\x8d\x01\x99\xb2\xdc\x64\x9e\x88\xca\xd0\x06\xbf\x30\x5b\x74\x5b\x89\xbb\xc5\xce\xb8\x7b\xa9\x37\x33\x52\xc4\x59\xd4\x04\xbf\xe0\xfc\x7f\xf1\x0d\xff\x79\x12\x15\x9c\xb5\x4f\x23\x85\x97\xfe\x19\xf4\xd0\x39\x7e\xf1\x68\x72\xea\x24\x71\x6a\x54\x7a\x7d\x13\x1d\x42\x76\xbc\x93\x4c\x62\x57\xc8\xac\xa8\x33\xfe\xd6\x7b\xf6\xda\x48\x29\xf2\xe8\x72\x85\xde\xbf\xb8\xda\x12\x23\x2b\x47\xf8\xec\x67\xa9\x02\x9d\x09\x9d\xd3\xff\x5e\xfa\x40\xf7\x1f\xcd\xb9\x4a\x4e\x30\xd2\x4f\x17\xef\xff\x1c\x65\xae\xe4\xa3\x2d\xb2\x37\xbb\xa5\xcb\x0a\x27\x4a\x0c\xe8\x1e\x24\x30\x32\x60\xd9\x93\xd5\x74\x68\xbe\xae\x77\x43\x26\x2c\x09\x24\xd5\x86\xc2\x49\xb1\x92\x4a\x86\x7d\x72\xcc\x7d\x55\x77\xf7\x5a\x21\x79\xf3\x98\x31\x06\xc9\x79\x27\x25\x0c\x6d\x12\xf9\xd0\xab\x8f\x25\x5f\x84\xe8\x5a\x54\x2a\xf4\x3d\x3a\xa2\xe2\x7d\x5c\x19\x8b\xad\xb4\x2d\xc5\xef\x18\x2a\x1b\x26\xd1\x12\xeb\xcc\x56\xe6\x03\x81\x16\x60\x15\xf3\xc6\x61\xac\xe1\x94\xb2\x75\xb1\x9b\x82\x7e\xf3\xa3\x15\x83\x00\x65\xf4\x06\xdd\xe1\x52\x92\x45\x61\x76\x03\x88\x13\xd6\x2d\xa1\x3b\xa9\x14\x17\x73\x1e\xf3\xa7\xd1\x20\xbd\x55\x62\x7f\x39\x10\x08\x8e\x69\x68\x57\xa7\x12\x23\x96\x14\xab\x3d\x7c\xba\xf0\x4f\x42\x60\x28\x04\x1d\x9d\xfc\xea\x32\x65\x3f\x44\xff\x61\xa5\x93\x52\xd7\x1a\x93\x14\xe7\xeb\xaa\x68\x80\x89\x6b\xa9\x90\xbb\x31\x87\x89\xe6\x5d\x6c\x5f\xbd\xbb\xfa\x74\x79\x7b\x47\x50\x34\x54\xbe\x2e\xdf\xa3\xad\x28\xd2\x98\x01\x98\x82\x13\xb3\x94\x4a\xfe\x5d\xc7\xa2\x94\xdd\xb9\x55\x32\x13\x7e\x09\xf0\xfb\xef\x30\x67\x5b\xf4\x73\x3e\x05\xfe\x18\xc8\x2e\x4f\xf0\x8c\x2a\x7a\x4a\xae\x27\xf0\xed\x63\x5a\xda\x64\x8d\xbe\xce\xa9\x3b\xd6\x52\x43\x84\x60\x86\x0c\x06\x1b\x93\x22\x71\x0b\xa5\x1a\xe3\xf1\xe7\x94\xae\xee\x0a\x0c\x05\xba\x03\xdb\x24\x0d\xf1\xd5\x7a\x2d\xc7\xed\x6b\x65\x8c\xc2\xde\x9a\x25\x65\xcb\x13\xc8\xbc\x8d\x2b\x41\xe6\x14\x62\x98\x4c\xa6\x51\x09\x1d\xd5\x64\x83\xc1\x03\x7e\xc6\xac\x22\x97\xb5\x2b\x70\x48\x8c\x31\x1f\x6e\x1d\x26\xa7\x94\xbe\xd6\xab\x8b\xa6\x54\x4e\xdd\xb1\x03\xa7\x74\x17\x3b\x2a\x77\x03\x80\x29\xae\x46\x84\x38\x05\x67\xac\x38\xa5\xc7\xcf\xd2\x07\xe2\x21\xb1\x6f\x27\x3d\x82\x0c\xaf\x3d\xdc\xe5\x68\x95\xd9\xdf\x3d\xc9\xaa\xd8\x2d\xce\x78\xd1\x04\xe6\xed\x2d\x1e\xe8\x47\x54\x77\x72\xab\xb4\xbf\x21\x91\xcb\x9b\xbb\x78\xe2\x53\x90\x1a\x8c\x6d\xf5\x23\xe1\x9c\xe8\x56\x1a\xc4\xad\x07\x41\x43\xe4\x39\x37\xb5\x85\xba\x1e\x09\x2c\xdd\xf8\x47\x5c\x6f\x09\x14\xe0\xd1\xd1\x3f\x66\x0d\xd7\x85\xf0\x4c\x33\x49\x03\x63\x77\x64\x45\x65\x1b\xb9\x85\xd0\xe7\x54\xc7\xc3\x99\x65\x78\x13\x98\x9e\x0e\x2e\x85\x25\x84\x78\x5b\x54\x07\xae\x6b\xf9\xe9\x68\x9d\x34\x10\xf7\x87\x4e\xea\x90\xaf\xa4\xe7\x52\xcc\x07\xb4\x89\xf6\xba\xf4\xff\xb1\xc9\x7a\x06\x40\xd7\xad\xca\x01\x6f\x7f\x8a\x3f\xf1\x1a\x76\xfa\xf1\x3a\xa1\xdd\xf1\x62\xec\xc7\xa0\x74\xb8\x70\xc3\xb4\x26\x76\xd3\xd6\x03\x6e\xd7\xfc\x68\x7a\x69\x23\x40\xe1\x80\x45\x35\x2b\xc0\x07\x43\xce\x53\xb4\x6d\xe0\x21\xee\xc0\x29\xd1\x0d\xa0\x7e\xd0\xf1\xf3\x75\xba\xef\x91\xb1\x8e\xdd\xb7\xe6\x35\xc7\xf0\xc5\x82\x36\x59\x56\x3d\xe8\xec\x75\xaf\x29\x12\x8c\xd7\x29\x39\xa6\x73\xa7\x48\x33\x2d\x15\xfe\xfe\xe4\xa9\x93\x38\xf8\xe8\xa3\x87\xdd\x50\xff\xba\x5e\x4f\xf6\x58\x70\x3e\x38\x11\x70\xb3\x9f\xac\xc6\x57\x2e\xc7\xd8\xb2\x6b\xec\xb9\x30\xbb\x98\x15\x55\x2b\xe6\x4b\xdd\xd5\x19\x97\xb1\x12\x7a\x11\xbd\x4e\x9b\x41\xf1\x5b\xc0\x1c\x4c\x35\xe0\x73\x0e\xe9\x1a\xe5\xe9\x49\x0e\x8d\xf3\x66\x9c\x2b\x4f\xe5\xc7\x01\xcd\x03\x68\x4f\xe7\xc4\x53\x23\xdd\x83\xb0\xd4\x46\x00\x8a\x51\xad\x4b\xa2\x9f\xfd\xfd\xf7\x1e\xb8\xbd\x56\xd3\x39\xea\x43\xf2\x70\x94\xb6\xf1\x6a\x10\x5b\x21\x55\xca\x68\x23\x7b\x46\xde\x7b\xc0\xc4\x42\xf3\x56\xf8\xfb\x58\x9f\x6d\x94\x59\x09\x75\x0e\xd6\xa8\x7d\x69\x9c\x2d\x64\x06\x92\x62\x6a\x59\xbf\x72\xac\xd1\xb1\xd5\x4a\xc9\xac\xb7\xc7\xd8\xe2\xc8\x38\x3f\x32\x14\x0f\xf5\xec\x9e\x51\x92\x9c\xd8\x78\xfc\x1e\x6a\x84\x4b\xfc\x1a\x0a\xcb\x15\xe6\x3e\x72\xc1\x78\x2f\x6b\x4a\x19\x90\x4f\x0d\x59\xa1\x94\xd9\x0d\x19\x73\x15\xfb\xe0\x5b\x23\x73\xd8\x39\xc9\x6f\x1b\x33\x9e\x02\x80\x4a\x2f\x4a\xe1\x7c\x21\x94\xe2\xde\x34\x39\xff\xd8\xfd\x36\x5a\xed\x29\xc5\x1d\xb4\x83\x0c\x1d\x27\x03\xdc\x63\xf5\x90\xa3\x45\x9d\x13\x68\x93\xaa\x2b\xc2\xf1\x47\xa9\x73\x42\x11\x21\x37\x3b\xed\x65\x8e\xf5\xbb\xe5\xa1\x02\xc9\x5a\x67\x44\x56\x80\xf4\xe7\x11\x1d\xa6\x9f\x0a\x0a\xee\x61\x72\xbd\xa0\x4d\x88\x5d\xe5\x74\x76\xca\x95\x07\x2d\x96\x0c\xe6\x37\x6f\xa2\xe9\x78\x8a\xc0\xb2\x26\x73\x85\x99\x29\x11\x44\xb9\x92\x9b\xca\x54\xbe\x79\xfd\x9e\xea\x93\xa1\xfc\x85\x18\xe3\xe6\xf0\x57\x84\x52\x6e\x8a\x00\x0e\xb7\xd2\xcb\x10\x8d\xa4\x25\xe2\xb0\xa1\x9c\x3c\xc7\x58\x49\x51\x63\xa3\x41\x7a\x5f\x0d\x14\x44\xa7\x23\xef\xf0\x9b\xc6\xf6\xea\x28\xdb\x41\xbe\x20\xac\x6d\x5e\x5f\x25\xf4\x0d\x55\x67\x52\x28\xaa\x52\xcd\xf9\x68\x10\x6a\x8b\x00\xda\x58\x1a\xcf\x33\x00\xc7\x2f\x9d\x3b\x74\x9f\x0e\x14\xb9\xd1\x23\x09\xc4\xa9\xfa\x90\xae\xb5\x08\x42\x3d\x0f\x44\x5d\xc7\x0e\x37\x41\x26\x52\x63\xba\x1e\xf4\xf1\x12\xaa\x01\xbc\x84\x98\xc0\xe8\x21\x2b\x9a\x44\x9c\x15\xd9\xbd\xd8\x8c\x30\xa8\x43\x14\x4a\x6e\x06\x10\xbe\xf5\x4e\x36\xd1\x73\x50\x26\x13\xaa\xb9\xb7\x36\x2a\x47\x37\x42\x91\x71\x64\x33\x9f\x3e\x7e\x20\x92\xea\x5d\x41\xb8\x95\xe0\x91\xa3\x67\x90\x93\x4a\xfc\x67\x49\xdc\xca\xa1\x0a\x1b\xa6\xa4\xa9\xdd\x74\x40\x5a\x4c\xd3\x09\x74\x77\x95\x52\x19\x11\xbb\x55\x1b\xd4\xa4\x1b\xdc\x7d\x19\xcd\xec\x04\x03\xaa\x23\x71\x5d\x99\xe8\xbc\xed\xff\x8f\x95\x37\x53\x73\x7f\xd4\xdb\x1f\xa4\x3a\x99\xfc\x4f\xce\xbf\xd7\x5f\x12\xd8\x3d\x8e\x26\xd4\x8f\x83\x35\x98\x45\x3c\x1a\xd8\xa4\xba\xe1\x74\xd5\x60\xcd\x08\x42\x13\x50\x21\x9b\xfc\xf4\xf1\xc3\x44\x43\x96\xeb\xf6\x85\xf3\x79\xa7\xdf\x4c\x06\xe9\xd0\x1b\xb5\xc5\xd4\x50\x64\xe3\x1c\x61\x40\x74\x00\x6d\x9f\x8b\xa7\xc5\xc8\x1d\xa8\x3d\x64\x46\xaf\xe5\xa6\xe2\x19\x0e\x8d\x9d\x83\x9e\x47\x6c\x1a\xf3\x7a\x86\x95\x4e\x12\xef\x69\xb9\xed\x84\x0e\xdf\xbb\x93\xfe\x66\x2c\x56\x8d\x6a\xd0\x13\x1a\x6c\x75\xc3\xe3\x89\x4d\xb6\x11\xc6\x74\x13\xde\xfa\x98\x58\x1b\xd4\xed\x96\xc6\x21\x71\xb8\xfa\xdb\xdb\x9f\x3e\xb4\x08\xb1\x5e\xf4\x75\x62\x8e\x1a\xfc\x3e\x45\x10\xf6\x70\x74\xc3\x1d\xf8\xb9\x34\x2b\x47\xb5\xc3\xe4\x5a\xaa\xb2\x1b\x27\x72\x2a\x3a\x7e\x70\xa6\x1c\x2d\xaa\x3e\x75\x96\x32\x59\x31\x99\x3f\xaa\xa4\x7c\x3b\x93\x15\xa1\x3f\xb4\x11\x0e\xd7\x5f\xa6\x06\xfb\x42\x73\x56\x4f\x9c\xb4\x3a\x15\x41\xc6\xa7\xad\x9e\x37\x6f\x75\xd2\x52\x47\x67\xae\x9e\x39\x75\x35\x94\xaf\xeb\xa7\xcf\x5d\x0d\x65\xbc\x3c\xf9\xf2\x02\x93\x57\x2f\x31\x7b\xf5\x32\xd3\x57\x2f\x34\x7f\xf5\x82\x13\x58\xcf\x9a\xc1\x1a\x2a\x37\xa8\x8e\x7f\xda\x14\xd6\xb0\xb6\xd6\x2f\xb6\x1e\x3b\x87\x35\x54\xa0\x1f\x4f\x67\x4d\x98\xc4\x3a\x69\xc7\xc3\x39\xd9\xd7\x3f\x8f\xf5\xa5\x5e\x88\x7f\x0d\x53\x59\x93\x68\x19\x9c\xcc\xfa\x4a\x67\xb3\x26\xbc\x7d\x3f\x39\x9f\xf5\xdc\x09\xad\xa1\x46\x85\xff\x8f\x98\xd1\x3a\xc9\xc1\x81\x39\xad\xaf\x6e\x52\xeb\xcb\xbf\xd1\xde\x3e\xea\x2b\x81\xfe\x8f\x1d\x82\x08\x95\x9f\xfc\xb9\x03\xaf\xee\x7c\xf0\x60\x56\x1e\xdd\x76\xe2\x17\x0f\x3d\x28\x1c\xdd\x6a\x3f\x10\x4b\xdf\xa2\x35\xb7\x18\xc9\x59\xfa\x2a\xac\x7d\x0a\x10\xcf\x5f\x42\x70\x55\x54\x75\x1f\x8c\x13\x1b\x4c\x77\x5a\x0a\x29\xc9\xb1\x01\xf3\xcb\xe3\xcf\xc3\x5e\xc5\x28\x5b\x7f\xef\xc5\x3f\x33\xa3\x63\xd1\xe2\x97\xf0\xcb\xaf\x67\x90\x7a\x1f\x75\x12\xce\x37\xff\x15\x00\x00\xff\xff\x74\xc8\x6d\x07\x56\x37\x00\x00")

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
