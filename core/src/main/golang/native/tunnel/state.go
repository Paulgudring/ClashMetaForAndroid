package tunnel

import (
	"github.com/MetaCubeX/mihomo/tunnel"
)

func QueryMode() string {
	return tunnel.Mode().String()
}
