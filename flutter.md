

# Extensions:
- [Dart](https://marketplace.visualstudio.com/items?itemName=Dart-Code.dart-code)
- [Flutter](https://marketplace.visualstudio.com/items?itemName=Dart-Code.flutter)
- [Snippets](https://marketplace.visualstudio.com/items?itemName=Nash.awesome-flutter-snippets)

## Helpers
- [Material Icon Theme](https://marketplace.visualstudio.com/items?itemName=PKief.material-icon-theme)
- [Bracket Pair Toggler](https://marketplace.visualstudio.com/items?itemName=dzhavat.bracket-pair-toggler)
- [Better Coments](https://marketplace.visualstudio.com/items?itemName=aaron-bond.better-comments)

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

// boton que ocupa todo el ancho de la pantalla => width: double.infinity
ElevatedButton(
        onPressed: () {},
        child: SizedBox(
        width: double.infinity,
        child: Text('Guardar'),
        ),
        )

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