# Prometheus Sandbox

This is a personal learning sandbox so I can poke around Prometheus.

## Running
To start prometheus
```sh
docker compose -f ./starter/docker-compose.yml up -d
```

To start the instrumented http servers. If building for the first time, this might take a bit because I didn't optimize the builds in any way...

```sh
docker compose -f ./go_instrumentation/docker-compose.yml up
```

## Interacting

### Prometheus
Visit `localhost:9090` to view the prometheus UI.

### Servers
Each server exposes a `hello` endpoint that you can interact with. Eg,

```json
curl -s localhost:8002/hello | jq
{
  "ip": "172.26.0.1:59192",
  "message": "Hello From Go!"
}

curl -s localhost:8003/hello | jq
{
  "ip": "172.26.0.1:63558",
  "message": "Hello From Go!"
}
```

Some have artificial timeouts to observe metric differences with prometheus.