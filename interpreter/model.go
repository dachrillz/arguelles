package interpreter

//Vocab vocabulary item
type Vocab struct {
	Source    string
	TargetMap map[string]TargetVocab
}

type TargetVocab struct {
	Language string
	Target   string
	Gender   string
	Number   string
	Desc     string
}

type statistics struct {
	fail  int
	succ  int
	Total int
}

func (s *statistics) AddSucc() {
	s.Total++
	s.succ++
}

func (s *statistics) AddFail() {
	s.Total++
	s.fail++
}

func (s *statistics) GetFraction() float32 {
	return float32(s.succ) / float32(s.Total)
}
