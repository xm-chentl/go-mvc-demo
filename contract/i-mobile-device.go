package contract

import channeltype "github.com/xm-chentl/go-mvc-demo/model/enum/channel-type"

type IMobileDevice interface {
	SetDevice(channelType channeltype.Value, uniqueCode string)
}
