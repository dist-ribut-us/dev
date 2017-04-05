package generator

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGenerator(t *testing.T) {
	generatorJson := `
	{
		"Package": "test",
		"Imports": ["foo","bar"],
		"TSMaps": [{
			"Key":"string",
			"Val":"int",
			"Name": "StringIntMap"
		}]
	}
	`
	g, err := Read(strings.NewReader(generatorJson))
	assert.NoError(t, err)
	if g == nil {
		t.Error("g should not be nil")
		return
	}
	gen := g.Generate()

	for _, part := range []string{
		"package test",
		"import (",
		"  \"foo\"",
		"  \"bar\"",
		"  \"sync\"",
		"func NewStringIntMap() *StringIntMap {",
	} {
		assert.Contains(t, gen, part)
	}
}
