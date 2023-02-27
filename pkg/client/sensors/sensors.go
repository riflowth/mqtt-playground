package sensors

import (
	"encoding/csv"
	"os"
	"strconv"
)

type SensorData struct {
	NodeId       string
	Time         string
	Humidity     float64
	Temperature  float64
	ThermalArray string
}

func NewSensors() *csv.Reader {
	file, err := os.Open("SampleInput.csv")
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(file)
	csvReader.Read()
	return csvReader
}

func Read(sensors *csv.Reader) SensorData {
	var data SensorData
	recv, err := sensors.Read()
	if err != nil {
		panic(err)
	}

	data.NodeId = recv[0]
	data.Time = recv[1]
	data.Humidity, err = strconv.ParseFloat(recv[2], 64)
	data.Temperature, err = strconv.ParseFloat(recv[3], 64)
	data.ThermalArray = recv[4]

	return data
}
