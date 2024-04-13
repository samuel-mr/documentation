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
    > git add --all   like git add . (I don't know the difference)
git add -u      stages modifications and deletions, without new files

```

## Envio y obtención de datos
```
git commit -m "xxxxxxxx"


git push
     -u origin [branch-name]    // (si es 1ra vez) crea la rama en remote
     origin +[branch]           // SI revierto algunos commit localmente, debo hacer FORCE Push, el signo + significa force
     origin HEAD             // HEAD = current branch
git pull origin [branch]                obtiene rama específica

git push 

git branch		                muestra rama actual
```

## Ramas
```
git checkout [commitId]         cambia los archivos a ese commit
git checkout main               cambia los archivos a main (a la posicion actual)
git checkout -b [newBranchName] crea un nuevo branch y se mueve ahi
git branch branch_name <commit-hash>  crea nuevo branch en base a otro commit
git checkout -b myFeature anotherBranch  crea nuevo branch en base a otro branch
git branch -d [branchName]              elimina branch local
git push origin --delete [branchName]   elimina branch remoto

```
## Commit 
```
git checkout [commitId]          detach HEAD para hacer cambios
git checkout [branch]            para retornar del estado 'detached'

```

## Fusiones
```
git merge [branch]              une una rama con la actual (es importante estar en MAIN antes de hacerlo)
```

## Micelaneos
```
git config --get remote.origin.url  // Get Url
```


## Info
```
git log --oneline //muestra donde está el HEAD
```

## Miscelaneos
```
git diff   # util cuando hago ejemplos
```
## Undo changes
```
git restore .               // deshace todos los cambios (q no están en el stage)
git checkout -- <file>     // si está modificandose (y no esté en el 'stagging area')
git reset HEAD <file>      // si está en 'stage' (y aun no fue commiteado)

// si está commiteado (no pusheado) -----------------
git reset [previuosCommitID]          // elimina los cambios del stage, pero los cambios en si están en modo local (??? aun pendiente probar)
git reset --hard [previuosCommitID]   // elimina los cambios del stage + archivos modificados = elimina todo el commit completo
git reset --soft [previousCommitID]   // elimina el commit superior, pero deja los cambios en el stage

```

## Save Temporaly
```
git stash save "message"  // save all changes
git stash list
git stash pop             // restore the first and remote it
```

## Move a commit to another branch
```
// 1. destination
git checkout[destination branch]
// 2. get
git cherry-pick [commit-id]
```
## Solving errors
```
// Obtener los commits de otras ramas a mi rama de trabajo actual
// 1) estando en la otra final a corregir ejem: [branch-001]
git pull origin [origin-source-branch] 
```
Clean (pending to ingestigate more)
```
git clean -fdx

# Clean algun archivo 'pegado' que supeustameinte tiene cambios pero no tiene naa'
git checkout pathtofile
# All files
git checkout .
```

// Renombrar: source: https://linuxize.com/post/how-to-rename-local-and-remote-git-branch/
git checkout <old_name>
git branch -m <new_name>

## Change repo
```
# in the current repo, only change the remote, then, push
git remote set-url origin <url.git>
```

# Visual Studio

```
dotnet new gitignore
```
Cambiar a otro commit antiguo en VS
```
# en Historial Commits, sobre el commit antiguo: click derecho / checkout (--detach)
Checkout (--detach)
# Para volver al ultimo commit real
```
git checkout [branchName]
```