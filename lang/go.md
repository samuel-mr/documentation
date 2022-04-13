me quede aqui: https://www.youtube.com/watch?v=6hDMfVykjo4&list=PLV-IZMSUqzrNkcKRCw2cjKrW0a8l5qvgb&index=6
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
```
package main

import "fmt"

func main() {
    fmt.Println("Hola Gouu")
}
```
# Variables
- Todas las variables tienen un valor inicial
    - int     -> 0
    - string  -> "" (vacío)
    - bool    -> false
    - []string-> []
- Declaraciones
    - Según naturaleeza
        ```
        var cambios int
        const constant = "This is a constant"
        ```
    ```
    var numero int      explícita 
    var x:=23           implícita (tipo inferido)
    x := 23             implícita (tipo inferido) - pero solo dentro de funciones
    ```
    según incialización
    ```
    var numero int      sin valor (implíctamente será 0)
    var numero int = 1  con valor
    ```
    según cantidad de variables
    ```
    var foo, bar int = 42, 1302     declaración e inicialización múltiple
    ```
- Visibilidad
    - variable pública: con mayúscula
    ```
    type Customer struct {}
    ```
    - variable privada: con minúscula
    ```
    type customer struct {}
    ```
- Las concatenaciones de diferentes tipos no está permitido
    ```
    "hola" + 5      // error!
    ```
- No se pueden volver a declarar en el mismo nivel
    ```
	var numero int
	var numero int  // error!
    ```

# Convertir
```
numero, _ := strconv.Atoi(texto)        // String -> int// el 2do valor es ignorado con _
texto := strconv.Itoa(numero)           // int -> String

// Otras conversiones
    entero a 
	f32 := float32(i)
	f64 := float64(i)
	var_u := uint(i)
```


# Imprimir
```
// simple
fmt.Println("numero a:", numero)

// con formato
fmt.Printf("numero b: %d \n", numero)
    %T => Tipo del valor
    %d => base 10
    
    %s   => formato texto plano
    %t   => formato del valor que representa, ejemplo boolean será: true/false
    %f   => formato flotante
    %.2f => formato flotante, precisión 2
    %e   => formato científico
    %p   => formato base 16 que inicia con 0x (para los punteros)
```
# Consola
Ingresar datos
```
//   numero
	var numero int
	fmt.Println("Escribir un numero")
	fmt.Scanf("%d\n", &numero)

//   texto
	var texto string
	fmt.Println("Ingresa tu nombre")
	fmt.Scanf("%s\n", &texto)
```

# Funciones
- por el retorno
    ```
    func name() {        sin retorno
    func name() int {    con retorno 
    ```
- tipo de los parámetros
    ```
    func name(param1 int, param2 int) {     explicito
    func name(param1    , param2 int) {     resumido
    ```
- según cantidad de variables retornadas
    ```
    func name() int {                       simple
    func name() (int, string) {             múltiple
    ```
    - ejem retorno múltiple
        ```
        func returnMulti2() (n int, s string) {
            n = 42
            s = "foobar"
            // n y s serán retornados
            return
        }
        var x, str = returnMulti2()
        ```
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
