package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
)

// 自定义哈希表，不再使用 Go map
func read7(inputPath string, output io.Writer) error {
	type stats struct {
		min, max, count int32
		sum             int64
	}
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// 自定义哈希表的元素类型定义
	type item struct {
		key  []byte
		stat *stats
	}
	const numBuckets = 1 << 17        // 预先桶的数量，不用编写逻辑来调整表的大小，2 的指数倍
	items := make([]item, numBuckets) // 哈希桶，采用开放寻址（线性探测）的方式来避免哈希冲突
	size := 0                         // 有效 item 数

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
			// 使用带有线性探测的 FNV-1a 哈希算法，如果发生冲突则使用下一空槽
			const (
				// FNV-1 64-bit constants from hash/fnv.
				offset64 = 14695981039346656037
				prime64  = 1099511628211
			)

			var station, after []byte
			hash := uint64(offset64)
			i := 0
			for ; i < len(chunk); i++ {
				c := chunk[i]
				if c == ';' {
					station = chunk[:i]
					after = chunk[i+1:]
					break
				}
				hash ^= uint64(c) // FNV-1a is XOR then *
				hash *= prime64
			}
			if i == len(chunk) {
				break
			}

			index := 0
			negative := false
			if after[index] == '-' {
				negative = true
				index++
			}
			temp := int32(after[index] - '0')
			index++
			if after[index] != '.' {
				temp = temp*10 + int32(after[index]-'0')
				index++
			}
			index++ // skip '.'
			temp = temp*10 + int32(after[index]-'0')
			index += 2 // skip last digit and '\n'
			if negative {
				temp = -temp
			}
			chunk = after[index:]

			hashIndex := int(hash & uint64(numBuckets-1))
			for {
				if items[hashIndex].key == nil {
					// Found empty slot, add new item (copying key).
					key := make([]byte, len(station))
					copy(key, station)
					items[hashIndex] = item{
						key: key,
						stat: &stats{
							min:   temp,
							max:   temp,
							sum:   int64(temp),
							count: 1,
						},
					}
					size++
					if size > numBuckets/2 {
						panic("too many items in hash table")
					}
					break
				}
				if bytes.Equal(items[hashIndex].key, station) {
					// Found matching slot, add to existing stats.
					s := items[hashIndex].stat
					s.min = min(s.min, temp)
					s.max = max(s.max, temp)
					s.sum += int64(temp)
					s.count++
					break
				}
				// Slot already holds another key, try next slot (linear probe).
				hashIndex++
				if hashIndex >= numBuckets {
					hashIndex = 0
				}
			}
		}

		readStart = copy(buf, remaining)
	}

	stationItems := make([]item, 0, size)
	for _, item := range items {
		if item.key == nil {
			continue
		}
		stationItems = append(stationItems, item)
	}
	sort.Slice(stationItems, func(i, j int) bool {
		return string(stationItems[i].key) < string(stationItems[j].key)
	})

	fmt.Fprint(output, "{")
	for i, item := range stationItems {
		if i > 0 {
			fmt.Fprint(output, ", ")
		}
		s := item.stat
		mean := float64(s.sum) / float64(s.count) / 10
		fmt.Fprintf(output, "%s=%.1f/%.1f/%.1f", item.key, float64(s.min)/10, mean, float64(s.max)/10)
	}
	fmt.Fprint(output, "}\n")
	return nil
}
