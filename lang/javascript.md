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
# Bucles
```js
const array = [1, 2, 3, 4, 5];

// for
for (let i = 0; i < array.length; i++) {
  console.log(array[i]);
}
// for ... of
for (const element of array) {
  console.log(element);
}
// forEach
array.forEach(element => {
  console.log(element);
});

// map
const duplicate = array.map(element => {
  return element * 2;
});

// filter
const onlyGreaterThan2 = array.filter(element => {
  return element > 2;
});

// find (return the first match)
const onlyGreaterThan2 = array.find(element => {
  return element > 2;
});


// reduce
const sum = array.reduce((accumulator, currentValue) => {
  return accumulator + currentValue;
}, 0);


```
 
# Sort
```javascript
    const people = [
        { name: 'John', age: 30, gender: 'male' },
        { name: 'Jane', age: 25, gender: 'female' },
        { name: 'Bob', age: 40, gender: 'male' },
        { name: 'Alice', age: 35, gender: 'female' },
        { name: 'Mike', age: 28, gender: 'male' }
      ];
    // sort by MALE first
    people.sort((a,b) => {
        // -1 : left need to be first
        //  1 : right need to be first
        //  0 : remain the same
        if(a.gender === 'male') return -1;
        return 1;
    });
    console.table(people);
┌─────────┬─────────┬─────┬──────────┐
│ (index) │  name   │ age │  gender  │
├─────────┼─────────┼─────┼──────────┤
│    0    │ 'Mike'  │ 28  │  'male'  │
│    1    │  'Bob'  │ 40  │  'male'  │
│    2    │ 'John'  │ 30  │  'male'  │
│    3    │ 'Jane'  │ 25  │ 'female' │
│    4    │ 'Alice' │ 35  │ 'female' │
└─────────┴─────────┴─────┴──────────┘
```