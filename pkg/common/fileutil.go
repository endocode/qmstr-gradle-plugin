package common

import (
	"os"
	"path/filepath"
)

func BuildCleanPath(base string, subpath string, abs bool) string {
	if filepath.IsAbs(subpath) {
		return filepath.Clean(subpath)
	}

	if abs && !filepath.IsAbs(base) {
		// ignore error and use non absolute path
		base, _ = filepath.Abs(base)
	}

	tmpPath := filepath.Join(base, subpath)
	return filepath.Clean(tmpPath)
}

// CheckExecutable checks the given file to be no directory and executable flagged
func CheckExecutable(file string) error {
	d, err := os.Stat(file)
	if err != nil {
		return err
	}
	if m := d.Mode(); !m.IsDir() && m&0111 != 0 {
		return nil
	}
	return os.ErrPermission
}

// exists checks if file exists and is not a directory
func exists(file string) bool {
	if d, err := os.Stat(file); err == nil {
		if d.IsDir() {
			return false
		}
		return true
	}
	return false
}