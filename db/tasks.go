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

// Initiates bolt.db, updates db with new or existing bucket
// Caution: Init is not init. The function is not package level.
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

// Creates task and stores it as a byteslice bucket in bolt.db
func CreateTask(task string) (int, error) {
  var id int
  err := db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket(taskBucket)
    id64, _ := b.NextSequence()
    id = int(id64)
    key := itob(int(id64))
    return b.Put(key, []byte(task))
  })
  // this is a way to check if there is anything else than nil or the expected
  // value. If err is not nil, err is returned.
  if err != nil {
    return -1, err
  }
  return id, nil
}

// Read all tasks in db.
func Alltasks() ([]Task, error) {
  var tasks []Task
  err := db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket(taskBucket)
    c := b.Cursor()
    for k, v := c.First(); k != nil; k, v = c.Next() {
      tasks = append(tasks, Task{
        Key:    btoi(k),
        Value:  string(v),
      })
    }
    // even if it's a read-only function, return *something*.
    return nil
  })
  if err != nil {
    return nil, err
  }
  return tasks, nil
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
