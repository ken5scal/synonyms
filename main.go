package main

import (
	"os"
	"github.com/ken5scal/thesaurus"
	"bufio"
	"log"
	"fmt"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurus := &thesaurus.BigHuge{APIKEy:apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word:= s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("Failed searching synonyms for %q: %v\n", word, err)
		}

		if len(syns) == 0 {
			log.Fatalf("No synonyms found for %q\n", word)
		}

		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
