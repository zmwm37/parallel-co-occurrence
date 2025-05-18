package cooc

import (
	"testing"
)

func TestLoadVocab(t *testing.T) {
	t.Run("V=5", func(t *testing.T) {
		filePath := "../util/vocab_map_5.txt"

		vocabMap := LoadVocab(filePath)
		if len(vocabMap) != 5 {
			t.Errorf("Head value incorrect.\ngot:%v\nexpected:%v", len(vocabMap), 5)
		}
	})
	t.Run("V=500", func(t *testing.T) {
		filePath := "../util/vocab_map_500.txt"

		vocabMap := LoadVocab(filePath)
		if len(vocabMap) != 500 {
			t.Errorf("Head value incorrect.\ngot:%v\nexpected:%v", len(vocabMap), 5)
		}
	})
	t.Run("V=1000", func(t *testing.T) {
		filePath := "../util/vocab_map_1000.txt"

		vocabMap := LoadVocab(filePath)
		if len(vocabMap) != 1000 {
			t.Errorf("Head value incorrect.\ngot:%v\nexpected:%v", len(vocabMap), 5)
		}
	})
	t.Run("V=5000", func(t *testing.T) {
		filePath := "../util/vocab_map_5000.txt"

		vocabMap := LoadVocab(filePath)
		if len(vocabMap) != 5000 {
			t.Errorf("Head value incorrect.\ngot:%v\nexpected:%v", len(vocabMap), 5)
		}
	})

}
