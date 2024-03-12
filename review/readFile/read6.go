package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
)

// 去掉 bufio.Scanner
func read6(inputPath string, output io.Writer) error {
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

	// 分配了一个 1MB 的缓冲区来读取大块文件，查找块中最后一个换行符来确保不会把单行截断，之后再处理这些单个块
	buf := make([]byte, 1024*1024)
	readStart := 0
	for {
		n, err := f.Read(buf[readStart:])
		if err != nil && err != io.EOF {
			return err
		}
		if readStart+n == 0 {
			break
		}
		chunk := buf[:readStart+n]

		newline := bytes.LastIndexByte(chunk, '\n')
		if newline < 0 {
			break
		}
		remaining := chunk[newline+1:]
		chunk = chunk[:newline+1]

		for {
			station, after, hasSemi := bytes.Cut(chunk, []byte(";"))
			if !hasSemi {
				break
			}
			var temp int32
			temp, chunk = parseBytes2Int2(after)
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
		readStart = copy(buf, remaining)
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

func parseBytes2Int2(after []byte) (int32, []byte) {
	negative := false
	index := 0
	if after[index] == '-' { // 判断正负
		index++
		negative = true
	}
	temp := int32(after[index] - '0') // 解析第一位（从左往右）数字
	index++
	if after[index] != '.' { // 解析第二位
		temp = temp*10 + int32(after[index]-'0')
		index++
	}
	index++                                  // skip '.'
	temp = temp*10 + int32(after[index]-'0') // 解析小数位，但会将最终结果转为整数
	index += 2
	if negative {
		temp = -temp
	}
	chunk := after[index:]
	return temp, chunk
}
