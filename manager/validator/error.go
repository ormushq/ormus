package validator

import "fmt"

type Error struct {
	Fields map[string]string `json:"error"`
	Err    error             `json:"message"`
}

func (v Error) Error() string {
	var err string

	for key, value := range v.Fields {
		err += fmt.Sprintf("%s: %s\n", key, value)
	}

	return err
}
