## Conversão de Tipos

Conversão de tipos é exatamente o que parece. Em Go, não se diz "casting", mas sim "conversion".

### Exemplo de Conversão

```go
package main

import (
  "fmt"
)

func main() {
  var a int
  var b float64 = 3.14

  a = int(b) // conversão de float64 para int
  fmt.Println(a) // saída: 3
}
```

Para mais detalhes, consulte a [documentação oficial](https://golang.org/ref/spec#Conversions).
