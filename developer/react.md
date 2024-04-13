# Comandos comunes
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
            
            By Label
            <label>Contador {value}</label>

            By Text
            <label>Text {Texto internolabel>

            By TestId
            <Element data-testid='uuu' />
            screen.getByTestId(('uuu'))
            // Forma directa
            <button onClick={() => setValue(value + 1)}> +1 </button>
            // A traves de una function
            <button onClick={increment}> +1 </button>
        </>
    )
```

# Effect
- No pueden ser Async
- Se refiere a los `efectos secundarios` que pueden ser:
  - llamar a un api
  - hacer un cambio en el Dom
  - log

### useEffect
  - Ejecutará una función interna luego que se renderice
  - Se puede indicar si se ejecutará siempre, solo una vez o dependiendo del cambio de un valor
```
useEffect(()=> {

    getGifs()
}, [])                // [] solo se ejecutará 1 vez, indica q no tiene dependencia para ejecutarse denuevo.

         // si no tiene un segundo argumento se ejecutará siempre que se renderice
[]       // indica que no depende de nadie, solo se ejecutará 1 vez
[name]   // indica que depende de name, si name cambia se volverá a ejecutar
```

- Nota: Envio una copia de todo el componente. Ventaja: ya no usar props.title, etc
    ```

    <Componente
        {...item}

    // esto posibilida la desestructuración:
    export const ListItemImage = ({ title, url }) 
    ```
- Limpieza: ejecutará lo que esté en 'return' cuando el componente se elimine de la UI (unmounted)
    ```
    useEffect(()=> {

        return ( () => { // pendiente verificar las veces q se llama })
    }) 
    ```
Lanzará error:
```
const SomeComponent() => {
    const [value, setValue] = useState('');
    setValue(); // error, infinite loop, because it will recreate all the 'SomeComponent'
}
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
- Guarda en cache algo para evitar volver a renderizarse. Puede ser:
    - Componente 
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
    - Variable: permitiendo configurar cuando actualizar el valor memorizado
        ```
        // Si cambia 'variable' se volverá a ejecutar la función
        const resultado = useMemo(() => procesoPesado(variable), [variable])
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

- Consejos generales:
    1. Cuando pruebes algo que tomará cierto tiempo en 'procesar' has un 'expect' para que se asegure que ya se procesó
        ```javascript
        await userEvent.click(await screen.findByText("SOME_NAME"));
        expect(await screen.findByText("SOME_NEW_CONTENT")).toBeInTheDocument();
        ```

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
#### Search > Search : lookup from parent component
```javascript
// get the edit form
let form = screen.queryByTestId("household-form-edit");
expect(form).toBeInTheDocument();

// click in Edit button
await userEvent.click(await within(form).getByText("Edit", { selector: "button" }));

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

# Styled components
Create new element (from stratch)
```
styled.div`
```
Create element based on another
```
styled(div)`
```
Chained styles (cascade)
```javascript
const AA = styled.div`
& > div > button {
    background: red;
  }
  ...
   <AA>
      <div>
        <button> // applied red background
```
Chained style (elements)
```javascript
const AA = styled.div`
 div[role="tablist"] > button {
    background: red;
  }
```



# Testing library
- [doc](https://testing-library.com/docs/queries/about/)
```
# Buscan solo 1 elemento. Lanza error si encuentra más de 1
getBy
    Busca
    - Falla:
        - Si no encuentra
        - Si encuentra más de 1

queryBy
    Busca. Si no encuentra retorna null

findBy
    Busca, permitiendo esperar

# Busca muchos elementos
getAllBy
queryAllBy
findAllBy
```
Common query
```
By Rol
screen.getByRole('button', { name: 'Submit'});

By Label
screen.getByLabelText('Mi label');

By Text
screen.getByText('Texto interno');

By TestId
<Element data-testid='uuu' />
screen.getByTestId(('uuu'))
```
// some samples
```
 benefitsInput = screen.getByRole('combobox', { name: /Пільги %/ })
 ```
 To see all generated code in UI
 ```
 screen.logTestingPlaygroundURL();
 ```
 # Good Practice
 
 # Debug
 Timer
```
console.time('filter array');
const visibleTodos = getFilteredTodos(todos, filter);
console.timeEnd('filter array');
```

# Tips
Props
```javascript

// Error (because props could override previous properties)
<SomeComponent
  property1=...
  property2=...
  {...props} 

// OK
<SomeComponent
  {...props}
  property1=...
  property2=...
```

Async nature of React's state
```javascript
// Wrong (because the 'value' will be updated later, its in a queue)
setValue(value+1);
onSendTheNewValue(value);

// OK
let newValue = value + 1;
setValue(newValue);
onSendTheNewValue(newValue)
```
useState: in Childre and what happen if the value is changed in fater
```c#
const Fater() => {

    const [text, setText] = 'A'
   return (
    <Child someProp={text} />
    <Button onClick={() => setText('B')}    // ERRor : this action won't change 'someProp' in the Child component
  )
}
const Child(someProp) => {
    const [prop, setProp] = useState(someProp); // useState only init once, future change in parent component won't be reflected
}
```
change property content
```
// value
setCurrentValue( prep => (!prep));
// objec
setItems(prevItems => ({
      ...prevItems,
      somePropA: "a",
      somePropB: "b"
    }));
```
Add properties conditionally
```
let myObject = {
      propA: "1",
      ...(condition) && {
        propB : "2",
```
# Enums
    - Object.freeze to be each item inmutable
```
  const states = Object.freeze({
    Default: Symbol("default"),
    Discarding : Symbol("discarding"),
    Disenrolling : Symbol("disenrolling"),
  });
```
Conditional rendering, getIsSorted return asc or desc
```ts
{
    {
        asc: <ArrowUpIcon boxSize={5} />,
        desc: <ArrowDownIcon boxSize={5} />,
    }[header.column.getIsSorted() as string] ?? null
}
```
# Best Practice
- destructuring over props
```
// bad practice
function SomeFunc(props){
    //maybe you have: props.someSecretKey
    {Json.Stringify(props)} // security risk!
}
// ok
function SomeFunc({param1,param2})
    ...
```