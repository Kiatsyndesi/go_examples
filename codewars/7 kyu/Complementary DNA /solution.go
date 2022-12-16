package Complementary_DNA_

import (
	"strings"
)

func DNAStrand(dna string) string {
	symbols := map[string]string{
		"A": "T",
		"T": "A",
		"G": "C",
		"C": "G",
	}

	result := ""
	dnaInput := strings.Split(dna, "")

	for _, value := range dnaInput {
		result += symbols[value]
	}

	return result
}
