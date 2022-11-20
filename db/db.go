package db

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func Write(task string) string {
	db, err := bolt.Open("todo.db", 0600, nil)
	var resp string

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return err
		}
		id, _ := b.NextSequence()

		err = b.Put([]byte(task), []byte(fmt.Sprint(id)))
		resp = fmt.Sprint(err)
		return err
	})
	return resp
}

func List() []string {
	db, err := bolt.Open("todo.db", 0600, nil)
	var tasks []string

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var byteTest []byte

	byteTest = []byte("anan")
	tasks = append(tasks, string(byteTest))

	err = db.View(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return err
		}

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
			tasks = append(tasks, fmt.Sprintf("key=%s, value=%s\n", k, v))
		}

		return nil
	})

	return tasks
}
