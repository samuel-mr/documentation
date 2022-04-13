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
