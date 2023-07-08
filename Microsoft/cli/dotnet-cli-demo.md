# Como usarlo?
# 0. Guardar el contenido como .ps1
# 1. ubicate en la carpeta donde crearás los proyectos
# 2. Invoke-Expression -Command "& 'C:\path\namefile.ps1'"

# Define la variable de texto
$empresa = "MyCompany"
$solution = "KataSolution" 
$solution_name = $solution + ".sln"
$framework = "net6.0"
 
function CreateSimpleSolution([TestingFrameworkEnum]$testingFramework) {
 

  $api_path = [Simplepath]::new($empresa, "Presentation", "Api")
  $domain_path = [Simplepath]::new($empresa, "Core", "Domain")
  $application_path = [Simplepath]::new($empresa, "Core", "Application")
  $infra_path = [Simplepath]::new($empresa, "Infrastructure", "Services")
  $test_domain_path = [Simplepath]::new($empresa, "Test", "DomainTest")
  $test_application_path = [Simplepath]::new($empresa, "Test", "ApplicationTest")

  Write-Output "Creando Api ---------------------"  
  dotnet new webapi --output $api_path.NameAsFileSystemFormat() --name $api_path.Namespaces() --framework $framework 
  Write-Output "Creando Domain ---------------------"  
  dotnet new classlib --output $domain_path.NameAsFileSystemFormat() --name $domain_path.Namespaces() --framework $framework 
  Write-Output "Creando Application ---------------------"  
  dotnet new classlib --output $application_path.NameAsFileSystemFormat() --name $application_path.Namespaces() --framework $framework 
  Write-Output "Creando Infra ---------------------"  
  dotnet new classlib --output $infra_path.NameAsFileSystemFormat() --name $infra_path.Namespaces() --framework $framework 
  
  if ($testingFramework -eq [TestingFrameworkEnum]::xunit) {
    Write-Output "Creando Test XUnit---------------------"
    dotnet new xunit --output $test_domain_path.NameAsFileSystemFormat() --name $test_domain_path.Namespaces() --framework $framework 
    dotnet new xunit --output $test_application_path.NameAsFileSystemFormat() --name $test_application_path.Namespaces() --framework $framework 
  }
  elseif ($testingFramework -eq [TestingFrameworkEnum]::nunit) {
    Write-Output "Creando Test NUnit---------------------"
    dotnet new nunit --output $test_domain_path.NameAsFileSystemFormat() --name $test_domain_path.Namespaces() --framework $framework 
    dotnet new nunit --output $test_application_path.NameAsFileSystemFormat() --name $test_application_path.Namespaces() --framework $framework 
  }
  else {
    Write-Output "Este tipo de test no está soportado---------------------"
  }
 
  Write-Output "Creando .sln ---------------------"
  dotnet new sln --name $solution
  dotnet sln $solution_name add $api_path.NameAsTextFormat()
  dotnet sln $solution_name add $domain_path.NameAsTextFormat()
  dotnet sln $solution_name add $application_path.NameAsTextFormat()
  dotnet sln $solution_name add $infra_path.NameAsTextFormat()
  dotnet sln $solution_name add $test_domain_path.NameAsTextFormat()
  dotnet sln $solution_name add $test_application_path.NameAsTextFormat()

  dotnet add $test_domain_path.NameAsTextFormat() reference $domain_path.NameAsTextFormat()
  dotnet add $test_application_path.NameAsTextFormat() reference $application_path.NameAsTextFormat()

  dotnet add $application_path.NameAsTextFormat() reference $domain_path.NameAsTextFormat()
  dotnet add $application_path.NameAsTextFormat() reference $infra_path.NameAsTextFormat()
  
  dotnet add $api_path.NameAsTextFormat() reference $application_path.NameAsTextFormat()

  dotnet build
}
# Tipo de proyecto de test
enum TestingFrameworkEnum{
  xunit
  nunit
}

function  run {
  $presentation = [Simplepath]::new("fff", "hola", "ffg")
  $presentation.NameAsTextFormat()
  $presentation.NameAsFileSystemFormat()
  $presentation.Namespace()
}

class Simplepath {

  Simplepath([string] $company, [string] $parent_folder, [string] $child_folder) {
    $this.company = $company
    $this.parent_folder = $parent_folder
    $this.child_folder = $child_folder
  }

  [string] $company
  [string] $parent_folder
  [string] $child_folder

  # Format to include as argument for cli dotnet
  [string] NameAsFileSystemFormat() {
    return $this.parent_folder + "/" + $this.child_folder  
  }

  # Format to include in the .sln file
  [string] NameAsTextFormat() {
    return $this.parent_folder + "\" + $this.child_folder + "\" 
  }
  
  # Format to include as namespance in each .proj
  [string] Namespaces() {
    return $this.company + "." + $this.child_folder
  }
}


# Llama al método y le envía la variable de texto como argumento
#CreateSimpleSolution 
CreateSimpleSolution([TestingFrameworkEnum]::nunit)