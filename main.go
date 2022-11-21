package main

import (
	"fmt"

	"github.com/kzmijak/zswod_api_go/internal/config"
)

func main() {
  var cfg config.Config
	if err:= config.LoadJson(&cfg); err != nil {
    fmt.Print(err)
  }
}