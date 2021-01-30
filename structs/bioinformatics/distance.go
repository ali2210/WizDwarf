package bioinformatics

// import libraries
import (
	"strings"

	"github.com/agnivade/levenshtein"
)

type LevenTable interface {
	// Calcualte Distance
	EditDistanceStrings(s1, s2 []string) int
}

type SequenceMatch struct{}

func NewMatch() LevenTable {
	return &SequenceMatch{}
}

func (*SequenceMatch) EditDistanceStrings(s1, s2 []string) int {
	rowString := strings.Join(s1, " ")
	colStrings := strings.Join(s2, " ")

	return levenshtein.ComputeDistance(rowString, colStrings)
}
