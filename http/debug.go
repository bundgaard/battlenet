// Copyright (C) 2021 David Bundgaard
package http

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

func Todo(r io.Reader, format string, v ...interface{}) error {
	fd, err := os.Create(fmt.Sprintf(format, v...) + ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	io.Copy(fd, r)

	pc, _, _, _ := runtime.Caller(1)
	c := runtime.FuncForPC(pc)

	whoCalledme := fmt.Sprintf("You are not done yet\nwe need to write the type\nfile and complete the %q", c.Name())
	return fmt.Errorf(whoCalledme)
}
