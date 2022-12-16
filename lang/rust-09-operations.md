## Iterators
```rust
max_by_key(X)                   obtiene el máximo según X
    max_by_key(String::len)     obtiene el máximo valor de 'len' (len será calculado en cada item)
```

## Extensions
Quiero `NewType` se convierta a `i32`
```rust
use std::convert::*;

struct NewType(pub i32);

impl From<NewType> for i32 {
    fn from(src: NewType) -> i32 {
        src.0
    }
}

fn main() {
    let a = NewType(5);
    println!("{}", i32::from(a));
}
```

## Modules
Call between moduls
```rust
//Shared.rs

pub fn CallMe() {

//Another.rs
crate::shared::CallMe();
```

## Atributos
```rust
#![...]         All module
#[...]          Only Method
```

## Convert
Char
```rust
/*
A-Z  65-90
a-z  97-122
*/
let number = some_char as u32;

```
&str / String
```rust
let some_str      = line.parse::<i32>().unwrap_or(0) // turbo-fish version
let some_str: i32 = line.parse().unwrap(); // parse infiere el tipo destino de la conversión según la declaración de la variable
```
 