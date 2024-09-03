# Basic Lang Features
- Markup
    ```html
    <script>
        let name = " "
    </script>
    <span>Hola {name}
    ```
    - Shorthand: Name and attributes are the same?
        ```
        <img src={src}
        <img {src} //the same
        ```
- Styles are scope to the component (even for child components)

    ```html
    <p>red...
    <style>
        p {
            color: red
        }
    ```

- Import components
    ```html
    <p>Principal</p>
    <Nested />
    <script>
        import Nested from './Nested.svelte'
    </script>
    ```

- Render Html
    ```html
    <script>
        let string = `contains <strong>HTML</strong>`;
    </script>

    <p>{@html string}</p>
    ```
- events
```html
<script>
	function increment() {	}
</script>

<button on:click={increment}>
```
# Reactivity
- `Reactivity`: re-run this $:code whenever any of the referenced values change
    - Execution order:
        -  script code => REACTIVE => markup render

    - Reactivity declarations: 
        ```html
            <script>
                let count = 0;
                $: double = count * 2 // Svelte inject a 'let' declaration
            </script>
            <span>{double}</span>
            <span>{count * 2}</span> #equals values
        ```
    - Reactivity statements
        ```html
        <script>
            let count = 0;

            $: if(count>2){
                console.log(`Por dos ${count*2}`)
                console.log(`Por tres ${count*3}`)
            }else{
                console.log('menor a dos')
            }
        </script>
        ```
    - Reactivity when update arrays and objects
        - same pathern for `pop, shift, unshift and splice`
        ```html
        <script>
            let numbers = [1, 2, 3, 4];

            function addNumber() {
                // IT DOESN'T WORK
                numbers.push(2);   
                // it works
                numbers = [...numbers,2] 
                // it works too (both)
                numbers.push(2);   
                numbers = numbers
                // it works
                numbers[0]= 9
            }
        ```
        - object:
            - doesn't work !!!
            ```html
                <script>
                let obj = {foo: {bar: 1}}

                function func() {
                    const foo = obj.foo
                    foo.bar = 2 // It doesn't trigger the reactivity
                    // one alternative solution is reasign
                        obj = obj // it'll work
            </script>
            ```
            - it works
            ```html
                <script>
                let obj = {foo: {bar: 1}}

                function func() {
                    obj.foo.bar = 2 
                } 
            </script>
            ```

# Props:
    ```
    <-- Child component -->
        <script>
            export let answer; // also, you could assign a default value
        </script>
    <- Parent component->
        <Nested answer={42} />
    ```

- `Spread props` : the prop names need to be equal
    ```html
    <PackageInfo
        name={pkg.name}
        speed={pkg.speed}
        website={pkg.website}
    />
    // OR 
    <PackageInfo {...pkg} />
    ```
# Logic
- if
    ```html
    {#if true}
    ...
    {:else if true}
    ...
    {:else}
    ...
    {/if}
    ```

- Iterators
    ```html
	{#each colors as color, i}
	<button	style="background: {color}">{i}</button>
	{/each}
    ```

    - To modify array : must add an ID
        ```html
        {#each items as item (item.id)}
        ```
- Await : eg: Promise runs on load and on click
    ```html
    <script>
        let promise = getRandomNumber();

        function handleClick() {
            promise = getRandomNumber();
        }
    </script>

    <button on:click={handleClick}>
        generate random number
    </button>

    {#await promise}
        <p>...waiting</p>
    {:then number}
        <p>Number is {number}</p>
    {:catch error}
        <p>{error.message}</p>
    {/await}
    ```
# Events
- DOM events
    ```html
    <script>
        function handleMove(event) {
            x = event.clientX;
            y = event.clientY;
            ...
        }
    </script>

    <div on:pointermove={handleMove}>
    ```
    - on:EVENT usecases
        - `keydown` : shortcuts, validate inputs
        - `submit` + `preventDefault` : validate form before send it
        - `input` : realtime search on keystroke
        - `change` : display selected value, filter based on selection
        - `mouse over`+`mouseout` : show tooltips, hover effects
        - `scroll` : infinite loading, animate some element at scrolling
        - `resize` : implement adaptative design
        - `touchstart,touchmove,touchend` : implement tactic interactions on movile
        - `dragstart, dragover, drop` : reorder, load file by droping
        - `online, offline`: Notify when connection changed
        - `pointermove` : advanced hover effect, hand-drawing apps
        - `volumechange` : update the volumen control
        - `error` : log error, manage loading resorce failures
- Inline handlers
    ```html
    <div 
        on:pointermove={(e) => {
        m = { x: e.clientX, y: e.clientY};
        }}
    >
    ```
    - Event modifier
        ```html
        <button on:click|once={() => ...}>
        ```
        - List:
            - `once` : remove the handler after the first time it runs
            - preventDefault
            - stopPropagation
            - passive
            - nonpassive
            - capture
            - self
            - trusted
        - Chain modifiers
            ```
            on:click|once|capture={...} 
            ```
- Component events : like `export function` ... _but it needs a dispatcher_
    - Child component
        ```html
        <script>
            import { createEventDispatcher} from 'svelte'
            const dispatch = createEventDispatcher();
            
            function internal_method_from_component() {
                dispatch('EVENT_NAME', {
                    CONTENT: 'Hello!'
                });
            }
        </script>

        <button on:click={internal_method_from_component}>
            Click to say hello
        </button>
        ```
    - Parent component
        ```html
        <script>
            import Inner from './Inner.svelte';

            function handleMessage(event) {
                alert(event.detail.CONTENT);
            }
        </script>

        <Child on:EVENT_NAME={handleMessage} />
        ```
    - Event forwarding (shorcut)
        - Traditional code
            ```html
            <script>
                import Inner from './Inner.svelte';
                import {createEventDispatcher} from 'svelte'

                const dispatcher = createEventDispatcher()
                
                // In this example the handler is just forwarding the same structure
                function help(event){
                    dispatcher('message', event.detail)
                }
            </script>
            <Inner on:message={help}/>
            ```
        - Shortcut (same behavior as above code) 
            ```html
            <script>
                import Inner from './Inner.svelte';
            </script>

            <Inner on:message/>
            ```
    
    - DOM Event forwarding
        ```html
        // Child component: notify the clic event to the parent
        <button on:click>
            ... 
        // Parent:handle it
        <BigRedButton on:click={handleClick}
        ```

    - Text Input
        ```html
        <input bind:value={texto} />
        <input bind:value={num} type="number" min="0" max="10" />
        <select bind:value={selected} on:change={() => (...)} />
        <textarea bind:value={value} />
        ```

    - Group inputs
        ```html  
            // Radio
            <input
                type="radio" name="radio_group_name" value={item}
                bind:group={selected_items} />

            // Checkbox
            <input
                type="checkbox" name="checkbox_group_name" value={item}
                bind:group={selected_items} />

            // Multiple 
            <select multiple bind:value={selected_items}>
                {#each items as item}
                    <option>{item}</option>
                {/each}
            </select>
        ```
# Lifecycle
- `onMount` : after first rendered
    ```html
        <script>
            import { onMount} from 'svelte'
            onMount(()=> {
                ...
                return () => {
                    ... clean ...
                }
            })
        </script>
    ```
- `beforeUpdate + afterUpdate`: lifecicle is `before` -> render -> `after`    
    - Use cases:
        - `Scroll` : updateing the scroll position https://learn.svelte.dev/tutorial/update
        - `Interaction with updated DOM`: adjust size or position of an element
        - `Mesurements` : that depends on a rendered content for instance: apply complex design that depends on children elements
        - `Animations`: post rendering
        - `Sync external data` : like update a map or chart, notify external services
        - `Validations`: verufy DOM integrity, check attributes
    - Don't use it on those scenarios, because there are build-in approaches
        - `reactive declarations`
            ```js
            // recalculated when count is updated
            $: doubledCount = count * 2;
            ```
        - `reactive effects`
            ```js
            $: {
                console.log(`El contador ahora es ${count}`);
                // run when count is updated
                updateSomeExternalLibrary(count);
            }
            ```
        - `Lifecicle functions`
            ```js
            let element;

            $: if (element) {
                // Este código se ejecuta cada vez que 'element' se monta o actualiza
                setupSomeExternalLibrary(element);
                return () => {
                    // Limpieza cuando el elemento se desmonta
                    cleanupExternalLibrary(element);
                };
            }
            ```
- `Tick` a promise that resolves ASAP as any pending changes have been applied to the DOM. So it await til the DOM changes are applied. [Oficial example](https://learn.svelte.dev/tutorial/tick)
    - before
        ```html
        <script>
            let count = 0
            $: double = count * 2

            function increase(){
                count++
                console.log(double) // print 0 ... why? because double will be evaluated after
            }
        </script>
        ```
    - after
        ```html
        <script>
            import {tick} from 'svelte'
            let count = 0
            $: double = count * 2

            // version 1: callback
            function increase(){
                count++
                console.log('before tick', double) // print 0
                tick().then(()=> {
                    console.log('after tick', double) // print 2
                    console.log(double)
                })
            }

            // version 2: async
            async function increase(){
                count++
                console.log('before tick', double) // print 0
                await tick()
                console.log('after tick', double) // print 2
            }
        </script>
        ```
    - others common use cases (not test ...jet)
        ```js
        // Ejemplo 1: Medir altura después de actualizar contenido
        async function updateAndMeasure() {
            content = newContent; // Actualiza el estado
            await tick();
            console.log(element.clientHeight); // Mide la altura después de la actualización
        }

        // Ejemplo 2: Enfocar un input después de mostrarlo
        async function showAndFocusInput() {
            showInput = true; // Muestra el input
            await tick();
            inputElement.focus(); // Enfoca el input después de que se haya renderizado
        }

        // Ejemplo 3: Inicializar una biblioteca de terceros después de renderizar
        async function initializeChart() {
            chartData = newData; // Actualiza los datos del gráfico
            await tick();
            new Chart(chartElement, { data: chartData }); // Inicializa el gráfico con los nuevos datos
        }

        // Ejemplo 4: Scroll a un elemento recién añadido
        async function addItemAndScroll() {
            items = [...items, newItem]; // Añade un nuevo item a la lista
            await tick();
            lastItemElement.scrollIntoView(); // Hace scroll al último item añadido
        }

        // Ejemplo 5: Actualizar un canvas después de cambiar su tamaño
        async function resizeAndUpdateCanvas() {
            canvasSize = newSize; // Cambia el tamaño del canvas
            await tick();
            const ctx = canvas.getContext('2d');
            ctx.fillRect(0, 0, canvas.width, canvas.height); // Dibuja en el canvas redimensionado
        }
        ```
# Stores
- `Writable stores` : like global variables
    ```js
    // store.js
    import { writable } from 'svelte/store';
    export const count = writable(0);

    // all components needs 
    import { count } from './stores.js';

    // Some 'component.svelte'
    count.update(n=> n+1) // includes current value
    count.set(0) // direct assignment

    // Manual-suscription
    const unsubscribe =count.subscribe((value) => {
		my_internal_variable = value;
	});
	onDestroy(unsubscribe)
    <span>{my_internal_variable}</span>

    // Auto-subscription (also, usable in script section)
    <span>{$count}</span>
    ```
> $ is a reserved character for a referenced store 
    
- `Readable stores` : read-only global variables like current time, geolocation, mouse possition
    ```js
    export const time = readable(
        // Initial value: could be null or undefined
        new Date(), 
        // Called when the store gets its first subscriber
        function start(set) {
            ...

            // Called when the las subscriber unsbscribes of the store
            return function stop() {
                ...
            };
    });

- `Derived stores` : stores based on another store. like when you want to display how long have you been in a webpage
    ```js
    export const elapsed = derived(
        time, // I'm not pretty sure about the first parameter
        (another_store_variable) => Math.round((another_store_variable - start) / 1000) // another_store_variable is int
    );
    ```html

- `Custom Stores` : stores + logic. To have explicit behavior and to avoid expose `set` and `update`
    ```js
    // store.js
    import { writable } from 'svelte/store';

    function createCount() {
        // extract wich one needs
        const { subscribe, set, update } = writable(0);

        return {
            subscribe, // must
            increment: () => update((n) => n + 1),
            decrement: () => update(n => n - 1),
            reset: () => set(0)
        };
    }

    export const count = createCount();

    // component.js
    <script>
        import { count } from './stores.js';
    </script>

    <span>is {$count}</span>

    <button on:click={count.increment}>+</button>
    <button on:click={count.decrement}>-</button>
    <button on:click={count.reset}>reset</button>
    ```
- `Store bindings` : Writable stores + Derived stores 
    ```js
    // store.js
    import { writable, derived } from 'svelte/store';
    export const name = writable('world');
    export const greeting = derived(name, (current_value) => `Hello ${current_value}!`);

    // component.svelte
    <script>
        import { name, greeting } from './stores.js'; // export both stores
    </script>

    <h1>{$greeting}</h1>
    <input bind:value={$name} />

    <button on:click={()=> name.update(value => value + '!' )}> // update the writable store
        Add exclamation mark
    </button>

        // another variant for on:click is
            on:click={()=> name.set($name + '!' )}
        // another variant is
        on:click={()=> $name += '!'}
    ```
    
    ```html

    ```
    
    ```html

    ```
    