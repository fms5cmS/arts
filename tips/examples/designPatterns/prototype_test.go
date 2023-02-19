package designPatterns

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Keyword 搜索关键字
type Keyword struct {
	Word      string
	visit     int
	UpdatedAt *string
}

// Clone 这里使用序列化与反序列化的方式深拷贝
func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	b, _ := jsoniter.Marshal(k)
	jsoniter.Unmarshal(b, &newKeyword)
	return &newKeyword
}

func TestPrototype(t *testing.T) {
	now := time.Now().String()
	keyword := &Keyword{
		Word:      "original word",
		visit:     2,
		UpdatedAt: &now,
	}
	newKeyword := keyword.Clone()
	// 该字段未对外暴露，无法被序列化
	assert.NotEqual(t, keyword.visit, newKeyword.visit)

	assert.Equal(t, keyword.Word, newKeyword.Word)
	assert.Equal(t, keyword.UpdatedAt, newKeyword.UpdatedAt)
	t.Log(keyword.UpdatedAt)
	t.Log(newKeyword.UpdatedAt)
}
