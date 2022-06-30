# Masstransit
templates
```
dotnet new -i MassTransit.Templates
```
exec
```
dotnet new mtworker -n GettingStarted
cd GettingStarted
dotnet new mtconsumer -n Hello

```
## Format for names
```
  public class MiMessage
  {
```
| Formater                          | Result     |
|-----------------------------------|------------|
| Default                           | mimessage  |
| SetKebabCaseEndpointNameFormatter | mi-message |
| SetSnakeCaseEndpointNameFormatter | mi_message |

 
### Cambiar el nombre del `Queue` (aplicar en Consumer)
1. En el  Consumer: (v 8.x ... pero no funciona bien con la v 7.x, )
```
  services.AddMassTransit(bus =>
      {
        bus.SetEndpointNameFormatter(new CustomNameFormatter("XXX", false));
```
```
  public class CustomNameFormatter :
    SnakeCaseEndpointNameFormatter
  {
    const char KebabCaseSeparator = '-';
    public CustomNameFormatter(bool includeNamespace)
      : base(KebabCaseSeparator, null, includeNamespace)
    {
    }
    public CustomNameFormatter(string prefix, bool includeNamespace)
      : base(KebabCaseSeparator, prefix, includeNamespace)
    {
    }
    protected CustomNameFormatter()
      : base(KebabCaseSeparator, null, false)
    {
    }
    public new static IEndpointNameFormatter Instance { get; } = new CustomNameFormatter();
    public override string SanitizeName(string name)
    {
      return base.SanitizeName(name).Replace(SnakeCaseSeparator, KebabCaseSeparator);
    }
  }
```

### Cambiar el nombre del 'Topic' (aplicar en Consumer y Sender)
(Solo como ejemplo, así es el formato por defecto que usa para ServiceBus)
```
bus.UsingAzureServiceBus((context, cfg) =>
{
  cfg.Message<MiMessage>((t) =>
  {
    IEntityNameFormatter formatter = new MessageNameFormatterEntityNameFormatter(new ServiceBusMessageNameFormatter());
    IMessageEntityNameFormatter<MiMessage> messageFo = new MessageEntityNameFormatter<MiMessage>(formatter);
    t.SetEntityNameFormatter(messageFo);
```
Nombre simple
```
bus.UsingAzureServiceBus((context, cfg) =>
{
  cfg.Message<MiMessage>((t) =>
  {
    t.SetEntityName($"TOPIC-{nameof(MiMessage)}");
```

# Clases principales
## 1)
Se obtiene el nombre de los entities: IMessageEntityNameFormatter (MassTransit.Abstraction)
```
IMessageEntityNameFormatter


// implementacion
  StaticEntityNameFormatter  : 
  MessageEntityNameFormatter<TMessage> : 
```
## 2)
(al parecer este  da el nombre del Queue)
```
IMessageNameFormatter

// implementations
 DefaultMessageNameFormatter :
 ActiveMqMessageNameFormatter :
 AmazonSqsMessageNameFormatter :
 ServiceBusMessageNameFormatter :
 RabbitMqMessageNameFormatter :
```

## 3)
```
IEntityNameFormatter

// implementations
PrefixEntityNameFormatter :
MessageUrnEntityNameFormatter :
MessageNameFormatterEntityNameFormatter : 

```
## 4) 
```
    public class ServiceBusReceiveEndpointBuilder

    // methodos:
    GenerateSubscriptionName() // !!!!
```

```
public static class AzureBusFactory

//metodos: tiene este metodo que almacena en cache, 
   // ServiceBusMessageNameFormatter es el formateador para los Topics
    _entityNameFormatter = new MessageNameFormatterEntityNameFormatter(new ServiceBusMessageNameFormatter());
```
