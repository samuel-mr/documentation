## Inicio
```
git init 
```
## Variables globales
```
git config --global user.name <nombre>
git config --global user.email <mail>
```

## Creas en la web un repositorio
```
git clone <url>
```

##  Add
```
git add -A      stages all changes (igual a ejecutar los 2 siguientes juntos)

git add .       stages new files and modifications, without deletions (on the current directory and its subdirectories).

git add -u      stages modifications and deletions, without new files
```

## Envio y obtención de datos
```
git commit -m "xxxxxxxx"
git push
git pull origin [branch]                obtiene rama específica

// SI revierto algunos commit localmente, debo hacer FORCE Push, el signo + significa force
git push origin +[branch]

git branch		                muestra rama actual
```

## Ramas
```
git checkout [commitId]         cambia los archivos a ese commit
git checkout main               cambia los archivos a main (a la posicion actual)
git checkout -b [newBranchName] crea un nuevo branch y se mueve ahi


git branch -d [branchName]              elimina branch local
git push origin --delete [branchName]   elimina branch remoto

```

## Fusiones
```
git merge [branch]              une una rama con la actual (es importante estar en MAIN antes de hacerlo)
```

## Micelaneos
```
git log --oneline //muestra donde está el HEAD
```



// Renombrar: source: https://linuxize.com/post/how-to-rename-local-and-remote-git-branch/
git checkout <old_name>
git branch -m <new_name>

# Visual Studio

```
dotnet new gitignore
```