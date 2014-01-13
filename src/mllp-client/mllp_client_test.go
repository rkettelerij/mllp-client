package main

import (
	"path"
	"path/filepath"
	"testing"
)

func TestSend(t *testing.T) {
	host := "localhost"
	port := 2575
	dir, _ := filepath.Abs(filepath.Dir("adt_a01.txt"))
	file := path.Join(dir, "adt_a01.txt")
	Send(&file, &host, &port)
}
