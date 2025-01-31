package fileutil

import "os"

// CreateDirIfNotExist : Make a directory if not exist
func CreateDirIfNotExist(dir string) error {
	if IsExist(dir) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteDir : Delete a directory
func DeleteDir(dir string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}

	return nil
}
