package doittypes

import (
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func TestDomainDBBasic(t *testing.T) {
	testDBFile := "_test_tmp/TestDomainDBBasic.db"
	//Remove testing file
	os.Remove(testDBFile)
	db, err := gorm.Open("sqlite3", testDBFile)
	if err != nil {
		t.Fatal(err)
	}
	db.DB()
	db.CreateTable(&Domain{})

	testDomain := &Domain{Name: "Potato"}
	db.NewRecord(testDomain)
	db.Create(&testDomain)
	db.Close()
}
