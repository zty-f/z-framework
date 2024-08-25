package zorm

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"testing"
	"zorm/session"
)

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite3", "gee.db")
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	return engine
}

func TestNewEngine(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
}

// 测试事务--使用Go原生标准库实现
func TestTransaction(t *testing.T) {
	db, _ := sql.Open("sqlite3", "gee.db")
	defer func() { _ = db.Close() }()
	_, _ = db.Exec("CREATE TABLE IF NOT EXISTS User(`Name` text);")

	tx, _ := db.Begin()
	_, err1 := tx.Exec("INSERT INTO User(`Name`) VALUES (?)", "Tom")
	_, err2 := tx.Exec("INSERT INTO User(`Name`) VALUES (?)", "Jack")
	if err1 != nil || err2 != nil {
		_ = tx.Rollback()
		log.Println("Rollback", err1, err2)
	} else {
		_ = tx.Commit()
		log.Println("Commit")
	}
}

type User struct {
	Name string `zorm:"PRIMARY KEY"`
	Age  int
}

// 测试事务--批量语句的使用形式
func TestEngine_Transaction(t *testing.T) {
	t.Run("rollback", transactionRollback)
	t.Run("commit", transactionCommit)
}

func transactionRollback(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		// 创造错误使其回滚
		return nil, errors.New("error")
	})
	if err == nil || s.HasTable() {
		t.Fatal("failed to rollback")
	}
}

func transactionCommit(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return
	})
	u := &User{}
	_ = s.First(u)
	if err != nil || u.Name != "Tom" {
		t.Fatal("failed to commit")
	}
}

// 常用形式使用事务
func TestEngine_Transaction2(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	s.Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, _ = s.Insert(&User{"Tom1", 18})
	testTransactionRollBack(t)
	user := make([]User, 0)
	_ = s.Find(&user)
	fmt.Printf("%v", user)
}
func testTransactionRollBack(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Begin()
	_, err := s.Insert(&User{"Tom1", 18})
	if err != nil {
		_ = s.Rollback()
		t.Error("failed to insert user", err)
		return
	}
	_, err = s.Insert(&User{"Tom3", 18})
	if err != nil {
		_ = s.Rollback()
		t.Error("failed to insert user", err)
		return
	}
	_ = s.Commit()
}
