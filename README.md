# C++


### Print
```
cout << "\ntexto" << endl;      Imprime a consola

string almacen;
cin >> almacen;                 Pide que el usuario ingrese dato y será almacenado en 'almacen'
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

### Instanciar Union
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
