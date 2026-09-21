package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// node represents a folder in the tree.
type node struct {
	children map[string]*node
	files    int32 // files directly in this folder
}

// newNode creates a new empty node.
func newNode() *node {
	return &node{children: make(map[string]*node)}
}

// insert adds a path into the tree, creating nodes as needed.
// if the last segment contains a "." it is treated as a file, otherwise a folder.
func insert(root *node, parts []string) {
	curr := root
	for i, part := range parts {
		isLast := i == len(parts)-1
		isFile := isLast && strings.Contains(part, ".")
		if isFile {
			curr.files++
			return
		}
		if curr.children[part] == nil {
			curr.children[part] = newNode()
		}
		curr = curr.children[part]
	}
}

// maxDirectChildren recursively finds the max number of direct children any node has.
func maxDirectChildren(n *node) int32 {
	max := int32(len(n.children))
	for _, child := range n.children {
		if v := maxDirectChildren(child); v > max {
			max = v
		}
	}
	return max
}

// countFiles returns the total number of files under n (recursively),
// and updates maxFiles if this count exceeds it.
func countFiles(n *node, maxFiles *int32) int32 {
	total := n.files
	for _, child := range n.children {
		total += countFiles(child, maxFiles)
	}
	if total > *maxFiles {
		*maxFiles = total
	}
	return total
}

func solve(paths []string) []int32 {
	root := newNode()

	for _, path := range paths {
		var parts []string

		// windows-style path: strip drive prefix (e.g. "C:\")
		if idx := strings.Index(path, `\`); idx != -1 {
			parts = strings.Split(path[idx+1:], `\`)
		} else {
			parts = strings.Split(path, `/`)
		}

		insert(root, parts)
	}

	// max direct subfolders any single folder has
	maxSubfolders := maxDirectChildren(root)

	// max files in any folder (recursively)
	var maxFiles int32
	countFiles(root, &maxFiles)

	return []int32{maxSubfolders, maxFiles}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	var paths []string
	for {
		line, err := readLine(reader)
		if line != "" {
			paths = append(paths, line)
		}
		if err == io.EOF {
			break
		}
	}

	result := solve(paths)
	fmt.Println(result[0])
	fmt.Println(result[1])
}

func readLine(reader *bufio.Reader) (string, error) {
	str, _, err := reader.ReadLine()
	return strings.TrimRight(string(str), "\r\n"), err
}
