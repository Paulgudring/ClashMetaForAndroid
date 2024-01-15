package tunnel

import "github.com/Paulgudring/clash/adapter/provider"

func Suspend(s bool) {
	provider.Suspend(s)
}
