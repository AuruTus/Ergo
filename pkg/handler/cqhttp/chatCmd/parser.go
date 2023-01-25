package chatcmd

import (
	"strings"
)

/*
	Raw command string always starts with '.', for example, ".help --verbose".
	The parser will fetch command and its options and aruguments without the
	leading characters ('.' for command and '-' or '--' for options).
*/

type optNode struct {
	opt string
	arg string
}

type cmdNode struct {
	cmd  string
	opts []optNode
	args []string
}

func initCmdNode(raw string) *cmdNode {
	c := &cmdNode{
		opts: make([]optNode, 0, 4),
		args: make([]string, 0, 1),
	}
	c.cmd = raw[1:]
	return c
}

func getValidOpt(raw string) string {
	start := 1
	if raw[start] == '-' {
		start = 2
	}
	return raw[start:]
}

func buildCmdNode(c *cmdNode, tokens []string) {
	const (
		ARG = iota
		OPT
	)

	pre := ARG
	restTs := tokens[1:]
	for _, t := range restTs {
		switch {
		// option: token starts with '-' or '--'
		case t[0] == '-':
			o := optNode{
				opt: getValidOpt(t),
			}
			c.opts = append(c.opts, o)
			pre = OPT
		// argument
		default:
			if pre == OPT {
				c.opts[len(c.opts)-1].arg = t
			} else {
				c.args = append(c.args, t)
			}
			pre = ARG
		}
	}
}

func CmdLexer(raw []byte) *cmdNode {
	// split tokens
	tokens := strings.FieldsFunc(string(raw), func(r rune) bool {
		switch r {
		case '\n', '\t', '\r', ' ':
			return true
		}
		return false
	})

	c := initCmdNode(tokens[0])
	buildCmdNode(c, tokens[1:])
	return c
}

func (c *cmdNode) Excute() []byte {
	return commands[c.cmd].handle(c)
}
