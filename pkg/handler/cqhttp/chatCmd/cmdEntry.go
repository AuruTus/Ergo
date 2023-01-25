package chatcmd

import (
	"sync"
)

type handleFunc (func(*cmdNode) string)

type cmdEntry struct {
	handle handleFunc
	desc   string
	info   string
	cmds   []string
}

var (
	commands         map[string]*cmdEntry
	commandsInitOnce sync.Once
)

func OnCommand(cmdFunc handleFunc, desc, info string, cmds ...string) {
	// lazy assignment in case of the undefined behaviour of the init() excution order
	if commands == nil {
		commandsInitOnce.Do(func() { commands = make(map[string]*cmdEntry) })
	}
	e := &cmdEntry{
		handle: cmdFunc,
		desc:   desc,
		info:   info,
		cmds:   cmds,
	}
	for _, cmd := range cmds {
		commands[cmd] = e
	}
}
