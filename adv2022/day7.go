package adv2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day7() {
	type dir struct {
		size      int
		children  map[string]*dir
		parent    *dir
		totalSize int
	}
	mkDir := func(s string) *dir {
		return &dir{children: make(map[string]*dir)}
	}
	mkChild := func(d *dir, cName string) {
		_, found := d.children[cName]
		if !found {
			ch := mkDir(cName)
			ch.parent = d
			d.children[cName] = ch
		}
	}
	root := mkDir("/")
	cwd := root
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		items := strings.Fields(sc.Text())
		switch items[0] {
		case "$":
			switch items[1] {
			case "cd":
				switch items[2] {
				case "..":
					cwd = cwd.parent
				case "/":
					cwd = root
				default:
					cwd = cwd.children[items[2]]
				}
			case "ls":
				continue
			default:
				panic("  invalid command")
			}
		case "dir":
			mkChild(cwd, items[1])
		default:
			fSize, _ := strconv.Atoi(items[0])
			cwd.size += fSize
		}
	}
	sum := 0
	var f func(node *dir)
	f = func(node *dir) {
		node.totalSize = node.size
		for _, d := range node.children {
			f(d)
			node.totalSize += d.totalSize
		}
		if node.totalSize <= 100000 {
			sum += node.totalSize
		}
	}
	f(root)
	fmt.Println(sum)
	free := 70000000 - root.totalSize
	toDelete := 30000000 - free
	minDelSize := root.totalSize
	var g func(node *dir)
	g = func(node *dir) {
		if node.totalSize < toDelete {
			return
		}
		if node.totalSize >= toDelete && node.totalSize < minDelSize {
			minDelSize = node.totalSize

		}
		for _, d := range node.children {
			g(d)
		}
	}
	g(root)
	fmt.Println(minDelSize)
}
