# Dart - Basics

void / fun
```
void main() { }
main() { }
```
comentarios
```
// 
/* */
```
declarar
```
var nombre = "gind";      //implícito
String nombre = "gind";   //explícito


```
 
Condicionales
```
if(ok)
if(!ok)
```

Nullable
```
  bool? flag = null;
  if(flag == null) ...
```

Operadores
```
==
!= 
??        // if is null, asign this ...
```
## Methods 
Parameters
```dart
// required
void run(String parameter)

// optional
void run([String parameter='default'])

// any order
void run({
    String first  = 'defualt1', 
    String second = 'default2'
  }){

run(second: '2', first: '1');
}

// any order without default values
void run({
    required String first, 
    required String second
  }){
```
## late
Cuando declaras sin inicializar,
pero aseguras que tendrá valor antes de usarse
```dart
late String nombre; // Declaración

void main() {
  nombre = "Juan"; // Inicialización posterior
  print(nombre); // Imprime: Juan
...

```