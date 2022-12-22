

# Extensions:
- [Dart](https://marketplace.visualstudio.com/items?itemName=Dart-Code.dart-code)
- [Flutter](https://marketplace.visualstudio.com/items?itemName=Dart-Code.flutter)
- [Snippets](https://marketplace.visualstudio.com/items?itemName=Nash.awesome-flutter-snippets)

## Helpers
- [Material Icon Theme](https://marketplace.visualstudio.com/items?itemName=PKief.material-icon-theme)
- [Bracket Pair Toggler](https://marketplace.visualstudio.com/items?itemName=dzhavat.bracket-pair-toggler)
- [Better Coments](https://marketplace.visualstudio.com/items?itemName=aaron-bond.better-comments)

## Top Controls
- MediaQuery: Obtiene alto, ancho de la pantalla
# Commands
```
flutter --version
flutter doctor
```

# Common methods

```
Columns / Rows
mainAxisAlignment: MainAxisAlignment.center,
```
# Containers
```
Stack                   // permite  uno sobre otro
```
## Definición de tamaños
| Control          | 'Height' definido por |
|------------------|-----------------------|
| Column           | hijos                 |
| ListView.builder | padre                 |

- Al menos el padre o el hijo debe saber 
- Escenario: ⚠️Padre e hijo con 'height indeterminado  <span style="color:red">no compilará</span>

         
        // El problema es que tanto el padre como el hijo esperan del uno al otro para definir su 'height'
        Column(
                children: [
                ListView.builder...
         

# Controls
```
SizedBox (like Box)

ListView.separated(
          itemCount:  ,
          itemBuilder:  ,
          separatorBuilder:  ,
        )

ListView(
        padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 10),

// Simula en Android el efecto de IOS cuando se hace scroll hasta llegar al final de la lista, la lista 'intentará bajar más' 
ListView.builder(
          physics: const BouncingScrollPhysics(),
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
## Controls - scroll when overflow
Genera un scroll si el hijo supera el tamaño de la pantalla
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
### Lists

| Shortcut              | When use it?                 |
|-----------------------|------------------------------|
| Column                | Hijos con el mismo 'height'  |
| SingleChildScrollView | Hijos con diferente 'height' |

### ListTile vs NO ListTile
- Checkbox vs CheckboxListTile : el segundo tiene todo el row como área clickeable para ser seleccionado
- Switch   vs SwitchListTile   : idem que el anterior

## Controls - adaptative
Algunos controles tienen la propiedad `adaptive` para que se muestre la versión automáticametne su versión IOS / Android
```
Slider.adaptive
```

# Controls - images
```
// network
Image( image: NetworkImage('https://'))

// local
FadeInImage(
          placeholder: const AssetImage('assets/images/jar-loading.gif'),
          image: NetworkImage(imageUrl),
          width: double.infinity,
          height: 230,
          fit: BoxFit.cover,
          fadeInDuration: const Duration(milliseconds: 300),
        )
```

# Controls - IOS
Remove notch space: `MediaQuery.removePadding`
```
  Widget build(BuildContext context) {
    return Scaffold(
      body: MediaQuery.removePadding(
        context: context,
        removeTop: true,
        removeBottom: true,
```
# Code snippets
```
mateapp  -> create starting code
```

# Platform specific
```
import 'dart:io';

if (Platform.isAndroid)
if (Platform.isIOS)
```

Quitar el foco del input actual:
```
FocusManager.instance.primaryFocus?.unfocus();

// (old) segunda forma: aunq esta me da problrmeas con el dropdown
FocusScope.of(context).requestFocus();
```
Obtener Ancho y Alto de la pantalla
```
    final size = MediaQuery.of(context).size;
```

# Theme
1. Modificacion en `build`
```
Widget build(BuildContext context) {
    return MaterialApp(
        theme: ThemeData.light()
            .copyWith(appBarTheme: AppBarTheme(color: Colors.red)));
```