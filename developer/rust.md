# Rust !
> Seguridad, concurrencia y rendimiento

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
```
+ - * /     Numéricos

> < == !=   Relacionales

&&          Lógicos
||


1u32 - 2    !error, desbordamiento, porque el cálculo será negativo y el tipo resultado se infiere q es unsigned 32
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

# TIPOS PRIMITIVOS
- Primitivos Escalares
- Primitivos Compuestos
    - Tuplas (1,true)
    - Arreglos [1,2,3]

## Escalares
```
i8, i16, i32, i64, i128  enteros positivos y negativos
u8, u16, u32, u64, u128  enteros positivos
f32, f64                 flotantes
char                     UTF-8, podemos colocar emojis
bool
()                       unit type
```

### String
```
.to_lowercase()
```

## Arreglos
- La longitud es fija
- almacenado en el stack
- Solo puede almacenar 1 solo tipo de dato
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

- Permite prestar una sección de un arreglo
- No se le conoce el tamaño en tiempo de ejecución
- Almacenado en el Heap
```
    let mensaje = String::from("hola mundo");

    let inicio = &mensaje[0..4];          // indice inicial .. indice final
    let inicio = &mensaje[..4];           // idem

    let fin= &mensaje[4..mensaje.len()];  // inicia en 4 hasta el final
    let fin= &mensaje[4..];               // idem

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
numero    print:1
boleano   print: true
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
    Response::Error( _, mensaje) =>  ,           // si la segunda variable es un string, ingorará la primera variable
};
```

## Estructuras
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
### Estructura tipo Tupla:
```
#[derive(Debug)]
struct Color(u32, u32, u32);      // no serán propiedades sino valores de la tupla

let white = Color(255,255,255);

println!("El color es: {:?}", white);
```


# TIPOS SEGUN LIBRERIA STD
## VECTORES
- Son arrays redimensionables
- El tamaño no es conocido en tiempo de compilación
- vec! es un macro para inicializar

Inicialización
```
let vector           = vec![1,2,3];     Tipo implícito con data
let vector: Vec<i32> = vec![1,2,3];     Tipo explícito con data

let mut vector: Vec<i32> = Vec::new();  Tipo explícito
 
let mut vector = Vec::new();            Tipo postergado: sin tipo en la inicialización, pero...
vector.push(1);                         gracias a esta línea, en tiempo de diseño se infiere el tipo i32 del vector
```

```    
let vector = Vec::new();
let vector = vec![1, 2, 3];         implícito
#####let vector: vec<i32> = [1, 2, 3];   explícito

let ...          no modificable
let mut ...      modificable 

vector[1]       obtiene el elemento de alguna [posición]

METODOS
.push(1)        agregar al final
.insert(0, -1)  agrega en (posición, elemento)
.pop()          elimina el último (por defecto retorna un Option con el elemento eliminado)
.pop().unwrap() es el pop y además desempaqueta ese elemento eliminado
.remove(0)      elimina en (posición)

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


// Metodos
.to_string()                                        convierte srt a String
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


# FLUJOS DE CONTROL

## Condicionales

```
if a == b {
    ...
} else if a == c {
    ...
} else {
    ...
}
```

## Ciclos
para iterar: loop, for, while
```
loop {  }

let numeros = [5, 4, 3, 2, 1];
for item in numeros.iter() {
                              // item será cada número
}
    
for item in 1..3 {            // genera un array del 1 al 2 (el último no se considera)
                              // item será cada número
}

while [condición] {

}

```
Alteraciones de los ciclos
```
break;      finaliza el ciclo actual
```

## Match
~ switch
```
match numero {
    1 =>        ,         si es 1
    2 =>        ,         
    3 =>        ,
    4 | 5 =>    ,          si es 4 o 5
    6..=10      ,          si está entre el 6 y el 10 (inclusivos)
    _ => 
};
```
match con return implícito
```
let resultado = match numero {  // el string final se asignará a la variable 'resultado'
    1 => "es uno",              // retorna str
    _ => "es otro numero"       // retorna str
};
```

# FUNCIONES
```
fn sumar(a: i32, b: i32) -> i32 {                  Tradicional
    return a + b;
}

fn sumar(a: i32, b: i32) -> i32 {                  Simplificado
    a + b
}

// ejemplo If Else
fn factorial(num: i32) -> i32 {                    Tradicional
    if num == 1 {
        return num;
    }
    return num * factorial_clasico(num - 1);
}

// ejemplos
fn factorial(num: i32) -> i32 {                     Simplificado
    if num == 1 {
        num
    } else {
        num * factorial_simplificado(num - 1)
    }
}
```

## Funciones en Estructuras

```
struct User {
    name: String,
}

impl User {                          // métodos para la estructura 'User'
    fn execGreeting(&self) {         // selft: para usar internamente las propiedades
        // self hace referencia a si mismo (como this)
        println!("Soy {}", self.name); 
    }
    fn changeName(&mut self, newName: String) {
        // &mut -> hace referencia a si mismo y especifica que es modificable
        self.name = newName;
    }
}

// Usando la esctructura y sus métodos
let mut usu = User {
    name: String::from("bach"),
};
usu.execGreeting();
usu.name = "mozart".to_string();
usu.execGreeting();
usu.changeName("beethoven".to_string());
usu.execGreeting();
```
