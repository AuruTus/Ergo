package chatcmd

import (
	"testing"
)

func _compareArray[T interface{ ~string | OptNode }](a, b []T) bool {
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

func _compareCNode(a, b CmdNode) bool {
	return a.Cmd == b.Cmd && _compareArray(a.Opts, b.Opts) &&
		_compareArray(a.Args, b.Args)
}

func TestPaser(t *testing.T) {
	type tArg struct {
		raw      string
		expected CmdNode
	}

	args := []tArg{
		{
			raw: ".hello -i tom tom",
			expected: CmdNode{
				Cmd: "hello",
				Opts: []OptNode{
					{"i", "tom"},
				},
				Args: []string{"tom"},
			},
		},
		{
			raw: ".hello -i tom --verbosetom",
			expected: CmdNode{
				Cmd: "hello",
				Opts: []OptNode{
					{"i", "tom"},
					{"verbosetom", ""},
				},
				Args: []string{},
			},
		},
		{
			raw: ".hello -i tom --verbose tom jack Da大Mi明ng 力LiNing[CQ:face,id=1,name=2]",
			expected: CmdNode{
				Cmd: "hello",
				Opts: []OptNode{
					{"i", "tom"},
					{"verbose", "tom"},
				},
				Args: []string{"jack", "Da大Mi明ng", "力LiNing[CQ:face,id=1,name=2]"},
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
