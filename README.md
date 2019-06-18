This repo is a dockerized version of the https://github.com/pressly/goose
which enables you to build and make go migrations very easily and also
integrate it's migrations into you CI easily

It only includes some docker-compose file which brings up a local 
database connected to the goose instance, so that enables you to run
your migrations locally in your development environment easily

Usage:
Just put your migration files in the migrations folder, or create new ones using the 
available create command through docker-compose:
docker-compose run goose ./goose create [name]
be aware that if you don't use the -config & -source flags, goose is going to use the
default postgres configured instance in the docker-compose which is only suitable for
local development.


Notes:
- should not use the word "test" at the end of table names!
- after creating any migrations written in go, run: make build
- also don't forget to run: make build after any change to the go migration files, before trying to apply the changes

List of some commands to be used:

To rebuild the goose binary
docker-compose run goose go build -o goose ./cmd
or
make build

To run goose commands with default config 
docker-compose run goose ./goose status

To run goose commands passing custom config through command params
docker-compose run goose ./goose -source "user=local password=local dbname=local sslmode=disable host=postgres port=5432" status

To run goose commands passing custom config through db.env file
docker-compose run goose ./goose -config=env -source db.env status

To run goose commands passing custom config through vault
docker-compose run goose ./goose -config=vault -source "vault config!" status