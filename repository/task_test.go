package repository

import (
	"context"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestAddNewTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open connection to DB: %v", err)
	}

	repo := NewRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `task_models` (`num_task`,`title`,`description`,`completed`) VALUES (?,?,?,?)")).
		WithArgs(
			15,
			"tarefa 15",
			"descricao tarefa 1",
			true,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Print the generated query

	mock.ExpectCommit()

	ok, err := repo.AddNewTask(context.Background(), 15, "tarefa 15", "descricao tarefa 1", true)
	if err != nil {
		t.Fatalf("error in creating task: %v", err)
	}

	if !ok {
		t.Fatalf("task not created")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestListTasks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open connection to DB: %v", err)
	}

	repo := NewRepository(gormDB)

	rows := sqlmock.NewRows([]string{"id", "num_task", "title", "description", "completed"}).
		AddRow(1, 1, "tarefa 1", "descricao tarefa 1", true).
		AddRow(2, 2, "tarefa 2", "descricao tarefa 2", false)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `task_models`")).WillReturnRows(rows)

	tasks, err := repo.ListTasks()
	assert.NoError(t, err)
	assert.Len(t, tasks, 2)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestDeleteTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open connection to DB: %v", err)
	}

	repo := NewRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `task_models` WHERE `task_models`.`id` = ?")).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	ok, err := repo.DeleteTask(context.Background(), 1)

	if err != nil {
		t.Fatalf("error in deleting task: %v", err)
	}

	if !ok {
		t.Fatalf("task not deleted")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

/*
func TestMarkTaskAsCompleted(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open connection to DB: %v", err)
	}

	repo := NewRepository(gormDB)

	{
		mock.ExpectBegin() // Set expectation for transaction begin
		mock.ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `task_models` WHERE num_task = ? ORDER BY `task_models`.`id` LIMIT 1")).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "num_task", "title", "description", "completed"}))

		mock.ExpectExec(
			regexp.QuoteMeta("UPDATE `task_models` SET `task_models`.`completed` = ? WHERE `task_models`.`num_task` = ?")).
			WithArgs(true, 1).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit() // Set expectation for transaction commit
	}

	ok, err := repo.MarkTaskAsCompleted(context.Background(), 1, true)
	if err != nil {
		t.Fatalf("error in marking task as completed: %v", err)
	}

	if !ok {
		t.Fatalf("task not marked as completed")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
*/
