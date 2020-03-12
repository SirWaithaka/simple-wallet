package registry

import "wallet/data"

type Channels struct {
	ChannelNewUsers data.ChanNewUsers
}

func NewChannels() *Channels {
	chanNewUsers := make(chan data.NewRegisteredUser, 10)

	return &Channels{
		ChannelNewUsers: data.ChanNewUsers{chanNewUsers, chanNewUsers, chanNewUsers},
	}
}
