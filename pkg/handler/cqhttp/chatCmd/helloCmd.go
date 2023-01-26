package chatcmd

func init() {
	OnCommand(helloHandle, helloDesc, helloInfo, "hello")
}

const (
	helloDesc = `"hello" is for comunication test`
	helloInfo = `
	===============================================
	  ______     ______     ______     ______    
	  /\  ___\   /\  == \   /\  ___\   /\  __ \   
	  \ \  __\   \ \  __<   \ \ \__ \  \ \ \/\ \  
	   \ \_____\  \ \_\ \_\  \ \_____\  \ \_____\ 
	    \/_____/   \/_/ /_/   \/_____/   \/_____/ 
												
	===============================================
	`
)

func helloHandle(c *CmdNode) string {
	return "Hello! This is Ergo"
}
