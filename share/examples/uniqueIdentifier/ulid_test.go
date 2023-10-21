package uniqueIdentifier

import (
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
	"testing"
	"time"
)

const ulidLen = 26

func TestULID(t *testing.T) {
	// 仅生成 ulid，而不关注性能、密码安全等问题
	id1 := ulid.Make()
	t.Log("ulid1:", id1.String())
	assert.Len(t, id1.String(), ulidLen)

	// 提供熵源时要小心，math/rand.Rand 是非并发安全的，可以使用 x/exp/rand 包替代
	// 对安全敏感的应用应始终使用 crypto/rand 提供的加密安全熵
	// 对性能敏感的应用在生成 ID 时应避免同步：
	// 如可以每个并发 goroutine 使用一个唯一的源，这样就没有了锁竞争，但无法保证随机数据的随机性以及给定毫秒内的单调性
	// 更通用的性能优化是使用 sync.Pool 来作为熵源池
	entropy := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	ms := ulid.Timestamp(time.Now())
	id2, _ := ulid.New(ms, entropy)
	t.Log("ulid2:", id2.String())
	assert.Len(t, id2.String(), ulidLen)
}
