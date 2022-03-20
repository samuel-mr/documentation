# Rust Configuration
> VS Code


## Cargo
Gestor de paquetes + administrador de proyectos

# VS Code - Config
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
# VS Code - Extensions
- [rust-analizer](https://marketplace.visualstudio.com/items?itemName=matklad.rust-analyzer) sucesor de Rust (official-plugin)
- [better TOML](https://marketplace.visualstudio.com/items?itemName=bungcip.better-toml) formateador y resaltador de sintaxis para los archivos de configuración de Rust
- [crates](https://marketplace.visualstudio.com/items?itemName=serayuzgur.crates) Da información sobre las dependencias que usan Cargo.toml
- [Search crates.io](https://marketplace.visualstudio.com/items?itemName=belfz.search-crates-io) Muestra sugerencias de bibliotecas mientras se escribe en Cargo.toml
- [Rust Test Explorer](https://marketplace.visualstudio.com/items?itemName=swellaby.vscode-rust-test-adapter) Muestra el explorador de test para Rust
- [CodeLLDB](https://marketplace.visualstudio.com/items?itemName=vadimcn.vscode-lldb) Debuguer nativo de C++, Rust y otros lenguajes compilados

