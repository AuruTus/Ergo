package tools

import (
	"fmt"
	"testing"
)

func TestKeyGen(t *testing.T) {
	type testArg struct {
		tokens   []any
		expected string
	}

	i := 321883312
	ip := &i
	args := []testArg{
		{
			tokens:   []any{"wsd呜dggrt中", i},
			expected: fmt.Sprintf("%s-%d", "wsd呜dggrt中", i),
		},
		{
			tokens:   []any{"abc", ip, i},
			expected: fmt.Sprintf("%s-%p-%d", "abc", ip, i),
		},
		{
			tokens:   []any{ip},
			expected: fmt.Sprintf("%p", ip),
		},
		{
			tokens:   []any{"abc", -1234},
			expected: fmt.Sprintf("%s-m%d", "abc", 1234),
		},
	}

	for i, a := range args {
		k := KeyGen(a.tokens...)
		t.Logf("key: %s\n", k)
		if k != a.expected {
			t.Errorf("case %d key unmatched: %v, %v\n", i, k, a.tokens)
		}
	}
}
