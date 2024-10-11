based on https://github.com/bitnami/containers/blob/main/bitnami/kafka/3.3/debian-12/docker-compose.yml
```ps
# Run
docker-compose -f docker-external.yml up -d
```
## Notes
- docker-external.yml : configuracion pc local
## Testing
```ps1
# go (previously downloaded kafka's binaries)
cd C:\Learn\Kotlin\kafka-binaries-3-3-2\bin\windows

# Create topic
.\kafka-topics.bat --bootstrap-server localhost:9092 --topic mi_topico --create --partitions 1 --replication-factor 1

# Create message (interactively)
.\kafka-console-producer.bat --broker-list localhost:9092 --topic mi_topico

# Consume message (realtime). Just new incoming message
.\kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic mi_topico
.\kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic mi_topico --from-beginning
```

## Troubleshouting
- Aveces, cuando tienes errores por previas configuraciones (en este caso, colisionan)
```
# Elimina el contenedor, el volumen y la red (util para limpiar)
docker-compose down -v
```