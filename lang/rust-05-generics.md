# Generics
Aplicado en -> functions, structs
- Simple
    ```
    fn process<T>(num: &[T]) -> T {
    ```
- Múltiple (ejem: PartialOrd + Copy) ... para este ejemplo las 2 funciones son idénticas
    ```
    fn process<T: std::cmp::PartialOrd + std::marker::Copy>(num: &[T]) -> T {

    fn process<T>(num: &[T]) -> T
    where
        T: std::cmp::PartialOrd + std::marker::Copy
    {
    ```

Restringir parámetros indicando que interfaz debe implementar
```
use std::string::ToString;

// item será de un tipo que ha implementado el 'Trait' ToSring
fn super_metodo<T: ToString> (item: &T) -> String {
    item.to_string()
}

// uso
println!("{}", super_metodo(&42));
println!("{}", super_metodo(&"hello"));
```