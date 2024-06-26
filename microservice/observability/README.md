## Go Microservice Observability

Stack used:
- [Grafana](https://grafana.com/)
- [Prometheus](https://prometheus.io/)
- [Loki](https://grafana.com/oss/loki/)
- [Promtail](https://grafana.com/oss/promtail/)
- [Tempo](https://grafana.com/oss/tempo/)

Additionally, this example uses the [Kafka](https://kafka.apache.org/) stack for sink and source of logs.

### Running the example

Pre-requisites:

- Set /etc/hosts with the following entries:
```bash
127.0.0.1 kafka1
127.0.0.1 kafka2
127.0.0.1 kafka3
```

1. Start the observability stack and Kafka(with KRaft mode):
```bash
docker-compose -f containers/docker-compose.yml up
```

3. Start the microservice:
```bash
go run main.go
```