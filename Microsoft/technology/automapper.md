# Automapper

### Custom Mappings
Replace the original mapping 'Phone' and send this value to 'NEW_COLUMN'
source (json)
```
    {
      "Id": "0",
      "Name": "Test",
      "Phone": "11111",
    }
```
execution
```
var data = readJson();
var customMappings = new List<Mapping>
{
    new() { Source = "Phone", Target = $"{nameof(Company)}.NEW_COLUMN", Action = "replace" }
};

var mappedCustomer = Mapper.Map<Company>(
        data,
    Guid.NewGuid(),
    customMappings);

mappedCustomer.Name.Should().Be("Test");
mappedCustomer.ExportId.Should().Be("0");
mappedCustomer.NEW_COLUMN.Should().Be("11111");
```