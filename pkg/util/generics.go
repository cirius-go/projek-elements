package util

// FirstOrDefault returns first non null value inside elems. If no item found,
// return default value instead.
func FirstOrDefault[T comparable](def T, elems ...T) T {
	var (
		zero T
	)

	for i := 0; i < len(elems); i++ {
		if elems[i] != zero {
			return elems[i]
		}
	}

	return def
}

// Something.
type Something struct{}

// SomethingHandler.
type SomethingHandler interface{}
