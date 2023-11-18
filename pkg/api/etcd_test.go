package api

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"changeme/pkg/etcd"
)

func TestDoAction(t *testing.T) {
	_, err := etcd.NewEtcdClient()
	assert.NoError(t, err)
	resp, err := DoAction(BaseRequest{
		Data:    "hci/phyNetworkPort/phynic-0097AC50A6",
		KeyType: DTypeWholeKey,
		Action:  ActionGet,
	})
	assert.NoError(t, err)
	assert.Equal(t, CodeSuccess, resp.Code)
	t.Log(resp)
}
