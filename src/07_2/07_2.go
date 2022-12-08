package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type inode struct {
	dir      bool
	size     int64
	children map[string]*inode
	name     string
	parent   *inode
}

func mkDir(name string, parent *inode) *inode {
	return &inode{dir: true, children: make(map[string]*inode), name: name, parent: parent}
}

func mkFile(size int64, name string, parent *inode) *inode {
	return &inode{dir: false, size: size, name: name, parent: parent}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	root := mkDir("/", nil)
	cwd := root
	pwd := "/"
	skipRead := false
	var s string
outer:
	for {
		if !skipRead {
			var err error
			s, err = r.ReadString('\n')
			if err == io.EOF {
				log.Print("EOF")
				break
			}
			if err != nil {
				panic("unable to read")
			}
		}
		skipRead = false
		log.Printf("current line: %q", s)

		if strings.HasPrefix(s, "$ ls") {
			// this is not useful for us, just skip
			log.Print("ls - skipping")
			continue
		}

		if strings.HasPrefix(s, "$ cd") {
			arg := strings.TrimSpace(s[strings.LastIndex(s, " ")+1:])
			if arg == "/" {
				cwd = root
				pwd = "/"
			} else {
				if arg == ".." {
					cwd = cwd.parent
				} else {
					if _, ok := cwd.children[arg]; !ok {
						cwd.children[arg] = mkDir(arg, cwd)
					}
					cwd = cwd.children[arg]
				}
				pwd = filepath.Join(pwd, arg)
			}
			log.Printf("cd into %q", pwd)
			continue
		}

		// ls output
		// TODO
		for {
			log.Printf("current line is a dir listing: %q", s)

			if strings.HasPrefix(s, "dir") {
				arg := strings.TrimSpace(s[strings.LastIndex(s, " ")+1:])
				log.Printf("making new dir entry %q (%q)", arg, filepath.Join(pwd, arg))
				if _, ok := cwd.children[arg]; !ok {
					cwd.children[arg] = mkDir(arg, cwd)
				}
			} else {
				var size int64
				var name string
				n, err := fmt.Sscanf(s, "%d %s\n", &size, &name)
				if n != 2 || err != nil {
					panic("invalid input")
				}
				log.Printf("making new file entry %q (%q) - size %d", name, filepath.Join(pwd, name), size)
				cwd.children[name] = mkFile(size, name, cwd)
			}

			var err error
			s, err = r.ReadString('\n')
			if err == io.EOF {
				log.Print("EOF")
				break outer
			}
			if err != nil {
				panic("unable to read")
			}
			if strings.HasPrefix(s, "$") {
				log.Printf("current line (%q) is not a dir listing - falling back to outer loop", s)
				skipRead = true
				continue outer
			}
		}
	}
	log.Printf("done parsing input")
	logFS(root)
	computeSizes(root)
	log.Printf("done recursively computing sizes")
	logFS(root)

	freeSpace := 70000000 - root.size
	minToFree := 30000000 - freeSpace
	log.Printf("have %d free; need to free up at least %d more", freeSpace, minToFree)
	log.Printf("delete the dir of size %d", getSmallestBigEnough(root, minToFree))
}

// same format as the webside (but maybe different order because map)
func logFS(root *inode) {
	logFSInternal(root, 0)
}

func logFSInternal(root *inode, depth int) {
	logInode(root, depth)
	for _, v := range root.children {
		logFSInternal(v, depth+1)
	}
}

func logInode(n *inode, depth int) {
	var sb strings.Builder

	for i := 0; i < depth; i++ {
		sb.WriteString("  ")
	}
	sb.WriteString(fmt.Sprintf("- %s (", n.name))
	if n.dir && n.size > 0 {
		sb.WriteString(fmt.Sprintf("dir, recursive size=%d", n.size))
	} else if n.dir {
		sb.WriteString("dir")
	} else {
		sb.WriteString(fmt.Sprintf("file, size=%d", n.size))
	}
	sb.WriteString(")\n")

	log.Print(sb.String())
}

func computeSizes(root *inode) int64 {
	var recursiveSize int64
	for _, v := range root.children {
		if v.dir {
			recursiveSize += computeSizes(v)
		} else {
			recursiveSize += v.size
		}
	}
	root.size = recursiveSize
	return recursiveSize
}

func getSmallestBigEnough(root *inode, minToFree int64) int64 {
	smallestBigEnough := root.size
	for _, v := range root.children {
		if v.dir {
			smallestBigEnough = betterSize(smallestBigEnough, getSmallestBigEnough(v, minToFree), minToFree)
		}
	}
	return smallestBigEnough
}

func betterSize(a, b, min int64) int64 {
	if b < a && b > min {
		return b
	}
	return a
}
