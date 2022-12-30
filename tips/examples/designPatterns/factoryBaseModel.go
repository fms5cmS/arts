package designPatterns

import "fmt"

const (
	jsonFormat = iota
	yamlFormat
	xmlFormat
)

// ConfigFileParser 产品接口
type ConfigFileParser interface {
	Parse(data []byte)
}

/** 不同的产品实现 */

type jsonConfigParser struct{}

func (j jsonConfigParser) Parse(data []byte) {
	fmt.Println("JSON format config file: ", string(data))
}

type yamlConfigParser struct{}

func (y yamlConfigParser) Parse(data []byte) {
	fmt.Println("YAML format config file: ", string(data))
}

type xmlConfigParser struct{}

func (x xmlConfigParser) Parse(data []byte) {
	fmt.Println("XML format config file: ", string(data))
}
