#!/bin/bash

echo "🔧 Setting up your Mac for development..."

# 1. Ensure Homebrew is installed
if ! command -v brew &>/dev/null; then
  echo "🚀 Installing Homebrew..."
  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi

## FORMULAE _____________________
brew install git
brew install wget
brew install curl
brew install tree
brew install go
brew install nvm # it need to run additional comands shown onto console
brew install pnpm

brew install jq    # https://jqlang.org/tutorial/
brew install fzf   # It's an interactive filter program
brew install ripgrep    # searches directories for a regex pattern while respecting your gitignore

# Pyton tools
brew install pipx # like nvn for python but only for CLI not libs
pipx ensurepath

## pipx commands:
<< 'MULTILINE-COMMENT'
pipx list  # list installed apps with pipx
pipx run <name> # run without installing
pipx upgrade <name>
pipx uninstall <name>
pipx info <name>
MULTILINE-COMMENT

# Crate for python
pipx install uv  # use pipx to be isolated

# App Store
brew install mas # Mac App Store cli
mas install 497799835 # Xcode

 
## CASKS_____________________-
# IDE
brew install --cask visual-studio-code
brew install --cask jetbrains-toolbox

# dev
brew install --cask postman
brew install --cask insomnia
brew install --cask github 
brew install --cask ngrok 

# design
brew install --cask figma 

# navigators
brew install --cask google-chrome
brew install --cask firefox
brew install --cask arc
brew install --cask duckduckgo

# cli
brew install --cask warp 

# productivity
brew install --cask rectangle   # like powertoys's screen tools
brew install --cask alt-tab #like in windows
brew install --cask stats   # performance metrics
brew install --cask monitorcontrol  # britnes and volume control
brew install --cask notion
brew install --cask spotify
brew install --cask vlc
brew install --cask altserver  # load apps into iPhone without App Store
brew install --cask breaktimer # pomodoro

# miselaneous
brew install --cask spotify
brew install --cask whatsapp
brew install --cask freetube
# on run freetube: run it at first time (or after update)
## xattr -d com.apple.quarantine /Applications/FreeTube.app

# containers - virtualization
brew install colima docker docker-compose
## colima start --memory 4

# fonts
brew install --cask font-fira-code
brew install --cask font-jetbrains-mono

# sec
brew install --casks protonvpn
brew install --cask qbittorrent

# chat
brew install --cask telegram

# games - win emulation
## porting kit

# Others
## BalenaEtcher : to create usb bootable
