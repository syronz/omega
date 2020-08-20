package types

import (
	"strconv"
	"strings"
	"time"
)

type Envkey string

// Envs holds all environments
type Envs map[Envkey]string

func (p Envs) ToBool(key Envkey) bool {
	return strings.ToUpper(p[key]) == "TRUE"
}

func (p Envs) ToUint64(key Envkey) uint64 {
	num, _ := strconv.ParseUint(p[key], 10, 64)
	return num
}

func (p Envs) ToByte(key Envkey) []byte {
	return []byte(p[key])
}

func (p Envs) ToDuration(key Envkey) time.Duration {
	num := p.ToUint64(key)
	return time.Duration(num)
}
