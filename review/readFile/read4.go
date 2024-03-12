package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
)

// 由于实际的数据格式都是包含一位小数的，可以将所有数据转为整数，仅在最后返回时再转为浮点数
// 浮点指令的执行速度比整数指令慢，把浮点转换成整数可以提高性能
func read4(inputPath string, output io.Writer) error {
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
		station, tempBytes, hasSemi := bytes.Cut(line, []byte(";"))
		if !hasSemi {
			continue
		}
		// 将所有数据处理为整数！
		temp := parseBytes2Int(tempBytes)

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

	// 返回时，计算结果需要转为浮点数
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

func parseBytes2Int(tempBytes []byte) int32 {
	negative := false
	index := 0
	if tempBytes[index] == '-' { // 判断正负
		index++
		negative = true
	}
	temp := int32(tempBytes[index] - '0') // 解析第一位（从左往右）数字
	index++
	if tempBytes[index] != '.' { // 解析第二位
		temp = temp*10 + int32(tempBytes[index]-'0')
		index++
	}
	index++                                      // skip '.'
	temp = temp*10 + int32(tempBytes[index]-'0') // 解析小数位，但会将最终结果转为整数
	if negative {
		temp = -temp
	}
	return temp
}
