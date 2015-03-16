package main

import (
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func TestVarDBBasic(t *testing.T) {
	testDBFile := "_test_tmp/TestVarDBBasic.db"
	//Remove testing file
	os.Remove(testDBFile)
	db, err := gorm.Open("sqlite3", testDBFile)
	if err != nil {
		t.Fatal(err)
	}
	db.DB()
	db.CreateTable(&Var{})

	testVar := &Var{Name: "Potato", Value: "Spud"}
	db.NewRecord(testVar)
	db.Create(&testVar)
	db.Close()
}
