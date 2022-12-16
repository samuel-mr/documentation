## Package

### DotNet-CLI
- dotnet add package [NAME]

### NugetPackage
```
FontAwesome5
Microsoft.Extensions.DependencyInjection
Prism.Core
```
```
Microsoft.EntityFrameworkCore
Microsoft.EntityFrameworkCore.SqlServer
```
Migrations:
```
dotnet tool install --global dotnet-ef
dotnet add package Microsoft.EntityFrameworkCore.Design
dotnet ef migrations add InitialCreate
dotnet ef database update
```