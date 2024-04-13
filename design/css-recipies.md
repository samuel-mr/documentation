# CSS Recipies

- Separete childrens
- Error: avoid to add margin/padding to the childrens.
- Solution: in their parent component add this:

```css
.somecss {
  display: flex;
  flex-direction: column;
  gap: 8px;
}
```
To pause the JS / CSS runtime : e.g. you wan't to see some element that only shown when an onOver method is raised
```cs
setTimeout(() => {debugger;}, 5000)
```