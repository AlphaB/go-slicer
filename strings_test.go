package slicer

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StringsSuite struct {
	suite.Suite
}

func TestStrings(t *testing.T) {
	suite.Run(t, new(StringsSuite))
}

func (s *StringsSuite) TestStrings() {
	slice := generateStringSlice(100)
	result := Strings(slice, 7)

	s.Assert().Len(result, 15)

	slice = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	result = Strings(slice, 2)
	s.Assert().Equal([][]string{
		{"1", "2"}, {"3", "4"}, {"5", "6"}, {"7", "8"}, {"9"},
	}, result)
	s.Assert().Len(result, 5)

	s.Assert().Equal([][]string{slice}, Strings(slice, 10))
}

func BenchmarkStrings(b *testing.B) {
	slice := generateStringSlice(1000)
	for n := 0; n < b.N; n++ {
		Strings(slice, 9)
	}
}

func generateStringSlice(size int) []string {
	result := make([]string, size)
	for i := 0; i < size; i++ {
		result[i] = strconv.FormatInt(rand.Int63n(100000), 10)
	}
	return result
}
