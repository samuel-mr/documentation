# TIPOS PRIMITIVOS
- Primitivos Escalares
- Primitivos Compuestos
    - Tuplas (1,true)
    - Arreglos [1,2,3]

## Escalares
```
i8, i16, i32, i64, i128  enteros positivos y negativos
u8, u16, u32, u64, u128  enteros positivos
f32, f64                 flotantes (float y double en C++)
char                     UTF-8, podemos colocar emojis
bool
()                       unit type

usize                    tipo entero sin signo del tamaño de puntero
```

### String
```
.to_lowercase()
.len()
.ends_with(...)
```

### str
```
let hello = "Hello, world!";                // implícito

let hello: &'static str = "Hello, world!";  // explícito

// CONVERT
String::from_str(mayor)     str a String
```


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


# TIPOS SEGUN LIBRERIA STD
## VECTORES
- Son arrays redimensionables
- El tamaño no es conocido en tiempo de compilación

INICIALIZACION
```
let vector = Vec::new();

let ...          no modificable
let mut ...      modificable 

let vector: Vec<i32> = Vec::new();  Tipo explícito

// inicializar usando la macro vec!
let vector           = vec![1,2,3];     implícito
let vector: Vec<i32> = vec![1,2,3];     explícito


let mut vector = Vec::new();            implícito sin data (pero igualment inferido el tipo por contexto)
vector.push(1);                         
   // Tipo postergado: sin tipo en la inicialización, pero...
   // gracias a que en la siguiente línea se agrega un '1', en tiempo de diseño se infiere el tipo i32 del vector
```
METODOS y PROPIEDADES
```    
vector[1]       obtiene el elemento de alguna [posición]

.push(1)        agregar al final
.insert(0, -1)  agrega en (posición, elemento)
.pop()          elimina el último (por defecto retorna un Option con el elemento eliminado)
.pop().unwrap() es el pop y además desempaqueta ese elemento eliminado
.remove(0)      elimina en (posición)
.windows(x)     retorna un iterador agrupando cada x elementos contiguos
    ['a','b','c'].windows(2)        // &['a','b']  &['b','c']
.sort() 
.append(&mut otro_vector);         agrega los valores de 'otro_vector', donde este último quedará vacio

["a", "b"].join("*"), "a*b"         concatena con un comodín
array[index..index + k].join("");   obtiene un rango de elementos interno y los concatena
```
RECORRER
```
    for i in vector {
      
    }
```

## Strings
```
str           inmutable (en stack)
String        mutable   (en heap)

let variable = "Texto";                          tipo 'str' inferido
let mut variable = String::new();                tipo string
let mut variable = String::from("iniciado");     tipo string con texto

variable.push('.');
variable.push('.');
variable.push('.');
variable.push_str(" !!!");                       print iniciado... !!!

```
METODOS
```
.to_string()                                        convierte srt a String
.split(" ")                                         divide según el texto del parámetro
```
CONCATENAR
```
let mut owned_string: String = "hello ".to_owned();
let borrowed_string: &str = "world";

owned_string.push_str(borrowed_string);
```

raw string
```
r#"
    sirve para poner fácilmente texto normal y espacial (como este: ")
    en diferentes líneas
"#
```

## Option

Permite manejar la falla desde un generic q tiene 2 posibilidades, ok y error
```
enum Option<T> {
    Some(T),
    None
}

let variable: Option<i32> = Some(32);     // inicializa con valor
let variable: Option<i32> = None;         // inicializa como 'sin valor'

match opcion {
    Some(variable) =>     ,       // si es ok
    None =>               ,       // si hay error
}
```
Ejemplo 1
```
fn main() {
    let result = dividir(2, 0);
    match result {
        Some(1) =>                ,       // si es 1
        Some(2) =>                ,       // si es 2
        Some(_) =>                ,       // si es cualquier valor ok
        Some(variable) =>         ,       // si es cualquier valor ok (pero rescato la variable)
        None =>                   ,       // si hay error
    }
}

fn dividir(a: i32, b: i32) -> Option<i32> {
    if b == 0 {
        None
    } else {
        Some(a / b)
    }
}
```
Ejemplo 2: dentro de una estructura
```
    #[derive(Debug)]
    struct User {
        edad: Option<i32>,
    };
    let entidad = User {
        edad: None,
    };

    println!("{:?}", entidad);
```

funciones
```
.unwrap()             // obtendrá el valor directamente. solo usar si está seguro que tiene valor.
                      // en caso no tenga valor lanzará un 'panick!!!'
.unwrap_or(T)         // como unwrap(), pero si no tiene valor retornará el valor del parámetro
.expect("..");        // como unwrap(), pero agrega un mensaje descriptivo si lanza el 'panick'
```

## Result
Como el Option pero con la adición que permite especificar el error
```
enum Result<T, E> {
    Ok(T),
    Err(E)
}

let result: Result<i32,String> = Ok(2);                         // inicializa como 'ok' con el valor 2
let result: Result<i32,String> = Err(String::from("mensaje"));  // inicializa como 'error' con el mensaje

match resultado {
    Ok(variable) =>     ,       // si es ok
    Err(error) =>       ,       // si hay error
}

// extensiones
.expect("Obtiene el valor OK en sí, pero si hay un Error lanzará un Panic y mostrará ESTE MENSAJE")
```
ejemplo 1:
```
enum DivisionErrorType {
    ByCero,
    ByNegative,
}
fn division2(a: i32, b: i32) -> Result<i32, DivisionErrorType> {
    if b == 0 {
        Err(DivisionErrorType::ByCero)
    } else if b < 0 {
        Err(DivisionErrorType::ByNegative)
    } else {
        Ok(a / b)
    }
}

fn main() {
    let result = division2(3, 2);
    match result {
        Ok(1) =>                     ,          // si es ok con valor 1
        Ok(variable) =>              ,          // si es ok, asigno el valor a la variable
        Ok(_) =>                     ,          // si es ok con cualquier valor
        Err(DivisionErrorType::ByCero) =>     , // si es error pero 'ByCero'
        Err(DivisionErrorType::ByNegative) => , // si es error pero 'ByNegative'
    };
}
```

funciones
```
.unwrap()             // obtendrá el valor directamente. solo usar si está seguro que tiene valor.
                      // en caso no tenga valor lanzará un 'panick!!!'
.unwrap_or(T)         // como unwrap(), pero si no tiene valor retornará el valor del parámetro
.expect("error");     // como unwrap(), pero agrega un mensaje descriptivo si lanza el 'panick'
```

## Conversiones

```
let texto   = "12";
let num:i32 = texto.parse().unwrap();  // parse infiere el tipo destino de la conversión según la declaración de la variable
```

# ITERATOR
````
max_by_key(X)                   obtiene el máximo según X
    max_by_key(String::len)     obtiene el máximo valor de 'len' (len será calculado en cada item)
```