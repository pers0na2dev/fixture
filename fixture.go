package fixture

import (
	"reflect"
	"unsafe"
)

type Fixture[T any] struct {
	value  T
	errors []error
}

type With struct {
	Name  string
	Value any
}

func NewFixture[T any](opts ...With) *Fixture[T] {
	var defaultValue T
	f := &Fixture[T]{value: defaultValue}

	for _, opt := range opts {
		f = with(f, opt.Name, opt.Value)
	}

	return f
}

func with[T any, V any](g *Fixture[T], name string, value V) *Fixture[T] {
	r := reflect.ValueOf(&g.value).Elem()
	f := r.FieldByName(name)

	if !f.IsValid() {
		g.errors = append(g.errors, &FieldError{FieldName: name, Message: "Field is not valid"})
		return g
	}

	if !reflect.TypeOf(value).AssignableTo(f.Type()) {
		g.errors = append(g.errors, &FieldError{FieldName: name, Message: "Type mismatch"})
		return g
	}

	if f.CanSet() {
		f.Set(reflect.ValueOf(value))
	} else {
		fieldPtr := unsafe.Pointer(f.UnsafeAddr())
		reflect.NewAt(f.Type(), fieldPtr).Elem().Set(reflect.ValueOf(value))
	}

	return g
}

func (g *Fixture[T]) With(fieldName string, value any) *Fixture[T] {
	return with(g, fieldName, value)
}

func (g *Fixture[T]) Build() (T, error) {
	if len(g.errors) > 0 {
		return g.value, g.errors[0]
	}

	return g.value, nil
}
