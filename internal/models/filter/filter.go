package filter

type TodoFilter int

const (
	All TodoFilter = iota
	OnlyCompleted
	NotCompleted
)
