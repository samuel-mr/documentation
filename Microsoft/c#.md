Null-forgiving
```
// null! => indica solo para compile-time que el parámetro NO ES NULO (para evitar que muestre un mensaje de error ). Aplica cuando el proyecto no acepta nulos
 var person = new Person(null!)     
 
```
__init__: Permite que una propiedad inmutable se inicialice sin necesitad de un constructor sino también por medio de la construcción
````
struct Point
{
    public int X { get; init; }
    public int Y { get; init; }
}

var p = new Point() { X = 42, Y = 13 };

// antes, para una prop inmutable, se tenía que poner los valores en el constructuor
```


# Tips
```
 private T[] _array = Array.Empty<T>();
```