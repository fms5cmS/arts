package designPatterns

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// NewConfigFileParserBySimpleFactory 简单工厂模式
func NewConfigFileParserBySimpleFactory(format int) ConfigFileParser {
	switch format {
	case jsonFormat:
		return jsonConfigParser{} // 这里实际应调用类型的构造函数
	case yamlFormat:
		return yamlConfigParser{}
	case xmlFormat:
		return xmlConfigParser{}
	default:
		return yamlConfigParser{}
	}
}

func TestSimpleFactory(t *testing.T) {
	tests := []struct {
		name   string
		format int
		want   ConfigFileParser
	}{
		{
			name:   "get JSON Parser",
			format: jsonFormat,
			want:   jsonConfigParser{},
		},
		{
			name:   "get YAML Parser",
			format: yamlFormat,
			want:   yamlConfigParser{},
		},
		{
			name:   "get XML Parser",
			format: xmlFormat,
			want:   xmlConfigParser{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parser := NewConfigFileParserBySimpleFactory(test.format)
			parser.Parse([]byte(test.name))
			assert.Equal(t, test.want, parser)
		})
	}
}
