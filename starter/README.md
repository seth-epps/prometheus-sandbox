# Setup Prometheus
To setup prometheus I set up some demo configuration and alerts and use the local directory for data persistence. To make sure this works, make sure you run `mkdir -p prom_data`

## Some extra context
This docker compose creates a network (`prom_net`) so other containers can easily attach.


## Running
```
docker compose up -d
```

visit `localhost:9090` in the browser to view the UI

## Node exporter
To understand the environment prometheus is running in I added a node exporter.

### Container setup
add the service to `docker-compose`
```yaml
  node-exporter:
    image: prom/node-exporter:v1.5.0
    ports:
      - 9100:9100
```

then run `docker compose up -d node-exporter` and check that it works by running `curl localhost:9100/metrics`.


### Configuring Prometheus to use scrape node_exporter
Add the node exporter job to the prometheus static scrape configs
```diff
scrape_configs:
  - job_name: services
    metrics_path: /metrics
    static_configs:
      - targets:
          - 'prometheus:9090'
          - 'idonotexists:564' # non-existent target in the network to demonstrate alert management

+  - job_name: node_exporter
+    static_configs:
+      - targets:
+        - 'node-exporter:9100'
```

Now reload prometheus to get the new config with `curl -X POST localhost:9090/-/reload`



