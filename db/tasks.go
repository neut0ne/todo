package db

import (
  "time"
  "github.com/boltdb/bolt"
)

// buckets need to be populated with BYTESLICES.
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

  // below fn is equal to writing:

  // return db.Update(func(tx *bolt.Tx) error{
  //   _, err := tx.CreateBucketIfNotExists(taskBucket)
  //   return err
  // })

  // which is more common, but the below gives a better visual when learning.
  fn := func(tx *bolt.Tx) error {
    _, err := tx.CreateBucketIfNotExists(taskBucket)
    return err
  }

  return db.Update(fn)
}
