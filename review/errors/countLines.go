package errors

import (
	"bufio"
	"io"
)

// CountLinesBefore 计算文件行数
func CountLinesBefore(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)

	for {
		_, err = br.ReadString('\n')
		lines++ // 在检查错误前增加行数
		// 因为有可能在遇到换行符之前先遇到了文件末尾 io.EOF
		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}

// Eliminate error handling by eliminating errors！

// CountLines 计算文件行数（使用 Scanner 改进）
func CountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	lines := 0
	for sc.Scan() {
		// 只有 scanner 缓冲区中有一行文本，才会自增行数
		// 这就意味着这个版本可以正确处理没有尾随换行符、文件为空的场景
		lines++
	}
	return lines, sc.Err()
}
