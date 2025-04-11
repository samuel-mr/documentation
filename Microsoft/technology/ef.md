# Entity Framework

## Package 
```
dotnet add package microsoft.entityframeworkcore       // prefiero tenerlo explicito
dotnet add package Microsoft.EntityFrameworkCore.Sqlite
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
      //// Aditionally: you are able to add this conditional
      //if (!optionsBuilder.IsConfigured)
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
example A : configure in class
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
example B : configure externally
```
public class BloggingContext : DbContext
{
    ...
    public string DbPath { get; }

    public BloggingContext(DbContextOptions<CityInfoContext> options) : base(options)
    {
    }
}
_________________________
// in Program.cs
builder.Services.AddDbContext<CityInfoContext>(
  dbContextOptions => dbContextOptions.UseSqlite("Data Source=blogging.db") //saved in root path of project
  );
```

in appsettings.json
```
 "ConnectionStrings": {
    # Sqllite
      "CityInfoDBConnectionString" :  "Data Source=CityInfo.db"
    # Postgresql
      "DefaultConnection" :  "Host=localhost;Port=5432;Database=simpleApiDb;Username=user;Password=passs;"
  } 
```
in Program.cs
```
 var connectionString = builder.Configuration["ConnectionStrings:CityInfoDBConnectionString"];
````

## Connection String
In the connection string may be needed to add:
```
// when certification is not trusted
TrustServerCertificate=true
```

# Migrations
- [documentation](https://docs.microsoft.com/en-us/ef/core/get-started/overview/first-app?tabs=netcore-cli#create-the-database)
- install (CLI): exec in data project
    ```
    # install Migration commands
    dotnet tool install --global dotnet-ef
    # install Migration Apis
    dotnet add package Microsoft.EntityFrameworkCore.Design

    # (only if you'll used Nuget Package Manager Console in VS)
    dotnet add package Microsoft.EntityFrameworkCore.Tools

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
    dotnet ef migrations remove     # remove the last one
    dotnet ef database update       # exec last migration
        -v                          # verbose : show details
    ```
    - Some required basic structure:
      ```
      var configuration = builder.Configuration;
      builder.Services.AddDbContext<HangContext>(options =>
      {
          options.UseSqlServer(configuration.GetConnectionString("DefaultConnection"));
          // for postgresql
          options.UseNpgsql(configuration.GetConnectionString("DefaultConnection"));
      });
      ...
      public class Persona
      {
          public int Id { get; set; }
          public string Name { get; set; }
      }
      public class HangContext : DbContext
      {
        // If you need to run Migrations in a Console APP add this empty constructor:
          public HangContext(){}
          public HangContext(DbContextOptions<HangContext> options) : base(options)
          {}
          public DbSet<Persona> Personas { get; set; }
      }
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

# Conventions
PK
```
[Key]
[DatabaseGenerated(DatabaseGeneratedOption.Identity)]
public int Id { get; set; }
```
FK
```
[ForeignKey(nameof(CityId))]
public City City { get; set; }
public int CityId { get; set; }
```


# Extensions
in Visual Studio
```
- SQLite and SQL Server Compact Toolbox
    (in VS) Menu: Tools / SQLite and SQL ...
```

# Querys
having `City` as `Entity`
```
var collection = _context.Cities as IQueryable<City>;

if(name != null)
{
  collection = collection.Where(c => c.Name == name);
}

return await collection
    .ToListAsync();
```