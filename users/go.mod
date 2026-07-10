module zerago/users

go 1.26.4

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-pg/pg v8.0.7+incompatible
	zerago/config v0.0.0-00010101000000-000000000000
	zerago/email v0.0.0-00010101000000-000000000000
	zerago/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/aws/aws-sdk-go v1.55.8 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	golang.org/x/crypto v0.53.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	mellium.im/sasl v0.3.2 // indirect
)

replace zerago/config => ../config

replace zerago/email => ../email

replace zerago/utils => ../utils
