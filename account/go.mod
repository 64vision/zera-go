module zerago/account

go 1.26.4

replace zerago/config => ../config

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-pg/pg v8.0.7+incompatible
	zerago/config v0.0.0-00010101000000-000000000000
	zerago/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/nxadm/tail v1.4.11 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/crypto v0.54.0 // indirect
	mellium.im/sasl v0.3.2 // indirect
)

replace zerago/utils => ../utils
