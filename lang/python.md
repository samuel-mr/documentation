## Range
```
numeros = [1, 2, 3, 4, 5]
numeros = list(range(1, 6))
numeros = [*range(1, 6)]

```
## Loops
```
# for
frutas = ["manzana", "banana", "cereza"]
for fruta in frutas:
    print(fruta)

# range
for i in range(5):
    print(i)

# while
contador = 0
while contador < 5:
    print(contador)
    contador += 1

# do-while
contador = 0
while True:
    print(contador)
    contador += 1
    if contador >= 5:
        break

# for + index
frutas = ["manzana", "banana", "cereza"]
for indice, fruta in enumerate(frutas):
    print(f"Índice: {indice}, Fruta: {fruta}")
```

# Variables
Dictionary
```
my_dict = {}  # Create an empty dictionary

# Insert values and keys dynamically
key1 = "name"
value1 = "John"
my_dict[key1] = value1
```