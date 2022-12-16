# Types

tipos de variables
```
List numbers =[1,true,'texto'];     // List dynamic
List<int>                           // List strong typed

Map persona = {}                    // Map dynamic
Map<String, dynamic> persona = {}   // Map strong typed
```

Listas
```
List            // dynamic
List<int>       // with type

final evenList = List.generate(5,(int index) => index*2);   // generaa 5 primeros numeros pares, inicia con 0

.add(...)            
list[1]     //get
```

final vs const
```
const PI = 3.14         // nunca cambia
final nombre = '...'    // asignable, luego no cambia
```
Map
```
  Map persona = {
    'texto' : 'do',
    'numero' : 2,
    'boleano' : true,
    1 : 100,
    2 : 200
  };

  persona['texto']      // get value stored as string
  persona[2]            // get value stored as number

// methods
  .addAll({3:300});
```
Strings
```
  var nombre = "gind";

// print
  'hola $nombre'       // direct
  'name: ${this.name}' // with two values
```
 # Import
 ```
import 'dart:math';
import 'dart:math' as mate;                           // with alias
import 'package:flutter/material.dart' show IconData; // import only IconData
 ```

 # Inheritance
 ```
 Extends
 Implements
 With
 ```
 
 # Future
 ```
 void main() {k 
  print('starting...');
  httpGet('https://api.nasa.com/aliens')
    .then((data){
      print(data.toString().toUpperCase());
    });
   print('finished');
}

Future httpGet(String url){
  return Future.delayed(Duration(seconds: 3), (){
    return "message delayed...";
  });
}
 ```
 # Async
 ```
 void main() async {
  print('1');
  final data = await get();
  print(data);
  print('3');
}

Future<String> get() async{
    return "2";  
}

...
// result
1
2
3
 ```