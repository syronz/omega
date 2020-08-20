package nodetype

// Node types
const (
	Online        = "online"
	Laptop        = "laptop"
	PC            = "PC"
	ServerPrivate = "server-private"
	ServerPublic  = "server-public"
)

// Types is list of account types for checking
var Types = []string{
	Online,
	Laptop,
	PC,
	ServerPrivate,
	ServerPublic,
}
