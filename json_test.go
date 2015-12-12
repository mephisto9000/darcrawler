package main

import (
	"os"
	"testing"
)

func TestExportImport(t *testing.T) {
	p1 := NewPageInfo("http://nur.kz", []string{"nur1", "nur2"})
	p2 := NewPageInfo("http://today.kz", []string{})
	p3 := NewPageInfo("http://tengrinews.kz", []string{"t1", ""})
	encodeInfos := []PageInfo{p1, p2, p3}

	filename := "tmp_.json"
	if err := exportToFile(encodeInfos, filename); err != nil {
		t.Error(err.Error())
	}

	// Do not forget to clean up after yourself
	defer os.Remove(filename)

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Error("File was not created")
	}

	// decodedInfos, err := importFromFile(filename)
	// if err != nil {
	// 	t.Error(err)
	// }
	// if len(encodeInfos) != len(decodedInfos) {
	// 	t.Error("Length of Encoded and Decoded values is not equal")
	// 	return
	// }

	// for i := 0; i < len(encodeInfos); i++ {
	// 	if encodeInfos[i].Title != decodedInfos[i].Title {
	// 		t.Errorf("Titles at at %v-th position are different", i)
	// 	}
	// 	// TODO compare keywords as well
	// }
}
