# Shorcuts
```
preventDefault = () => { }
preventDefault(){ }
```
If (one line)
```
// el mismo efecto
item.isActive ? 'SI' : ''
item.isActive && 'SI'
```
Copiar toda la clase pero cambiar 'propiedad'
```
let instanciaClase = {
    ...todo,
    propiedad: 'nuevo valor',
}
```

Map
```
// Forma 1
return state.map((todo) => {
    if (1 == 1) {
        return todo;
    } else {
        return todo2;
    }
});

// Forma 2
return state.map(todo => 
    (1 == 1) ? todo : todo2
);
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
Array
```
Array.from({ length: 5 }, () => console.log("Mensaje"));
```
Find
```
    const foundPerson = persons.find(person => person.name === name);
```