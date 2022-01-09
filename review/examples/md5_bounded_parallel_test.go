package examples

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"sort"
	"sync"
	"testing"
)

func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errCh := make(chan error, 1)
	go func() {
		defer close(paths)
		// No select needed here, since errCh is buffered.
		errCh <- filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return paths, errCh
}

// digester 不会关闭其输出通道，因为多个 goroutine 在共享通道上发送。
func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {
		file, err := ioutil.ReadFile(path)
		select {
		case c <- result{path: path, sum: md5.Sum(file), err: err}:
		case <-done:
			return
		}
	}
}

func md5AllBounded(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	paths, errCh := walkFiles(done, root)
	// Start a fixed number of goroutines to read and digest files.
	c := make(chan result)
	var wg sync.WaitGroup
	const numDigesters = 2
	wg.Add(numDigesters)
	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, paths, c)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()

	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}
	// Check whether the Walk failed.
	if err := <-errCh; err != nil {
		return nil, err
	}
	return m, nil
}

func TestMD5BoundedParallel(t *testing.T) {
	m, err := md5AllBounded("./")
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}
