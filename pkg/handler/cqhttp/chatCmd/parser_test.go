package chatcmd

import (
	"testing"
)

func _compareArray[T ~string | OptNode](a, b []T) bool {
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

const _s = "test-string-for-unpacking"

func _appendString() []byte {
	b := make([]byte, 0)
	for range [100]struct{}{} {
		b = append(b, _s...)
	}
	return b
}

func _appendByteSlice() []byte {
	b := make([]byte, 0)
	for range [100]struct{}{} {
		b = append(b, []byte(_s)...)
	}
	return b
}

// This pattern is a tiny bit faster.
func BenchmarkStringUnpack1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_appendString()
	}
}

func BenchmarkStringUnpack2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_appendByteSlice()
	}
}

const _mapRangeTime = 10

func _initEmptyMap() any {
	m := make(map[int]struct{}, _mapRangeTime)
	// m := make(map[int]struct{})
	// for i := range [_mapRangeTime]struct{}{} {
	// 	m[i] = struct{}{}
	// }
	return m
}

func _initIntMap() any {
	m := make(map[int]int, _mapRangeTime)
	// m := make(map[int]int)
	// for i := range [_mapRangeTime]struct{}{} {
	// 	m[i] = i
	// }
	return m
}

// This pattern indeed has a smaller mem usage.
func BenchmarkEmptyStructSize1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_initEmptyMap()
	}
}

func BenchmarkEmptyStructSize2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_initIntMap()
	}
}

const _forRangeTime = 100000

func _forEmptyStructRange() {
	a, b := 1, 2
	for range [_forRangeTime]struct{}{} {
		a ^= b
		b &= a
	}
}

func _forIntRange() {
	a, b := 1, 2
	for range [_forRangeTime]int{} {
		a ^= b
		b &= a
	}
}

func _forCounterLoop() {
	a, b := 1, 2
	for i := 0; i < _forRangeTime; i++ {
		a ^= b
		b &= a
	}
}

// No evident differences between them, both time and mem.
// Pick preferred one.
func BenchmarkForRange1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_forEmptyStructRange()
	}
}

func BenchmarkForRange2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_forIntRange()
	}
}

func BenchmarkForRange3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_forCounterLoop()
	}
}
