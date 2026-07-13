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
	zerago/auth v0.0.0-00010101000000-000000000000
	zerago/data v0.0.0-00010101000000-000000000000
	zerago/users v0.0.0-00010101000000-000000000000
)

require (
	github.com/aws/aws-sdk-go v1.55.8 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-pg/pg v8.0.7+incompatible // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	golang.org/x/crypto v0.53.0 // indirect
	mellium.im/sasl v0.3.2 // indirect
	zerago/config v0.0.0-00010101000000-000000000000 // indirect
	zerago/email v0.0.0-00010101000000-000000000000 // indirect
	zerago/utils v0.0.0-00010101000000-000000000000 // indirect
)

replace zerago/data => ../data
