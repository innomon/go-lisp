package lisp

// AB: 19-May-17

import (
	"reflect"
)

// Functor Extension functor
type Functor func(v []reflect.Value) (Value, error)

var registry = make(map[string]Functor)

// RegisterFunctor  Register a Functor name
func RegisterFunctor(name string, ftor Functor) {
	registry[name] = ftor
}
