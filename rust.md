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
Capacidad para declara una variable en un mismo escope
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
