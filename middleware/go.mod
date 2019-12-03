module github.com/digitalbitbox/bitbox-base/middleware

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/digitalbitbox/bitbox-base/tools/bbbsupervisor v0.0.0-20191204122151-73e6c6c772f3
	github.com/digitalbitbox/bitbox02-api-go v0.0.0-20191122093321-5bacb3c08094
	github.com/flynn/noise v0.0.0-20180327030543-2492fe189ae6
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/websocket v1.4.1
	github.com/stretchr/testify v1.4.0
	github.com/tarm/serial v0.0.0-20180830185346-98f6abe2eb07
	github.com/tidwall/gjson v1.3.4
	golang.org/x/crypto v0.0.0-20191117063200-497ca9f6d64f
)

replace github.com/digitalbitbox/bitbox02-api-go => /home/b10c/shift/bitbox02-api-go // TODO: remove before merge
