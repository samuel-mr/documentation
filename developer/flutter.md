

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
# Important
flutter clean
flutter pub get
flutter build apk

# Run
flutter run -d chrome

#General
flutter --version
flutter doctor
flutter doctor --android-licenses  # just to confirm some aditional licences
flutter devices  # you must see your phone (usb)
flutter upgrade

# pubspec
flutter pub add ...
flutter pub get       # install current v
flutter pub outdated  # display outdated libs
flutter pub upgrade   # To update to the latest compatible versions of all the dependencies listed in the pubspec.yaml
flutter pub upgrade --major-versions # the same + upgrade the constraints in the pubspec.yaml

## how to update
flutter pub outdated
- (for mayor version, modify it manually)
flutter pub upgrade
flutter pub get
flutter pub outdated
```

# New Proj
```
In VS Code : Ctrl + Shift + P
- 
```
# Firebase configuration
```ps1
npm install -g firebase-tools
firebase login
dart pub global activate flutterfire_cli
```
# Containers
```
Stack                   // permite  uno sobre otro
```
 
# Tamaños + distancias
- MediaQuery: Obtiene alto, ancho de la pantalla
- ⚠️ Al menos el padre o el hijo debe saber su 'height'

| Control          | 'Height' definido por | Full height? |
|------------------|-----------------------|--------------|
| Column           | hijos                 |              |
| ListView.builder | padre                 |              |
| Expanded         |                       | Si           |
### Escenarios:
- Padre e hijo con 'height' indeterminado. ❗no compilará❗
  ```
  // El problema es que tanto el padre como el hijo esperan del uno al otro para definir su 'height'
  Column(
    children: [
      ListView.builder...
  ```
  - Solución: Envolver al hijo en un control que si defina el 'height' (ejem: Expanded)
    ```
    Column(
        children: [
          Expanded(
            child: ListView.builder...
    ```  
 
# Components
```
class SomeComponent ... {
  const SomeComponent({super.key, required this.title});

  final String title;
```

### Escenarios:        
- Quitar el foco del input actual:
  ```
  FocusManager.instance.primaryFocus?.unfocus();

  // (old) segunda forma: aunq esta me da problrmeas con el dropdown
  FocusScope.of(context).requestFocus();
  ```
- Obtener Ancho y Alto de la pantalla
  ```
      final size = MediaQuery.of(context).size;
  ```
## Http random image
```
https://picsum.photos/id/{id}/500/300
https://via.placeholder.com/300x400
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

## Theme
1. Modificacion en `build`
```
Widget build(BuildContext context) {
    return MaterialApp(
        theme: ThemeData.light()
            .copyWith(appBarTheme: AppBarTheme(color: Colors.red)));
```

## Gesture
```
GestureDetector(
  onTap: () {
    Navigator.pushNamed(context, 'name_of_some_route',
        arguments: 'some_id');
  },
  child:
```

## Package
```
flutter pub add http
```
list
```
http                        https://pub.dev/packages/http
provider (gestor de estado) https://pub.dev/packages/provider

```

# Apis
- [Movies](https://developers.themoviedb.org/)

# Maps
> Instructions: https://pub.dev/packages/google_maps_flutter
- build.gradle
  ```
  defaultConfig {
          minSdkVersion 21
          ...
      }
  ```
- Get Google Maps Platform Key
  - Ensure are enabled:
    - Maps JavaScript API
    - Maps SDK for Android
    - Maps SDK for iOS

# Riverpod
ref.watch(...) // as listener
ref.read(...)  // as getter (eg. when click in a button)

// use internal method
ref.read(myProvider.notifier).someMethod();