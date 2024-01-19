package filter

type TodoFilter int

const (
	ALL TodoFilter = iota
	ONLY_COMPLETED
	NOT_COMPLETED
)

func (tf TodoFilter) IsValid() bool {
	switch tf {
	case ALL, ONLY_COMPLETED, NOT_COMPLETED:
		return true
	}

	return false
}
