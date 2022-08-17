# Net CORE

## Request
```
Input formatter ->      Header: Content-Type
Output formater ->      Header : Accept
```
examples:
```
Accept : */*     // todo
Accept : xml 

```
scenario 1: Return 406 error if format returned is diferent than format requested ( in header 'Accept')
```
builder.Services.AddControllers(options =>
{
  options.ReturnHttpNotAcceptable = true;
});
```
scenario 2: Change between requested format (json and xml)
```
builder.Services.AddControllers()
    .AddXmlDataContractSerializerFormatters()
```

| Lib                    | Descrip                                           |
|------------------------|---------------------------------------------------|
| Swashbuckle.AspNetCore | OpenApi implementation for Asp.Net Core (swagger) |

> To Host: Kestrel (prefered) over IISExpress

## Response 

ìn `header`
```
 Response.Headers.Add("X-NameOfHeader", JsonSerializer.Serialize(someEntity));
```
## Params
```
 [Route("api/customname/{cityId}/[controller]")]
 public class UserController : ControllerBase
 {
   [HttpGet]
   public ActionResult<T> Get(int cityId)
```
## Result
Generic
```
    public async Task<IActionResult> Get(...          # we can use Ok(), NotFound(), etc
    public async Task<ActionResult<Entidad>> Get(...  # recommended for Swagger 
```
JsonResult
```
[HttpGet()]
public JsonResult GetCities()
{
    return new JsonResult(CitiesDataStore.Current.Cities);
}
```
ActuonResult (allow return multiple status code, recommended)
```
[HttpGet("{id}")]
public ActionResult<CityDto> GetCitie(int id)
{
    var result = CitiesDataStore.Current.Cities.FirstOrDefault(x => x.Id == id);
    if (result == null)
        return NotFound();

    return Ok(result);
}
```

## Send Data
```
From:
    - Form
    - Body
          [FromBody] T parameter  // explicit
          T parameter             // implicit: infered only of T is a class
    - Header
          [FromHeader]
    - Query (string parameters)
          [FromQuery] string? name                        // default
          [FromQuery(Name ="anotherName")] string? name   // custom name
    - Route (route data)
    - Services (the service injected as action parameter)
```
# DI

request directly: (but constructor injection is preferred)
```
HttpContext.RequestServices.GetService(...)
```

### Reading 'AppSettings' values
appsettings.json
```
 "node": {
    "property": "valueToField",
  }
```
in program.cs
```
var prop = builder.Configuration["node:property"];
```
as DI
```
public Service(IConfiguration configuration)
{
  _field = configuration["node:property"];


```
# Login

in `appsettings.json`
```
// Set minimun: Critical, Error, Warning, Information, Debug, Trace
{
  "Logging": {
    "LogLevel": {
      "Default": "Information",                                  // set for all
      "LearnWebApi.Controllers.CitiesController": "Information", // set to specific class (or namespaces)
      "Microsoft.AspNetCore": "Warning",
      "Microsoft.EntityFrameworkCore.Database.Command" :  "Information" // see EF commands applying
    }
  }
}
```

# Mapper

lib
```
AutoMapper.Extensions.Microsoft.DependencyInjection
```
### use
```
// in program.cs
builder.Services.AddAutoMapper(AppDomain.CurrentDomain.GetAssemblies());

// in each entity to map:
  public class CityProfile : Profile
  {
    public CityProfile()
    {
      CreateMap<Entities.City, Models.CityWithoutPointsOfInterestDto>();
    }
  }
```
# Authentication
package
```
   // to generate
Microsoft.IdentityModel.Tokens
System.IdentityModel.Tokens.Jwt

    // to consume
Microsoft.AspNetCore.Authentication.JwtBearer
```
# source-code
```
https://github.com/dotnet/aspnetcore/blob/v6.0.0/src/DefaultBuilder/src/WebHost.cs
```
- [providers](https://docs.microsoft.com/en-us/ef/core/providers/?tabs=dotnet-core-cli)

# Documentation

### Swashbuckle.AspNetCore
- Generate OpenAPI specification from your api
- Wraps Swagger-UI and provides an embedded version of it
- [repo](https://github.com/domaindrivendev/Swashbuckle.AspNetCore)
```
dotnet add package --version 6.4.0 Swashbuckle.AspNetCore
```

### NSwag
- [official](https://docs.microsoft.com/en-us/aspnet/core/tutorials/getting-started-with-nswag?view=aspnetcore-6.0&tabs=visual-studio)