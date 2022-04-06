# FLUJOS DE CONTROL

## Condicionales

```
if a == b {
    ...
} else if a == c {
    ...
} else {
    ...
}
```
### Declaracion de arrays simplificado
- cada uno generará un array
```

    1..3               // 1,2
    1..=3              // 1,2,3
```
## Ciclos
para iterar: loop, for, while
```
loop {  }

let numeros = [5, 4, 3, 2, 1];
for item in numeros.iter() {
                              // item será cada número
}
    
for item in 1..3 { 

}


while [condición] {

}

```
Alteraciones de los ciclos
```
break;      finaliza el ciclo actual
```

## Ciclos - for
tipos de iteración
```
 let names = vec!["Bob", "Frank", "Ferris"];

 for name in names.into_iter() {  // recorre cada item y muda la propiedad 
 for name in names.iter()         // recorre cada item, pero solo se presta, names sigue siendo propietario
 for name in names.iter_mut()    // como iter() pero permitiendo mutar cada valor

 for item in numeros {}          // internalmente es: numeros.into_iter()
 
```

## Match
~ switch
```
match numero {
    1 =>        ,         si es 1
    2 =>        ,         
    3 =>        ,
    4 | 5 =>    ,          si es 4 o 5
    6..=10      ,          si está entre el 6 y el 10 (inclusivos)
    _ => 
};
```
match con return implícito
```
let resultado = match numero {  // el string final se asignará a la variable 'resultado'
    1 => "es uno",              // retorna str
    _ => "es otro numero"       // retorna str
};
```

# FUNCIONES
```
fn sumar(a: i32, b: i32) -> i32 {                  Tradicional
    return a + b;
}

fn sumar(a: i32, b: i32) -> i32 {                  Simplificado
    a + b
}

// ejemplo If Else
fn factorial(num: i32) -> i32 {                    Tradicional
    if num == 1 {
        return num;
    }
    return num * factorial_clasico(num - 1);
}

// ejemplos
fn factorial(num: i32) -> i32 {                     Simplificado
    if num == 1 {
        num
    } else {
        num * factorial_simplificado(num - 1)
    }
}
```

## Funciones en Estructuras

```
struct User {
    name: String,
}

impl User {                          // métodos para la estructura 'User'
    fn execGreeting(&self) {         // selft: para usar internamente las propiedades
        // self hace referencia a si mismo (como this)
        println!("Soy {}", self.name); 
    }
    fn changeName(&mut self, newName: String) {
        // &mut -> hace referencia a si mismo y especifica que es modificable
        self.name = newName;
    }
}

// Usando la esctructura y sus métodos
let mut usu = User {
    name: String::from("bach"),
};
usu.execGreeting();
usu.name = "mozart".to_string();
usu.execGreeting();
usu.changeName("beethoven".to_string());
usu.execGreeting();
```
