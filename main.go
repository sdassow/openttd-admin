package main

import (
	"log"

	"github.com/sdassow/openttd-admin/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
