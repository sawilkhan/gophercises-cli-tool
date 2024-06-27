package db

import (
	"encoding/binary"

	"go.etcd.io/bbolt"
)

var db *bbolt.DB

const taskBucket = "tasks"

type Task struct{
	ID int
	Name string
}

func InitDB() error{
	var err error
	db, err := bbolt.Open("tasks.db", 0600, nil)
	if err != nil {
        return err
    }

	return db.Update(func(tx *bbolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte(taskBucket))
        return err
    })
}

func AddTask(task Task) error{
	return db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(taskBucket))

		id, _ := b.NextSequence()
		task.ID = int(id)
		buf := make([]byte, binary.MaxVarintLen64)
        binary.PutVarint(buf, int64(task.ID))
		return b.Put(buf, []byte(task.Name))
	})
}

func ListTasks() ([]Task, error) {
    var tasks []Task
    err := db.View(func(tx *bbolt.Tx) error {
        b := tx.Bucket([]byte(taskBucket))
        return b.ForEach(func(k, v []byte) error {
            id, _ := binary.Varint(k)
            tasks = append(tasks, Task{
                ID:   int(id),
                Name: string(v),
            })
            return nil
        })
    })
    return tasks, err
}