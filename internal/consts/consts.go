package consts

import "math"

// constants which used inside the app
const (
	SecretKeyAES = "83nd81lodhg2kd9243jlqa12345jcDsk32021"
	IVAES        = "7D5Jle$9c2R7vFwL"

	MinimumPasswordChar = 8

	MinPin = 10000000000000
	MaxPin = 99999999999999

	// TemporaryTokenDuration = 100 * 100000 //in seconds
	TemporaryTokenDuration = 10

	HashTimeLayout = "060102150405.000000"

	MaxRowsCount = 1 << 62

	// MinFloat64 = k
	MinFloat64 = -1 * math.MaxFloat64

	DefaultLimit = 10000

	TimeLayout     = "2006-01-02 15:04:05"
	TimeLayoutZone = "2006-01-02 15:04:05 -0700"
)
