package FileManager

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {
	result := GetData("../cmd/urls.txt")
	assert.IsType(t, []string{}, result, "GetData should return []string")
}

func TestParseFile(t *testing.T) {
	file := new(File)
	f, err := os.Open("../cmd/urls.txt")
	if err != nil {
		panic(err)
	}
	file.Data = f
	res := ParseFile(file)
	assert.IsType(t, []string{}, res, "GetData should return []string")
}
