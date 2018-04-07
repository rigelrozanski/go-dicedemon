package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

func main() {

	i := bip39.ReverseWordMap["figure"]
	fmt.Println(i)
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, uint16(i))
	fmt.Printf("%b\n", bs)

	sRep := strconv.FormatInt(688, 2)
	sRep = fmt.Sprintf("%011v", sRep)
	fmt.Println(sRep)

	var input string
	for i, s := range sRep {
		input += string(s)
		if (i+1)%8 == 0 {
			input += " "
		}
	}
	fmt.Println(input)
	//input := "01010101 101"
	//input := "01010100 01100101 01110011 01110100"
	b := make([]byte, 0)
	for _, s := range strings.Fields(input) {
		n, _ := strconv.ParseUint(s, 2, 8)
		b = append(b, byte(n))
	}
	fmt.Printf("%b", b)

	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter Words: ")
	//text, _ := reader.ReadString('\n')
	//words := strings.Split(text, " ")

	//if len(words) != 23 {
	//panic("I only take 23 words")
	//}

	//entropy := make([]byte, 256)
	//for i, word := range words {
	//entropyWord, err := bip39.MnemonicToByteArray(word)
	//if err != nil {
	//panic(fmt.Sprintf("bad word \"%v\" err: %v\n", word, err))
	//}

	//copy(entropy[i:(i+10)], entropyWord[:])
	////entropy[i:(i + 10)] = entropyWord
	//}

	//// Generate a mnemonic for memorization or user-friendly seeds
	////entropy, _ := bip39.NewEntropy(256)
	//mnemonic, _ := bip39.NewMnemonic(entropy)

	//// Display mnemonic and keys
	//fmt.Println(entropy)
	//fmt.Printf("%x\n", entropy)
	//fmt.Println(mnemonic)
}
