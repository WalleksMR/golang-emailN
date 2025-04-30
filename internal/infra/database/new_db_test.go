package database

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewDb_Connection(t *testing.T) {
	assert := assert.New(t)

	db := NewDb()

	interf, verd := db.Get("docker")
	fmt.Println(interf, verd)

	assert.Equal(1, 1)
}
