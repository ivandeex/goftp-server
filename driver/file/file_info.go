// Copyright 2018 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package file

import "os"

type fileInfo struct {
	os.FileInfo

	mode  os.FileMode
	owner string
	group string
}

func (f *fileInfo) Mode() os.FileMode {
	return f.mode
}

func (f *fileInfo) Owner() string {
	return f.owner
}

func (f *fileInfo) Group() string {
	return f.group
}

func (f *fileInfo) UID() int {
	return -1
}

func (f *fileInfo) GID() int {
	return -1
}
