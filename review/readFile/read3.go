package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
)

// 针对 read2 中使用 strconv.ParseFloat 解析数据，根据实际情况下遇到的数据格式，调整解析的方法
// 这里示例的数据格式为：2 到 3 位数字（不超过 100，含小数），有些带负号
func read3(inputPath string, output io.Writer) error {
	type stats struct {
		min, max, sum float64
		count         int64
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
		station, tempBytes, hasSemi := bytes.Cut(line, []byte(";"))
		if !hasSemi {
			continue
		}
		// 根据实际数据格式调整转换函数！
		temp := parseBytes2Float(tempBytes)

		s := stationStats[string(station)]
		if s == nil {
			stationStats[string(station)] = &stats{
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

func parseBytes2Float(tempBytes []byte) float64 {
	negative := false
	index := 0
	if tempBytes[index] == '-' { // 判断数据的正负
		index++
		negative = true
	}
	temp := float64(tempBytes[index] - '0') // 解析第一位（从左往右）数据
	index++
	if tempBytes[index] != '.' { // 解析第二位数据
		temp = temp*10 + float64(tempBytes[index]-'0')
		index++
	}
	index++                                    // skip '.'
	temp += float64(tempBytes[index]-'0') / 10 // 解析小数位
	if negative {
		temp = -temp
	}
	return temp
}
