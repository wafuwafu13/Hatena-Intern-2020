package fetcher

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fetch_Normaly(t *testing.T) {

	url := "https://example.com/"
	expected := "Example Domain"

	title, err := Fetch(context.Background(), url)
	assert.NoError(t, err)
	assert.Equal(t, expected, title)
}

func Test_Fetch_Anormaly(t *testing.T) {

	url := "https://hogehoge/"
	expected := "https://hogehoge/"

	title, err := Fetch(context.Background(), url)
	assert.Error(t, err)
	assert.Equal(t, expected, title)
}
