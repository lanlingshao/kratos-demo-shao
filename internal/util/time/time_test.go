package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDateToTimestamp(t *testing.T) {
	s := "2020-01-01 01:01:00"
	tt, err := DateToTimestamp(s)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, tt, int64(1577811660))
}

func TestTimestampToDate(t *testing.T) {
	// s := "2020-01-01 01:01:00"
	d := TimestampToDate(1577811660)
	assert.Equal(t, d, "2020-01-01")
}
