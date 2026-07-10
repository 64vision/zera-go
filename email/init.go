package email

import (
	"zerago/dbconfig"

	"github.com/go-pg/pg"
)

var DBM *pg.DB

var (
	// Replace sender@example.com with your "From" address.
	// This address must be verified with Amazon SES.
	SenderEmail = ""
	SenderName  = ""
	SesID       = ""
	SesSecret   = ""
	// The character encoding for the email.
	CharSet = "UTF-8"
)

func init() {
	SenderEmail = dbconfig.CONFIG.Emailconfig.SenderEmail
	SenderName = dbconfig.CONFIG.Emailconfig.SenderName
	SesID = dbconfig.CONFIG.Emailconfig.SesID
	SesSecret = dbconfig.CONFIG.Emailconfig.SesSecret
	CharSet = dbconfig.CONFIG.Emailconfig.CharSet
	DBM = dbconfig.DBM
}
