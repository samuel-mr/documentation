# Entity Framework

## Package 
```
dotnet add package Microsoft.EntityFrameworkCore.Sqlite
dotnet add package Microsoft.EntityFrameworkCore.Design
```
## Base
```
  public class MyContext : DbContext
  {
    public DbSet<Car> Cars { get; set; }
  }
```

simple console test 
```
_db.Database.EnsureCreated();

var hyundai = new Brand() { Name = "Hyundai" };
_db.Brands.Add(hyundai);
_db.SaveChanges();

Cant = _db.Brands.Count();
```
## LocalDb
example
```
    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
    {
      optionsBuilder.UseSqlServer("Data Source=(localdb)\\MSSQLLocalDB; Initial Catalog=CarflowData");
    }
```
db path stored
```
%userprofile%
```

## SQLite
package
```
dotnet add package Microsoft.EntityFrameworkCore.Sqlite
```
example
```
public class BloggingContext : DbContext
{
    ...
    public string DbPath { get; }

    public BloggingContext()
    {
        var folder = Environment.SpecialFolder.LocalApplicationData;
        var path = Environment.GetFolderPath(folder);
        DbPath = System.IO.Path.Join(path, "blogging.db");
    }

    protected override void OnConfiguring(DbContextOptionsBuilder options)
        => options.UseSqlite($"Data Source={DbPath}");
}
```

# Migrations
- [documentation](https://docs.microsoft.com/en-us/ef/core/get-started/overview/first-app?tabs=netcore-cli#create-the-database)
- install (CLI): exec in data project
    ```
    # install Migration commands
    dotnet tool install --global dotnet-ef
    # install Migration Apis
    dotnet add package Microsoft.EntityFrameworkCore.Design

    dotnet ef migrations add InitialCreate
    dotnet ef database update
    ```
- info & util
    ```
    dotnet ef                       # show description of common commands
    dotnet ef migrations script     # show sql code to execute
    dotnet ef migrations list       # show all migrations available

    ```
- flow
    ```
    dotnet ef migrations add [name] # add new migration  
    dotnet ef database update       # exec last migration
        -v                          # verbose : show details
    ```
- tooling
    ```
    dotnet ef dbcontext scaffold        # reverse engineer
    ```
## Visual Studio
- Sql Server Object Explorer -> to visualize like Sql Managment Studio


# EF Tools
- Install in VS
```
DGML editor
```