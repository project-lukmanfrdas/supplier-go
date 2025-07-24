package config

import (
	"io"
	"os"
)

func SaveFile(src io.Reader, filepath string) error {
	dst, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}
