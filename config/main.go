package config

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
)

var (
	Port      int
	Directory string
	LogLevel  string
)

func initFlags() {
	flag.Usage = usage
	flag.IntVar(&Port, "p", 5030, "Port to use (default 5030)")
	flag.StringVar(&Directory, "d", ".", "Directory to serve (default '.')")
	flag.StringVar(
		&LogLevel,
		"log-level",
		"info",
		"Set the logging level ('info', 'debug', 'warn', 'error') (default 'info')",
	)
}

func usage() {
	fmt.Println("usage: restatic [options]")
	fmt.Println("")
	fmt.Println("A simple HTTP server that serves a local directory over HTTP.")
	fmt.Println("")
	fmt.Println("options:")
	fmt.Println("  -p --port         Port to use (default 5030)")
	fmt.Println("  -d --directory    Directory to serve (default '.')")
	fmt.Println(
		"     --log-level    Set the logging level ('info', 'debug', 'warn', 'error') (default 'info')",
	)
	fmt.Println("")
	fmt.Println(
		"Read README.md for more help on how to use restatic",
	)
}

func initLog() {
	switch LogLevel {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

func init() {
	initFlags()
	initLog()
}
