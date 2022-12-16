

## Big-O notation:
- Describe la complejidad computacional.
- Representa el límite superior del costo del algoritmo. Describe la complejidada para entender el costo. 
- Intenta saber como el algoritmo escalará cuando el tamaño de inputs crezca.

### Basic concepts:
- __Operaciones__: numero de veces para realizar una operacion
- __Memoria__: cuánta memoria consume el algoritmo

### Complexity
- Time Complexity: Tiempo usado para ejecutarse relativo al tamaño del input
- Space Complexity: Memoria usada relativo al tamaño del input

> Usualmente: Time > Space

Cases:
- Best case scenario
- Avg  case
- Worst case scenario

> Normalmente serán iguales, pero hay algunos algoritmos donde en los 3 casos difiere. El Big(O) será el peor escenario.

## Time Complexity

> O(1) : Ideal, siempre será su uso de recursos constante, idiferentemente al input.

| Notación    | Crecimiento | Ejemplo                 | Descripción                                                                             |
|-------------|-------------|-------------------------|-----------------------------------------------------------------------------------------|
| O (1)       | Constante   | x                       | Siempre toma el mismo tiempo                                                            |
| O (Log n)   | Logarítmica | 10x items ~ 2x tiempo   | cantidad de trabajo se divide % 2 x cada iteración  . (ejem: binary search)             |
| O (n)       | Lineal      | 10x items = 10xtiempo   | tiempo de ejecución proporcional al tamaño de los datos. (ejem: for loop)               |
| O (n+m)     | Lineal      |                         | tiempo de ejecución proporcional al tamaño de los datos. (ejem: 2 for loop NO anidados) |
| O (n Log n) |             | 10x items ~ 20x tiempo  | tiempo de ejecución disminuye en cada iteración (ejem: Quicksort,mergesort, heapsort)   |
| O ( n*m )   |             |                         | Cuando hay 2 loops anidados, pero son 2 arrays distintos                                |
| O ( n^2 )   | Cuadrática  | 10x items = 100x tiempo | Cuando hay 2 loops anidados de un mismo array (ejem: dos for anidados, Bubble  sort)    |
| O ( 2^n )   | Exponencial | imagínate!              | Duplica su complejidad por cada elemento. ejem: Fibonaci F(n)=F(n-1)+F(n-2)             |
| O ( n! )    | Explota!    | ya ni lo pienses        | Crecimiento desmedido (ejem: traveling salesman)                                        |

![image](https://user-images.githubusercontent.com/2782135/160618696-d9478e46-e314-42c6-a00c-40502f937b37.png)


[Tabla general](https://www.bigocheatsheet.com/). Otra vista en [google docs](https://docs.google.com/spreadsheets/d/1hyxEEFvF5zBcpC3ALPVPyE8kJ1Soiwd4jpwKjHgzG6o/edit#gid=0)

https://medium.com/@mendozabridget777/big-o-notation-time-space-complexity-9bc31952052c

## Space Complexity
- No considerar la memoria usada por el input ni output