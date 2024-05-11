package parser

func Or[T any](parsers ...Parser[T]) Parser[T] {
	return func(in *Input) (T, bool) {
		for _, p := range parsers {
			res, ok := p(in)
			if ok {
				return res, true
			}
		}
		return *new(T), false
	}
}
