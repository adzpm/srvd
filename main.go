package main

import (
	"flag"
	"os"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	dir string
	prt string
	adr string
)

func main() {
	flag.StringVar(&dir, "d", "./", "directory to serve")
	flag.StringVar(&prt, "p", "8080", "port to serve on")
	flag.StringVar(&adr, "a", "127.0.0.1", "address to serve on")

	flag.Parse()

	srv := fiber.New(fiber.Config{
		StrictRouting:         true,
		CaseSensitive:         true,
		DisableStartupMessage: true,
		AppName:               "srvd",
	})

	srv.Use(logger.New(logger.Config{
		Format:        "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
		TimeFormat:    time.TimeOnly,
		TimeZone:      "Local",
		DisableColors: false,
		Output:        os.Stdout,
	}))

	srv.Static("/", dir, fiber.Static{
		Browse: true,
	})

	if err := srv.Listen(adr + ":" + prt); err != nil {
		panic(err)
	}
}
