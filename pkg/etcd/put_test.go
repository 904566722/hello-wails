package etcd

import (
	"testing"
)

func Test_compact(t *testing.T) {
	test := `
{
  "createBy": "admin",
  "modifiedBy": "admin",
  "createAt": "2023-11-08T14:41:38.157348768+08:00",
  "updateAt": "2023-11-08T14:41:38.157351281+08:00",
  "phyNetworkPortId": "phynic-00B43B8B87",
  "hostId": "39532218-09c1-42af-9f43-c693402fabe9",
  "portName": "eno5",
  "displayName": "eno5",
  "index": 2,
  "portType": "Physical",
  "gatewayIp": "",
  "state": "Up",
  "aggregation": {
    "portNames": null,
    "portIds": null
  },
  "portIp": [
    "10.16.207.141/2"
  ],
  "usedForAggNetPort": false,
  "mtu": 1500,
  "hwAddr": "a4:ae:12:ff:22:e0",
  "speed": 1000,
  "maxSpeed": 1000,
  "healthStatus": "health",
  "healthStatusReason": ""
}

`
	s, err := compact([]byte(test))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(s)
}
