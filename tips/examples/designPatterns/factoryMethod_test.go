package designPatterns

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// ConfigFileParserFactory 工厂接口
type ConfigFileParserFactory interface {
	CreateParser() ConfigFileParser
}

/** 不同的工厂实现 */

type jsonConfigFileParserFactory struct{}

func (j jsonConfigFileParserFactory) CreateParser() ConfigFileParser {
	return jsonConfigParser{}
}

type yamlConfigFileParserFactory struct{}

func (y yamlConfigFileParserFactory) CreateParser() ConfigFileParser {
	return yamlConfigParser{}
}

type xmlConfigFileParserFactory struct{}

func (x xmlConfigFileParserFactory) CreateParser() ConfigFileParser {
	return xmlConfigParser{}
}

// NewConfigFileParserFactory 工厂方法！
func NewConfigFileParserFactory(format int) ConfigFileParserFactory {
	switch format {
	case jsonFormat:
		return jsonConfigFileParserFactory{}
	case yamlFormat:
		return yamlConfigFileParserFactory{}
	case xmlFormat:
		return xmlConfigFileParserFactory{}
	default:
		return yamlConfigFileParserFactory{}
	}
}

func TestFactoryMethod(t *testing.T) {
	tests := []struct {
		name   string
		format int
		want   ConfigFileParserFactory
	}{
		{
			name:   "get JSON factory",
			format: jsonFormat,
			want:   jsonConfigFileParserFactory{},
		},
		{
			name:   "get YAML factory",
			format: yamlFormat,
			want:   yamlConfigFileParserFactory{},
		},
		{
			name:   "get XML factory",
			format: xmlFormat,
			want:   xmlConfigFileParserFactory{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			factory := NewConfigFileParserFactory(test.format)
			parser := factory.CreateParser()
			parser.Parse([]byte(test.name))
			assert.Equal(t, test.want, factory)
		})
	}
}
