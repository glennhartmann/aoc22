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
	_, final := computeSizes(root)
	log.Printf("done recursively computing sizes")
	logFS(root)
	log.Printf("final answer: %d", final)
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

func computeSizes(root *inode) (recursiveSize int64, cumulativeRelevantSize int64) {
	for _, v := range root.children {
		if v.dir {
			rs, crs := computeSizes(v)
			cumulativeRelevantSize += crs
			recursiveSize += rs
		} else {
			recursiveSize += v.size
		}
	}
	root.size = recursiveSize
	if recursiveSize <= 100000 {
		log.Printf("adding %d (<= 10000) to crs (%d) to get %d (node %s)", recursiveSize, cumulativeRelevantSize, recursiveSize+cumulativeRelevantSize, root.name)
		cumulativeRelevantSize += recursiveSize
	}
	return recursiveSize, cumulativeRelevantSize
}
