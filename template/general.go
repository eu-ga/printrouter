package template

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/helper"
)

func LineSeparator(commands []command.PrinterCommand) []command.PrinterCommand {
	commands = append(commands, command.NewLine{})
	commands = append(commands, command.FontB)
	commands = append(commands, command.DashedLine{Font: command.FontB})
	commands = append(commands, command.NewLine{})
	return commands
}

func Footer(commands []command.PrinterCommand) []command.PrinterCommand {
	commands = append(commands, command.NewLine{})
	commands = append(commands, command.NewLine{})
	commands = append(commands, command.Text(helper.Center("Thank You!", " ", 38)))
	commands = append(commands, command.NewLine{})
	commands = append(commands, command.Text(helper.Center("Powered by Rockspoon", " ", 38)))
	commands = append(commands, command.NewLine{})
	commands = append(commands, command.Text(helper.Center("www.rockspoon.com", " ", 38)))
	commands = append(commands, command.NewLine{})
	return commands
}
