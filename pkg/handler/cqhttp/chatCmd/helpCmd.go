package chatcmd

func init() {
	OnCommand(helpHandle, helpDesc, helpInfo, "h", "help")
}

const (
	helpDesc = `"help" lists available commands on this bot. Use "-v", "--verbose" to get more infomation.`
	helpInfo = `-v --verbose: list the detailed information of available commands`
)

var (
	descLine = []byte("@ DESC> ")
	infoLine = []byte("\n@ INFO-->\n")
)

func helpHandle(c *cmdNode) string {
	msg := make([]byte, 0, 512)

	verboseFlag := false
	for _, o := range c.opts {
		switch {
		case o.opt == "v" || o.opt == "verbose":
			verboseFlag = true
		}
	}

	visited := make(map[*handleFunc]struct{})
	for _, c := range commands {
		if _, ok := visited[&c.handle]; ok {
			continue
		}
		visited[&c.handle] = struct{}{}

		// get segment string like: "@ DESC-> .h .help: "
		cmds := append([]byte{}, descLine...)
		for _, name := range c.cmds {
			cmds = append(cmds, '.')
			cmds = append(cmds, []byte(name)...)
			cmds = append(cmds, ' ')
		}
		cmds[len(cmds)-1] = ':'
		cmds = append(cmds, ' ')

		cmdBlock := []byte(c.desc)
		if verboseFlag {
			cmdBlock = append(cmdBlock, infoLine...)
			cmdBlock = append(cmdBlock, []byte(c.info)...)
		}
		cmdBlock = append(cmdBlock, '\n')

		msg = append(msg, cmds...)
		msg = append(msg, cmdBlock...)
	}
	return string(msg)
}
