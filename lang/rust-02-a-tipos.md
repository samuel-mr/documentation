Resumen:
- TIPOS PRIMITIVOS
    - Primitivos Escalares
    - Primitivos Compuestos
        - Tuplas (1,true)
        - Arreglos [1,2,3]
- TIPOS PERSONALIZADOS
    - Structs
    - Enums
    - constats

# Primitivos Escalares
```
i8, i16, i32, i64, i128  enteros positivos y negativos
u8, u16, u32, u64, u128  enteros positivos
f32, f64                 flotantes (float y double en C++)
char                     UTF-8, podemos colocar emojis
bool
()                       unit type

usize                    tipo entero sin signo del tamaño de puntero
```

## String
Common methods
```
.to_lowercase()
.len()
.ends_with(...)
```
Loop over
```
let texto = "abc";

let mut text_as_chars = texto.chars();
text_as_chars.next()  // Option('a')
text_as_chars.next()  // Option('b')
text_as_chars.next()  // Option('c')

//solo tenía 3
text_as_chars.next()  // Option(None)
...
```
convert
```
String::from_str(mayor)     str a String
```

## Str
Declare
```
let hello = "Hello, world!";                // implícito

let hello: &'static str = "Hello, world!";  // explícito
 
```
Loop
```
    for (size, item) in text.chars().enumerate() { ...
```
Concatenar
```
let concatenated = ["hola", "mundo"].concat();
let concatenated = format!("{}{}", "hola","mundo");
let concatenated = concat!("hola", "mundo");
let concatenated = "hola".to_string() + "mundo";

// holamundo
```
### UTF-8
- Los textos se guardan en UTF-8
- Cada caracter ocupará de 1-4 bytes
- [As String works (GetsRusty)](https://www.youtube.com/watch?v=Mcuqzx3rBWc)

```
let letter = "a";
letter.bytes().count()  // 1 bytes
letter.chars().count()  // 1 item

let crabs = "🦀";
crabs.bytes().count()) // 4 bytes
crabs.chars().count()  // 1 item

```

# Primitivos Compuestos
## Arreglos
- La longitud es fija
- almacenado en el stack
- Solo puede almacenar 1 solo tipo de dato
- Los arrays pueden ser prestados ( como `param `en otra `function` por ejem) como `slice`
```
definicion: [x...z]

let numeros     = [1,2,3];        arreglo solo lectura
let mut numeros = [1,2,3];        arreglo de valores modificables

let numeros:[i32:4] = [1,2,3,4];  especifica tipo y tamaño
let numeros = [2;5];              [valor;repetir x veces] print: 2,2,2,2,2


println!("{:?}", numeros);   imprime en una sola línea
println!("{:#?}", numeros);  imprime cada item en una nueva línea

arreglo[i] = 1               asignar

len()                        cantidad elementos
```
ENVIO DE REFERENCIA DE ARRAY
```
let numeros = [5,4,3,2,1];
fn metodo(slice: &[i32]){
  ...
}
metodo(&numeros);             TODO:                                          print 5,4,3,2,1
metodo(&numeros[1 .. 3]);     PARCIAL: [posición inicial .. posición final]  print 4,3
```

## Slices
- Es como un array, pero no se sabe el tamaño en tiempo de compilación
- Son usados para 'prestar' un corte de un arreglo: Declaración = &[T]
- Almacenado en el Heap
```
    let mensaje = String::from("hola mundo");

    let inicio = &mensaje[0..4];          // indice inicial .. indice final
    let inicio = &mensaje[..4];           // idem

    let fin= &mensaje[4..mensaje.len()];  // inicia en 4 hasta el final
    let fin= &mensaje[4..];               // idem

    let fin= &mensaje[..];               // todo el string


    fn procesa_slice(slice: &[i32]) {
```

## TUPLAS
- Puede almacenar variables de distintos tipos
```
definicion: (T1...Tx)

let tupla = (1,true, 3.4);                implícito
let tupla: (i32,bool,f64) = (1,true,3.4); explícito

let tupla     = (1,true);           tupla solo lectura
let mut tupla = (1,true);           tupla lectura + escritura

let tupla: ((i32,bool),(f64,i32));  tuplas anidadas

tupla.0                             accede a cada elemento por su posición

let (numero, boleano) = (1, true);  desestructuración
numero    // 1
boleano   // true
```


# TIPOS PERSONALIZADOS

## Enumeradores
```
enum Response {
    Ok,
    Error(u32, String),                         // tupla
}

let result = Response::Error(501, String::from("ocurrió algun error !"));
match result {
    Response::Ok =>                  ,          // si es 'Ok'
    Response::Error(403, _ ) =>      ,          // si la tupla es Error con el primero parámetro como '403'
    Response::Error(404, _ ) =>      ,
    Response::Error(405, _ ) =>      ,
    Response::Error( _, mensaje) =>  ,           // si la segunda variable es un string
};
// El guión '_' quiere decir: ignora este parámetro
```
compare: necesita implementar 'PartialEq'
```
#[derive(PartialEq)]
enum OptionGame{
    Rock,
    Sissor
}
...
assert_eq!(OptionGame::Sissor , value);
```
## Estructuras
- Como las Tuplas pero con nombres, de manera que ya no dependemos del orden sino de nombres
- Tipos
    - Tuple struct: es una tupla con nombre
    ```
    struct Pair(i32, f32); 
    
    //ejemplo: 
    #[derive(Debug)]
    struct Color(u32, u32, u32);      // no serán propiedades sino valores de la tupla

    let white = Color(255,255,255);

    println!("El color es: {:?}", white);
    ```
    - Struct: la clásica
    ```
    struct Point {
        x: f32,
        y: f32,
    }
    ```
    - Unit struct: no tiene campos (útil para generics)
    ```
    struct Unit;
    ```
Uso
```
struct User {
    username: String
}

// 1. inicializar las propiedades directamente
let variable = User {                            
    username: String::from("texto")
};

// 2. inicializar las propiedades por medio de otras variables
let username = String::from("texto");
let mut us = User { username }; // como la variable se llama 'username' buscará una propiedad con el mismo nombre
```
permitir que la información sea mostrada en Debug, ejem: para imprimir los datos
```
#[derive(Debug)]    // permite que se puede imprimir todos los datos tan solo indicando la estructura
struct User {
   ...
```

