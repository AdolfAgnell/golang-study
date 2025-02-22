### O pacote fmt

## Setup: strings, ints, bools

### Strings: interpreted string literals vs. raw string literals

```go
// Interpreted string literal
interpreted := "Hello, World!"

// Raw string literal
raw := `Hello,
World!`
```

### Rune literals

```go
r := 'a' // rune literal
```

Em ciência da computação, um literal é uma notação para representar um valor fixo no código fonte.

## Format printing: documentação

### Grupo #1: Print → standard out

```go
fmt.Print("Hello, World!")
fmt.Println("Hello, World!")
fmt.Printf("Hello, %s!", "World")
```

Format verbs: `%v`, `%T`

### Grupo #2: Print → string, pode ser usado como variável

```go
s := fmt.Sprint("Hello, World!")
s1 := fmt.Sprintf("Hello, %s!", "World")
s2 := fmt.Sprintln("Hello, World!")
```

### Grupo #3: Print → file, writer interface, e.g. arquivo ou resposta de servidor

```go
file, _ := os.Create("example.txt")
defer file.Close()

fmt.Fprint(file, "Hello, World!")
fmt.Fprintf(file, "Hello, %s!", "World")
fmt.Fprintln(file, "Hello, World!")
```
