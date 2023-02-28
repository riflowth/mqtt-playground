package sensors

import (
	"encoding/csv"
	"fmt"
	"os"
)

func NewSensors() *csv.Reader {
	file, err := os.Open("SampleInput.csv")
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(file)
	csvReader.Read()
	return csvReader
}

func Read(sensors *csv.Reader) string {
	data, err := sensors.Read()
	if err != nil {
		panic(err)
	}

	dataCombined := fmt.Sprintf("%v %v %v %v %v\n", data[0], data[1], data[2], data[3], data[4])
	return dataCombined
}
