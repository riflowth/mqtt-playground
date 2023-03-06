package repositories

import (
	"context"
	"time"

	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type SensorData struct {
	NodeId       string
	Time         string
	Humidity     float64
	Temperature  float64
	ThermalArray string
}

type SensorRepository interface {
	Save(metric SensorData) error
}

type sensorRepository struct {
	influxHostName string
	token          string
	org            string
	bucket         string
	writeAPI       api.WriteAPIBlocking
}

func (s sensorRepository) Save(sensorData SensorData) error {
	time, err := time.Parse("2006-01-02 15:04:05", sensorData.Time)

	tempPoint := influxdb.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("value", sensorData.Temperature).
		SetTime(time)

	humPoint := influxdb.NewPointWithMeasurement("stat").
		AddTag("unit", "humidity").
		AddField("value", sensorData.Humidity).
		SetTime(time)

	thermalPoint := influxdb.NewPointWithMeasurement("stat").
		AddTag("unit", "thermalArray").
		AddField("array", sensorData.ThermalArray).
		SetTime(time)

	err = s.writeAPI.WritePoint(context.Background(), tempPoint, humPoint, thermalPoint)

	return err
}

func NewSensorRepository(
	influxHostName string,
	token string,
	org string,
	bucket string,

) SensorRepository {
	client := influxdb.NewClient(influxHostName, token)
	writeAPI := client.WriteAPIBlocking(org, bucket)

	return &sensorRepository{
		influxHostName: influxHostName,
		token:          token,
		org:            org,
		bucket:         bucket,
		writeAPI:       writeAPI,
	}
}
