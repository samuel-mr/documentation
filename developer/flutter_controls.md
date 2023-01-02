# Controls
Texto
```
Text(
    "Starwars: el retorno del nuevo Jedi",
    textAlign: TextAlign.center,
    maxLines: 2,
    overflow: TextOverflow.ellipsis,  // coloca ... si sobrepasa el límite visual
    style: Theme.of(context).textTheme.headline5,
  )
```

otros
```
SizedBox (like Box)


Padding(
    padding: const EdgeInsets.only(right: 8.0)
    
Container(
        alignment: AlignmentDirectional.centerEnd,
        padding: const EdgeInsets.only(right: 20, top: 10, bottom: 10),
        child:

ElevatedButton(
        style: ElevatedButton.styleFrom(
        backgroundColor: Colors.red,
        shape: const StadiumBorder(),
        elevation: 0),
        child: const Text('presss'),
        onPressed: () {},
)

Container(
        margin: EdgeInsets.only(right: 15),

          width: 100,
          height: 100,

        // para usar el borde + color
          decoration: BoxDecoration(
            color: _color,
            borderRadius: _borderRadius,
          ),

```
## Add Borders
```
ClipRRect(
            borderRadius: BorderRadius.circular(5),
            child: FadeInImage(
              fit: BoxFit.cover //  WARNING: no se aun porque pero solo con esta propiedad se nota el 'borde circular'
```
## Recipies
Ocupa todo el ancho de la pantalla => width: double.infinity
```
SomeControl(
        width: double.infinity,
```
## AppBar

- CustomScrollView: scroll view personalizados
  ```
  CustomScrollView(
        slivers: [... only Silvers control, like SliverAppBar...],
      )
  ```
  - ejem:
    ````
    // tanto el header como el body deben ser componentes 'Silvers'
    slivers: [
        // HEADER
          _CustomAppBar(),  //retorna en algun lugar un return SliverAppBar(
        // BODY
          SliverList(
            delegate: SliverChildListDelegate(
              [
                // Cualquier tipo de control
                Text('Hola mundo'),
              ],
            ),
          )
        ],
    ```
- SliverAppBar: customized width + heigth. El bar en si no es fijo, sino será parte del body. 
  - Diferencias entre plataformas:
    - IOS: a pesar que no hay contenido en el body, siempre el bar hace la animación.
    - Android: solo hace la animación si hay contenido que supera la pantalla.
  - ejem: (app peliculas)
    ```
    SliverAppBar(
        backgroundColor: Colors.red,
        expandedHeight: 200,
        floating: false,
        pinned: true,
        flexibleSpace: FlexibleSpaceBar(
          centerTitle: true,
          titlePadding: EdgeInsets.only(bottom: 0),
          title: Container(
            width: double.infinity,
            alignment: Alignment.bottomCenter,
            color: Colors.black38, //alpha 38%
            child: Text(
              '$movie_title',
              style: TextStyle(fontSize: 16),
            ),
          ),
          background: FadeInImage(
            placeholder: AssetImage('assets/loading.gif'),
            image: NetworkImage("https://picsum.photos/id/1/500/200"),
            fit: BoxFit.cover,
          ),
        ),
      );
    ```
# Containers

### Columns & rows
```
crossAxisAlignment: CrossAxisAlignment.start,
mainAxisAlignment: MainAxisAlignment.center,

margin: EdgeInsets.symmetric(horizontal: 10, vertical: 10)
padding:  EdgeInsets.only(top:  20),
```
## List

| Shortcut              | When use it?                 |
|-----------------------|------------------------------|
| Column                | Hijos con el mismo 'height'  |
| SingleChildScrollView | Hijos con diferente 'height' |

- scroll when overflow: Genera un scroll si el hijo supera el tamaño de la pantalla
```

// 1. SingleChildScrollView a toda la pantalla:
        SingleChildScrollView(
           child: Column(
                children: [
                        Image(

// 2. SingleChildScrollView solo a la imagen
        SingleChildScrollView(
           child: Column(
                children: [

                // para que ocupete todo lo alto de la pantalla
                // motivo: está contenido dentro de un Column, y su tamaño es relativo
                Expanded(
                    child: SingleChildScrollView(
                        child: Image(
```

### ListTile vs NO ListTile
- Checkbox vs CheckboxListTile : el segundo tiene todo el row como área clickeable para ser seleccionado
- Switch   vs SwitchListTile   : idem que el anterior

## Controls - adaptative
Algunos controles tienen la propiedad `adaptive` para que se muestre la versión automáticametne su versión IOS / Android
```
Slider.adaptive
```
## ListView
```
ListView(
        padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 10),

ListView.separated(
          itemCount:  ,
          itemBuilder:  ,
          separatorBuilder:  ,
          scrollDirection: Axis.,
        )

ListView.builder(
          physics: const BouncingScrollPhysics(), // Simula en Android el efecto de IOS cuando se hace scroll hasta llegar al final de la lista, la lista 'intentará bajar más' 
          scrollDirection: Axis.horizontal,
          itemCount: 5,
          itemBuilder: (context, index) => ...,
```
# Images
```
// network
Image( image: NetworkImage('https://'))

// local
FadeInImage(
    placeholder: const AssetImage('assets/images/jar-loading.gif'),
    image: NetworkImage('https://picsum.photos/id/${imagesIds[index]}/500/300'),
    width: double.infinity,
    height: 230,
    fit: BoxFit.cover,
    fadeInDuration: const Duration(milliseconds: 300),

// rounded border
ClipRRect(
    borderRadius: BorderRadius.circular(20),
    child: FadeInImage...
```
## Avatares
```
Icon(Icons.arrow_forward_outlined, color: Colors.indigo)

CircleAvatar(
              child: const Text('SL'),
              backgroundColor: Colors.indigo[900]

CircleAvatar(
    maxRadius: 100,
    backgroundImage: NetworkImage(
        'https://images.cdn2.buscalibre.com/fit-in/360x360/0a/9a/0a9aaadabc53da4aca0ea95734bd5554.jpg'),
```