# Running services in the compose network
Each service can be run together if you specify the `APP_LISTEN_PORT` environment variable/

Add each to the docker compose services with
```yaml
  <service-name>:
    build: ./<service-name-dir>
    ports:
      - <port_num>:<port_num>
    environment:
      - APP_LISTENING_PORT=<port_num>
```

## Add monitoring services
Add each service as a target in `prometheus.yml`
```yaml
  - job_name: <whatever you want>
    static_configs:
      - targets:
        - '<service-name>:<service-port>'
```

then reload prometheus to get the new config with `curl -X POST localhost:9090/-/reload`