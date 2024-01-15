package tunnel

import (
	"github.com/Paulgudring/clash/tunnel"
)

func QueryMode() string {
	return tunnel.Mode().String()
}
