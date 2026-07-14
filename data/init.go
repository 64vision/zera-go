package data

import (
	"zerago/config"

	"github.com/go-pg/pg"
)

var DBM *pg.DB

var Prohibiteds = []string{"balance", "delete", "level", "truncate", "commission", "agent", "created_at", "password"}

func init() {
	DBM = config.DBM
}
