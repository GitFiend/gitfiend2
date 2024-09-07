package option

type T[V any] struct {
	value *V
}

func New[V any](value *V) T[V] {
	return T[V]{value: value}
}

func (o *T[V]) Get() (*V, bool) {
	if o.value == nil {
		return nil, false
	}
	return o.value, true
}

func (o *T[V]) Set(value *V) {
	o.value = value
}
