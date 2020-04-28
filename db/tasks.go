package db

import (
  "time"
  "github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
  Key   int
  Value string
}

func Init(dbPath string) error {
  var err error
  db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
  if err != nil {
    return err
  }
  fn := func(tx *bolt.Tx) error {
    _, err := tx.CreateBucketIfNotExists(taskBucket)
    return err
  }

  return db.Update(fn)
}
