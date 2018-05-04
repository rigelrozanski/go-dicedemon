package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rigelrozanski/go-dicedemon/maker"
)

func main() {

	// get the words
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter all words (besides checksum word) seperated by spaces:\n")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	text = strings.TrimRight(text, "\n") // trim the enter
	words := strings.Split(text, " ")

	mnenomics := maker.PartialMnemonicToAllMnemonic(words)
	for _, m := range mnenomics {
		fmt.Println(m)
	}
}
