package name

import (
	"fmt"

	"github.com/botlabs-gg/yagpdb/commands"
	"github.com/jonas747/dcmd/v4"
)

var Command = &commands.YAGCommand{
	CmdCategory:     commands.CategoryDebug,
	Name:            "Name",
	Description:     "Temp command",
	LongDescription: "Temp command",

	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		return fmt.Sprintf("Your name is Ciphereck"), nil
	},
}
