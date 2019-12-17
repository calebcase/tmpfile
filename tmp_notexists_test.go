// +build !windows

package tmp

import (
	"errors"
	"os"
)

func notExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return errors.New("File exists.")
	} else if !os.IsNotExist(err) {
		return err
	}

	return nil
}
