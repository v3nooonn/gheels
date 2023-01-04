package constant

type ContextKey string

func (r ContextKey) String() string {
	return string(r)
}
