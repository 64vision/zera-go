package email

import (
	"zerago/config"

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
	SenderEmail = config.CONFIG.Emailconfig.SenderEmail
	SenderName = config.CONFIG.Emailconfig.SenderName
	SesID = config.CONFIG.Emailconfig.SesID
	SesSecret = config.CONFIG.Emailconfig.SesSecret
	CharSet = config.CONFIG.Emailconfig.CharSet
	DBM = config.DBM
}
