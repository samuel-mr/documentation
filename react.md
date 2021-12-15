'Effect' se refiere a los *efectos secundarios*, y pueden ser:
- llamar a un api
- hacer un cambio en el Dom
- log

useEffect: 
  - Se ejecutará luego que se renderice
```
useEffect(()=> {

    getGifs()
}, [])                // [] indica q no tiene dependencia, solo se ejecutará 1 vez
```
