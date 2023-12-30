Null-forgiving
```
// null! => indica solo para compile-time que el parámetro NO ES NULO (para evitar que muestre un mensaje de error ). Aplica cuando el proyecto no acepta nulos
 var person = new Person(null!)     
 
```
__init__: Permite que una propiedad inmutable se inicialice sin necesitad de un constructor sino también por medio de la construcción
```c#
struct Point
{
    public int X { get; init; }
    public int Y { get; init; }
}

var p = new Point() { X = 42, Y = 13 };

// antes, para una prop inmutable, se tenía que poner los valores en el constructuor
```


# Tips
```c#
 private T[] _array = Array.Empty<T>();
```
## Implicit operator
```c#
public struct Temperature
{
    public double Fahrenheit { get; set; }
    public double Celsius { get; set; }

    public Temperature(double fahrenheit)
    {
        Fahrenheit = fahrenheit;
        Celsius = (Fahrenheit - 32) * 5 / 9;
    }
    
    public static implicit operator Temperature(double fahrenheit)
    {
        return new Temperature(fahrenheit);
    }
}
...
Temperature temp = 100.0;  // Allowed because of the implicit operator.
```

# Nullable true
- To avoid warnings after verification add [NotNull]
```c#
    string? property
    Guard.NullOrEmpty(property, nameof(property));
 
    property, // Not warnings at this point


public static class Guard
{
    public static void NullOrEmpty(
        [NotNull] string? input, // Add 'NotNull'
        [CallerArgumentExpression("input")] string? parameterName = null)
    {
        if (String.IsNullOrWhiteSpace(input))
            throw new ArgumentException("Should have a value", parameterName);
    }
}
```
public static void NullOrEmpty(
        [NotNull] string? input,