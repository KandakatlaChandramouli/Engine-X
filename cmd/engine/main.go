package main

import (
	"fmt"

	"engine-x/core"
)

const (
	DBSize = 16 * 1024 * 1024
)

func main() {

	db, err :=
		core.Open(
			"enginex.db",
			DBSize,
		)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println(
		"Engine-X booted",
	)

	fmt.Printf(
		"pages=%d\n",
		db.PageCount(),
	)
}
