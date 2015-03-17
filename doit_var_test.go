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

func TestVarDBUpdate(t *testing.T) {
	testDBFile := "_test_tmp/TestVarDBUpdate.db"
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

	v := &Var{ID: testVar.ID}
	db.First(&v)
	if v.Name != "" {
		v.Value = "Clock"
		db.Save(&v)
	} else {
		t.Fatal("Unable to find document")
	}

	db.Close()
}
