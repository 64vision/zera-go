package users

import (
	"zerago/config"

	"github.com/go-pg/pg"
)

var DBM *pg.DB

func init() {
	DBM = config.DBM
}
