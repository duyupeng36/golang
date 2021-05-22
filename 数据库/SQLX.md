# sqlx
在项目中我们通常可能会使用`database/sql`连接`MySQL`数据库。
本文借助使用`sqlx`实现批量插入数据的例子，介绍了`sqlx`中可能被你忽视了的`sqlx.In`和
`DB.NamedExec`方法。

## sqlx介绍
在项目中我们通常可能会使用`database/sql`连接`MySQL`数据库。
`sqlx`可以认为是`Go`语言内置`database/sql`的超集，
它在优秀的内置`database/sql`基础上提供了一组扩展。
这些扩展中除了大家常用来查询的`Get(dest interface{}, ...) error`
和`Select(dest interface{}, ...) error`外还有很多其他强大的功能。


## 下载
```shell
go get github.com/jmoiron/sqlx
```

## 基本使用
`sqlx`的基本使用方式和`database/sql`相差不大，下面几个示例

### 连接数据库
```go
package main

import "github.com/jmoiron/sqlx"
import "fmt"

var db *sqlx.DB

func initDB() (err error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}
```
### 查询
**`Get`方法查询一条数据**
```go
package main

import "fmt"
import "github.com/jmoiron/sqlx"

var db *sqlx.DB
// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u struct{
		ID int
		Name string
		Age int
    }
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", u)
}
```
**`Select`函数用于查询多条数据**
```go
package main

import "fmt"
import "github.com/jmoiron/sqlx"

var db *sqlx.DB
// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var users  []struct{
		ID int
		Name string
		Age int
	}
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}
```
### 插入、更新和删除
`sqlx`中的`exec`方法与原生`sql`中的`exec`使用基本一致：
```go
package main

import "fmt"
import "github.com/jmoiron/sqlx"

var db *sqlx.DB

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "沙河小王子", 19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 6)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
```

### NamedExec
`DB.NamedExec`方法用来绑定`SQL`语句与结构体或`map`中的同名字段。
```go
package main
import "github.com/jmoiron/sqlx"

var db *sqlx.DB
func insertUserDemo()(err error){
	sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
	_, err = db.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "七米",
			"age": 28,
		})
	return
}
```

### NamedQuery
与`DB.NamedExec`同理，这里是支持查询。
```go
package main
import "github.com/jmoiron/sqlx"
import "fmt"

var db *sqlx.DB

func namedQuery(){
	sqlStr := "SELECT * FROM user WHERE name=:name"
	// 使用map做命名查询
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "七米"})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}

	u := user{
		Name: "七米",
	}
	// 使用结构体命名查询，根据结构体字段的 db tag进行映射
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}
```

### 事务操作
对于事务操作，我们可以使用sqlx中提供的db.Beginx()和tx.Exec()方法。
```go
package main
import "github.com/jmoiron/sqlx"
import "fmt"
import "errors"

var db *sqlx.DB

func transactionDemo2()(err error) {
	tx, err := db.Beginx() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()

	sqlStr1 := "Update user set age=20 where id=?"

	rs, err := tx.Exec(sqlStr1, 1)
	if err!= nil{
		return err
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	sqlStr2 := "Update user set age=50 where i=?"
	rs, err = tx.Exec(sqlStr2, 5)
	if err!=nil{
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	return err
}
```
