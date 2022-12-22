# Init

```
 
dotnet new webapi 
    --output "Products/App"     # custom path
    --name "AQ.Products.App"    # custom project name

# Add reference
dotnet sln [.\Solution.sln] add [.\ProjectoA.csproj]

 
```
example: create a 3 projects related
```
dotnet new webapi --output "Products/App" --name "AQ.Products.App"
dotnet new classlib --output "Products/Domain" --name "AQ.Products.Domain"
dotnet new classlib --output "Products/Test" --name "AQ.Products.Test"

dotnet sln AQ.Middleware.sln add ".\Products\App\AQ.Products.App.csproj"
dotnet sln AQ.Middleware.sln add ".\Products\Domain\AQ.Products.Domain.csproj"
dotnet sln AQ.Middleware.sln add ".\Products\Test\AQ.Products.Test.csproj"

```
Referenciar
```

// implícito
dotnet add .\Resocentro.NewMedicalApp\ reference .\Resocentro.UI.Common\

// explícito
dotnet add .\Resocentro.NewMedical.Infraestructure\Resocentro.NewMedical.Infraestructure.csproj reference .\Resocentro.NewMedical.Application\Resocentro.NewMedical.Application.csproj

```
```
# verbosity
dotnet build -v 
    q[uiet], m[inimal], n[ormal], d[etailed], and diag[nostic]
```

## PowerShell
- Para ejecutar en ruta relativa
    ```powershell
    # - go final path
    # - edit .ps1 file
    # - run
    Invoke-Expression -Command "& 'W:\Learn\PowerShell\dotnet-create-projects.ps1'"
    ```