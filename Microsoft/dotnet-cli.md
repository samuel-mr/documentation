# Init

```
Create in folder: 'Common'
 dotnet new classlib --output "Common/ProjectName"

 
```
Referenciar
```

// implícito
dotnet add .\Resocentro.NewMedicalApp\ reference .\Resocentro.UI.Common\

// explícito
dotnet add .\Resocentro.NewMedical.Infraestructure\Resocentro.NewMedical.Infraestructure.csproj reference .\Resocentro.NewMedical.Application\Resocentro.NewMedical.Application.csproj

```