module zerago/route

go 1.26.4

replace zerago/users => ../users

replace zerago/auth => ../auth

replace zerago/email => ../email

replace zerago/utils => ../utils

replace zerago/config => ../config

require (
	github.com/gorilla/mux v1.8.1
	github.com/rs/cors v1.11.1
	zerago/account v0.0.0-00010101000000-000000000000
	zerago/auth v0.0.0-00010101000000-000000000000
	zerago/data v0.0.0-00010101000000-000000000000
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-pg/pg v8.0.7+incompatible // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	golang.org/x/crypto v0.54.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	mellium.im/sasl v0.3.2 // indirect
	zerago/config v0.0.0-00010101000000-000000000000 // indirect
	zerago/utils v0.0.0-00010101000000-000000000000 // indirect
)

replace zerago/data => ../data

replace zerago/account => ../account
