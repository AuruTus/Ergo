package chatcmd

import (
	"strings"
)

/*
	Raw command string always starts with '.', for example, ".help --verbose".
	The parser will fetch command and its options and aruguments without the
	leading characters ('.' for command and '-' or '--' for options).
*/

type OptNode struct {
	Opt string
	Arg string
}

type CmdNode struct {
	Cmd  string
	Opts []OptNode
	Args []string
}

func initCmdNode(raw string) *CmdNode {
	c := &CmdNode{
		Opts: make([]OptNode, 0, 4),
		Args: make([]string, 0, 1),
	}
	c.Cmd = raw[1:]
	return c
}

func getValidOpt(raw string) string {
	start := 1
	if raw[start] == '-' {
		start = 2
	}
	return raw[start:]
}

func buildCmdNode(c *CmdNode, tokens []string) {
	const (
		ARG = iota
		OPT
	)

	pre := ARG
	for _, t := range tokens {
		switch {
		// option: token starts with '-' or '--'
		case t[0] == '-':
			o := OptNode{
				Opt: getValidOpt(t),
			}
			c.Opts = append(c.Opts, o)
			pre = OPT
		// argument
		default:
			if pre == OPT {
				c.Opts[len(c.Opts)-1].Arg = t
			} else {
				c.Args = append(c.Args, t)
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

func cmdParser(tokens []string) *CmdNode {
	c := initCmdNode(tokens[0])
	buildCmdNode(c, tokens[1:])
	return c
}

func Parse(raw string) *CmdNode {
	return cmdParser(cmdLexer(raw))
}

func (c *CmdNode) Excute() string {
	if h, ok := commands[c.Cmd]; ok {
		return h.handle(c)
	}
	return defaultHandle(c)
}
