package main

import (
	"docked/goose/internal/run"
	_ "docked/goose/migrations"
	_ "github.com/lib/pq"
)

func main() {
	run.Run()
}
