<img width="200" src="https://github.com/marcusolsson/gophers/blob/master/gophernotes-gopher.png?raw=true" alt="" />

# MQTT Project
> This is a part of CPE314 Computer Networks project for learning/practicing about the MQTT

## Prequisite
- **[GO](https://go.dev/) 1.19**
- **[Docker](https://www.docker.com/)**
- **Make**
  - *For window: GNU make, MinGW*

## Architecture
- **Broker**, we use broker implementation from [mochi-co/mqtt](https://github.com/mochi-co/mqtt/)
- **Subscriber** and **Publisher**, we implemented by [esclipe/paho.mqtt.golang](https://github.com/eclipse/paho.mqtt.golang) with GO
- **Database**, we use [influxdb](https://www.influxdata.com/) (time-series database)
- **Visualization**, our visualization dashboard on [grafana](https://grafana.com/)

## Getting Started

1. Run **InfluxDB** for storing sensors data on Docker by using the command below on your terminal:
```bash
docker compose up -d
```

2. Build broker/publisher/subscriber by just typing:
```bash
make build
```

3. Run broker from built output directory:
```bash
./dist/broker
```

4. Run subscriber:
- Required flags:
  - `id`: unique id to identify the subscriber
  - `topic`: topic to subscribe
- Example:
  ```bash
    ./dist/sub --id sub-1 --topic hello-world
  ```

5. Run publisher:
- Required flags:
  - `id`: unique id to identify the publisher
  - `topic`: topic to publish
- Example:
  ```bash
    ./dist/pub --id pub-1 --topic hello-world
  ```

6. *Chill with your coffee* ☕️
