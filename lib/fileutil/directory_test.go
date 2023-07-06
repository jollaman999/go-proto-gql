package fileutil

import (
	"testing"
)

func Test_CreateDirIfNotExist(t *testing.T) {
	err := CreateDirIfNotExist("/tmp/go-proto-gql_test")
	if err != nil {
		t.Fatal("Failed to create dir!")
	}

	err = CreateDirIfNotExist("/proc/go-proto-gql_test")
	if err != nil {
		t.Log("Tried to create dir in /proc folder")
	}
}

func Test_DeleteDir(t *testing.T) {
	err := DeleteDir("/tmp/go-proto-gql_test")
	if err != nil {
		t.Fatal("Failed to delete dir!")
	}

	err = DeleteDir("/proc/cpuinfo")
	if err != nil {
		t.Log("Tried to delete file in /proc folder")
	}
}
