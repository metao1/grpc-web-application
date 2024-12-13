package data

import (
	"encoding/json"
	"fmt"

	"github.com/metao1/creativefabrica/backend/internal/file"
)

type Creator struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type Product struct {
	ID         string `json:"id"`
	CreatorID  string `json:"creatorId"`
	CreateTime string `json:"createTime"`
}

type Payload struct {
	Creators []Creator `json:"Creators"`
	Products []Product `json:"Products"`
}

func ReadData(filePath string) (*Payload, error) {
	file, err := file.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("%w", err)

	}
	defer file.Close()
	var data *Payload

	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON data: %w", err)
	}
	return data, nil
}
