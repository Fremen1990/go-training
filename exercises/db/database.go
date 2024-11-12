package db

import (
	"fmt"
	"os"
	"sync"
)

const stateFileSuffix = ".state"

type Record struct {
	Id     int64
	Offset int64
	Length int64
}

type Database struct {
	file        *os.File
	lock        sync.RWMutex
	records     []Record
	idGenerator IdGenerator
}

func (d *Database) Close() {
	err := d.file.Close()
	if err != nil {
		fmt.Println("Error closing database")
		return
	}
}

func (d *Database) saveState() {
	bytes, err := ToBytes(d.records)
	if err != nil {
		fmt.Println("Error converting records to bytes")
	}
	err = os.WriteFile(d.file.Name()+stateFileSuffix, bytes, 0644)
	if err != nil {
		fmt.Println("Error saving state file")
	}
}

func Db(filePath string) (*Database, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	var records []Record
	bytes, err := os.ReadFile(filePath + stateFileSuffix)
	if err != nil {
		records = make([]Record, 0)
	} else {
		err = FromBytes(bytes, &records)
		if err != nil {
			return nil, err
		}
	}

	var lastId int64
	if len(records) > 0 {
		lastId = records[len(records)-1].Id
	}
	generator := &Sequence{lastId}

	db := &Database{file, sync.RWMutex{}, records, generator}
	return db, nil
}

func Run() {
	db, err := Db("users.db")
	fmt.Println(err)
	defer db.Close()
	db.saveState()
}
