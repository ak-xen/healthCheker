package Query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryManager(t *testing.T) {
	res := QueryManager([]string{"https://www.google.com",

		"https://www.youtube.com",

		"https://www.wikipedia.org",

		"https://www.github.com",

		"https://www.amazon.com",
		"https://www.spotify.com",

		"https://www.nytimes.com",
	})

	assert.IsType(t, map[string]int{}, res, "GetData should return []string")
}
