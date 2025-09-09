package handler

import "encoding/json"

func JsonError(msg string) []byte {
	// Criando uma struct temporária apenas em tempo de execução
	error := struct {
		Message string `json:"message"`
	}{
		msg,
	}

	result, err := json.Marshal(error)
	if err != nil {
		return []byte(err.Error())
	}

	return result
}
