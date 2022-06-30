# Testing

## Exceptions
```
//Act
var exception = await Record.ExceptionAsync(async () => await instancia.Get(default));

//Arrange
exception.Should().NotBeNull();


//Act
var act = () => instancia.Get(default);

//Assert
await Assert.ThrowsAnyAsync<Exception>(async () => await act());
```

## Many
```
https://blog.ndepend.com/guide-code-coverage-tools/
```
# NSubstitute

ReturnsForAnyArgs
```
_interface.Query<string>()
        .ReturnsForAnyArgs(
            Task.FromResult(DATA_VERSION_A),
            Task.FromResult(DATA_VERSION_B));
...
// two consecutive calls
 var resultA = await _interface.Query<string>(); // will return DATA_VERSION_A
 var resultB = await _interface.Query<string>(); // will return DATA_VERSION_B
```

# Coverlet
- [Coverlet integration with VSTest](https://raw.githubusercontent.com/tonerdo/coverlet/master/Documentation/VSTestIntegration.md)
## Basic
```
dotnet tool install -g coverlet.console
```
example: this solution has 2 projects Logic and LogicTest
```

├───Logic
│   ├───bin
│   │   └───Debug
│   │       └───net6.0
│   └───obj
│       └───Debug
│           └───net6.0
│               ├───ref
│               └───refint
├───LogicTest
│   ├───bin
│   │   └───Debug
│   │       └───net6.0


(tree command)
```
in CMD go to testproject: `cd [project path]`
```
coverlet .\bin\Debug\net6.0\LogicTest.dll --target "dotnet" --targetargs "test --no-build"
    --exclude "[*]Logic.NameSpaceToSkip"
    --exclude "[*]Logic.NameSpaceToSkip*"
```
## 2) Collect

```
 cd [Test Project path]
 dotnet add package coverlet.collector
 dotnet test --collect:"XPlat code coverage"    # will create a TestResults folder
 notepad coverlet.runsettings                   # create a new empty configuration file

```

## 3) Generate Reporte
- install [ReportGenerator](https://github.com/danielpalme/ReportGenerator)
```
dotnet tool install -g dotnet-reportgenerator-globaltool
```

```
cd [Test Project path]
reportgenerator 
    -reports:"[path]\coverage.cobertura.xml" 
    -targetdir:"coverateresults" 
    -reporttypes:Html
```
example
```
reportgenerator -reports:".\TestResults\1618188d-df8e-4d13-bd3f-331e3551bace\coverage.cobertura.xml" -targetdir:"coverateresults" -reporttypes:Html
```