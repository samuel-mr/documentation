#RUST!
```
let      => inmutable
let mut  => mutable
const PI => inmutable constante

const PI = 3.14
let mut edad = 54

```

```
let numero = 2      # implícito
let numero:i32 = 2  # explícito

```

## Shadowing
Capacidad para declara una variable en un mismo escope o en un scope interno. La nueva variable reemplazaría a la antigua
``````
  let x = 1;
  {
      x             // print 1
      let x = "abc"; 
      x             // print 'abc'
  }
   x                // print 1
  let x = true;
  x                 // print true
``````


``````


``````
# Visual Studio Code Extensions
- [rust-analizer](https://marketplace.visualstudio.com/items?itemName=matklad.rust-analyzer) sucesor de Rust (official-plugin)
- [better TOML](https://marketplace.visualstudio.com/items?itemName=bungcip.better-toml) formateador y resaltador de sintaxis para los archivos de configuración de Rust
- [crates](https://marketplace.visualstudio.com/items?itemName=serayuzgur.crates) Da información sobre las dependencias que usan Cargo.toml
- [Search crates.io](https://marketplace.visualstudio.com/items?itemName=belfz.search-crates-io) Muestra sugerencias de bibliotecas mientras se escribe en Cargo.toml
- [Rust Test Explorer](https://marketplace.visualstudio.com/items?itemName=swellaby.vscode-rust-test-adapter) Muestra el explorador de test para Rust
- [CodeLLDB](https://marketplace.visualstudio.com/items?itemName=vadimcn.vscode-lldb) Debuguer nativo de C++, Rust y otros lenguajes compilados
