package sensors

import (
	"encoding/csv"
	"os"

)

type SensorData struct {
	NodeId       string
	Time         string
	Humidity     string
	Temperature  string
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
	data.Humidity = recv[2]
	data.Temperature = recv[3]
	data.ThermalArray = recv[4]

	return data
}
