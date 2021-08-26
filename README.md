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
docker run --name mini-node-exporter
           -v     /proc:/proc:ro
           -e     23333:23333
            mostafaelmenbawy/mini-node-exporter:latest
```
