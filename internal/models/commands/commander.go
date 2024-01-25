package commands

type Commander interface {
	Execute()
	ParseOutput()
}
