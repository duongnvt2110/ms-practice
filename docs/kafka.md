1. How to choose the no of partitions for a kafka topic?
- http://stackoverflow.com/questions/50271677/how-to-choose-the-no-of-partitions-for-a-kafka-topic
- Consider starting
  - 3-6 partition for a topics

2. Kafka Partition Strategies
- https://www.redpanda.com/guides/kafka-tutorial-kafka-partition-strategy
- https://www.confluent.io/blog/apache-kafka-producer-improvements-sticky-partitioner/
- Without Key
  - Default Partionner
  - Round-robin partitioner
  - Uniform sticky partitioner
- Key
  - Hash Algorithm