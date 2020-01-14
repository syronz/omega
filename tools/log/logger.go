// Package log TODO: this type of logging is depricated, maybe I should delete it later
package log

import (
	"log"
)

// Logger is used for logging the information
func Logger(str string) {
	log.Println(str)
}
