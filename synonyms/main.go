package main

import (
	"bufio"
	"fmt"
	"github.com/sk62793/thesaurus"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	t := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := t.Synonyms(word)
		if err != nil {
			log.Fatalf("%qの類語検索に失敗しました: %v\n", word, err)
		}
		if len(syns) == 0 {
			log.Fatalf("%q二塁後は有馬線でhした\n", word)
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
