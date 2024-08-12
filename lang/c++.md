# C++


### Print
```
cout << "\ntexto" << endl;      Imprime a consola

string almacen;
cin >> almacen;                 Pide que el usuario ingrese dato y será almacenado en 'almacen'
```

### Operadores
```
break                     Salta hacia la siguiente instrucción fuera del bucle
continue                  Salta hacia la siguiente iteración del bucle
```

### Punteros
```
& -> obtener la dirección de memoria de alguna variable
* -> obtener el valor en si de una dirección de memoria

char variable = 'a';
char *puntero = &variable;

(int*)&variable                convierte a int para que se la dirección de memoria sea legible
                                con otros tipos no necesita convertir

# Conversión
(int) VARIABLE                  convierte a int una variable
(int *) VARIABLE_TIPO_MEMORIA   convierte a int una variable tipo memoria
```

## Struct
```
struct Persona{ string nombre;    int edad; };

  Persona p;                          Instancia (forma 1)
  Persona p = Persona();              Instancia (forma 2)
  Persona *puntero = new Persona();   Crea puntero a futura instancia
  
  p.nombre                            Accede a nombre
  puntero->nombre;                    Accede a nombre desde una referencia
```

###  Union
```
 union numero_letra{                  
    int numero;
    char letra;
}

Uniones x = {'A'};                     Declara una variable que almacena un caracter                   
x.numero       // 65                   Obtiene el valor almancenado en su representación 'int'
x.letra        // A                    Obtiene el valor almacenado en su representación 'char'
(char)x.numero // A                    Como el primero + convierte a char
```

### Enum
```
    enum Workdays { lunes, martes, miercoles };
    Workdays midia = martes;
    midia;     // 1 
    
    
    enum Workdays { lunes = 'l', martes='m', miercoles'mm' };
    Workdays midia = martes;
    midia;     // 'm'
```

### Class
```
class MMM { }
public: string mipropiedad = "Hawking";    Propiedad pública
string mipropiedad = "Hawking";            Propiedad privada

// CONSTRUCTORES
string nombre;

Constructor(string nombre){               Constructor (forma 1)
  this->nombre = nombre;                  Para asignar a la propiedad
}

Constructor(string nam) : nombre(nam) {}  Constructor inline (forma 2)

// THIS
MiClase &miMetodo()                       & me sirve para indicar que el metodo retorna una posición de memoria
{
    return *this;                         *this retorno un puntero que apunta al objeto instanciado
}

// ABSTRACTOS
class Clase {
  void metodoAbstracto();                 Metodo sin implementar
}
void Clase::metodoAbstracto()             Implementándo el método
{
    ... implementación del método
}

// VARIABLES ESTATICAS
public static int numeroIdeas;            variable estática (definido dentro de la clase)
Clase::varibleEstatica= 1;                asignar un valor por default (hecho fuera de la clase que lo contiene)
void Clase::metodoNoImplementado()
{
    varibleEstatica+= 1;                  acceder a la variable estática (solo por método sobreestricto)
}

```
HWND -> Handle to a Window : Identificador único para una ventana.  Permite cambiar su título, posición, tamaño o estilo, enviar mensajes por el SO (como click del mouse), etc.
HINSTANCE -> instance Handle: Identificador único asignado a una instancia en ejecución de una app de Windows. Cuando se ejecuta un app el SO le asigna un HINSTANCE único. Se usa para identificar múltiples instancias de la misma app o entre diferentes apps
