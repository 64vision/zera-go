package auth

import (
	"zerago/config"
)

var AllowNoAuth []string

func init() {
	AllowNoAuth = config.CONFIG.PathPermission.NoAuth
}
