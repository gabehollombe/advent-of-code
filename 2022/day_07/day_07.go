package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
*/

type File struct {
	name string
	size int
}

type Dir struct {
	name      string
	parent    *Dir
	dirs      map[string]*Dir
	files     []File
	totalSize int
}

type Filesystem struct {
	currentDir *Dir
}

func (f *Filesystem) pushDir(dirName string) {
	if dir, ok := f.currentDir.dirs[dirName]; ok {
		// If the current dir already has dirName in it, just point to it
		f.currentDir = dir
	} else {
		// Otherwise, make it
		d := newDir(dirName, f.currentDir)
		// and add the dir to the parent's children
		f.currentDir.dirs[dirName] = d
		// and finally make it the current dir
		f.currentDir = d
	}

}

func (f *Filesystem) popDir() {
	f.currentDir = f.currentDir.parent
}

func (f *Filesystem) cdRoot() {
	for f.currentDir.name != "/" {
		f.currentDir = f.currentDir.parent
	}
}

func (f *Filesystem) addDir(name string) {
	f.currentDir.dirs[name] = newDir(name, f.currentDir)
}

func (f *Filesystem) addFile(file File) {
	f.currentDir.files = append(f.currentDir.files, file)
}

func (d Dir) size() int {
	fSum := 0
	for _, f := range d.files {
		fSum += f.size
	}
	if len(d.dirs) == 0 {
		return fSum
	} else {
		sum := fSum
		for _, d := range d.dirs {
			sum += d.size()
		}
		return sum
	}
}

func (d Dir) allDirsRecursive() []*Dir {
	dirs := make([]*Dir, 0)

	for _, mydir := range d.dirs {
		dirs = append(dirs, mydir)
	}

	for _, subdir := range d.dirs {
		for _, d2 := range subdir.allDirsRecursive() {
			dirs = append(dirs, d2)
		}
	}

	return dirs
}

func (f *Filesystem) totalSizeForDirsUnder(size int) int {
	f.cdRoot()
	dirs := f.currentDir.allDirsRecursive()

	sum := 0
	s := f.currentDir.size()
	if s < size {
		sum += s
	}
	for _, d := range dirs {
		s = d.size()
		if s < size {
			sum += s
		}
	}

	return sum
}

func (f *Filesystem) getDirToDelete(diskSize int, updateSize int) *Dir {
	f.cdRoot()
	totalSpaceLeftOnDisk := diskSize - f.currentDir.size()
	minToDelete := updateSize - totalSpaceLeftOnDisk

	bestCandidate := f.currentDir
	dirs := f.currentDir.allDirsRecursive()
	for _, d := range dirs {
		s := d.size()
		if s >= minToDelete && s < bestCandidate.size() {
			bestCandidate = d
		}
	}

	return bestCandidate
}

func newDir(name string, parent *Dir) *Dir {
	return &Dir{name: name, parent: parent, dirs: make(map[string]*Dir, 0), files: make([]File, 0)}
}

func newFile(name string, size int) File {
	return File{name: name, size: size}
}

func processCmd(filesystem *Filesystem, cmd string, args []string, output []string) {
	if cmd == "cd" {
		name := args[0]

		if name == "/" && filesystem.currentDir.name == "/" {
			// Don't push / if we are already in / (the default when we init the Filesystem)
			return
		}
		if name == ".." {
			// Pop the dir
			filesystem.popDir()
			return
		}
		filesystem.pushDir(name)
		return
	}

	if cmd == "ls" {
		for _, l := range output {
			parts := strings.Split(l, " ")
			if parts[0] == "dir" {
				// this is a dir
				filesystem.addDir(parts[1])
			} else {
				// this is a file
				size, _ := strconv.Atoi(parts[0])
				name := parts[1]
				f := newFile(name, size)
				filesystem.addFile(f)
			}
		}
	}
}

func loadInput(filename string) *Filesystem {
	content, _ := os.ReadFile(filename)
	lines := strings.Split(string(content), "\n")

	filesystem := &Filesystem{currentDir: newDir("/", nil)}
	cmd := ""
	var args []string
	var output []string
	for i, l := range lines {
		if l[0] == byte('$') {
			//if cmd != "" {
			if i != 0 {
				processCmd(filesystem, cmd, args, output)
				output = make([]string, 0)
			}
			parts := strings.Split(l, " ")
			cmd = parts[1]
			args = parts[2:]
		} else {
			output = append(output, l)
		}
	}
	processCmd(filesystem, cmd, args, output)

	return filesystem
}

func main() {
	fs := loadInput("input.txt")

	fmt.Printf("Part 1: %d\n", fs.totalSizeForDirsUnder(100000))

	diskSize := 70000000
	updateSize := 30000000
	dir := fs.getDirToDelete(diskSize, updateSize)
	fmt.Printf("Part 2: %d\n", dir.size())
}
