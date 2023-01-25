package chatcmd

func init() {
	OnCommand(helpHandle, helpDesc, helpInfo, "h", "help")
}

var (
	helpDesc = "`help` lists available commands on this bot"
	helpInfo = `
		-v --verbose: list the detailed information of available commands
	`
)

func helpHandle(c *cmdNode) []byte {
	msg := make([]byte, 0, 128)

	var (
		verboseFlag = false
	)

	for _, o := range c.opts {
		switch {
		case o.opt == "-v" || o.opt == "--verbose":
			verboseFlag = true
		}
	}

	for _, cmd := range commands {
		cLine := []byte(cmd.desc)
		if verboseFlag {
			cLine = append(cLine, '\n')
			cLine = append(cLine, []byte(cmd.info)...)
		}
		cLine = append(cLine, '\n')
		msg = append(msg, cLine...)
	}
	return msg
}
