me quede  aqui: https://app.pluralsight.com/course-player?clipId=837af101-3c33-4e88-ada9-766d0d36a1e3

# reglas
- El nombre de la función debe iniciar con la palabra **Test** 
- Los nombres de los archivos deben terminar con `_test.go` (estos archivos serán excluidos en el release)
- Agregar siempre este parámetro `t *testing.T`
### Estilos (pendiente entender como se diferencian)
- Caja negra:
    ```
    package main_test
    ```
- Caja blanca:
    ```
    package main 
    ```

# Simple
ejemplo simple (para ejecutarlo: )
- 1. crear módulo
    ```
    go mod init [name]
    ```
- 2. crear test
    ```
    package main_test

    import "testing"

    func TestAddition(t *testing.T) {
        result := 2 + 2
        expected := 4
        if result != expected {
            t.Errorf("No son iguales. Esperado: '%v' Obtenido: '%v'", expected, result)
        }
    }

    ```
- 3. ejecutar
    - `go test`
    - click derecho: "Go: test function at cursor"



## Paquetes
### I) Standar Library

| use               | Descripción                                         |
|-------------------|-----------------------------------------------------|
| testing           | paquete base para las pruebas                       |
| testing/quick     | casos específicos para los readers/writers          |
| testing/iotest    | serializa y deserializa                             |
| net/http/httptest | pruebas rest, simula request/response, test servers |

		
	
### II ) Comunidad
- [testify](https://github.com/stretchr/testify) Asserts
- [ginkgo](https://github.com/onsi/ginkgo) Testing con un estilo orientados al comportamiento
- [goconvey](https://github.com/smartystreets/goconvey) Genera Http con los resultados (muy visual)
- [httpexpect](https://github.com/gavv/httpexpect) Api test (muy popular en este campo)
- [mock](https://github.com/golang/mock) Mocking framework
- [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) Testing sql en memoria