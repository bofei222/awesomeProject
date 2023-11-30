创建topic
./kafka-topics.sh --zookeeper zookeeper:2181 --create --topic test --replication-factor 1 --partitions 3
创建消息
./kafka-console-producer.sh -- bootstrap-server kafka:9092  --topic test


分割线-----------------------------------------------------------------------------------------------------------------------------

创建topic

bin/kafka-topics.sh --create zookeeper:2181 --replication-factor 2 --partitions 2 --topic testtopic

查看topic

bin/kafka-topics.sh --describe --zookeeper 192.168.243.149:2181,192.168.243.149:2182,192.168.243.149:2183 --topic testtopic

查询topic内容
bin/kafka-console-consumer.sh --bootstrap-server 192.168.243.149:9092 --topic kafka-springboot-001 --from-beginning




生产消息
bin/kafka-console-producer.sh --bootstrap-server 192.168.243.149:9092 --topic realtime

消费消息
bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic test --from-beginning




写入测试
bin/kafka-producer-perf-test.sh --topic test --num-records 10 --record-size 1024  --throughput 20 --producer-props bootstrap.servers=kafka:9092
bin/kafka-producer-perf-test.sh --topic test --num-records 100 --record-size 1024000  --throughput 2000 --producer-props bootstrap.servers=kafka:9092
1024000
1048576

报错
org.apache.kafka.common.errors.RecordTooLargeException: The message is 2024088 bytes when serialized which is larger than 1048576, which is the value of the max.request.size configuration.


更新配置
bin/kafka-configs.sh --bootstrap-server 192.168.243.149:9092 --entity-type brokers --entity-name 0 --alter --add-config log.cleaner.threads=2