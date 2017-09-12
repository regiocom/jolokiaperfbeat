package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/regiocom/jolokiaperfbeat/beater"
)

func main() {
	err := beat.Run("jolokiaperfbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
