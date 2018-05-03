package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEntropyOptions(t *testing.T) {

	res := getEntropyOptions(1)
	exp := []string{"0", "1"}
	require.Equal(t, exp, res)

	res = getEntropyOptions(2)
	exp = []string{"00", "10", "01", "11"}
	require.Equal(t, exp, res)

	res = getEntropyOptions(3)
	exp = []string{"000", "100", "010", "110", "001", "101", "011", "111"}
	require.Equal(t, exp, res)
}
