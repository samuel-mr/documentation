
#  Ownership
## Memoria
- Estuctura de una variable en el Stack:

    | Name     | value              |
    |----------|--------------------|
    | ptr      | (memory possition) |
    | len      | 6                  |
    | capacity | 6                  |

- Estuctura de una variable en el Heap:

    | Index | value |
    |-------|-------|
    | 0     | h     |
    | 1     | e     |
    | 2     | l     |
    | 3     | l     |
    | 4     | o     |


## Assign values
- Si un tipo implementa el 'trait' 'Copy', esta se copia implícitamente (con en el caso de los i32)
- Si un tipo implementa el 'trait' 'Drop', rust no dejará que se implemente 'Copy'

1. Stack only: Se copia la información del stack (en un integer todo está en el stack)
    ```rust
    let x = 5;
    let y = x;
    ```
2. Stack + Heap: Se copia solo la información del stack, la información del heap no, solo se mueve el puntero
    ```rust
    let s1 = String::from("hello");
    let s2 = s1;  // en puntero de 's1' al heap se mueve hacia el 's2', apuntando a la misma posición de memoria
    ```
- ejem: <span style="color:red">example with error:</span>
    ```rust
    let s1 = String::from("hello");
    let s2 = s1; // ownership pasado a 's2'

    println!("{}, world!", s1); // error, porque s1 ya no 'existe'
    ```
    Para solventarlo podemos 'copiarlo'
    ```rust
    let s1 = String::from("hello");
    let s2 = s1.clone();
    ```
Stack Data: la data que está en el stack es copiado automáticamente, ejem: integers
```rust
    let x = 5;
    let y = x;
    println!("{}", x); //compila ok
```
¿Qué tipos implementan 'Copy'?
- All the integer types, such as u32.
- The Boolean type, bool, with values true and false.
- All the floating point types, such as f64.
- The character type, char.
- Tuples, if they only contain types that also implement Copy. For example, (i32, i32) implements Copy, but (i32, String) does not.

## Reference and Borrowing
```rust
fn read(variable: &String)  // solo lectura

fn write(variable: &mut String)  // lectura + escritura
```
Reglas:
1. Solo se puede tener un &mut en el mismo scope. <span style="color:red">example with error:</span>
    ```rust
    let mut s = String::from("hello");

    let r1 = &mut s;
    let r2 = &mut s; // Error !!!

    println!("{}, {}", r1, r2);
    ```
    - Para solventarlo podemos crear otro scope:
        ```rust
        let mut s = String::from("hello");
        {
            let r1 = &mut s;
        } // r1 sale del scope, este código si es válido

        let r2 = &mut s;
        ```
2. No se puede tener juntas una referencia normal y otra referencia mutable en el mismo scope. <span style="color:red">example with error:</span>
    ```rust
        let mut s = String::from("hello");
        let r1 = &s;  
        let r2 = &mut s; // Error !!!
        println!(..., r1, r2);
    ```
    - Para solventarlo puede asegurarse de usarse antes de, ejem:
        ```rust
        let mut s = String::from("hello");

        let r1 = &s; // no problem
        let r2 = &s; // no problem
        println!(..., r1, r2); // r1 y r2 ya no serán usadas, así que r3 será valido
        // aquí termina el scope de r1 y r2.

        let r3 = &mut s; // no problem
        println!(..., r3);
        ```
3. No deben ocurrir 'Dangling reference'.
<span style="color:red">example with error:</span>
    ```rust
    fn method() -> &String { // retorna una referencia a un String
        let s = String::from("hello"); // declara 's'
        &s // se retorna una referencia a 's'
    } // 's' intenta ser eliminado ya que su scope finaliza, pero también se intenta retornar una referencia a esta.
    // Error !!!
    ```
    - Solución: retornar el 'String' completo. (aquí el ownership se reasignará)
        ```rust
        fn method() -> String {
            let s = String::from("...");
            s
        }
        ```
## Lifetime
- Inicia cuando un objeto es creado, termina cuando es destruido.
- Lifetime != Scope
- No modifica el tiempo de vida, solo indica explícitamente el tiempo de vida
- Anotación de tiempo de vida genérico
    ```
    'a
    ```
- Declarando un parámetro lifetime llamado 'a' y 'b'
    ```
    foo<'a>
    foo<'a, 'b>
    ```
- Un lifetime corto no puede asignarse a uno más largo. ejem
<span style="color:red">example with error</span>
    ```rust
    fn failed_borrow<'a>() {
        let _x = 12;
        // ERROR: `_x` does not live long enough
        let y: &'a i32 = &_x;

    ```
- ejem 2 
<span style="color:red">example with error</span>
    ```rust
    fn run() {
        let r ;
        {
            let x = 5;
            // error: la referencia de x no se puede asignar a r, motivo: x vive menos que r.
            r = &x;
        }
    }
    ```

## in functions
- El lifetime del valor de retorno estará asociado siempre al lifetime de algún parámetro de la función.
- Nomenclatura
    ```rust
    &i32        // a reference
    &'a i32     // a reference with an explicit lifetime
    &'a mut i32 // a mutable reference with an explicit lifetime
    ```
- Ejem: ¿Que lifetime se le asignará al valor de retorno del método?: el lifetime que se le asignará será el parámetro que tenga el menor lifetime de todos los parámetros. ejem:
    ```rust
    fn method() {
        let x = String::from("xxx");

        let result;
        {
            let y = String::from("yyy");
            // Lifetime asignado final será: 'y'
            result = metodo(x.as_str(), y.as_str());
    ...

    fn metodo<'a>(x: &'a str, y: &'a str) -> &'a str {
    ```
    caso: 2 parámetros con lifetime
    ```rust
    pub fn run() {
        let x = "aaa";
        let y = "bb";

        let result = textos(&x, &y);
        print!("{result}");
    }

    // Si un método retorna un valor, este valor tendrá un 'lifetime' igual al menor 'lifetime' de un parámetro. Por eso, el tipo de retorno necesita que se especifique el tiempo de vida asociado ya que por si solo no puede inferirlo de los parámetros
    fn textos<'a>(x: &'a str, y: &'a str) -> &'a str {
        if x.len() > y.len() {
            return x;
        } else {
            return y;
        }
    }
    
    ```

    caso: 1 parámetro con lifetime. Solo se analizan los lifetime de los parámetros donde está indicado explícitamente:
    ```rust
    // solo x tiene lifetime
    fn metodo<'a>(x: &'a str, y: &str) -> &'a str {
            return x;
    }
    ...
    // este codigo compilará bien, porque 'y' no es considerado para el cálculo del lifetime
    fn otro(){
        let x = String::from("xxx");

        let result;
        {
            let y = String::from("yyy");
            // lifetime final = "x" ("y" no sería considerado)
            result = metodo(x.as_str(), y.as_str());
        }
        print!("{result}");
    ```
- Transferencia de ownership

    ```rust
    // La variable 'texto' transferirá el ownership hacia quién invoque a 'metodo'
    fn  metodo() -> String{
        let texto = String::from("aaaa");
        return texto; // se transfiere el ownership hacia el nivel superior
    }

    // igual con un Struct
    struct  Person {}
    fn  metodo() -> Person{
        let persona = Person{};
        return persona;
    }
    ```

# Samples
- puntero modificable
    ```rust
    struct Directory {
        name: String,
        children: Option<Box<Vec<Directory>>>,
        size: i32,
    }
    impl Directory {
        fn some_method(&mut self, name: &str) {}
    }
    fn run2()   {
        let mut root = Directory { // let mut!!!
            name: "/".to_string(),
            children: None,
            size: 0,
        };

        let mut current_node = &mut root; // &mut !!!!
        current_node.some_method("test");
    }
    ```