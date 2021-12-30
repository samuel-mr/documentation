#
Comandos comunes
```
npm install // instala todas las dependencias del archivo package.json
npm start
npm run build
```
# paquetes interesantes
server http para testing
```
npm install --global http-server
// in build folder
http-server 
  -o         // launch new browser
```

___
# Hooks
___
- Son funciones que nos ayuda a manejar los estados entre los componentes
- Solo usables en componentes como funciones, no como clases
- [Usando el Hook Effect](https://es.reactjs.org/docs/hooks-effect.html)
- [Preguntas frecuentes](https://es.reactjs.org/docs/hooks-faq.html#is-it-safe-to-omit-functions-from-the-list-of-dependencies)

# State
```
    const [value, setValue] = useState(0)
    const increment = () => setValue(value + 1)


    return (
        <>
            <label>Contador {value}</label>
            // Forma directa
            <button onClick={() => setValue(value + 1)}> +1 </button>
            // A traves de una function
            <button onClick={increment}> +1 </button>
        </>
    )
```

# Effect
- No pueden ser Async
- se refiere a los `efectos secundarios` que pueden pueden ser:
  - llamar a un api
  - hacer un cambio en el Dom
  - log

### useEffect
  - Ejecutará una función luego que se renderice
```
useEffect(()=> {

    getGifs()
}, [])                // [] indica q no tiene dependencia, solo se ejecutará 1 vez

[]       // indica que no depende de nadie, no volverá a ejecutarse
[name]   // indica que depende de name, si name cambia se volverá a ejecutar
```

Envio una copia de todo el componente. Ventaja: ya no usar props.title, etc
```

  <Componente
      {...item}

// esto posibilida la desestructuración:
export const ListItemImage = ({ title, url }) 
```

# Test (Enzyme)

### Cycles : Se reejecuta por cada test (para limpiar variables globales)
```
beforeEach(() =>{
    jest.clearAllMocks();
    wrapper = shallow(<AddCategory setCategory={setCategories} />)
})
```
### Declaration
```
// jest.fn es un wrapper para analizar un función
const FUNCION = jest.fn()

// uso
expect(FUNCION).not.toHaveBeenCalled()
expect(FUNCION).toHaveBeenCalledTimes(1)
expect(FUNCION).toHaveBeenCalledWith(expect.any(Function)) //q el argumento pasado halla sido una función
```

## Search
```
const input = wrapper.find('input')

```
inserta texto en input
```
input.simulate('change', { target: { value: INPUT_INGRESADO } })


```

Propiedades de element donde `element = wrapper.find('input')`
```
.text()           // obtiene el texto interno
.prop('PROPERTY') // obtiene una propiedad del elemento
.exists()         // que exista: true or false
```

### Test
```
expect(...).toEqual([])     // igualar a array vacio
expect(...).toBeTruthy()    // q sea verdadero
expect(...).toBeFalsy()     // q sea falso


expect(wrapper).toMatchSnapshot()

expect(HTML_ELEMENT.prop('PROPERTY')).toBe(...)
```

# Test (React Hooks)
```
npm install --save-dev @testing-library/react-hooks
```
in each file
```
import { renderHook, act } from '@testing-library/react-hooks'
```


uso (ejemplo pequeño y completo)
```
const {result} = renderHook(() => useFetchGifs())
const {data, loading} = result.current

expect(data).toEqual([])
expect(loading).toBeTruthy()
```

# Publish GitHub
