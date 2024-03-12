
Go 读取大文件（含/不含换行符），见 [readFile](../../review/readFile/README.md)

文件分隔（并未实际生成子文件，而是计算出每个子文件的偏移、大小），见 [splitFile](../../review/readFile/read8.go)

在上面不实际生成子文件的分隔前提下，**仅读取每个子文件的内容**！见 [r8ProcessPart](../../review/readFile/read8.go) 前面部分
