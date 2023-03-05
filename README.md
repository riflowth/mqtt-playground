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

1. Run **InfluxDB** for storing sensors data and **Grafana** for visualize the data on Docker by using the command below on your terminal:
```bash
cd ./.docker
docker compose up -d
```
*(grafana) http exposed at port 4000*
> default username `admin` default password `admin`

*(influxdb) http exposed at port 8086*
> default username `admin` default password `adminadmin`

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
  - `hostname`: hostname to subscribe to the broker
  - `influx-token`: token for authenticate to influxDB
  - `influx-org`: org name for influxDB
  - `influx-bucket`: bueckt name to connect to influxDB
- Example:
  ```bash
    ./dist/sub --id sub-1 --topic hello-world --hostname localhost:1883 --influx-token "tokentoken" --influx-org admin --influx-bucket mqtt
  ```

5. Run publisher:
- Required flags:
  - `id`: unique id to identify the publisher
  - `topic`: topic to publish
  - `hostname`: hostname to publish to the broker
  - `interval`: interval to publish the next message to the broker (second)
- Example:
  ```bash
    ./dist/pub --id pub-1 --topic hello-world --hostname localhost:1883 --interval 180
  ```

6. *Chill with your coffee* ☕️

## Preview
### Dashboard
![grafana dashboard](/docs/preview-grafana-dashboard.gif)
