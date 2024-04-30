package wc_common

import (
	"strings"

	"github.com/howeyc/gopass"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

func ValidatePassword(password string) {
	err := passwordvalidator.Validate(password, MinEntropyBits)
	CheckError(err, "Password does not meet the required security standards"+
		"\nIf you want to create keys for testing with weak/no password, use --insecure flag. Do NOT use those keys in production")
}

func ReadHiddenInput() string {
	password, err := gopass.GetPasswdMasked()
	CheckError(err, "Error reading input")
	return strings.TrimSpace(string(password))
}
