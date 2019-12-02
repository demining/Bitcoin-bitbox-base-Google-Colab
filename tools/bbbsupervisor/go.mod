module github.com/digitalbitbox/bitbox-base/tools/bbbsupervisor

go 1.13

require (
	github.com/digitalbitbox/bitbox-base/middleware v0.0.0-20191203134058-7c461ef0407d
	github.com/digitalbitbox/bitbox02-api-go v0.0.0-20191122093321-5bacb3c08094
	github.com/tidwall/gjson v1.3.4
)

replace github.com/digitalbitbox/bitbox02-api-go => /home/b10c/shift/bitbox02-api-go // TODO: remove before merge
