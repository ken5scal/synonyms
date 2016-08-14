package main

import (
	"os"
	"github.com/ken5scal/thesaurus"
	"bufio"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurus := &thesaurus.BigHuge{APIKEy:apiKey}

	s := bufio.NewScanner(os.Stdin)
}
