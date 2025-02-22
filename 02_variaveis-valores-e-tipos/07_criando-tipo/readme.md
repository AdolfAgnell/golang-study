Revisando: tipos em Go são extremamente importantes. (Veremos mais quando chegarmos em métodos e interfaces.)
Tem uma história que Bill Kennedy dizia que se um dia fizesse uma tattoo, ela diria "type is life."
Grande parte dos aspectos mais avançados de Go dependem quase que exclusivamente de tipos.
Como fundação para estas ferramentas, vamos aprender a declarar nossos próprios tipos.
Revisando: tipos são fixos. Uma vez declarada uma variável como de um certo tipo, isso é imutável.
```go
// Declarando um novo tipo
type hotdog int

// Usando o novo tipo
var b hotdog
b = 10

// Isso causará um erro, pois b é do tipo hotdog, não int
// var x int = b
```

Uma variável de tipo `hotdog` não pode ser atribuída com o valor de uma variável tipo `int`, mesmo que este seja o tipo subjacente de `hotdog`.

### Exemplos Adicionais

```go
// Declarando um novo tipo
type temperatura float64

// Usando o novo tipo
var t temperatura
t = 36.6

// Isso causará um erro, pois t é do tipo temperatura, não float64
// var temp float64 = t
```

```go
// Declarando um novo tipo
type idade uint

// Usando o novo tipo
var i idade
i = 30

// Isso causará um erro, pois i é do tipo idade, não uint
// var years uint = i
```

Esses exemplos mostram como a declaração de novos tipos em Go pode ajudar a garantir a segurança de tipos e evitar erros de atribuição.
