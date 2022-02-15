package request

import channeltype "github.com/xm-chentl/go-mvc-demo/model/enum/channel-type"

type MobileDevice struct {
	ChannelType  channeltype.Value
	DeviceUnique string
}

func (d *MobileDevice) SetDevice(channelType channeltype.Value, deviceUnique string) {
	d.ChannelType = channelType
	d.DeviceUnique = deviceUnique
}
