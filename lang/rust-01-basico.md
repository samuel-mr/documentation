# Rust !
> Seguridad, concurrencia y rendimiento

## Dictionary
```
::<SomeType>        turbo-fish, sirve para indicarle al compilador el tipo de dato final
explicacion: https://techblog.tonsser.com/posts/what-is-rusts-turbofish
```

## Declaración
```
let      => inmutable
let mut  => mutable
const PI => inmutable constante

const PI = 3.14
let mut edad = 54

```

```
let numero = 2            # implícito
let numero:i32 = 2        # explícito

let entero = 1            # i32 inferido
let real   = 1.1          # f64 inferido

let millon = 1_000_000u32 # número más legible

```
varias en 1 línea
```
    let (a, b, c) = (0, 0, 0);
```

## Shadowing
Capacidad para declara una variable con el mismo nombre en un mismo escope o en un scope interno. La nueva variable reemplazaría a la antigua.
```
  let x = 1;
  {
      x             // print 1
      let x = "abc"; 
      x             // print 'abc'
  }
   x                // print 1
  let x = true;
  x                 // print true
```


## Operadores
- Los tipos deben ser iguales
```
+ - * /     Numéricos

> < == !=   Relacionales

&&          Lógicos
||


1u32 - 2    !error, desbordamiento, porque el cálculo será negativo y el tipo resultado se infiere q es unsigned 32
```
```
 3 + 2f64;  // error! Operaciones de diferentes tipos no está permitido
```

## Bloques
- Las variables se crean, modifican y destruyen dentro de su bloque
- Una variable está disponible dentro de su bloque y los bloques anidados
- Una variable no está disponible en un bloque superior

Alcance de las variables
```
fn main() {
    // bloque 'A'
    {
        // bloque 'B' ... aquí también están disponibles las variables de 'A'
        // al finalizar este bloque se eliminarán todas las variables creadas en este bloque 'B'
    }
    // aquí ya no están disponible ninguna variable del bloque 'B'
}
```
Bloques con return (como funciones anónimas internas a un fn)
```
fn main() {
  let resultado = {
      let variable: i32 = 10;
      variable                // la última línea quiere decir que de esa variable se retornará su VALOR
  };
  println!("{}", resultado);  // print: 10
}
```
ejemplo 2
```
let resultado = 10;
let mensaje = if resultado == 10 {
    String::from("El resultado es 10")  // sin 'return' ni 'punto y coma'
} else {
    String::from("El resultado no es 10") // sin 'return' ni 'punto y coma'
};

println!("{}", mensaje);  // print: El resultado es 10
```
## Clousure
- Es un valor que puede ser llamado como una función. 
- | | aquí dentro pueden ir los parámetros
- { } aquí estará el cuerpo del clousure
- ejem:
    ```
        || {                                   
            App::new()
                .route("/",web::get().to(get_index))
        } 
    ```
- version sin el cuerpo (simplemente excluye las llaves)
    ```
        || App::new().service(get_index)
    ```

## Macros
- Son como programas externos
- No son funciones
> Assert!(...) : el símbolo ! indentifica que Assert es un Macro

## Panic
- Suceden cuando el programa finaliza abruptamente
- Es lanzado cuando una condición Assert es false


> debug_assert! solo se compila en debug-time, caso contrario es omitido

## Atributos
- Agregan información adicional a una función
- También le dicen a Rust como poder interactuar con otros lenguajes
- ejem:
    ```
    #[test]
    fn ...
    ```

## Test
- Omitidos en la compilación
- Ejem:
    ```
    #[test]
    fn test_mcd(){
    ```
    runing test
    ```
    cargo test
    ```
- Test panic?
    ```
    #[test]
    #[should_panic]
    fn ...

    #[test]
    #[should_panic(expected = "Divide result is zero")]
    fn ...
    ```
- Ignorar
    ```
    #[ignore]
    ```
# Importar

## Formas de importar
A) Implícito
```
use actix_web;

actix_web::HttpResponse
```
B) Explícito
```
use actix_web::{HttpResponse, ...};

HttpResponse
```
## Trait
- Interfaz (aunq más como una clase abstracta).
- Ejem: todos los tipos que implementan `FromStr` tienen el método `from_str`
    ```
    use std::str::FromStr;

    fn main() {
        numbers.push(u64::from_str(&arg)
    ```

## otros ---

```
.filter(|x| x % 3 == 0 )    // x=será cada item de la iteración, la expresión debe retornar un bool
```

# Ejecución del terminal

## Terminal
```
cargo run                       # ejecuta el programa en ese contexto 
cargo run param1 param2 etc     # ...y envía parámetros
```
## Ciclo de vida del programa
```
std::process::exit(1);      // finaliza el programa con el estado 'error': 1= error
```

## Helpers
```
return Default::default();
```
