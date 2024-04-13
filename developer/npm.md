https://docs.npmjs.com/cli/v9/configuring-npm/npmrc?v=true
- Its a file to store npm  configurations
- Steps:
    - Create a config file in %userprofile%\.npmrc
        ```bash
        # automatically create the file
        npm config edit
        ```
Content: (Example)

```
trict-ssl=false
//npm.pkg.github.com/:_authToken=${READ_PKG_TOKEN}
@imX1:registry=https://npm.pkg.github.com
@imX2:registry=https://npm.pkg.github.com
```

## using shell 
- If you need multiples files, it helps if you need some automatation, sino create it manually.
```powershell 
# Using the command line
# internally it will create a new folter to store all profiles (.npmrc files)
# with a direct access file

npm install -g npmrc

# create a new profile
npmrc -c <profile>
# switch
npmrc <profile
```

## Tips
```powershell
npm ls [name] # For looking by package name into our package folder. Returns a tree
```

## Version manager
```
npm install -g npm-check-updates
ncu
ncu --upgrade  #to run the upgrade
ncu --interactive
ncu --interactive --format group

```