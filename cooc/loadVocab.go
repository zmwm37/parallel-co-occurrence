package cooc

import (
	"encoding/json"
	"fmt"
	"os"
)

// Load vocabularly from pre-created file.
func LoadVocab(vocabPath string) map[string]int {
	var vocabMap map[string]int

	f, err := os.Open(vocabPath)
	if err != nil {
		fmt.Println("vocabPath:", vocabPath)
		panic("Could not read filepath")
	}
	reader := json.NewDecoder(f)
	err = reader.Decode(&vocabMap)
	if err != nil {
		panic("Could not decode string")
	}

	return vocabMap
}
