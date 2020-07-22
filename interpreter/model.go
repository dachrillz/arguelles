package interpreter

type gender int
type number int

const (
	singular number = iota
	plural
)

const (
	none gender = iota
	male
	female
)

//Vocab vocabulary item
type Vocab struct {
	Source    string
	TargetMap map[string]TargetVocab
}

type TargetVocab struct {
	Target string
	Gender gender
	Number number
	Desc   string
	Stats  *statistics
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
