package tunnel

import "github.com/MetaCubeX/mihomo/adapter/provider"

func Suspend(s bool) {
	provider.Suspend(s)
}
