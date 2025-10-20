package repo

import (
	"encoding/json"
	"os"
)

// Lưu đối tượng (object graph) vào file JSON
func SaveObjects(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

// Tải lại object từ file JSON (deserialize)
func LoadObjects(path string, data any) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(data)
}
