package db

import (
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
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
	lock        sync.Mutex
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

func (d *Database) Insert(input interface{}) (*Record, error) {
	bytes, err := ToBytes(input)
	if err != nil {
		return nil, err
	}
	d.lock.Lock()
	defer d.lock.Unlock()
	offset, err := d.file.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, err
	}
	length, err := d.file.WriteAt(bytes, offset)
	if err != nil {
		return nil, err
	}
	record := Record{d.idGenerator.next(), offset, int64(length)}
	d.records = append(d.records, record)
	d.saveState()
	return &record, nil
}

func (d *Database) findRecordIndex(id int64) int {
	return slices.IndexFunc(d.records, func(r Record) bool { return r.Id == id })
}

func (d *Database) FindById(id int64, output interface{}) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	idx := d.findRecordIndex(id)
	if idx != -1 {
		record := d.records[idx]
		buffer := make([]byte, record.Length)
		_, err := d.file.ReadAt(buffer, record.Offset)
		if err != nil {
			return err
		}
		err = FromBytes(buffer, output)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("record not found")
}

func (d *Database) DeleteById(id int64) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	idx := d.findRecordIndex(id)
	if idx != -1 {
		d.records = append(d.records[:idx], d.records[idx+1:]...)
		d.saveState()
		return nil
	}
	return errors.New("record not found")
}

func (d *Database) UpdateById(id int64, input interface{}) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	idx := d.findRecordIndex(id)
	if idx != -1 {
		bytes, err := ToBytes(input)
		if err != nil {
			return err
		}
		offset, err := d.file.Seek(0, io.SeekEnd)
		if err != nil {
			return err
		}
		length, err := d.file.WriteAt(bytes, offset)
		if err != nil {
			return err
		}
		d.records[idx].Offset = offset
		d.records[idx].Length = int64(length)
		d.saveState()
		return nil
	}
	return errors.New("record not found")
}

func getMaxId(records []Record) int64 {
	if len(records) == 0 {
		return 0
	}
	maxId := records[0].Id
	for _, record := range records {
		if record.Id > maxId {
			maxId = record.Id
		}
	}
	return maxId
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

	generator := &Sequence{getMaxId(records)}

	db := &Database{file, sync.Mutex{}, records, generator}
	return db, nil
}

type User struct {
	FirstName string
	LastName  string
	IsActive  bool
}

func Run() {
	db, _ := Db("users.db")
	defer db.Close()

	/*record, err := db.Insert(&User{"Jan", "Kowalski", true})
	if err != nil {
		fmt.Println("Error inserting user")
	}
	fmt.Println(record)*/

	_ = db.DeleteById(1)

	//_ = db.UpdateById(1, &User{"Jan", "Nowak", false})

	user := User{}
	if db.FindById(1, &user) == nil {
		fmt.Println("User: ", user)
	}
}
