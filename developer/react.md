#
Comandos comunes
```
npm install // instala todas las dependencias del archivo package.json
npm start
npm run build
```
# paquetes interesantes
local http server
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

# useState
Cuando se cambia una variable el componente se renderiza
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
- se refiere a los `efectos secundarios` que pueden ser:
  - llamar a un api
  - hacer un cambio en el Dom
  - log

### useEffect
  - Ejecutará una función interna luego que se renderice
  - Se puede indicar si se ejecutará siempre, solo una vez o dependiendo del cambio de un valor
```
useEffect(()=> {

    getGifs()
}, [])                // [] indica q no tiene dependencia, solo se ejecutará 1 vez

         // si no tiene un segundo argumento se ejecutará siempre que se renderice
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
Limpieza: ejecutará lo que esté en 'return' cuando el componente se elimine
```
useEffect(()=> {

    return ( () => {  })
}) 
```

## useRef
- Agrega referencias en el código a controles del DOM
- La referencia persiste a los renderizados
- los cambios de valor no ejecutan el renderizado
```
// DECLARACION
const isMounted = useRef(true)
isMounted.current = false


// USO
const ref = useRef()
const onClickHandler = () => {
    ref.current.select()
}

return (
    <>
        <input
            ref={ref}
        />
        <button onClick={onClickHandler}>Focus</button>
    </>
)
```

## Memoization
- Guarda en cache un componenete para evitar volver a renderizarse
```
// SHOT-CUT
export const Small = React.memo(({ value }) => {
})

// USO
import { memo } from 'react'

export const Componenete = memo(() => {
    return (
        ...
    )
})
```

useMemo: lo mismo, permitiendo configurar cuando actualizar el valor memorizado
```
// variable: si cambia, volverá a ejecutar la función
const resultado = useMemo(() => procesoPesado(variable), [variable])
...

<p>{resultado}</p>
```


## useCallback
Casos de uso 1: Cuando envio una función como parámetro a un componente Memorizado. El useCallback evitará que el componente se vuelva a renderizar ya que la funcion 'increment' es declarada como 'useCallback'

en `App.js`
```
const increment = useCallback((PARAMETRO_ENVIADO) => {
    setCounter(x => x + PARAMETRO_ENVIADO)
}, [setCounter])

return (
    <>
        <span>{counter}</span>
        <ShowIncrement increment={increment} />
    </>
)
```
en `Componente.js`
```
export const ShowIncrement = React.memo(({ increment }) => {

    return (
        <button onClick={() => {
            increment(5)
        }}>Incrementar</button>
    )
})
```
# Función Pura
1. No tiene efectos secundarios (no manipula una bd, un archivo, algun estado de otro componente, etc)
2. Es síncrono
3. Retorna siempre un nuevo estado
4. Solo requier una acción que pudiera requerir un argumento

# Reducer
Contiene la lógica de todas la acciones centralizadamente, es llamado por un `Action`

### Caracteristicas:
1. Debe ser una función pura
2. Debe retornar un nuevo estado
3. Recibe 2 args (estado incial + función a ejecutar)


```
useReducer(toodReducer, initialState, init );

toodReducer : funcion que recibe el estado y la accion
initialState: recibe el estado inicial del componente
init        : recibe el estado para inicializar 

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
