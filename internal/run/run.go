package run

import (
	cfg "docked/goose/internal/config"
	shellparams "docked/goose/internal/config/shell-params"
	"flag"
	"fmt"
	"github.com/pressly/goose"
	"log"
	"os"
)

var (
	flags   = flag.NewFlagSet("goose", flag.ExitOnError)
	dir     = "./migrations/"
	config  = flags.String("config", "shell-params", "read database config from(shell-params, env, vault")
	source  = flags.String("source", "", "source of the config")
	verbose = flags.Bool("v", false, "enable verbose mode")
	help    = flags.Bool("h", false, "print help")
	version = flags.Bool("version", false, "print version")
)

func Run() {
	flags.Usage = usage
	_ = flags.Parse(os.Args[1:])

	if *version {
		fmt.Println(goose.VERSION)
		return
	}
	if *verbose {
		goose.SetVerbose(true)
	}

	args := flags.Args()
	if len(args) == 0 || *help {
		flags.Usage()
		return
	}

	var configurator cfg.Configurator

	switch *config {
	case "shell-params":
		configurator = shellparams.NewShellParamConfig(*source)

	default:
		configurator = shellparams.NewShellParamConfig(*source)
	}

	command := args[0]
	driver, _ := configurator.GetDriver()
	dbString, _ := configurator.GetDbString()

	switch command {
	case "create":
		if err := goose.Run("create", nil, dir, args[1:]...); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	case "fix":
		if err := goose.Run("fix", nil, dir); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	}

	db, err := goose.OpenDBWithDriver(driver, dbString)
	if err != nil {
		log.Fatalf("-dbstring=%q: %v\n", dbString, err)
	}

	if err := goose.Run(command, db, dir, args[1:]...); err != nil {
		log.Fatalf("goose run: %v", err)
	}
}

func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
	fmt.Println(usageCommands)
}

var (
	usagePrefix = `Usage: goose [OPTIONS] COMMAND

Drivers:
    postgres

Examples:

    goose -c "user=postgres dbname=postgres sslmode=disable" status

Options:
`

	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations
`
)
