![](icon.png)

# hydration

A simple type registration and instance hydration library.

## Usage Example

This library was originally created to solve type hydration from message buses,
as loosely shown in the example below, but has a wide variety of usages in other
problem spaces.

```go
// During setup.
hydration.Register[event.CustomerCreated]()

// During message bus serialization.
message.TypeKey = hydration.DerivedTypeKey[event.CustomerCreated]()

// During message bus deserialization.
typeKey := message.TypeKey
e := hydration.ResolveWithKey[event.Event](typeKey)
```

## Core Concepts

The unit of isolation in this package is the `Registry`.  Users of this package
can either create their own registry or use the global, default registry. Types
are registered in registries using Type Keys, which are just strings used
to look-up types.

There are 3 core operations exposed by this package:

- `Register` a new type in a registry.
- `Resolve` a concrete instance from a registry.
- `Know` if a registry has a registration for a type key.

For each of these operations, there are 4 different variants:

- Derived type key, default registry
- Custom type key, default registry
- Derived type key, custom registry
- Custom type key, custom registry

For example, `Register` uses derived type keys and the default registry, but
`RegisterIntoWithKey` allows a custom registry and type key to be specified.
