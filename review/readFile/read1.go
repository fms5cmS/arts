package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read1(inputPath string, output io.Writer) error {
	type stats struct {
		min, max, sum float64
		count         int64
	}
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	stationStats := make(map[string]stats)
	// 读取数据行
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// 分隔数据，从而便于解析
		station, tempStr, hasSemi := strings.Cut(line, ";")
		if !hasSemi {
			continue
		}
		temp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			return err
		}
		// 每一行数据，都会对字符串执行两次哈希处理（见下），为了避免这种情况，可以使用 map[string]*stats，见 read2
		s, ok := stationStats[station] // 1. 从 map 中获取值
		if !ok {
			s.min, s.max, s.sum, s.count = temp, temp, temp, 1
		} else {
			s.min, s.max = min(s.min, temp), max(s.max, temp)
			s.sum += temp
			s.count++
		}
		stationStats[station] = s // 2. 更新 map
	}

	stations := make([]string, 0, len(stationStats))
	for station := range stationStats {
		stations = append(stations, station)
	}
	sort.Strings(stations)

	fmt.Fprint(output, "{")
	for i, station := range stations {
		if i > 0 {
			fmt.Fprint(output, ", ")
		}
		s := stationStats[station]
		mean := s.sum / float64(s.count)
		fmt.Fprintf(output, "%s=%.1f/%.1f/%.1f", station, s.min, mean, s.max)
	}
	fmt.Fprint(output, "}\n")
	return nil
}
