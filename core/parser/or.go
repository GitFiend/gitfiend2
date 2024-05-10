package parser

func Or[T any](parsers ...Parser[T]) Parser[T] {
	return func(in *Input) Result[T] {
		for _, p := range parsers {
			res := p(in)
			if !res.Failed {
				return res
			}
		}
		return Result[T]{Failed: true}
	}
}
