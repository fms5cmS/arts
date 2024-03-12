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

// 针对 read1 中对 map 的两次操作，改变 value 的类型，从而优化为一次操作
func read2(inputPath string, output io.Writer) error {
	type stats struct {
		min, max, sum float64
		count         int64
	}
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()
	// 调整了 map 的类型
	stationStats := make(map[string]*stats)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		station, tempStr, hasSemi := strings.Cut(line, ";")
		if !hasSemi {
			continue
		}
		temp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			return err
		}
		// 处理 value 类型为指针的 map
		s := stationStats[station]
		if s == nil {
			stationStats[station] = &stats{
				min:   temp,
				max:   temp,
				sum:   temp,
				count: 1,
			}
		} else {
			s.min, s.max = min(s.min, temp), max(s.max, temp)
			s.sum += temp
			s.count++
		}
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
