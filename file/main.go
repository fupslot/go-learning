// reads the content of a file and prints its content to the stdout
// Notes:
//   - The file nto open should be provider as an argument os.Args
//   -  use io.Copy function
//   - "go run main.go in.txt"
package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"
)

var usageStr = `

Usage "go run main.go in.txt"
`

type FileResource struct {
	name  string
	umask string
	owner string
	group string
	perm  os.FileMode
	octal string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usageStr)
		os.Exit(1)
	}

	fileName := os.Args[1]

	fmt.Printf("%s\n", filepath.Base(fileName))

	if info, err := os.Stat(fileName); err == nil {
		// fmt.Printf("%d\n", int(info.Mode()))
		// fmt.Printf("%o", info)

		stat := info.Sys().(*syscall.Stat_t)
		uid := stat.Uid
		gid := stat.Gid

		// Retriving file's User and Group
		u, _ := user.LookupId(strconv.FormatUint(uint64(uid), 10))
		g, _ := user.LookupGroupId(strconv.FormatUint(uint64(gid), 10))

		umask := info.Mode().Perm().String()
		octal := strconv.FormatUint(uint64(info.Mode().Perm()), 8)

		fmt.Printf("%s %s %s %s %s\n", fileName, umask, octal, u.Username, g.Name)

		fileResource := &FileResource{
			name:  fileName,
			umask: umask,
			octal: octal,
			perm:  info.Mode().Perm(),
			owner: u.Username,
			group: g.Name,
		}

		fmt.Printf("%+v", fileResource)

	} else if os.IsNotExist(err) {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// f, err := os.Open(fileName)
	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	os.Exit(1)
	// }

	// defer f.Close()

	// io.Copy(os.Stdin, f)
}
