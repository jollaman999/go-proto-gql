package fileutil

import "os"

// IsExist : Return true if file or directory is existed return false otherwise.
func IsExist(fileLocation string) bool {
	if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
		return false
	}

	return true
}

// WriteFile : Write input string to fileLocation
func WriteFile(fileLocation string, input string) error {
	file, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	_, err = file.WriteString(input)
	if err != nil {
		return err
	}

	return nil
}

// WriteFileAppend : Write input string to fileLocation with append mode
func WriteFileAppend(fileLocation string, input string) error {
	file, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	_, err = file.WriteString(input)
	if err != nil {
		return err
	}

	return nil
}

// DeleteFile : Delete a file from given fileLocation
func DeleteFile(fileLocation string) error {
	err := os.Remove(fileLocation)
	if err != nil {
		return err
	}

	return nil
}
