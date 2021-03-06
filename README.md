# mini-node-exporter

A web server that expose prometheus metrics on `/metrics`

## Endpoints

- `/info/hostname`: show the hostname with plain text
- `/info/uptime`: show the uptime of the system in seconds with plain text`
- `/info/load`: show the load average in 1m, 5m and 15m with JSON, example {"1m": 0.57, "5m":0.80, "15m":0`77}
- `/metrics`: expose metrics that could be scraped by prometheus

## Metrics

- `node_load`: load average, with a label for three duration
- `node_uptime`: uptime of the system in seconds

## Docker

### Pull image

`docker pull mostafaelmenbawy/mini-node-exporter:latest`

### Run container

```sh
docker run --name     mini-node-exporter \
           --network  host               \
           --rm                          \
            mostafaelmenbawy/mini-node-exporter:latest
```

## Monitoring Helm Charts

### Prometheus

```sh
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install [RELEASE_NAME] prometheus-community/prometheus -f ./kube/helm-values/prometheus.yml
```

### Grafana

```sh
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm install [RELEASE_NAME] grafana/grafana -f ./kube/helm-values/grafana.yml
```
