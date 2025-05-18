package cooc

type CooccurrenceMatrix struct {
	C         [][]int
	VocabMap  map[string]int
	VocabSize int
}

// Initialize a Co-occurrence matrix struct
func NewCoocMatrix(vocabSize int) *CooccurrenceMatrix {
	return &CooccurrenceMatrix{VocabSize: vocabSize}
}

// Initialize the co-occurrence matrix C
func (cm *CooccurrenceMatrix) InitC() {
	cm.C = make([][]int, cm.VocabSize)
	for i := range cm.C {
		cm.C[i] = make([]int, cm.VocabSize)
	}
}

// Add a context map to cumulative counts in co-occurrence matrix
func (cm *CooccurrenceMatrix) ReduceDocMap(docMap map[string]map[string]int) {
	for iWord, iMap := range docMap {
		for jWord, val := range iMap {
			i := cm.VocabMap[iWord]
			j := cm.VocabMap[jWord]
			cm.C[i][j] = cm.C[i][j] + val
		}
	}
}
