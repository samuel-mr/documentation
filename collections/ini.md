# Buenas prácticas
- Preferir la colección encapsulada sobre caulquier otra (favor encapsulated collection over any other????)
- Ocultar colecciones sobre una interfaz absctracta
## División Lógica
- Linear
    - Array
    - List
    - Stack
    - Queue

    - IList
- Asociative
    - Dictionary
    - HashSet

    - IDictionary
- Graphs
    - ProperGraphs
    - Trees

## Comparación en tiempo de ejecución
| Implementaciones List                                  | Implementaciones Dictionary                              |
|--------------------------------------------------------|----------------------------------------------------------|
| Está un Item al lado del otro                          | Asocia cada item con Keys                                |
| puede tener tiempo lectura constante                   | puede tener tiempo lectura constante                     |
| lectura/escritura aleatorio puede tener tiempo linear  | lectura/escritura aleatorio les toma el mismo tiempo     |
| insertar/eliminar al final usualmente tiempo constante | agregar/eliminar al final no es diferente a los demás en |
| Permite cualquier item | Puede evitar los duplicados |