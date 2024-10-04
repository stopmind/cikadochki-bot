package bot

import (
	"encoding/json"
	"io"
	"os"
)

type data struct {
	Channels []int64
}

func tryReadData(dataPath string) (data, error) {
	file, err := os.Open(dataPath)
	if err != nil {
		return newData(), err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return newData(), err
	}

	var result data
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return newData(), err
	}

	return result, file.Close()
}

func (d *data) write(dataPath string) error {
	file, err := os.OpenFile(dataPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(d)
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	return err
}

func newData() data {
	return data{
		Channels: []int64{},
	}
}
