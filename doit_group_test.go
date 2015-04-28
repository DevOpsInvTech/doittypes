package doittypes

import (
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func TestGroupDBBasic(t *testing.T) {
	testDBFile := "_test_tmp/TestGroupDBBasic.db"
	//Remove testing file
	os.Remove(testDBFile)
	db, err := gorm.Open("sqlite3", testDBFile)
	if err != nil {
		t.Fatal(err)
	}
	db.DB()
	db.CreateTable(&Host{})
	db.CreateTable(&Var{})
	db.CreateTable(&Domain{})
	db.CreateTable(&Group{})

	testGroup := &Group{Name: "Potato"}
	db.NewRecord(testGroup)
	db.Create(&testGroup)
	db.Close()
}
