package main

import (
	"goprinter/logger"
)

func main() {

	/*
	 * In contrast to the C stdlib writing to stdout
	 * or stderr using the Go stdlib is never buffered
	 * (no matter if using fmt.Printf or log.Logger).
	 */

	for i := 0; i < 99_999; i++ {
		logger.Info.Printf("%d", i)
	}
	logger.Error.Fatal("99")
}
