package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// 定义一个初始化数据库的函数
func initDB(dsn string) (err error) {
	// DSN:Data Source Name
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return
	}
	return nil
}

type student struct {
	stuNo   string
	stuName string
	stuSex  string
	stuAge  int
}

func queryOne(stuNo string) (err error) {
	var stu student
	err = db.QueryRow(`select stuNo, stuName, stuSex, stuAge from stu where stuNo=?;`, stuNo).Scan(&stu.stuNo, &stu.stuName, &stu.stuSex, &stu.stuAge)
	if err != nil {
		return
	}
	fmt.Printf("%#v\n", stu)
	return nil
}

func queryMany() {
	sqlStr := "select stuNo, stuName, stuSex, stuAge from stu where stuAge > ?"
	rows, err := db.Query(sqlStr, 20)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u student
		err := rows.Scan(&u.stuNo, &u.stuName, &u.stuSex, &u.stuAge)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("%#v\n", u)
	}
}

func insert() {
	sqlStr := "insert into stu(stuNo, stuName, stuSex, stuAge, stuSeat,stuAddress, ch, math) values (?,?,?,?,?,?,?,?)"
	ret, err := db.Exec(sqlStr, "s25322", "王冶", "男", 28, 7, "武当山", 90, 87)
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

func update() {
	sqlStr := `update stu set stuNo=? where stuNo=?;`
	ret, err := db.Exec(sqlStr, "s25322", "s25323")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 受影响的行数
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, rows %d.\n", n)
}

func deleteMySQL() {
	sqlStr := `delete from stu where stuNo=?;`
	ret, err := db.Exec(sqlStr, "s25322")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 受影响的行数
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, rows %d.\n", n)
}

func main() {
	dsn := "root:dyp1996@tcp(127.0.0.1:3306)/test"
	var err error
	err = initDB(dsn)
	if err != nil {
		fmt.Println("创建连接失败")
		return
	}
	fmt.Println("数据库连接成功")
	//err = queryOne("s25304")
	queryMany()
	//insert()
	//update()
	deleteMySQL()
	queryMany()
}
