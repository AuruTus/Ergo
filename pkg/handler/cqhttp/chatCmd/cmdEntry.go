package chatcmd

import "sync"

type cmdEntry struct {
	handle func(*cmdNode) []byte
	desc   string
	info   string
}

var (
	commands         map[string]*cmdEntry
	commandsInitOnce sync.Once
)

func OnCommand(cmdFunc func(*cmdNode) []byte, desc, info string, cmds ...string) {
	// lazy assignment in case of the undefined behaviour of the init() excution order
	if commands == nil {
		commandsInitOnce.Do(func() { commands = make(map[string]*cmdEntry) })
	}
	e := &cmdEntry{
		handle: cmdFunc,
		desc:   desc,
		info:   info,
	}
	for _, cmd := range cmds {
		commands[cmd] = e
	}
}
