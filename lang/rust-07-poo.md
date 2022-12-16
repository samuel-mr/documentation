# Modules
Split code in logical units
 ### Structure visualizer
 ```
 cargo install cargo-modules
 ```
 ```
 cargo modules generate tree
    --with-types
 ```
# Class

> Function != Method

## Method
```
Struct Rectangle{ ...

impl Rectangle{ ...
    fn methodo(){ ...

```

# Traits (intefaces)
### PartialEq (equality comparison)
- Permite hacer comparaciones == y !=
- Implementará el método 'eq'
- Si es aplicado sobre una estructura, será igual solo si todos los campos son iguales a los campos de la otra estructura
```rust
#[derive(PartialEq)]
enum BookFormat {
```
## Common traits

Debug
```rust
#[derive(Debug)]
struct SomeName {
...
println!("{:?}", variable);
```

Display
```rust
impl Display for Directory {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        write!(f,...)
    }
}
...
// 1: 
println!("{variable}");
// 2:
variable.to_string();
```