package context

type GoContext struct {
	Mode             string
	CorpusPath       string
	WindowSize       int
	VocabMap         map[string]int
	Capacity         int
	NDocuments       int
	WorkLeft         bool
	ReduceChannel    chan map[string]map[string]int
	BalanceThreshold float64
}
