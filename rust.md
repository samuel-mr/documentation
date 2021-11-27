# Rust !
> Seguridad, concurrencia y rendimiento

### Bloques
- Las variables se crean, modifican y destruyen dentro de su bloque
- Una variable está disponible dentro de su bloque y los bloques anidados
- Una variable no está disponible en un bloque superior

Alcance de las variables
```
fn main() {
    // bloque 'A'
    {
        // bloque 'B' ... aquí también están disponibles las variables de 'A'
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
Declaración
```
let      => inmutable
let mut  => mutable
const PI => inmutable constante

const PI = 3.14
let mut edad = 54

```

```
let numero = 2        # implícito
let numero:i32 = 2    # explícito

let entero = 1        # i32 inferido
let real   = 1.1      # f64 inferido

let millon = 1_000_000u32 # número más legible

```
# Tipos
### Tipos Escalares
```
i8, i16, i32, i64, i128  enteros positivos y negativos
u8, u16, u32, u64, u128  enteros positivos
f32, f64                 flotantes
char                     UTF-8, podemos colocar emojis
bool
()                       unit type
```

### Tipos Compuestos
```
[1,2,3]       arrays
(1,true)      truplas
```

### String
```
.to_lowercase()
```

### Arreglos
- La longitud es  fija
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
push(1)        agregar al final
insert(0, -1)  agrega en (posición, elemento)
pop()          elimina el último (por defecto retorna un Option con el elemento eliminado)
pop().unwrap() es el pop y además desempaqueta ese elemento eliminado
remove(0)      elimina en (posición)

```
# Operadores
```
+ - * /     Numéricos

> < == !=   Relacionales

&&          Lógicos
||


1u32 - 2    !error, desbordamiento, porque el cálculo será negativo y el tipo resultado se infiere q es unsigned 32
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

## Conversiones

```
let texto   = "12";
let num:i32 = texto.parse().unwrap();  // parse infiere el tipo destino de la conversión según la declaración de la variable
```


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
match con return 
```
let resultado = match numero {  // el string final se asignará a la variable 'resultado'
    1 => "es uno",
    _ => "es otro numero"
};
```



## Visual Studio Code Extensions
- [rust-analizer](https://marketplace.visualstudio.com/items?itemName=matklad.rust-analyzer) sucesor de Rust (official-plugin)
- [better TOML](https://marketplace.visualstudio.com/items?itemName=bungcip.better-toml) formateador y resaltador de sintaxis para los archivos de configuración de Rust
- [crates](https://marketplace.visualstudio.com/items?itemName=serayuzgur.crates) Da información sobre las dependencias que usan Cargo.toml
- [Search crates.io](https://marketplace.visualstudio.com/items?itemName=belfz.search-crates-io) Muestra sugerencias de bibliotecas mientras se escribe en Cargo.toml
- [Rust Test Explorer](https://marketplace.visualstudio.com/items?itemName=swellaby.vscode-rust-test-adapter) Muestra el explorador de test para Rust
- [CodeLLDB](https://marketplace.visualstudio.com/items?itemName=vadimcn.vscode-lldb) Debuguer nativo de C++, Rust y otros lenguajes compilados
