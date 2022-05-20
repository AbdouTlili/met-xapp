package test_nb

import (
	"testing"

	"github.com/AbdouTlili/met-xapp/test/northbound"
)

func Test_CreateParameter(t *testing.T) northbound.Parameter {
	return northbound.CreateParameter("test", "test")
}
