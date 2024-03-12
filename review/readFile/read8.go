package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type r8Stats struct {
	min, max, sum float64
	count         int64
}

// 并行处理，数据处理和 read1 中一致
func read8(inputPath string, output io.Writer) error {
	// 分隔文件（并没有直接分割成多个文件，而是计算出每个子文件的偏移和长度）
	parts, err := splitFile(inputPath, maxGoroutines)
	if err != nil {
		return err
	}

	resultsCh := make(chan map[string]r8Stats)
	for _, part := range parts {
		go r8ProcessPart(inputPath, part.offset, part.size, resultsCh)
	}

	totals := make(map[string]r8Stats)
	for i := 0; i < len(parts); i++ {
		result := <-resultsCh
		for station, s := range result {
			ts, ok := totals[station]
			if !ok {
				totals[station] = r8Stats{
					min:   s.min,
					max:   s.max,
					sum:   s.sum,
					count: s.count,
				}
				continue
			}
			ts.min = min(ts.min, s.min)
			ts.max = max(ts.max, s.max)
			ts.sum += s.sum
			ts.count += s.count
			totals[station] = ts
		}
	}

	stations := make([]string, 0, len(totals))
	for station := range totals {
		stations = append(stations, station)
	}
	sort.Strings(stations)

	fmt.Fprint(output, "{")
	for i, station := range stations {
		if i > 0 {
			fmt.Fprint(output, ", ")
		}
		s := totals[station]
		mean := s.sum / float64(s.count)
		fmt.Fprintf(output, "%s=%.1f/%.1f/%.1f", station, s.min, mean, s.max)
	}
	fmt.Fprint(output, "}\n")
	return nil
}

func r8ProcessPart(inputPath string, fileOffset, fileSize int64, resultsCh chan map[string]r8Stats) {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 找到当前块起点（根据相对位置和偏移得到，这里根据文件原点 io.SeekStart）
	_, err = file.Seek(fileOffset, io.SeekStart)
	if err != nil {
		panic(err)
	}
	// 从文件源 R 读取数据，但会限制返回的数据量为 N 个字节
	f := io.LimitedReader{R: file, N: fileSize}

	stationStats := make(map[string]r8Stats)

	scanner := bufio.NewScanner(&f)
	for scanner.Scan() {
		line := scanner.Text()
		station, tempStr, hasSemi := strings.Cut(line, ";")
		if !hasSemi {
			continue
		}

		temp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			panic(err)
		}

		s, ok := stationStats[station]
		if !ok {
			s.min, s.max, s.sum, s.count = temp, temp, temp, 1
		} else {
			s.min, s.max = min(s.min, temp), max(s.max, temp)
			s.sum += temp
			s.count++
		}
		stationStats[station] = s
	}

	resultsCh <- stationStats
}

type part struct {
	offset, size int64
}

func splitFile(inputPath string, numParts int) ([]part, error) {
	const maxLineLength = 100

	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	st, err := f.Stat()
	if err != nil {
		return nil, err
	}
	size := st.Size()
	splitSize := size / int64(numParts)

	buf := make([]byte, maxLineLength)

	parts := make([]part, 0, numParts)
	offset := int64(0)
	for offset < size {
		seekOffset := max(offset+splitSize-maxLineLength, 0)
		if seekOffset > size {
			break
		}
		_, err := f.Seek(seekOffset, io.SeekStart)
		if err != nil {
			return nil, err
		}
		n, _ := io.ReadFull(f, buf)
		chunk := buf[:n]
		newline := bytes.LastIndexByte(chunk, '\n')
		if newline < 0 {
			return nil, fmt.Errorf("newline not found at offset %d", offset+splitSize-maxLineLength)
		}
		remaining := len(chunk) - newline - 1
		nextOffset := seekOffset + int64(len(chunk)) - int64(remaining)
		parts = append(parts, part{offset, nextOffset - offset})
		offset = nextOffset
	}
	return parts, nil
}
