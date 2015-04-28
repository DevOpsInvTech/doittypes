package doittypes

import (
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func TestHostDBBasic(t *testing.T) {
	testDBFile := "_test_tmp/TestHostDBBasic.db"
	//Remove testing file
	os.Remove(testDBFile)
	db, err := gorm.Open("sqlite3", testDBFile)
	if err != nil {
		t.Fatal(err)
	}
	db.DB()
	db.CreateTable(&Host{})

	testHost := &Host{Name: "Potato"}
	db.NewRecord(testHost)
	db.Create(&testHost)
	db.Close()
}

func TestHostVarDBBasic(t *testing.T) {
	testDBFile := "_test_tmp/TestHostVarDBBasic.db"
	//Remove testing file
	os.Remove(testDBFile)
	db, err := gorm.Open("sqlite3", testDBFile)
	if err != nil {
		t.Fatal(err)
	}
	db.DB()
	db.CreateTable(&Host{})
	db.CreateTable(&Var{})

	testVar := &Var{Name: "Potato", Value: "Spud"}
	db.NewRecord(testVar)
	db.Create(&testVar)

	testHost := &Host{Name: "Potato"}
	db.NewRecord(testHost)
	db.Create(&testHost)
	db.Close()
}
