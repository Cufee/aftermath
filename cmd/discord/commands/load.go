package commands

import "github.com/cufee/aftermath/cmd/discord/commands/builder"

type commandInit []builder.Builder

func (c *commandInit) add(cmd builder.Builder) {
	*c = append(*c, cmd)
}

func (c *commandInit) Compose() []builder.Command {
	var commands []builder.Command
	for _, builder := range *c {
		commands = append(commands, builder.Build())
	}
	return commands
}

var LoadedPublic commandInit
var LoadedInternal commandInit
