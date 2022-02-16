package showMeBug

import (
	"fmt"
	"strings"
	"testing"
)

// GenerateXmasTree 生成圣诞树
// height = 5
//     *
//    ***
//   *****
//  *******
// *********
// 用空格填充，使每行长度相同。最后一行只有星星，没有空格
func GenerateXmasTree(height int) string {
	maxLength := 2*height-1
	spaceNum := maxLength/2
	strBuild := strings.Builder{}
	for i := 1; i <= height; i++ {
		strBuild.WriteString(strings.Repeat(" ", spaceNum))
		strBuild.WriteString(strings.Repeat("*", 2*i-1))
		strBuild.WriteString(strings.Repeat(" ", spaceNum))
		if i != height {
			strBuild.WriteString("\n")
		}
		spaceNum--
	}
	return strBuild.String()
}

func TestGenerateXmasTree(t *testing.T) {
	fmt.Println(GenerateXmasTree(4))
}
