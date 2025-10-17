package util

import (
	"log"
)

var L = log.New(log.Writer(), "order-service ", log.LstdFlags|log.Lshortfile)
