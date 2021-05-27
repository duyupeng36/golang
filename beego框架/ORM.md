# ORM介绍

`beego ORM` 是一个强大的 `Go` 语言 `ORM` 框架。她的灵感主要来自 `Django ORM` 和 `SQLAlchemy`。

目前该框架仍处于开发阶段，可能发生任何导致不兼容的改动。

**已支持数据库驱动：**

- `MySQL`：[github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- `PostgreSQL`：[github.com/lib/pq](https://github.com/lib/pq)
- `Sqlite3`：[github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)

以上数据库驱动均通过基本测试，但我们仍需要您的反馈。

**`ORM` 特性：**

- 支持 Go 的所有类型存储
- 轻松上手，采用简单的 CRUD 风格
- 自动 Join 关联表
- 跨数据库兼容查询
- 允许直接使用 `SQL` 查询／映射
- 严格完整的测试保证 `ORM `的稳定与健壮

更多特性请在文档中自行品读。



`Beego`中内嵌了`ORM`框架，用来操作数据库。那么`ORM`框架是什么呢？`ORM`框架是`Object-RelationShip Mapping`的缩写，中文叫**关系对象映射**，他们之间的关系，我们用图来表示：

![](./.img/orm.png)

**作用**

1. 通过对象操作数据表
2. 通过对象生成与对象相同属性的数据表



## 简单使用

```go
package main

import (
    "fmt"
    "github.com/beego/beego/v2/client/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
    Id   int
    Name string
}

func init() {
    // set default database
    orm.RegisterDataBase("default", "mysql", "username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8&loc=Local")  // 连接数据库，给定一个别名

    // register model
    orm.RegisterModel(new(User))  // 注册表到orm

    // create table
    orm.RunSyncdb("default", false, true)  // 创建表
}

func main() {
    o := orm.NewOrm()  // 获取orm对象

    user := User{Name: "slene"}  // 获取数据对象

    // insert
    id, err := o.Insert(&user)  // 插入数据
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    // update
    user.Name = "astaxie"  // 修改对象
    num, err := o.Update(&user)  // 修改表
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)

    // read one
    u := User{Id: user.Id}  // 生成一个对象包含对象的id
    err = o.Read(&u)  // 查询出一个对象
    fmt.Printf("ERR: %v\n", err)

    // delete
    num, err = o.Delete(&u)  // 删除数据库中数据
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
```

* `orm.RegisterDataBase("default", "mysql", "username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8&loc=Local") ` 连接数据库

  * 第一个参数`"default"`: 连接数据库的别名，默认必须要有一个`default`数据库
  * 第二个参数`"mysql"`: 要使用的驱动名
  * 第三个参数，就是`dsn`数据资源地址

* ` orm.RegisterModel(new(User))`: 注册表给数据库。传递的参数是结构体的指针

* `orm.RunSyncdb("default", false, true)` 创建表

  * 第一个参数`"default"`: 之前指定的数据库别名
  * 第二个参数`false`: 指定是否强制创建表，通常指定为`false`

  * 第三个参数`true`:  指定是否显示创建过程。通常指定为`true`

# orm正常使用

**连接数据库**
```go
package models

import (
   "fmt"
   "github.com/beego/beego/v2/client/orm"
   _ "github.com/go-sql-driver/mysql"
   "os"
)

// 结构体，用于与数据表进行映射
type user struct {
   Id int
   Name string
   Age int
}

func init() {
   // 注册一个默认数据库，必须且有且只有一个default数据库
   err := orm.RegisterDataBase("default", "mysql", "root:dyp1996@tcp(127.0.0.1:3306)/beego?charset=utf8")
   if err != nil {
      fmt.Println("连接数据库失败，失败原因为: ", err)
      os.Exit(1)
   }
   // 注册表到数据库
   orm.RegisterModel(new(user))
   
    // 创建表，如果表存在则跳过，否则创建
   err = orm.RunSyncdb("default", false, true)
   if err != nil {
      fmt.Println("表更新失败，失败原因为: ", err)
      os.Exit(1)
   }
}
```
**需要在`main.go`中引入**
```go
package main

import (
   _ "beego/models"  // 引入models包，确保它的init方法执行以得到orm
   _ "beego/routers"
   beego "github.com/beego/beego/v2/server/web"
)

func main() {
   beego.Run()
}
```


## 简单的增删查改
```go
package controllers

import (
   "beego/models"
   "fmt"
   "github.com/beego/beego/v2/client/orm"
   beego "github.com/beego/beego/v2/server/web"
   "strconv"
)

type ORMController struct {
   beego.Controller
}

// OrmInsert 插入数据
func (c *ORMController) OrmInsert()  {
   o := orm.NewOrm()

   // 生成一个对象
   u := models.User{
      Name: "dyy",
      Age:  13,
   }
   // 写入到数据库
   id, err:= o.Insert(&u)
   if err != nil {
      fmt.Println("插入失败，失败原因: ", err)
      c.Data["ERR"] = "插入数据失败"
      c.TplName = "err.html"
      return
   }
   c.Data["MSG"] = "插入数据成功，成功id为" + strconv.FormatInt(id, 10)
   c.TplName = "test.html"
}

func (c *ORMController) OrmDelete() {
   o := orm.NewOrm()
   // 构造对象查询条件
   u := models.User{
      Id: 1,
   }
   // 查询
   err := o.Read(&u)
   if err != nil {
      fmt.Println("查询失败，失败原因: ", err)
      c.Data["ERR"] = "没有id为1的记录"
      c.TplName = "err.html"
      return
   }
   // 删除
   i, err := o.Delete(&u)
   if err != nil {
      fmt.Println("删除失败，失败原因: ", err)
      c.Data["ERR"] = "删除数据失败"
      c.TplName = "err.html"
      return
   }
   c.Data["MSG"] = "删除数据成功，成功id为" + strconv.FormatInt(int64(u.Id), 10) + "影响条数"+strconv.FormatInt(i,10)
   c.TplName = "test.html"
}

func (c *ORMController) OrmUpdate() {
   o := orm.NewOrm()

   // 构造查询条件
   u := models.User{Id: 1}
   err := o.Read(&u)
   if err != nil {
      fmt.Println("查询失败，失败原因: ", err)
      c.Data["ERR"] = "没有id为1的记录"
      c.TplName = "err.html"
      return
   }

   // 修改
   u.Name = "小不点"
   var i int64
   i, err = o.Update(&u)
   if err != nil {
      fmt.Println("修改数据，失败原因: ", err)
      c.Data["ERR"] = "修改数据失败"
      c.TplName = "err.html"
      return
   }
   c.Data["MSG"] = "修改数据成功，成功条数为" + strconv.FormatInt(i, 10)
   c.TplName = "test.html"
}

func (c *ORMController) OrmSelect() {
   o := orm.NewOrm()
   // 构造查询条件
   u := models.User{Id: 1}
   // 查询
   err := o.Read(&u)
   if err != nil {
      fmt.Println("查询失败，失败原因: ", err)
      c.Data["ERR"] = "没有id为1的记录"
      c.TplName = "err.html"
      return
   }
   c.Data["MSG"] = fmt.Sprintf("查询到id为%d的记录为%#v\n", 1, u)
   c.TplName = "test.html"
}
```
* `o := orm.NewOrm()`: 获得默认数据库
* `id, err:= o.Insert(&u)`: 添加数据
  * 返回插入数据的`id`和错误
* `i, err = o.Update(&u)`: 修改数据
  * 返回影响的行数和错误
* `err := o.Read(&u)`: 查询数据
* `i, err := o.Delete(&u)`: 删除数据
  * 返回影响的行数和错误

## 约束条件指定

`orm`指定每个字段的约束条件是通过结构体标签(`tag`)进行指定的。其格式如下

`orm:"限制条件"`，常用约束条件如下表

|         限制条件         |                             作用                             |
| :----------------------: | :----------------------------------------------------------: |
|           `pk`           |                       设置该字段为主键                       |
|          `auto`          |           这只该字段自增，但是要求该字段必须为整型           |
|       `default(0)`       |     设置该字段的默认值，需要注意字段类型和默认值类型一致     |
|       `size(100)`        |      设置该字段长度为100个字节，一般用来设置字符串类型       |
|          `null`          |              设置该字段允许为空，默认不允许为空              |
|         `unique`         |                      设置该字段全局唯一                      |
| `digits(12);decimals(4)` | 设置浮点数位数和精度。比如这个是说，浮点数总共12位，小数位为四位。 |
|        `auto_now`        |          针对时间类型字段，作用是保存数据的更新时间          |
|      `auto_now_add`      |          针对时间类型字段,作用是保存数据的添加时间           |

> 注意：**当模型定义里没有主键时，符合`int, int32, int64, uint, uint32, uint64` 类型且名称为`Id` 的 `Field` 将被视为主键，能够自增.** 
>
> `Mysq`l中时间类型有`date`和`datetime`两种类型，但是我们`go`里面只有`time.Time`一种类型，如果项目里面要求精确的话，就需要指定类型，指定类型用的是`type(date)`或者`type(datetime)`

