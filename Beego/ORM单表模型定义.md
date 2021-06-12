# 模型定义
默认的表名规则，使用驼峰转蛇形：
```
AuthUser -> auth_user
Auth_User -> auth__user
DB_AuthUser -> d_b__auth_user
```
除了开头的大写字母以外，遇到大写会增加 `_`，原名称中的下划线保留。

## 自定义表名
实现接口`TableNameI`:
```go
package main
type User struct {
	Id   int    // 默认主键
	Name string `orm:"size(100)"`
}

func (u *User) TableName() string {
	return "main_user"
}
```
方法 `TableName`返回什么，创建的表明就是什么

## 自定义索引
实现接口`TableIndexI`，为单个或多个字段增加索引

```go
package main

type User struct {
	Id    int    // 默认主键
	Name  string `orm:"size(100)"`
	Email string
}
// TableIndex 切片中的字符串切片为联合索引
func (u *User) TableIndex() [][]string {
	return [][]string{
		[]string{"Id", "Name"},
	}
}

// TableUnique 切片中的字符串切片为联合唯一
func (u *User) TableUnique() [][]string {
	return [][]string{
		[]string{"Name", "Email"},
	}
}
```

## 设置参数（约束条件）
使用结构体`Tag`，用于设置字段的约束条件
```
`orm:"限制条件"`
```

### 忽略字段
设置 `-` 即可忽略 `struct` 中的字段
```go
package main

type User struct {
  // ...
    AnyField string `orm:"-"`
  //...
}
```

### 限制条件

|限制条件|描述|
|:---:|:---:|
|`auto`|自增长|
|`pk`|主键|
|`null`|允许为空|
|`index`|为单个字段增加索引|
|`unique`|为单个字段增加 唯一 键|
|`column`|为字段设置`db`字段的名称|
|`size`|设置 `size` 以后，`db type` 将使用 `varchar(size)`|
|`digits / decimals`|设置 `float32, float64` 类型的浮点精度|
|`auto_now / auto_now_add`|设置时间添加方式|
|`type`|设置为 `date` 时，`time.Time` 字段的对应 `db` 类型使用 `date/datetime`|
|`default`|默认值|
|`description`|字段说明描述|
|`precision`|为`datetime`字段设置精度值位数|

* `auto`: 当 `Field` 类型为 `int, int32, int64, uint, uint32, uint64` 时，
  可以设置字段为自增健
  
    * 当模型定义里没有主键时，符合上述类型且名称为 `Id` 的 `Field` 将被视为自增健。
    
    ```
    Id int `orm:"auto"`
    ```
  
* `pk`: 主键约束,设置为主键，适用于自定义其他类型为主键
  
    ```
    Id int `orm:"pk;auto"`
    ```

* `null`: 数据库表默认为 `NOT NULL`，设置 `null` 代表 `ALLOW NULL`
  
    ```
    Name string `orm:"null"
    ```

* `index`: 为单个字段增加索引

    ```
    Name string `orm:"index"`
    ```

* `unique`: 唯一键，为单个字段增加 唯一约束

    ```
    Name string `orm:"unique"`
    ```

* `column`: (数据库中的字段名称)为字段设置 `db` 字段的名称  
  
    ```
    Name string `orm:"column(user_name)"`
    ```

* `size`: 长度, `string` 类型字段默认为 `varchar(255)`

    设置 `size` 以后，`db type` 将使用 `varchar(size)`

    ```
    Title string `orm:"size(60)"`
    ```

* `digits/decimals`：设置 `float32`, `float64` 类型的浮点精度

    ```
    Money float64 `orm:"digits(12);decimals(4)"`
    ```
    * 总长度 `12` 小数点后 `4` 位 eg: `99999999.9999`

* `auto_now/auto_now_add`

    ```
    Created time.Time `orm:"auto_now_add;type(datetime)"`
    Updated time.Time `orm:"auto_now;type(datetime)"`
    ```
    * `auto_now`: 每次 `model` 保存时都会对时间自动更新
    * `auto_now_add`: 第一次保存时才设置时间

* `type`: go类型对应数据库中有多个类型时使用
  
  * 设置为 `date` 时，`time.Time` 字段的对应 `db` 类型使用 `date`
        ```
        Created time.Time `orm:"auto_now_add;type(date)"`
        ```
  * 设置为 `datetime` 时，`time.Time` 字段的对应 `db` 类型使用 `datetime`

    ```
    Created time.Time `orm:"auto_now_add;type(datetime)"`
    ```

* `default`: 设置默认值
    
    ```
    Age           int `default:"12"`
    AgeInOldStyle int `orm:"default(13);bee()"`
    ```
* `description`: 字段描述

    ```
    Status int `orm:"default(1);description(这是状态字段)"`
    ```

* `precision`: 为datetime字段设置精度值位数

    ```
    Created time.Time `orm:"type(datetime);precision(4)"`
    ```
