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
let vector           = vec![1i32,2,3];  explícito B


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
.sort_by(|a, b| b.cmp(a))   // sort invertido
.append(&mut otro_vector);     agrega los valores de 'otro_vector', donde este último quedará vacio
.swap(i, j);                   intercambia ambos
.filter(|x| x % 2 == 0 )    'like where' ejem:solo números pares

["a", "b"].join("*"), "a*b"         concatena con un comodín
array[index..index + k].join("");   obtiene un rango de elementos interno y los concatena
```

Recorrerlos
```
    for i in vector {
```
Como parámetro
```
fn run(v: &mut Vec<char>) {
```
Modificando valores en otro método
```
fn run(arr: &mut Vec<Vec<char>>) {
    // get_mut : obtener una referencia mutable
    let mut item = arr.get_mut(i).unwrap();
    // es posible hacer esto porque 'item' es una referencia mutable
    item.pop()
```
## Str + Strings
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

- Permite manejar la falla desde un generic q tiene 2 posibilidades, ok y error
- Definición
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
- Ejemplo 1
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
- Ejemplo 2: dentro de una estructura
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
- if statement
    ```rust
        if let Some(value) = array.last() {     // si '.last()' tiene un valor
    ```
- if statement + `mut`
    ```rust
    fn print_type_of<T>(_: &T) {
        println!("{}", std::any::type_name::<T>())
    }

    pub fn version_1() {
        let mut foo: Option<Box<Vec<i32>>> = Some(Box::new(vec![1,2,3]));
        if let Some(mut x) = foo {
            println!("{:?}", x);
            print_type_of(&x); // alloc::boxed::Box<alloc::vec::Vec<i32>>
        }
    }
    pub fn version_2() {
        let mut foo: Option<Box<Vec<i32>>> = Some(Box::new(vec![1,2,3]));
        // debido al simbolo = asumirá que a la izquierda quieres un '&mut'
        if let Some(x) = &mut foo {
            println!("{:?}", x);
            print_type_of(&x); // &mut alloc::boxed::Box<alloc::vec::Vec<i32>>
        }
    }
    ```
- Funciones
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
