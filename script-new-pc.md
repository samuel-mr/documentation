> PowerShell AS admin
- If for some reason some installation is stopped, close powershell, and start again (For the last truncated installation probably you need to use --force)
- Some package needs to be called by cmd, to avoid create a new environment variable, choco do this:
  - It created a special folder aim to store all the 'direct access' in C:\ProgramData\chocolatey\bin (with an Env variable pointing to this one). 
    - e.g.: for Sysinternals, but not for Davinci Resolve or another big software
 
# Actualizar Chocolatey
choco upgrade chocolatey

# Desarrollo - Herramientas de Desarrollo
```ps1
choco install -y googlechrome
choco install -y visualstudio2022community --package-parameters "--add Microsoft.VisualStudio.Workload.CoreEditor --add Microsoft.VisualStudio.Workload.NetWeb --includeRecommended"
choco install -y vscode
choco install -y git.install
#choco install -y poshgit
choco install -y github-desktop
choco install -y nvm
### choco install -y nodejs ! no more individual installation!
choco install pyenv-win # python version manager for windows
###
# pyenv.bat install 3.5.2
# pyenv.bat global 3.5.2
# Win10 or 11 : Disable in : "app execution aliases" / "Pyhon" alias
# To test if it's working: python -c "import sys; print(sys.executable)"
#
# pyenv.bat versions
###
# choco install -y python
  # choco install -y awscli
  # choco install -y terraform
choco install -y postman
```
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