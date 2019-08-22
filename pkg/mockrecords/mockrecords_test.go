package mockrecords_test

import (
	"testing"

	"github.com/chanchailee/golang-grpc-credit-risk-api/pkg/mockrecords"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRecords(t *testing.T) {
	assert.Nil(t, mockrecords.Generate(10, "test.csv"))
	assert.FileExists(t, "test.csv")
}
