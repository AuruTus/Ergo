package chatcmd

const (
	defaultDesc = `Unknown command. See ".help" for more information.`
)

func defaultHandle(c *CmdNode) string {
	return defaultDesc
}
