package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	TotalDiskSpace = 70000000
	DiskRequired   = 30000000
	MaxDirSize     = 100000
	RootFolder     = "/"
	CommandSign    = "$"
	CdCommand      = "cd"
	DirCommand     = "dir"
	LsCommand      = "ls"
	UpperDir       = ".."
)

var totalSum int
var dirSizes []int

type File struct {
	IsDir    bool
	Name     string
	Parent   *File
	Children []*File
	Size     int
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("error converting string to int: %v", err)
		os.Exit(1)
	}
	return i
}

func processFile(commandLine []string, curDir *File) {
	if strings.HasPrefix(commandLine[0], DirCommand) {
		dir := &File{
			IsDir:    true,
			Name:     commandLine[1],
			Parent:   curDir,
			Children: []*File{},
			Size:     -1,
		}
		curDir.Children = append(curDir.Children, dir)
	} else {
		f := &File{
			IsDir:    false,
			Name:     commandLine[1],
			Parent:   curDir,
			Children: nil,
			Size:     stringToInt(commandLine[0]),
		}
		curDir.Children = append(curDir.Children, f)
	}
}

func cd(commandLine []string, curDir *File) *File {
	dirName := commandLine[2]
	if dirName == UpperDir {
		curDir = curDir.Parent
	} else {
		for _, f := range curDir.Children {
			if f.Name == dirName {
				curDir = f
				break
			}
		}
	}
	return curDir
}

func iterateDir(f *File) int {
	var sum int
	for _, file := range f.Children {
		if file.IsDir == false {
			sum += file.Size
		} else {
			sum += iterateDir(file)
		}
	}

	fmt.Printf("Total size of \"%s\" directory: %d\n", buildPath(f), sum)
	dirSizes = append(dirSizes, sum)
	if sum < MaxDirSize {
		totalSum += sum
	}
	return sum
}

func buildPath(f *File) string {
	if f != nil {
		path := f.Name
		if f.Parent != nil {
			if strings.HasSuffix(path, RootFolder) {
				path = buildPath(f.Parent) + path
			} else {
				path = buildPath(f.Parent) + path + "/"
			}
		}
		return path
	}
	return ""
}

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	root := File{true, RootFolder, nil, []*File{}, 0}
	// assuming current dir is "/"
	curDir := &root
	for i, line := range strings.Split(string(file), "\n") {
		if line == "" || i == 0 {
			continue
		}
		fmt.Printf("curDir: %s\n", buildPath(curDir))
		fmt.Println(line)
		commandLine := strings.Split(line, " ")
		if strings.HasPrefix(line, CommandSign) {
			command := commandLine[1]
			switch command {
			case LsCommand:
				// do nothing
			case CdCommand:
				curDir = cd(commandLine, curDir)
			}
		} else {
			processFile(commandLine, curDir)
		}
	}
	totalUsed := iterateDir(&root)
	// Part 1
	fmt.Println("-------------------")
	fmt.Printf("Part 1: %d\n", totalSum)

	// Part 2
	fmt.Println("-------------------")
	totalUnused := TotalDiskSpace - totalUsed
	fmt.Printf("Currently free: %d\n", totalUnused)
	spaceNeeded := DiskRequired - totalUnused
	fmt.Printf("Space needed: %d\n", spaceNeeded)
	sort.Ints(dirSizes)
	for _, size := range dirSizes {
		if size >= spaceNeeded {
			fmt.Printf("Part 2. Size of the folder to remove: %d\n", size)
			break
		}
	}
}
