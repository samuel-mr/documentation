https://www.youtube.com/watch?v=JWu33wDT1kk


- Imagen de la empresa `wurstmeister`
    - (url)[https://hub.docker.com/r/wurstmeister/kafka]
    - Enfocada en simplicidad: está pensado para desarrolladores en local
    - Compatible con Docker Compose y puede trabajar : Kafka + Zookeeper
    - Soporte integrado para múltiples brokers (lo que permite simular un entorno distribuido de múltiples brokers)
    - Alta configurabilidad externa: permite enviar diferentes parametros de entorno en vez de modifiar la configuracion interna

    >  Confluent o Bitnami: sus imagenes están más pensadas para Kubernetes

- Start
    ```ps1
    docker-compose -f docker-compose.yml up -d
    # Viewers
    a) http://localhost:9000/
    b) Offset Explorer:
        - Properties:
            - Zookeeper Host: localhost
            - Zookeeper Port: 2182
        - Advanced: Bootstra servers : localhost:9092
    ```
- Docker test
    ```ps1
    # run
    docker-compose -f docker-compose.yml up -d
    # list running images
    docker ps
    # Run an interactive 'sh' shell over container called 'kafka' 
    docker exec -it  kafka /bin/sh
    # Navigate to command's folder
    cd opt/kafka_2.13-2.8.1/bin     #check the kafka version folder
    # Create 'quickstart' topic 
    kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic quickstart
    # Produce a message
    kafka-console-producer.sh --topic quickstart --bootstrap-server localhost:9092
    kafka-console-producer.sh --topic test-topic --bootstrap-server localhost:9092
    # ... add some words + press enter...

    # Consume a mesage
    kafka-console-consumer.sh --topic quickstart --from-beginning --bootstrap-server localhost:9092
    ```
- Kafka veresion:
    ```ps1
    docker exec -it kafka bash
    opt/kafka_2.13-2.8.1/bin/kafka-topics.sh --version
    ```
Some aditional utils
```ps1
# Individual launch
docker-compose up [serviceName]
docker-compose up zookeeper
docker-compose up kafka
```
# Tools
- (Offset Explorer)[https://www.kafkatool.com/download.html] : Administra, monitorea y debug de cluster de kafka
     - choco install offset-explorer