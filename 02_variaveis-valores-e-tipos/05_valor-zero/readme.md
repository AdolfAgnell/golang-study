### Valor zero

Declaração vs. inicialização vs. atribuição de valor. Variáveis: caixas postais.
O que é valor zero?
Os zeros:
ints: 0
floats: 0.0
booleans: false
strings: ""
pointers, functions, interfaces, slices, channels, maps: nil
Use := sempre que possível.
Use var para package-level scope.

#### Exemplos

```go
package main

import "fmt"

func main() {
    var i int
    var f float64
    var b bool
    var s string
    var p *int
    var fn func()
    var iface interface{}
    var sl []int
    var ch chan int
    var m map[string]int

    fmt.Printf("int: %v\n", i)
    fmt.Printf("float64: %v\n", f)
    fmt.Printf("bool: %v\n", b)
    fmt.Printf("string: %q\n", s)
    fmt.Printf("pointer: %v\n", p)
    fmt.Printf("function: %v\n", fn)
    fmt.Printf("interface: %v\n", iface)
    fmt.Printf("slice: %v\n", sl)
    fmt.Printf("channel: %v\n", ch)
    fmt.Printf("map: %v\n", m)
}
