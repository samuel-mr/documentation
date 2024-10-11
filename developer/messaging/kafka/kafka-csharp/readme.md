Curso: https://developer.confluent.io/courses/apache-kafka-for-dotnet/producing-messages-hands-on/
- Your Kafka server need to be in: localhost:9092
- Create a topic: `purchases`
- dotnet run c


https://developer.confluent.io/courses/apache-kafka-for-dotnet/producing-messages/

top properties
```c#
new ProducerConfig
{
    // La lista de brokers (servidores) de Kafka a los que conectarse
    BootstrapServers = "localhost:9092",

    // ID único para identificar al productor en los logs del broker
    ClientId = "example-producer",

    // Número de reintentos si falla el envío del mensaje
    MessageSendMaxRetries = 3,

    // Tiempo de espera entre reintentos (en milisegundos)
    RetryBackoffMs = 1000,

    // Nivel de garantía de entrega del mensaje (puede hacer que demore o no el retorno de respuesta)
    // 0 = Sin garantía, 1 = Garantía de escritura en el líder, all = Garantía de escritura en todos los replicas
    Acks = Acks.All,

    // Compresión de los mensajes para reducir el uso de la red
    CompressionType = CompressionType.Snappy,

    // Tamaño máximo del lote de mensajes en bytes
    BatchSize = 16384,

    // Tiempo máximo de espera antes de enviar un lote (en milisegundos)
    LingerMs = 5,

    // Tiempo máximo de espera para una respuesta del servidor (en milisegundos)
    RequestTimeoutMs = 30000,

    // Habilita la idempotencia para garantizar que los mensajes se entreguen exactamente una vez
    EnableIdempotence = true,

    // Número máximo de mensajes sin confirmar que el productor mantendrá en vuelo
    MaxInFlight = 5,
};
```
Load Config
```c#
var ss = builder.Services.Configure<ProducerConfig>(
    builder.Configuration.GetSeection("kafka")
 );
...
{
    "kafka" :
    {
        "BootstrapServers" : "localhost:9092",
        "ClienteId" : "MyPc"
        ...
```