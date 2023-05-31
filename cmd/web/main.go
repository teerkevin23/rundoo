package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"sync"

	"github.com/teerkevin23/rundoo/cmd/web/repository"
	"github.com/teerkevin23/rundoo/cmd/web/usecase"

	"github.com/teerkevin23/rundoo/internal/version"
	"github.com/teerkevin23/rundoo/cmd/web/domain"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)

	err := run(logger)
	if err != nil {
		trace := debug.Stack()
		logger.Fatalf("%s\n%s", err, trace)
	}
}

type config struct {
	baseURL  string
	httpPort int
}

type application struct {
	config config
	logger *log.Logger
	wg     sync.WaitGroup
	ProductUsecase domain.ProductUsecase
}


func doSomething(ctx context.Context) {
	fmt.Println("Doing something!")
}


func run(logger *log.Logger) error {
	var cfg config

	flag.StringVar(&cfg.baseURL, "base-url", "http://localhost:4444", "base URL for the application")
	flag.IntVar(&cfg.httpPort, "http-port", 4444, "port to listen on for HTTP requests")

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}
	productRepository := repository.NewProductRepository()
	app := &application{
		config: cfg,
		logger: logger,
		ProductUsecase: usecase.NewProductUsecase(productRepository),
	}

	ctx := context.TODO()
	doSomething(ctx)

	return app.serveHTTP()
}
