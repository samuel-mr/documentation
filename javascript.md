# Shorcuts
```
preventDefault = () => { }
preventDefault(){ }
```
Copiar toda la clase pero cambiar 'propiedad'
```
let instanciaClase = {
    ...todo,
    propiedad: 'nuevo valor',
}
```
### Ver si hay data
```
null     // es null
!null    // es true
!!null   // es false

```

Null + ejecutar código
```
// Aquí SI ejecutará el código (negación simple)
!null && 'codigo a ejecutar'

// Aquí NO ejecutará el código (doble negación)
!!null && 'codigo a ejecutar'

```