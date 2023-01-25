package chatcmd

import (
	"testing"
)

func _compareArray[T interface{ ~string | optNode }](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, n := 0, len(a); i < n; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func _compareCNode(a, b cmdNode) bool {
	return a.cmd == b.cmd && _compareArray(a.opts, b.opts) &&
		_compareArray(a.args, b.args)
}

func TestPaser(t *testing.T) {
	type tArg struct {
		raw      string
		expected cmdNode
	}

	args := []tArg{
		{
			raw: ".hello -i tom tom",
			expected: cmdNode{
				cmd: "hello",
				opts: []optNode{
					{"i", "tom"},
				},
				args: []string{"tom"},
			},
		},
		{
			raw: ".hello -i tom --verbosetom",
			expected: cmdNode{
				cmd: "hello",
				opts: []optNode{
					{"i", "tom"},
					{"verbosetom", ""},
				},
				args: []string{},
			},
		},
		{
			raw: ".hello -i tom --verbose tom jack Da大Mi明ng 力LiNing[CQ:face,id=1,name=2]",
			expected: cmdNode{
				cmd: "hello",
				opts: []optNode{
					{"i", "tom"},
					{"verbose", "tom"},
				},
				args: []string{"jack", "Da大Mi明ng", "力LiNing[CQ:face,id=1,name=2]"},
			},
		},
	}

	for i, a := range args {
		nd := Parse(a.raw)
		t.Logf("result: %+v\n", nd)
		if !_compareCNode(a.expected, *nd) {
			t.Errorf("case %d failed %+v\n", i+1, nd)
		}
	}
}
