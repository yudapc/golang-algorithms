package connectredis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectRedis(t *testing.T) {
	err := ConnectRedis()
	assert.Equal(t, err, nil, "Connected with redis")
}
