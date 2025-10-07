# Kafka
## Kafka in docker 
### Problems
1. The Kafka client and the server may not share the same network in Docker
- Reason: 
  - Background:
    - `LISTENERS`: are what interfaces Kafka binds to.
    - `KAFKA_ADVERTISED_LISTENERS`: are how clients can connect.
  - Kafka uses the `metadata` from `KAFKA_ADVERTISED_LISTENERS` to resolve the `host` and `port`
- Solution:
  - Expose a port for External connections:
    - `LISTENERS: LISTENER_HOST_MACHINE://:29092,LISTENER_LOCALHOST://:29093`
    - `KAFKA_ADVERTISED_LISTENERS:LISTENER_HOST_MACHINE://host.docker.internal:29092,LISTENER_LOCALHOST://localhost:29093`
    - Note:
      - Use `host.docker.interal` instead of `localhost` on macOS because of the VM Network setup.
- References: 
  - https://rmoff.net/2018/08/02/kafka-listeners-explained/
  - https://www.confluent.io/blog/kafka-listeners-explained/

2. Single or multiple node `_consumer_offsets`
- Config `KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR` to <= the number of the broker
- References:
  - https://stackoverflow.com/questions/49490835/kafka-server-offsets-topic-has-not-yet-been-created/49493121
  - https://strimzi.io/blog/2021/06/08/broker-tuning
  - https://docs.confluent.io/platform/current/kafka/post-deployment.html#changing-the-replication-factor
## Kafka cheatsheet
- cd `/opt/kafka/bin`
- List Topic 
```
./kafka-topics.sh --bootstrap-server kafka:9092 --list
```