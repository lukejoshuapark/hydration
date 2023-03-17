package hydration

import (
	"fmt"
	"reflect"
)

type Registry struct {
	typeMapping map[string]func() any
}

// The DefaultRegistry that is used when a specific registry is not specified.
var DefaultRegistry = NewRegistry()

// Creates a new, empty registry.
func NewRegistry() *Registry {
	return &Registry{
		typeMapping: map[string]func() any{},
	}
}

// --

// Registers a type using the derived type key into the default registry.
// If a registration for the type key already exists, it is replaced.
func Register[T any]() {
	typeKey := DerivedTypeKey[T]()
	RegisterWithKey[T](typeKey)
}

// Registers a type using the provided type key into the default registry.
// If a registration for the type key already exists, it is replaced.
func RegisterWithKey[T any](typeKey string) {
	RegisterIntoWithKey[T](DefaultRegistry, typeKey)
}

// Registers a type using the derived type key into the provided registry.
// If a registration for the type key already exists, it is replaced.
func RegisterInto[T any](registry *Registry) {
	typeKey := DerivedTypeKey[T]()
	RegisterIntoWithKey[T](registry, typeKey)
}

// Registers a type using the provided type key into the provided registry.
// If a registration for the type key already exists, it is replaced.
func RegisterIntoWithKey[T any](registry *Registry, typeKey string) {
	registry.typeMapping[typeKey] = func() any {
		return Hydrate[T]()
	}
}

// --

// Resolves a type using the derived type key from the default registry.
// If a registration for the type key does not exist, the call will panic.
func Resolve[T any]() T {
	typeKey := DerivedTypeKey[T]()
	return ResolveWithKey[T](typeKey)
}

// Resolves a type using the provided type key from the default registry.
// If a registration for the type key does not exist, the call will panic.
func ResolveWithKey[T any](typeKey string) T {
	return ResolveFromWithKey[T](DefaultRegistry, typeKey)
}

// Resolves a type using the derived type key from the provided registry.
// If a registration for the type key does not exist, the call will panic.
func ResolveFrom[T any](registry *Registry) T {
	typeKey := DerivedTypeKey[T]()
	return ResolveFromWithKey[T](registry, typeKey)
}

// Resolves a type using the provided type key from the provided registry.
// If a registration for the type key does not exist, the call will panic.
func ResolveFromWithKey[T any](registry *Registry, typeKey string) T {
	return registry.typeMapping[typeKey]().(T)
}

// --

// Returns true if the default registry has a registration for the derived type
// key of T.
func Knows[T any]() bool {
	typeKey := DerivedTypeKey[T]()
	return KnowsTypeKey(typeKey)
}

// Returns true if the default registry has a registration for the provided type
// key.
func KnowsTypeKey(typeKey string) bool {
	return KnowsTypeKeyIn(DefaultRegistry, typeKey)
}

// Returns true if the provided registry has a registration for the derived type
// key of T.
func KnowsIn[T any](registry *Registry) bool {
	typeKey := DerivedTypeKey[T]()
	return KnowsTypeKeyIn(registry, typeKey)
}

// Returns true if the provided registry has a registration for the provided
// type key.
func KnowsTypeKeyIn(registry *Registry, typeKey string) bool {
	_, ok := registry.typeMapping[typeKey]
	return ok
}

// --

// Returns the derived type key for a given type.  For example, if T was
// io.Reader, then the string "io.Reader" would be returned.
func DerivedTypeKey[T any]() string {
	t := Hydrate[T]()
	return fmt.Sprintf("%v", reflect.TypeOf(t))
}

// Creates a new instance of the type T and returns it.
func Hydrate[T any]() T {
	var t T
	return t
}
