Generics provide a way to programatically (in runtime) instantiate functions and types

```go
func F(x T1, y T2) {...}
// becomes...
func F[T1 Constraint1, T2 Constraint2](x T1, y T2) {...}
// F is a Generic Function!

Constraints are interfaces:

type Constraint1 interface {
    Method(a, b int) int
}
```
Can be also be declared inline together with function declaration:
```go

func F[T1 interface { Method(a, b int) int }, T2 Constraint2](x T1, y T2) {...}
```
Generic functions must always be instantiated before being called!
```go
type SomeType interface {
    Less(c SomeType) bool
}

func Sort[Elem interface{}](list []Elem) {...}

func main() {
    var list = []SomeType{...}
    
    Sort[SomeType](list)
  //-------------||----
  // instatiation  invocation
}

Step #1: Instantiation:
    Sort[SomeType]
            ↓					-- pass type argument
    Sort[Elem interface{}]
            ↓					-- substitute SomeType for Elem
    Sort[SomeType interface{}]
            ↓					-- verify that SomeType satisfies type parameter constraint
    `#Sort[SomeType]`           -- instantiated function (normal function)
    |-> func (list []SomeType) {...}

Step #2: Invocation (as usual):
    #Sort[SomeType](bookshelf)
```
However, for usability, this process of instantiation can often be called implicitly, by using type inference under the hood. This is called type unification:
```go
func main() {
    var list = []SomeType{...}

    Sort(list)
}

`list` underlying type is compared against that of `Elem`:
            list -> []SomeType
func Sort[Elem ...]([]Elem)
                    => Elem == SomeType
```
What are generic types?
- Just as generic functions, generic types allows us to instantiate types in runtime to create more complex constraints
- Generic types can only define constraints and can not be used outside this context
```go

type Comparable[T interface{}] interface {
    Less (el T) bool
} 

The `Comparable` generic type expresses a constraint of types 
that can be compared to themselves

Generic types need also be instantiated with the [] syntax:

func Sort[Elem Comparable[Elem]](list []Elem) {...}

```
Type constraints may also have a list of types. Examples:
```go
type PointerT[T any] interface {
    *T
}
// Types satisfying PointerT must have their type equal to *T

type Float interface {
    float32 | float64
}
// Types satisfying Float must have their type
// equal to one of float32, float64

type Bytes interface {
    []byte | ~string
}
// The ~ operator includes types' underlying types, so a custom type
// MyString would also satisfy the Bytes interface
func Index[bytes Bytes](s, sep bytes) int
```
