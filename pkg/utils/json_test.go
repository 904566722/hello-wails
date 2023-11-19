package utils

import "testing"

func TestJQ(t *testing.T) {
	jsonCompact := "{\"actions\":[\"InitClusterInfo\",\"AddHost\",\"AggNicCreate\",\"ConfigStorageNetwork\",\"ConfigStorage\"],\"currentStep\":1,\"status\":\"failed\",\"installReport\":{\"0\":{\"no\":0,\"params\":{\"clusterName\":\"local-hci\"},\"status\":\"succeed\",\"installReport\":null},\"1\":{\"no\":1,\"params\":{\"10.16.0.1\":\"CDAC5F0C\",\"10.16.0.2\":\"ACE08313\"},\"status\":\"failed\",\"installReport\":{\"10.16.0.1\":\"status: failed msg: 添加节点任务执行结束，但数据库中没有host记录\",\"10.16.0.2\":\"status: failed msg: 添加节点任务执行结束，但数据库中没有host记录\"}}},\"hostNum\":0}"
	jsonExpand, err := JQ(jsonCompact)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(jsonExpand)
}

func TestIsJsonFormat(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "{}",
			},
			want: true,
		}, {
			args: args{
				s: "{\"createBy\":\"admin\",\"modifiedBy\":\"admin\",\"createAt\":\"2023-11-17T17:06:03.354246+08:00\",\"updateAt\":\"2023-11-17T17:06:03.354249+08:00\",\"phyNetworkPortId\":\"phynic-0033930FC0\",\"hostId\":\"39532218-09c1-42af-9f43-c693402fabe9\",\"portName\":\"eno4\",\"displayName\":\"eno4\",\"index\":7,\"portType\":\"Physical\",\"gatewayIp\":\"\",\"state\":\"Down\",\"aggregation\":{\"portNames\":null,\"portIds\":null},\"usedForAggNetPort\":false,\"mtu\":1500,\"hwAddr\":\"e8:b4:70:09:cb:6f\",\"portPurpose\":[\"Manager\",\"NorthSouth\"],\"maxSpeed\":-1,\"healthStatus\":\"health\",\"healthStatusReason\":\"\"}",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsJsonFormat(tt.args.s); got != tt.want {
				t.Errorf("IsJsonFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
