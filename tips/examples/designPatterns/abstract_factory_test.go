package designPatterns

import (
	"fmt"
	"testing"
)

type ConfigFileTemplateCreator interface {
	Create(filename string) error
}

type jsonConfigFileTemplateCreator struct {
	template []byte
}

func (j jsonConfigFileTemplateCreator) Create(filename string) error {
	fmt.Println("JSON config file name is ", filename)
	return nil
}

// ConfigFileAbstractFactory 抽象工厂接口
type ConfigFileAbstractFactory interface {
	CreateParser() ConfigFileParser
	CreateTemplateCreator() ConfigFileTemplateCreator
}

type jsonConfigFile struct {
	template []byte
}

func (j jsonConfigFile) CreateParser() ConfigFileParser {
	return jsonConfigParser{}
}

func (j jsonConfigFile) CreateTemplateCreator() ConfigFileTemplateCreator {
	return jsonConfigFileTemplateCreator{}
}

func Test_jsonConfigParser_Parse(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name     string
		template []byte
		data     []byte
		filename string
	}{
		{
			name:     "json",
			template: []byte("test"),
			data:     []byte("test data"),
			filename: "testFile",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := jsonConfigFile{template: tt.template}
			parser := c.CreateParser()
			parser.Parse(tt.data)
			creator := c.CreateTemplateCreator()
			creator.Create(tt.filename)
		})
	}
}
