package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	start, err := makeSystem()
	if err != nil {
		log.Fatal(err)
	}

	rd1, err := round1(start)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Round 1: %d\n", rd1)

	rd2, err := round2(start)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Round 2: %d", rd2)

}

func round2(start *Folder) (int, error) {
	usedSize, err := start.GetSize()
	if err != nil {
		return 0, err
	}
	available := 70000000 - usedSize
	needed := 30000000 - available
	return FindMinFolderLarger(start, needed)
}

func round1(start *Folder) (int, error) {
	return DfsSizeUnder(start, 100000)
}

type Folder struct {
	Name    string
	Files   map[string]*File
	Parent  *Folder
	Folders map[string]*Folder
}

func NewFolder(name string, parent *Folder) *Folder {
	return &Folder{
		Name:    name,
		Parent:  parent,
		Files:   make(map[string]*File),
		Folders: make(map[string]*Folder),
	}
}

func (f *Folder) GetSize() (int, error) {
	total := 0
	for _, folder := range f.Folders {
		temp, err := folder.GetSize()
		if err != nil {
			return 0, err
		}
		total += temp
	}
	for _, file := range f.Files {
		total += file.Size
	}
	return total, nil
}

type File struct {
	Size   int
	Name   string
	Parent *Folder
}

func NewFile(name string, parent *Folder, size int) *File {
	return &File{
		Size:   size,
		Parent: parent,
		Name:   name,
	}
}

func DfsSizeUnder(head *Folder, max int) (int, error) {
	total := 0
	for _, folder := range head.Folders {
		temp, err := DfsSizeUnder(folder, max)
		if err != nil {
			return 0, err
		}
		total += temp
	}

	if size, _ := head.GetSize(); size <= max {
		total += size
	}
	return total, nil
}

func FindMinFolderLarger(head *Folder, min int) (int, error) {
	curr := int(^uint(0) >> 1)
	for _, folder := range head.Folders {
		temp, err := FindMinFolderLarger(folder, min)
		if err != nil {
			return 0, err
		}
		if temp >= min && temp < curr {
			curr = temp
		}
	}

	if size, _ := head.GetSize(); size >= min && size < curr {
		curr = size
	}

	return curr, nil
}

func makeSystem() (*Folder, error) {
	f, err := os.Open("in.txt")

	if err != nil {
		return nil, err
	}

	defer f.Close()

	start := NewFolder("/", nil)

	curr := start

	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i++ {
		text := lines[i]
		if strings.Contains(text, "$ cd") {
			folderName := text[5:]
			fmt.Println(folderName)
			if folderName == ".." {
				curr = curr.Parent
			} else {
				if folder := curr.Folders[folderName]; folder != nil {
					curr = folder
				} else {
					new := NewFolder(folderName, curr)
					curr.Folders[folderName] = new
					curr = new
				}
			}
		} else if text == "$ ls" {
			for j := i + 1; j < len(lines); j++ {
				line := lines[j]
				if strings.HasPrefix(line, "$") {
					i = j - 1
					break
				} else if strings.HasPrefix(line, "dir") {
					folderName := line[4:]
					if curr.Folders[folderName] == nil {
						folder := NewFolder(line[4:], curr)
						curr.Folders[folderName] = folder
					}
				} else {
					split := strings.Split(line, " ")
					size, _ := strconv.Atoi(split[0])
					name := split[1]
					if curr.Files[name] == nil {
						curr.Files[name] = NewFile(name, curr, size)
					}
				}
			}
		}
	}
	return start, nil
}
