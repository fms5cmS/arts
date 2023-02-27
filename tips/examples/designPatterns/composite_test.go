package designPatterns

import (
	"fmt"
	"testing"
)

// 定义接口
type Component interface {
	search(string)
}

// 单个对象
type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}

// 单个对象，也是组合对象
type Folder struct {
	components []Component // 组合对象可以包含其他对象，而这些对象都有相同的接口
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}

func TestComposite(t *testing.T) {
	file1 := &File{name: "file1"}
	file2 := &File{name: "file2"}
	file3 := &File{name: "file3"}
	folder1 := &Folder{name: "folder1"}
	folder1.add(file1)

	folder2 := &Folder{name: "folder2"}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("file1")
}
