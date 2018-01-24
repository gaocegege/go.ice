package migration

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

var initialTask = NewTask(
	time.Date(2018, 1, 21, 23, 43, 01, 0, time.UTC),
	initTaskName, "create migration table to track future migration tasks",
	createMigrationTable, dropMigrationTable)

type TaskFunc func(tx *sql.Tx) error

type Task interface {
	CreateTime() time.Time
	Name() string
	Description() string
	Up(tx *sql.Tx) error
	Down(tx *sql.Tx) error
}

type BasicTask struct {
	createTime  time.Time
	name        string
	description string
	up          TaskFunc
	down        TaskFunc
}

func InitTask() Task {
	return initialTask
}

func NewTask(t time.Time, name string, desc string, up TaskFunc, down TaskFunc) Task {
	return &BasicTask{
		createTime:  t,
		name:        name,
		description: desc,
		up:          up,
		down:        down,
	}
}

func (t *BasicTask) CreateTime() time.Time {
	return t.createTime
}

func (t *BasicTask) Name() string {
	return t.name
}

func (t *BasicTask) Description() string {
	return t.description
}

func (t *BasicTask) Up(tx *sql.Tx) error {
	return t.up(tx)
}

func (t *BasicTask) Down(tx *sql.Tx) error {
	return t.down(tx)
}

func createMigrationTable(tx *sql.Tx) error {
	// we need to use ` to quote the table name `_ice_migration`, which is why we concat string instead of using literal
	c := "CREATE TABLE " + migrationTableNameQuoted + " (" +
		` name VARCHAR(255),
		  description TEXT,
		  create_time INT,
		  apply_time INT
		)`
	if _, err := tx.Exec(c); err != nil {
		return errors.Wrap(err, "can't create migration table")
	}
	return nil
}

func dropMigrationTable(tx *sql.Tx) error {
	d := "DROP TABLE " + migrationTableNameQuoted
	if _, err := tx.Exec(d); err != nil {
		return errors.Wrap(err, "can't drop migration table")
	}
	return nil
}