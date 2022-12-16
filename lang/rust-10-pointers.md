# Box
CASOS DE USO:
- Cuando tienes un valor grande, y quieres asegurarte que se use por referencia (porq quieres asegurar q no sea copiado)
- Cuando quieres usar un struct recursivo
- Cuando se quiere usar para declarar una interfaz

```rust
// 5 : será alojado en 'Heap'
// b : será alojado en 'Stack'
let b = Box::new(5);
```
with Struct
```rust
struct Truck {
    next_truck : Option<Box<Truck>>
}

// 2) 'a : agrego un indicador 'lifetime'
struct AnotherSuper<'a> {
    Data: (&'a str, i32), // 1) debido a str necesito agregar un inidicador de 'lifetime'
    Files: Box<Vec<AnotherSuper<'a>>>, // 3) como inicialmente al Struct se le agregó un 'a aquí también debo indicárselo
}

// 3) Otra alternativa para el punto '2': usar String (ya no se necesitará 'a)

struct AnotherSuper {
    Data: (String, i32),
    Files: Option<Vec<Box<AnotherSuper>>>,
}

```
with Enum
```rust

enum Directory {
    Cons(i32, Box<Directory>), // Con Box el compilador ya sabe el tamaño que ocupará ya que es un puntero y eso tiene tamaño conocido en el Stack
    Nil,
}

use Directory::{Cons, Nil};

...
// inicializando
let list = Box::new(Cons(1, Box::new(Cons(2, Box::new(Nil)))));
```