### Recursos:
- __Operaciones__: numero de veces para realizar una operacion
- __Memoria__: cuánta memoria consume el algoritmo

## Big-O notation:
Representa el límite superior del costo del algoritmo. Describe la complejidada para entender el costo. Intenta saber como el algoritmo escalará cuando el tamaño de inputs crezca.


| Notación   | Crecimiento      | Ejemplo                     | Descripción                                                                               |
|-------------|-------------|-------------------------|-------------------------------------------------------------------------------------------|
| O (1)       | Constante   | x                       | Siempre toma el mismo tiempo                                                              |
| O (n)       | Lineal      | 10x items = 10xtiempo   | tiempo de ejecución proporcional al tamaño de los datos                                   |
| O (Log n)   | Logarítmica | 10x items ~ 2x tiempo   | cantidad de trabajo se divide % 2 x cada iteración                                        |
| O (n Log n) |             | 10x items ~ 20x tiempo  | Divide el problema en varios y luego los recompone (ejem: Quicksort,mergesort, heapsort)  |
| O ( n^2 )   | Cuadrática  | 10x items = 100x tiempo | Con Loops anidados. Cuando cada elemento debe recorrer todos el array Ejem:(Bubble  sort) |
| O ( 2^n )   | Exponencial | imagínate!              | Duplica su complejidad por cada elemento. ejem: Fibonaci F(n)=F(n-1)+F(n-2)               |
| O ( n! )    | Explota!    | ya ni lo pienses        | Crecimiento desmedido. Nunca debería existir.                                                                    |

![image](https://user-images.githubusercontent.com/2782135/160618696-d9478e46-e314-42c6-a00c-40502f937b37.png)


[Tabla general](https://www.bigocheatsheet.com/). Otra vista en [google docs](https://docs.google.com/spreadsheets/d/1hyxEEFvF5zBcpC3ALPVPyE8kJ1Soiwd4jpwKjHgzG6o/edit#gid=0)

https://medium.com/@mendozabridget777/big-o-notation-time-space-complexity-9bc31952052c