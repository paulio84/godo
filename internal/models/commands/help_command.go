package commands

type HelpCommand struct {
	parser func()
}

func NewHelpCommand(parser func()) HelpCommand {
	return HelpCommand{
		parser: parser,
	}
}

func (hc HelpCommand) Execute() {}

func (hc HelpCommand) ParseOutput() {
	hc.parser()
}
