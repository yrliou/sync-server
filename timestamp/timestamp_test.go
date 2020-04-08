package timestamp

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTimestamp(t *testing.T) {
	rsp, err := GetTimestamp()
	assert.Nil(t, err)

	// Unmarshal to get the timestamp value
	time := Timestamp{}
	err = json.Unmarshal(rsp, &time)
	assert.Nil(t, err)

	expectedJSON := "{\"timestamp\":\"" + time.Timestamp + "\"}"
	assert.Equal(t, expectedJSON, string(rsp))
}
