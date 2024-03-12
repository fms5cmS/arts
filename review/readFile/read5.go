package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

// 根据实际的数据格式，优化其解析方法
// 这里的实际数据格式为包含 1 位小数，含小数点在内最长不超过 5 位，如 32.7、-32.7 的数字
func read5(inputPath string, output io.Writer) error {
	type stats struct {
		min, max, count int32 // 整数类型
		sum             int64
	}

	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	stationStats := make(map[string]*stats)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()
		station, temp := parseLine(line)

		s := stationStats[string(station)]
		if s == nil {
			stationStats[string(station)] = &stats{
				min:   temp,
				max:   temp,
				sum:   int64(temp),
				count: 1,
			}
		} else {
			s.min = min(s.min, temp)
			s.max = max(s.max, temp)
			s.sum += int64(temp)
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
		mean := float64(s.sum) / float64(s.count) / 10
		fmt.Fprintf(output, "%s=%.1f/%.1f/%.1f", station, float64(s.min)/10, mean, float64(s.max)/10)
	}
	fmt.Fprint(output, "}\n")
	return nil
}

// 从后往前查找 ; 来解析温度的速度更快
func parseLine(line []byte) (station []byte, temp int32) {
	end := len(line)
	tenths := int32(line[end-1] - '0')
	ones := int32(line[end-3] - '0') // line[end-2] is '.'
	var semicolon int
	if line[end-4] == ';' { // >10 的正数
		temp = ones*10 + tenths
		semicolon = end - 4
	} else if line[end-4] == '-' { // <-10 的负数
		temp = -(ones*10 + tenths)
		semicolon = end - 5
	} else {
		tens := int32(line[end-4] - '0')
		if line[end-5] == ';' {
			temp = tens*100 + ones*10 + tenths
			semicolon = end - 5
		} else { // '-'
			temp = -(tens*100 + ones*10 + tenths)
			semicolon = end - 6
		}
	}
	station = line[:semicolon]
	return
}
