package chainerr

import (
	"fmt"
	"github.com/cockroachdb/errors"
)

var ForeignKey = errors.New("foreign-key")

// NewForeignKey is used for initiating an error
func NewForeignKey(err error, domain string) error {
	// err = fmt.Errorf("original: %v; %w", err.Error(), ForeignKey)
	// return fmt.Errorf("domain: %v; %w", domain, err)
	err = errors.Wrapf(ForeignKey, "original: %v;", err.Error())
	return err
}

func AddCode(err error, code string) error {
	return fmt.Errorf("code: %v; %w", code, err)
}

func AddPath(err error, path string) error {
	return fmt.Errorf("path: %v; %w", path, err)
}

func AddMessage(err error, message string, params ...interface{}) error {
	return fmt.Errorf("mesage: %v; message_params: %v; %w", message, params, err)
}
