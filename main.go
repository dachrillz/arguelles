package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	i "github.com/dachrillz/arguelles/interpreter"
)

func onExit(v []i.Vocab) {
	for _, v := range v {
		fmt.Printf("%s:\t %f (total: %d)\n", v.Target, v.Stats.GetFraction(), v.Stats.Total)
	}
}

func checkVocab(input *string, v *i.Vocab) bool {
	return *input == v.Target
}

func handleVocab(input *string, v *i.Vocab) {
	if checkVocab(input, v) {
		v.Stats.AddSucc()
		fmt.Print("Correct!\n")

	} else {
		fmt.Printf("Wrong: %s\n", v.Target)
		v.Stats.AddFail()
	}
}

func play(v []i.Vocab) {
	text := ""

	reader := bufio.NewReader(os.Stdin)
	for {
		i := rand.Intn(len(v))
		curr := v[i]
		fmt.Printf(":> %s : ", curr.Source)
		text, _ = reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if strings.HasPrefix(text, "quit") {
			onExit(v)
			break
		} else {
			handleVocab(&text, &curr)
		}
	}
}

func main() {
	f, err := os.Open("vocab")
	if err != nil {
		panic(err.Error())
	}

	//v, err := parse(f)
	if err != nil {
		panic(err.Error())
	}

	defer f.Close()
	//play(*v)
}
