package db

import (
  "time"
  "encoding/binary"
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
  return db.Update(func(tx *bolt.Tx) error{
    _, err := tx.CreateBucketIfNotExists(taskBucket)
    return err
  })
}

func CreateTask(task string) (int, error) {
  var id int
  err := db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket(taskBucket)
    id64, _ := b.NextSequence()
    id = int(id64)
    key := itob(int(id64))
    return b.Put(key, []byte(task))
  })
  if err != nil {
    return -1, err
  }
  return id, nil
  return 0, nil
}
// Key converter integers to byteslices
func itob(v int) []byte{
  b := make([]byte, 8)
  binary.BigEndian.PutUint64(b, uint64(v))
  return b
}
// Key converter byteslices to integers
func btoi(b []byte) int {
  return int(binary.BigEndian.Uint64(b))
  }
