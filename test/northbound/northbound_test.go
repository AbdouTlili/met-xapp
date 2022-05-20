package ha

import (
	"testing"

	"github.com/AbdouTlili/met-xapp/pkg/northbound"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func Test_CreateParameter(t *testing.T) {

	param := northbound.Parameter{}

	protoBytes, err := proto.Marshal(&param)

	assert.NoError(t, err)

	decodedParam := &northbound.Parameter{}
	err = proto.Unmarshal(protoBytes, decodedParam)

	assert.NoError(t, err)
	assert.Equal(t, decodedParam.Name, param.Name)
	assert.Equal(t, param.Value, decodedParam.Value)
}
