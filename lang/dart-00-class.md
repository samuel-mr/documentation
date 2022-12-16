
# Dart - Class
// create instance
```
final instance = Heroe();      // implicit
final instance = new Heroe();  // explicit

```
override
```

class Heroe{
  
  @override
  String toString() {
    return 'name: ${this.name}';
  }
}
```
### constructors
strict order
```
class Heroe{
  String Name;  
  String Power;
  
  Heroe(this.Name, this.Power);
}

final wolverine = new Heroe('Logan', 'Regeneration');
```
unnorder parameters
```
class Heroe{
  String Name;  
  String Power;
  
  Heroe({
    required this.Name,
    required this.Power
  });
}
... new Heroe(Name:'Logan',Power: 'Regeneration');
}
```

custom parameters
```
class Person{
  String? name;
  String? last;
  
  Person.fromJson(Map<String,String> json){
    this.name = json['Name']!;
    this.last = json['LastName']!;
  }
}

// using
  final musician = {
    'Name' : 'bettwoven',
    'LastName' : 'ludwing'
  };

  final myMusician = Person.fromJson(musician);
```

custom parameters (version 2)
```
class Person{
  String name;
  String last;
  
  // : mean this will be excecuted at the begining
  Person.fromJson(Map<String,String> json) :
    this.name = json['Name']!,
    this.last = json['LastName']!;
  
}
```
Getters
```
class Person{
  
  String get texto{
    return '111';
  }

  set texto(String newText){
    print('ingresando nuevo texto');
  }
}

person.texto          // 111
person.texto = 'aaa'  // setter
```