> PowerShell AS admin
- If for some reason some installation is stopped, close powershell, and start again (For the last truncated installation probably you need to use --force)
- Some package needs to be called by cmd, to avoid create a new environment variable, choco do this:
  - It created a special folder aim to store all the 'direct access' in C:\ProgramData\chocolatey\bin (with an Env variable pointing to this one). 
    - e.g.: for Sysinternals, but not for Davinci Resolve or another big software
 
# Actualizar Chocolatey
choco upgrade chocolatey

# WinGet
```ps1
winget install WhatsApp
winget install --id JetBrains.Toolbox -e
winget install --id Microsoft.VisualStudioCode -e
winget install --id Mozilla.Firefox -e
winget install --id Microsoft.Edge -e
winget install --id Google.Chrome -e

winget install --id 7zip.7zip -e
winget install --id VideoLAN.VLC -e
winget install --id Notepad++.Notepad++ -e
winget install --id ShareX.ShareX -e

winget install --id Microsoft.PowerToys -e
winget install --id CPUID.CPU-Z -e
winget install --id Piriform.CCleaner -e

winget install --id Docker.DockerDesktop -e

```
# Desarrollo - Herramientas de Desarrollo
```ps1
choco install -y googlechrome
choco install -y visualstudio2022community --package-parameters "--add Microsoft.VisualStudio.Workload.CoreEditor --add Microsoft.VisualStudio.Workload.NetWeb --includeRecommended"
choco install -y vscode
choco install -y git.install
#choco install -y poshgit
choco install -y github-desktop
#####choco install -y nvm (currently not working) instead use -> https://github.com/coreybutler/nvm-windows/releases
### choco install -y nodejs ! no more individual installation!

choco install -y postman
choco install -y offset-explorer # Kafka UI admin
choco install -y winscp #transfer files
# sdkman (see the sdkman.md file)
# jabba (install manually): https://github.com/shyiko/jabba
choco install -y k6 # load testing
choco install -y jetbrainstoolbox
```
- Configure jabba (instead of sdkman)
  ```bash
  # Path -> C:\Users\<USER_NAME>\.jabba\jdk\
  # list
  jabba ls-remote
  # list local
  jabba ls
  jabba install openjdk@1.17.0
  jabba use openjdk@1.17.0
  jabba alias default openjdk@1.17.0

  ```

### Python Launcher (py.exe) temporal replace for Pyenv
```ps1
winget install --id Python.Python.3.10 -e
winget install --id Python.Python.3.12 -e
py -3.12
py --list

pip3 install instaloader
instaloader --no-captions -- -XXXXXXX
``` 
### Pyenv (in windows 11 it has some issues in Agust '25...)
- Win10 or 11 : Disable in : "app execution aliases" / "Pyhon" alias
```ps1
choco install pyenv-win # python version manager for windows
pyenv.bat update # Good practice, first, update the local list of available versions

pyenv.bat install --list
pyenv.bat versions # display all installed versions

pyenv.bat install 3.5.2
pyenv.bat global 3.5.2

pyenv.bat install 3.12.5
pyenv.bat global 3.12.5
```

- To test if it's working
  ```ps1
  python -c "import sys; print(sys.executable)"
  ```
- Configuracion guardada en `pyvenv.cfg`


- Create `Env` : In the project's root 
  ```ps1
  # Create virtual environment
  python3 -m venv env
  # Activate : as result in the console, you should see the name of the env at the beginning
  .\env\Scripts\activate
  # Desactivate : will work if it was activated previously
  deactivate
  # List all installed package
  python -m pip list
  # Install dependencies using txt (con en env activo me funciono usando 'python', no 'python3')
  python -m pip install -r requirements.txt
  # Uninstall
  python -m pip uninstall -r requirements.txt -y
  # Upgrade (it works only if you don't specify the version of your package)
  python -m pip install --upgrade -r requirements.txt
  ```
- To update a package just change to values in the requirement.txt file and run the pip install...


### Oh-My-Phosh:
```ps1
choco install -y oh-my-posh ## I have test it with powershell 5
oh-my-posh font install FireCode
```
- Windows Terminal : Terminal / Configuracion / Valores Predeterminados
  - Fuente: 'FiraCode Nerd Font Mono'
- VS Code Terminal:
  - Go to : Settings / User / Features / Terminal
  - Add to Font Family: 'FiraCode Nerd Font Mono'

```ps1
notepad $PROFILE
# Add this line in the opened file:
    oh-my-posh init pwsh | Invoke-Expression
# Preview all themes in console if you wanna use another
Get-PoshThemes
# Export the installed default theme into %userprofile%/ and copy it to %POSH_THEMES_PATH%
oh-my-posh config export --output ~/samutheme.omp.json
# Add this line to : notepad $PROFILE
oh-my-posh init pwsh --config "$env:POSH_THEMES_PATH\samutheme.omp.json" | Invoke-Expression
```

- POSH_THEMES_PATH = Environment variable name for installed Themes
- in `samutheme.omp.json` : Change the background and foreground values like this:
```json
  "blocks": [
    {
      "segments": [
        {
          "type": "session",
          ...
          "background": "p:transparent",
          "foreground": "p:white"
        },
        {
          "type": "path",
          ...
          "background": "p:black",
          "foreground": "p:orange"
          ...
          ...
          ...
  "transient_prompt": {
    "template": "<p:black,transparent>\ue0b6</><,p:black> {{ .Folder }} </><p:yellow,transparent>\ue0b0</> ",
    "background": "transparent",
    "foreground": "p:yellow"
  },
  "secondary_prompt": {
    "template": "<p:black,transparent>\ue0b6</><,p:black> > </><p:yellow,transparent>\ue0b0</> ",
    "background": "transparent",
    "foreground": "p:yellow"
  },
```
# Desarrollo - Herramientas de Bases de Datos
```ps1
  #choco install -y mysql
  #choco install -y pgadmin4
  #choco install -y dbeaver
```
# Diseño - Herramientas de Diseño
```ps1
  #choco install -y gimp
  #choco install -y inkscape
  #choco install -y krita
choco install -y blender
choco install -y paint.net
choco install -y davinci-resolve
```
# 3d
```
# https://www.pureref.com    # to organize images
```

# Utilidades - Herramientas de Productividad y Utilidades
```ps1
choco install -y googlechrome
choco install -y firefox
choco install -y 7zip
choco install -y notepadplusplus
choco install -y slack
choco install -y telegram.install
choco install -y zoom
# choco install -y teamviewer
choco install -y spotify
choco install -y cpu-z.install
choco install -y notion
choco install -y powertoys
choco install -y qbittorrent
choco install -y treesizefree

```
# Utilidadoes (manualy)
```
Keystroke display (for tutorials)
https://github.com/bfritscher/carnac
Firewall
https://safing.io/
*/
```
# To Complete the setup!
```ps1
choco install -y sysinternals
choco install -y discord.install
# everything (has params)
choco install docker-desktop
cwrsync (win version of a linux tool to create backup and allow configure gitignore riles)
```

```ps1
# Manually
Geforce Experience
NSwagStudio
Ollama
Oracle VM Virtual Box
Popcorn-Time
```

# Limpiar archivos temporales de Chocolatey
choco clean

# Mensaje final
Write-Host "Instalación completada. Por favor, reinicia tu computadora para completar el proceso."

# Manual installation
```
# win11 customization
winaerotweaker #many options for UI and behavior
O&O ShoutUp 10 ++  # privacy configurations


```
