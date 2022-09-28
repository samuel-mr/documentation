https://github.com/markodenic/web-development-resources

# fastest fingers
- https://www.keybr.com // usarlo como entrenamiento general inicial
- https://monkeytype.com //usarlo para palabras lentas o erradas y finalmente palabras reales
- https://ohkeycaps.com/ // comprar!!!

## win10
- Mapas mentales (nodos) https://gitmind.com/
- Text to Speech (varias voces) https://www.naturalreaders.com/online/
- Muestra shorcut presionado (para Presentaciones) http://carnackeys.com/
- (Windows) Busca archivos con expresiones regulares  https://github.com/stefankueng/grepWin

# Developer
- [Source Tree](https://www.sourcetreeapp.com/) Git UI free (win + mac)

## VS
- https://marketplace.visualstudio.com/items?itemName=martinw.SmartPaster2013
- https://marketplace.visualstudio.com/items?itemName=ErikEJ.EFCorePowerTools
- https://marketplace.visualstudio.com/items?itemName=TeamXavalon.XAMLStyler
- https://marketplace.visualstudio.com/items?itemName=MadsKristensen.LearntheShortcut
- https://marketplace.visualstudio.com/items?itemName=MadsKristensen.AddNewFile
- https://marketplace.visualstudio.com/items?itemName=MadsKristensen.MarkdownEditor
- http://erikej.github.io/SqlCeToolbox/
- https://marketplace.visualstudio.com/items?itemName=MadsKristensen.ExtensibilityEssentials2019
- https://www.vsixgallery.com/extension/0c8bd9fa-77d5-4563-ab57-9e01608c3d04#collapseattributes (varios)
- https://marketplace.visualstudio.com/items?itemName=Diagnostics.DiagnosticsConcurrencyVisualizer2019
- https://marketplace.visualstudio.com/items?itemName=stefankert.BuildVision
- https://marketplace.visualstudio.com/items?itemName=TomEnglert.ResXManager
- https://marketplace.visualstudio.com/items?itemName=ConnieYau.NETPortabilityAnalyzer
- https://marketplace.visualstudio.com/items?itemName=sergeb.GhostDoc
- [Inline Color XAML](https://marketplace.visualstudio.com/items?itemName=NikolaMSFT.InlineColorPicker)
### Testing
- [OZ-Code](https://marketplace.visualstudio.com/items?itemName=CodeValueLtd.OzCode)

***
## VS Code
- (Collection) [Awesome plugins](https://viatsko.github.io/awesome-vscode/#better-align)
- [Beetter Align](https://marketplace.visualstudio.com/items?itemName=wwm.better-align) Code and comments alignment
- [Bookmarks](https://marketplace.visualstudio.com/items?itemName=alefragnani.Bookmarks) Add Bookmarks
- [Auto Rename Tag](https://marketplace.visualstudio.com/items?itemName=formulahendry.auto-rename-tag)  Rename automatically editing tags
- [TabNine](https://marketplace.visualstudio.com/items?itemName=TabNine.tabnine-vscode) AI code completion assistent
- [FireCode](https://github.com/tonsky/FiraCode) Fuente con ligaduras
- [Auto Close Tag](https://marketplace.visualstudio.com/items?itemName=formulahendry.auto-close-tag)
- [VS Keymap](https://marketplace.visualstudio.com/items?itemName=ms-vscode.vs-keybindings) Shortcuts like Visual Studio
- [Icons](https://marketplace.visualstudio.com/items?itemName=vscode-icons-team.vscode-icons) VS Code Icons
- [Color Picker](https://marketplace.visualstudio.com/items?itemName=anseki.vscode-color) Para Css
- [Git History](https://marketplace.visualstudio.com/items?itemName=donjayamanne.githistory)

- [PowerShell](https://marketplace.visualstudio.com/items?itemName=ms-vscode.PowerShell)

- [Sass](https://marketplace.visualstudio.com/items?itemName=Syler.sass-indented) Resaltador de sintaxis, autocompletado, formateador
- [ESLint](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint) Resalta en el Workspace los errores y advertencias al ejecutar el linter ESLint
- [Huachao Mao](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) Rest client as code
- [Excel Viewer](https://marketplace.visualstudio.com/items?itemName=GrapeCity.gc-excelviewer)
### configuraciones:
1. Configuración de bindings sobre-escritos
```
// Place your key bindings in this file to override the defaults
[
    {
        "key": "ctrl+alt+a",
        "command": "wwm.aligncode",
        "when": "editorTextFocus && !editorReadonly"
    },
    {
        "key": "ctrl+r ctrl+t",
        "command": "editor.action.changeAll",
        "when": "editorTextFocus && !editorReadonly"
    },
    {
        "key": "ctrl+f2",
        "command": "explorer.newFile"
    }
]
```
2. configuración para colorizar los brackets:
```
 "editor.bracketPairColorization.enabled": true,
    "workbench.colorCustomizations": {
        "editorBracketHighlight.foreground1": "#5caeef",
        "editorBracketHighlight.foreground2": "#dfb976",
        "editorBracketHighlight.foreground3": "#c172d9",
        "editorBracketHighlight.foreground4": "#4fb1bc",
        "editorBracketHighlight.foreground5": "#97c26c",
        "editorBracketHighlight.foreground6": "#abb2c0",
        "editorBracketHighlight.unexpectedBracket.foreground": "#db6165"
      }
```
3. configuración para github copilot
```
    {
        "key": "alt+k alt+j",
        "command": "editor.action.inlineSuggest.trigger",
        "when": "config.github.copilot.inlineSuggest.enable && editorTextFocus && !editorHasSelection && !inlineSuggestionsVisible"
    },
    {
        "key": "alt+oem_5",
        "command": "-editor.action.inlineSuggest.trigger",
        "when": "config.github.copilot.inlineSuggest.enable && editorTextFocus && !editorHasSelection && !inlineSuggestionsVisible"
    },
    {
        "key": "alt+k alt+i",
        "command": "github.copilot.generate",
        "when": "editorTextFocus && github.copilot.activated"
    },
    {
        "key": "ctrl+enter",
        "command": "-github.copilot.generate",
        "when": "editorTextFocus && github.copilot.activated"
    }

```
***
# IntelliJ
- [Fill Class](https://plugins.jetbrains.com/plugin/10942-kotlin-fill-class) completa constructor o función con valores por defecto
- [Key Promoter](https://plugins.jetbrains.com/plugin/9792-key-promoter-x) para aprender los shortcuts

*** 
# Google Chrome Extensions:
### Software & Design
- VisBug (herramientas para ver y modificar html en vivo)
- JSON Viewer Pro
- Wappalyzer
- Markdown Here (agrega un menu contextual para convertir el MD en Html) util para correos rápidos
- Image Downloader (Pact Interactive) 
- Octotree (show files as explorer)
- WhatFont
- ColorZilla
### Productivity
- feedly (fav admin)
### Learn Language
- Mate Translate (traduce al click)
- Toucan

***

## Wpf
- Tab que mantiene el VisualTree https://github.com/ControlzEx/ControlzEx#tabcontrolex
- 
## VS extensibility Samples
- https://github.com/microsoft/VSSDK-Extensibility-Samples

## MSIX
https://github.com/microsoft/MSIX-Labs/releases/tag/1.1

# Learn
#### Software
- [Free Code Camp](https://github.com/florinpop17/app-ideas)
- [Desde 0](https://github.com/danistefanovic/build-your-own-x)
- [Javascript 30](https://javascript30.com/)
- https://github.com/TheAlgorithms
- https://pragprog.com/titles/jwdsal2/a-common-sense-guide-to-data-structures-and-algorithms-second-edition/

## Algoritmos
- https://www.academia.edu/32790423/An%C3%A1lisis_de_Algoritmos_para_Ingenier%C3%ADa_de_Sistemas_y_Computaci%C3%B3n
- https://jvmr2012.files.wordpress.com/2014/12/manual-analisis-de-algoritmos.pdf
- http://www.cua.uam.mx/pdfs/conoce/libroselec/Notas_Analisis_AlgoritmosVF.pdf
- https://github.com/raywenderlich/swift-algorithm-club
- https://www.campusmvp.es/recursos/post/Ser-mejor-programador-aprendiendo-algoritmos-(y-de-paso-Swift).aspx

#### DataScience
- https://www.datasource.ai/es/data-science-articles/21-recursos-para-aprender-matematicas-para-data-science
- [16 Data Science Projects with Source Code to Strengthen your Resume](https://data-flair.training/blogs/data-science-project-ideas/)
- [Top 9 Data Science Projects for a Beginner in 2020](https://towardsdatascience.com/top-9-data-science-projects-for-a-beginner-in-2020-26eb7d42b116?gi=159b758c8f95)
- [6 Interesting Data Science Project Ideas & Examples](https://www.springboard.com/blog/data-science-projects/)

## Architecture
- https://github.com/ddd-crew/ddd-starter-modelling-process
- https://github.com/joelparkerhenderson/architecture_decision_record
- https://martinfowler.com/articles/branching-patterns.html
- https://pbs.twimg.com/media/Eoyc4ZvXMAAI7dy?format=jpg&name=4096x4096 (Mapa a un producto evolucionable)

### Music
- https://orquestacionmidi.com

### Books
- https://www.kevinrooke.com/book-genres/psychology
- https://soenglish.me/books-to-learn-english/#ielts

### Roadmaps
- https://github.com/MoienTajik/AspNetCore-Developer-Roadmap
- https://roadmap.sh

### Platform Online
- https://www.microsoft.com/es-xl/latampartneruniversity

- 
## Blazor
- Plugins Dinámicos https://stevenknox.net/exploring-assembly-unloading-in-net-core-3-0-by-building-a-simple-plugin-architecture/
## BlazorWeb
- https://blazorhelpwebsite.com

## Proyects
- https://www.juanbarrios.com/algoritmo-para-la-deteccion-del-uso-de-la-mascarilla/
- https://ekobit.com/blog/designing-an-entity-attribute-value-model-with-ef-core/ 

## C# Post
- https://johnthiriet.com/configure-await

## Q
- https://www.pramp.com/#/
- https://leetcode.com/

See Syntax Tree
- https://sharplab.io/
Source Generator
- https://sourcegen.dev/

## Shortcut
//select up scope
Shift + Alt + +
//select down scope
Shift + Alt + -
//select all similar names
Shift + Alt + ´

## Clouds
- https://www.ovhcloud.com/en/vps/ (cheap)

## Xamarin forms samples
- https://github.com/syncfusion/xamarin-shopping-cart
- https://github.com/syncfusion/Xamarin-ExpenseAnalysis
- https://github.com/roubachof
- https://github.com/davidortinau/Xappy

## Teacher
- https://coggle.it/diagram/X-ZbLirTPnv7D4zP/t/220-recursos-tic-docentes-a5ba75c14f530e3db2bd6ffdfa76afde

## Apps interesantes
- https://www.drools.org/

## IA
- https://towardsdatascience.com/whats-new-in-yolov4-323364bb3ad3
- https://jonathan-hui.medium.com/ssd-object-detection-single-shot-multibox-detector-for-real-time-processing-9bd8deac0e06#:~:text=SSD%20is%20a%20single%2Dshot,offsets%20to%20default%20boundary%20boxes.

## IA - Apps
- Ejemplos configurable de algoritmos genéticos https://www.blazor.ai/
- Crear varias cosas con IA https://metacreation.net/
## Android
- Analisis de seguridad https://github.com/MobSF/Mobile-Security-Framework-MobSF
- Lo mismo pero web: https://www.immuniweb.com/mobile/
- Otro analizador https://www.hybrid-analysis.com/
- GenyMotion in Cloud (probar muchos ndroids) https://azuremarketplace.microsoft.com/en-us/marketplace/apps/genymobile.genymotion-cloud?tab=overview

## Seguridad y Firewalls
- Gente Microsoft https://seguridadit.blogspot.com/

## Email
- Plantilla email Outlook https://adaptivecards.io/

# Benchmark
- [Cinebench 15](https://www.guru3d.com/files-get/cinebench-15-download,1.html) benchmark tool usado por Nate - for windows
