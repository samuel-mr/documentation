
me quede aqui: https://www.youtube.com/watch?v=6hDMfVykjo4&list=PLV-IZMSUqzrNkcKRCw2cjKrW0a8l5qvgb&index=6


https://youtu.be/8uiZC0l4Ajw?si=JZ74aQsHXid2C5jR&t=1112
https://www.youtube.com/watch?v=E6KVIUJ50y0
# Sintaxis
- Lenguaje imperativo
- Tipado estático
- compila a lenguaje máquina (no maquinas virtuales)
- No tiene clases, pero si estructuras con métodos
- Funciones son ciudadanos de primera clase y pueden retornar múltiples valores
- puntos y coma al final son opcionales (el compilador los agrega)
- Tiene punteros, pero no punteros aritméticos
- go no tiene excepciones

# Hello word
```go
package main

import "fmt"

func main() {
    fmt.Println("Hola Gouu")
}
```
# Commands
```bash
# compile
go build ./main.go
# to run compiled binary
./main
# compile + run
go run ./main.go


```
# Variables
- Todas las variables tienen un valor inicial

| Tipo         | Valor Inicial (Zero Value)       |
|-------------|---------------------------------|
| Booleano    | false                           |
| Enteros     | 0 (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr) |
| Flotantes   | 0 (float32, float64)            |
| Complejos   | 0 + 0i (complex64, complex128)  |
| Texto       | "" (cadena vacía)               |
| Caracteres  | 0 (rune, byte)                  |
| Array       | Elementos con su zero value     |
| Referencias | nil (slice, map, chan, func, interface, pointer) |
| Struct      | Campos con su zero value        |


- Declaraciones
    ```go
    var numero int
    const contante = "This is a constant"

    // type declaration
    var numero int      // explícito
    var x:=23           // implícita (tipo inferido)
    x := 23             // implícita (tipo inferido) - pero solo dentro de funciones
    x,y := 1,2          // implícita múltiple

    // initial values
    var numero int  // sin valor (implíctamente será 0)
    var numero int = 1 // con valor

    // declaración e inicialización múltiple
    var foo, bar int = 42, 1302
    ```
- Visibilidad
    ```go
    // variable pública: con mayúscula
    type Customer struct {}
    // variable privada: con minúscula
    type customer struct {}
    ```
- Common erros
    ```go
    // not allowed (different types)
    "hola" + 5      ❌

	var numero int
	var numero int  ❌


    len("世界")                    // 6  ❌
    import "unicode/utf8"
    utf8.RuneCountInString("世界") // 9  ✅
    ```

# Convertir
```go
numero, _ := strconv.Atoi(texto)        // String -> int// el 2do valor es ignorado con _
texto := strconv.Itoa(numero)           // int -> String

// Otras conversiones
    entero a 
	f32 := float32(i)
	f64 := float64(i)
	var_u := uint(i)
```

# Imprimir
```go
// simple
fmt.Println("numero a:", numero)

// con formato
fmt.Printf("numero b: %d \n", numero)
// args:
    %T => Tipo del valor
    
    %s   => formato texto plano
    %t   => boolean será: true/false
    %f   => formato flotante
    %.2f => formato flotante, precisión 2
```

| Categoría      | Formato  | Descripción                              | Ejemplo de uso               | Resultado            |
|---------------|---------|------------------------------------------|------------------------------|----------------------|
| **Genéricos** | %v      | Valor en formato predeterminado         | fmt.Printf("%v", 42)         | 42                   |
|               | %+v     | Estructura con nombres de campo         | fmt.Printf("%+v", struct{X int}{10}) | {X:10}       |
|               | %#v     | Representación completa de Go           | fmt.Printf("%#v", struct{X int}{10}) | main.struct{X:10} |
| **Números**   | %d      | Entero decimal                         | fmt.Printf("%d", 42)         | 42                   |
|               | %b      | Binario                                 | fmt.Printf("%b", 42)         | 101010               |
|               | %o      | Octal                                   | fmt.Printf("%o", 42)         | 52                    |
|               | %x / %X | Hexadecimal (min/may)                   | fmt.Printf("%x", 255)        | ff / FF              |
|               | %f      | Flotante                                | fmt.Printf("%.2f", 3.1416)   | 3.14                 |
|               | %e / %E | Notación científica                     | fmt.Printf("%e", 1234.56)    | 1.234560e+03         |
| **Texto**     | %s      | Cadena de texto                        | fmt.Printf("%s", "Hola")     | Hola                 |
|               | %q      | Cadena con comillas                     | fmt.Printf("%q", "Hola")     | "Hola"               |
|               | %c      | Carácter Unicode                        | fmt.Printf("%c", 65)         | A                    |
|               | %U      | Código Unicode                          | fmt.Printf("%U", 'A')        | U+0041               |
| **Booleanos** | %t      | true o false                            | fmt.Printf("%t", true)       | true                 |
| **Punteros**  | %p      | Dirección de memoria                    | fmt.Printf("%p", &variable)  | 0xc0000140a0         |
| **Múltiples formatos** | %T | Tipo de la variable                | fmt.Printf("%T", 42)         | int                  |

# Consola
- Ingresar datos
    ```go
    // numero
    var numero int
    fmt.Println("Escribir un numero")
    fmt.Scanf("%d\n", &numero)

    // texto
    var texto string
    fmt.Println("Ingresa tu nombre")
    fmt.Scanf("%s\n", &texto)
    ```

# Funciones
- por el retorno
    ```go
    func name() {        // sin retorno
    func name() int {    // con retorno 

    func name(param1 int, param2 int) {     // explicito
    func name(param1    , param2 int) {     // resumido

    func name() int {                      // retorno simple
    func name() (int, string) {            // retorno múltiple
    ```

- ejem retorno múltiple
    ```go
    // explicit
    func method() (int, string) {
        var n = 42
        var s = "foobar"
        return n, s
    }
    // implicit
    func method() (n int, s string) {
        n = 42
        s = "foobar"
        return
    }
    
    var number, text = method()
    ```

- según cantidad de variables retornadas
## VS Code - extensions
- [Go language support](https://marketplace.visualstudio.com/items?itemName=golang.Go)

# Compilar
```
go build <name.go>
```
# Paquetes
```
import "fmt"
```

| nombre | descripcion |
|--------|-------------|
| fmt   | imprimir         |
| strconv   | convertir         |

# Cheatsheet
- https://devhints.io/go
- https://github.com/a8m/golang-cheat-sheet#arithmetic
- https://quickref.me/golang

# Modulos

```
go mod init [name]      # inicializa un módulo
```
