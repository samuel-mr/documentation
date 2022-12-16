# Install

- A) [rustup](https://rustup.rs/) Install + automatic updates. Like NVM (for Node) (recommended)
- B) [Official](https://www.rust-lang.org/learn/get-started) Official and manual installation

 ## Rustup tool
 ```
rustup show        # General info
rustup update      # update Rust installation
rustup doc --std   # open offline documentation
 ```

 [Managing diferent version](https://doc.bccnsoft.com/docs/rust-1.36.0-docs-html/edition-guide/rust-2018/rustup-for-managing-rust-versions.html)

## Check current installation
```
cargo --version
rustc --version
rustdoc --version
```

- `Cargo`: Adminsitrador de compilación, gestor de paquetes y administrador de proyectos
- `rustc`:  Es el compilador en sí. Llamado por `Cargo`
- `rustdoc`: Herramienta de documentación de Rust. Genera un HTML de los comentarios que agreguemos al código

# Definiciones
- Crate: Librería o ejecutable. Es una unidad de compilación. Reutilizable.
- Módulos: Administra ámbitos de código dentro de un crate
- Rutas: Dan nombre a los elementos del código. ejem: vector, función, módulo

# VS Code
## Config
Ocultar archivos generados de compilación 
```
{
    "files.exclude": {
        "**/*.exe": true,
        "**/*.pdb": true,
        "**/Cargo.lock": true,
        "**/Cargo.toml": true,
        "**/target": true
    }
}
```
## Extensions
- [rust-analizer](https://marketplace.visualstudio.com/items?itemName=matklad.rust-analyzer) sucesor de Rust (official-plugin)
- [CodeLLDB](https://marketplace.visualstudio.com/items?itemName=vadimcn.vscode-lldb) Debuguer nativo de C++, Rust y otros lenguajes compilados
- [Error Lens](https://marketplace.visualstudio.com/items?itemName=usernamehw.errorlens) Resealta en la misma línea de código los mensajes de error y advertencias
- [better TOML](https://marketplace.visualstudio.com/items?itemName=bungcip.better-toml) formateador y resaltador de sintaxis para los archivos de configuración de Rust (Cargo.toml)
- [crates](https://marketplace.visualstudio.com/items?itemName=serayuzgur.crates) Da información sobre las dependencias que usan Cargo.toml
- [Search crates.io](https://marketplace.visualstudio.com/items?itemName=belfz.search-crates-io) Muestra sugerencias de bibliotecas mientras se escribe en Cargo.toml
- [Rust Test Explorer](https://marketplace.visualstudio.com/items?itemName=swellaby.vscode-rust-test-adapter) Muestra el explorador de test para Rust

