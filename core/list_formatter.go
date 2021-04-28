// Copyright 2018 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package core

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type listFormatter []FileInfo

// Short returns a string that lists the collection of files by name only,
// one per line
func (formatter listFormatter) Short() []byte {
	var buf bytes.Buffer
	for _, file := range formatter {
		fmt.Fprintf(&buf, "%s\r\n", file.Name())
	}
	return buf.Bytes()
}

// Detailed returns a string that lists the collection of files with extra
// detail, one per line
func (formatter listFormatter) Detailed() []byte {
	var buf bytes.Buffer
	for _, file := range formatter {
		fmt.Fprint(&buf, file.Mode().String())
		fmt.Fprintf(&buf, " 1 %s %s ", file.Owner(), file.Group())
		fmt.Fprint(&buf, lpad(strconv.FormatInt(file.Size(), 10), 12))
		fmt.Fprint(&buf, file.ModTime().Format(" Jan _2 15:04 "))
		fmt.Fprintf(&buf, "%s\r\n", file.Name())
	}
	return buf.Bytes()
}

func lpad(input string, length int) (result string) {
	if len(input) < length {
		result = strings.Repeat(" ", length-len(input)) + input
	} else if len(input) == length {
		result = input
	} else {
		result = input[0:length]
	}
	return
}

// RFC3659 returns a string that lists the collection of files
// according to RFC3659, one per line
func (formatter listFormatter) RFC3659() []byte {
	buf := &bytes.Buffer{}
	for _, file := range formatter {
		kind := "file"
		sizeField := "size"
		if file.Mode().IsDir() {
			switch file.Name() {
			case ".":
				kind = "cdir"
			case "..":
				kind = "pdir"
			default:
				kind = "dir"
			}
			sizeField = "sizd"
		}
		fmt.Fprintf(buf, "type=%s;", kind)
		fmt.Fprintf(buf, "%s=%d;", sizeField, file.Size())
		fmt.Fprintf(buf, "modify=%s;", file.ModTime().Format("20060102150405"))
		fmt.Fprintf(buf, "UNIX.mode=%04o;", file.Mode().Perm())
		if file.UID() != -1 {
			fmt.Fprintf(buf, "UNIX.uid=%d", file.UID())
		}
		if file.GID() != -1 {
			fmt.Fprintf(buf, "UNIX.gid=%d", file.GID())
		}
		fmt.Fprintf(buf, " %s\r\n", file.Name())
	}
	return buf.Bytes()
}
