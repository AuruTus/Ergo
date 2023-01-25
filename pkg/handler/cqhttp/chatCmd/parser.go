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
	for _, t := range tokens {
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

func cmdLexer(raw string) []string {
	tokens := strings.FieldsFunc(raw, func(r rune) bool {
		switch r {
		case '\n', '\t', '\r', ' ':
			return true
		}
		return false
	})
	return tokens
}

func cmdParser(tokens []string) *cmdNode {
	c := initCmdNode(tokens[0])
	buildCmdNode(c, tokens[1:])
	return c
}

func Parse(raw string) *cmdNode {
	return cmdParser(cmdLexer(raw))
}

func (c *cmdNode) Excute() string {
	if h, ok := commands[c.cmd]; ok {
		return h.handle(c)
	}
	return defaultHandle(c)
}
