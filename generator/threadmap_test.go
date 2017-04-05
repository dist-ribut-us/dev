package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreadSafeMapGenerate(t *testing.T) {
	mp := &ThreadsafeMap{
		Key:  "string",
		Val:  "int",
		Name: "StringIntMap",
	}
	gen := mp.Generate()

	for _, part := range []string{
		"Map map[string]int",
		"func NewStringIntMap() *StringIntMap {",
		"func (t *StringIntMap) Get(key string) (int, bool) {",
	} {
		assert.Contains(t, gen, part)
	}
}
