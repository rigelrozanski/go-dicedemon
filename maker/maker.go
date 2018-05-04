package maker

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rigelrozanski/go-bip39"
)

// get all possible mnemonics from a mnemonic (without the checksome word)
func PartialMnemonicToAllMnemonic(words []string) (mnemonics []string) {

	l := len(words)
	if !(l == 2 || l == 5 || l == 8 || l == 11 ||
		l == 14 || l == 17 || l == 20 || l == 23) {
		panic("I will only take 2, 5, 8, 11, 14, 17, 20, or 23 words")
	}

	var bitsStr string
	for _, word := range words {
		i := bip39.ReverseWordMap[word]
		if i == 0 && word != "abandon" {
			panic(fmt.Sprintf("bad word: %v\n", word))
		}
		wordBits := strconv.FormatInt(int64(i), 2)
		wordBits = fmt.Sprintf("%011v", wordBits)
		bitsStr += wordBits
	}

	// get all possible entropy that needs to be added
	groupsOfThree := (l + 1) / 3
	entropyLength := (32 * groupsOfThree) - (11 * (groupsOfThree*3 - 1))
	checksumEntropys := GetEntropyOptions(entropyLength)

	for _, checksumEntropy := range checksumEntropys {
		fullBitsStr := bitsStr + checksumEntropy

		// add a space to the bits every 8 characters
		// to process the string bits to actual bytes
		var spaced string
		for i, s := range fullBitsStr {
			spaced += string(s)
			if (i+1)%8 == 0 {
				spaced += " "
			}
		}

		lenEntropy := len(fullBitsStr) / 8
		//fmt.Printf("debug len: %v fullBitsStr: %v\n", lenEntropy, fullBitsStr)

		// process the string bits to bytes
		entropy := make([]byte, lenEntropy)
		for i, s := range strings.Fields(spaced) {
			n, _ := strconv.ParseUint(s, 2, 8)
			b := byte(n)
			entropy[i] = b
		}

		// Generate a mnemonic for memorization or user-friendly seeds
		mnemonic, _ := bip39.NewMnemonic(entropy)

		// Display mnemonic and keys
		mnemonics = append(mnemonics, mnemonic)
	}

	return mnemonics
}

// get the set of all possible entropy with a length of length :)
func GetEntropyOptions(length int) []string {
	if length == 1 {
		return []string{"0", "1"}
	}

	less1 := GetEntropyOptions(length - 1)
	var res []string
	for _, resultless1 := range less1 {
		res = append(res, "0"+resultless1, "1"+resultless1)
	}
	return res
}
