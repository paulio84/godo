package filter

type TodoFilter int

const (
	All TodoFilter = iota
	OnlyCompleted
	NotCompleted
)

func (tf TodoFilter) IsValid() bool {
	switch tf {
	case All, OnlyCompleted, NotCompleted:
		return true
	}

	return false
}
