# 零 连接服务器
在命令行窗口输入命令启动客户端连接数据库服务器
需要指定下面的参数

1. `host`：主机          `-h`
2. `username`：用户名	    `-u`
3. `password`：密码		`-p`
4. `port`：端口			`-P`(大写)

示例
```shell
msyql -h127.0.0.1 -P3306 -uroot -p
```
回车输入密码登录服务端。

# 一 库操作语句
## 1.1 查看存在的库
```mysql
show  databases;
```
```
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
4 rows in set (0.00 sec)
```
安装`MySQL`后，`MySQL`自带了`4`个数据库
1. `information_schema`：存储了`MySQL`服务器管理数据库的信息。
2. `performance_schema`：`MySQL5.5`新增的库，用来保存数据库服务器性能的参数
3. `mysql`：`MySQL`系统数据库，保存的登录用户名，密码，以及每个用户的权限等等
4. `test`：给用户学习和测试的数据库。

## 1.2 创建数据库
```mysql
create database if not exists  `databaseName`  charset=utf8;
```
只有当`databaseName`数据库不存在则创建数据库
* `if not exists`: 可以不写的。
* \` `databaseName` \`: 当`databaseName`为关键字时，可以加上\`\`(反引号)
* `charset=utf8`: 指定默认字符编码集，如果不指定，默认为服务端默认指定的编码

## 1.3 删除数据库
```mysql
drop database if exists `databaseName`;
```
只有当`databaseName`数据库存在才会执行删除
* `if exists`: 可以不写的。

## 1.4 显示创建数据库的sql语句
```mysql
show create database databaseName;
```

## 1.5 更新数据库
修改数据库的字符编码
```mysql
alter database databaseName charset=utf8;
```

## 1.6 选择要使用的数据库
```mysql
use databaseName;
```

# 二 表操作语句
一定要先创建库，再创建表
## 2.1 查看当前数据库

```mysql
select database();
```

## 2.2 创建表
```mysql
create table if not exists 表名(
    字段名 数据类型 约束条件,
    字段名 数据类型 约束条件
) engine=存储引擎;
```
* `engine=存储引擎`:指定表的存储引擎

**MySQL主要的存储引擎**
* `InnoDB`: `MySQL5.5`之后的默认存储引擎； 支持事务、行级锁和外键(数据存储安全)

  > 对应文件，表结构(`.frm`)和表数据(`.ibd`)。
* `MyIsam`: `MySQL5.5`之前的默认存储引擎； 数据存取速度块

  > 对应文件，表结构(`.frm`)、表数据(`.MYD`)、表索引(`.MYI`)(基于目录查找数据)
* `Memory`: 基于散列， **存储在内存中** ，对临时表很有用。 **断电数据丢失**

  > 对应文件，只有表结构，**数据保存在内存**
* `BlackHole`: 无论存什么，都会立刻消失

  > 对应文件，只有表结构，**数据不会保存**
  

**约束条件**

|约束条件|说明|
|:---:|:---:|
|`null`|空|
|`not null`|非空|
|`default 默认值`|默认值|
|`auto_increment`|自动增长|
|`primary key`|主键|
|`comment`|字段备注|

## 2.3 查看表的创建SQL语句
```mysql
show create table `tableName`;
```

## 2.4 查看表结构
```mysql
describe `tableName`;
desc `tableName`;
```

## 2.5 删除表
```mysql
drop table if exists `tableName`;
```
* `tableName`: 可以指定多个

## 2.6 修改表
```mysql
alter table `tableName`;
```

### 2.6.1 添加字段
```mysql
alter table `tableName` add column 字段名 数据类型 约束条件 [位置];
```
* `column`: 列，可以不写
* `位置`: 默认添加在表最后一列
    * `first`: 表的第一列
    * `after 字段1`: 在`字段1`之后


### 2.6.2 删除列
```mysql
alter table `tableName` drop column 字段名;
```
* `column`: 列，可以不写

### 2.6.3 修改字段名
数据类型可以同时修改
```mysql
alter table `tableName` change column 原字段名 新字段名 数据类型 约束条件;
```
* `column`: 列，可以不写

### 2.6.4 修改字段类型
不修改字段名
```mysql
alter table `tableName` modify 字段名 新的数据类型;
```

### 2.6.5 修改存储引擎
```mysql
alter table `tableName` engine=引擎名;
```

### 2.6.6 修改表名
```mysql
alter table `tableName` rename to `newTableName`;
```

# 三 数据操作语句
一定要先创建库，再创建表，最后在操作数据

```mysql
create table stu
(
  id    int auto_increment primary key comment '主键',
  name  varchar(20) not null,
  addr  varchar(50) default '地址不详',
  score int comment '成绩'
);
```
## 3.1 添加数据

**指定字段添加一条**
```mysql
insert into `tableName` (字段名, 字段名,…) value (值1, 值1,…)
```
**示例**
```mysql
insert into stu (id,name,addr,score) values (1,'tom','上海',88);
```
**指定字段添加多条**
```mysql
INSERT INTO `tableName`(字段名, 字段名,…)  VALUES (值1, 值1,…),  (值1, 值1,…), ...;
```
**示例**
```mysql
insert into stu (id,name,addr,score) values (2,'jack','上海',90),(3,'amy','北京',77);
```
**全字段添加**
```mysql
INSERT INTO `tableName` VALUES (值1, 值1,…),  (值1, 值1,…), ...; -- 可以插入多行数据
```
**示例**
```mysql
insert into stu values (4,'小小','上海',90),(5,'兰兰','天津',77);
```
* 字段省略后，插入数据必须与数据库字段顺序一致。
* 对于自增长数据，插入`null`即可

## 3.2 更新数据
**全表更新**
```mysql
-- 全文档修改
update `tableName` set field_name_1=field_value_1,field_name_2=field_value_2; -- 可以修改多个
```
**满足条件更新**
```mysql
update `tableName` set field_name=field_value where conditions;
```
**示例**
```mysql
update stu set score=70 where name='jack';
```

**where条件说明**
1. 支持 `and or not` 逻辑
2. `=`: 等于

## 3.3 查询数据
**全字段查询**
```mysql
-- 全字段查询
SELECT * FROM `tableName`;  -- 从tb_name表中查询出全部字段的数据
```
**指定字段查询**
```mysql
-- 指定字段查询
SELECT field_name_1, field_name_2 FROM `tableName`;
```

## 3.4 删除数据
**全表删除**
```mysql
delete from `tableName`;  -- 删除tb_name表中的所有数据
```
**满赠条件删除**
```mysql
delete from `tableName` where conditions;  -- 满足条件删除数据
```
* `conditions`: 条件

## 3.5 清空表
```mysql
truncate table `tableName`;
```
>脚下留心：delete from 表和truncate table 表区别？
> * `delete from 表名`：遍历表记录，一条一条的删除
> * `truncate table 表名`：将原表销毁，再创建一个同结构的新表。
> 就清空表而言，这种方法效率高。

## 3.6 复制表
```mysql
create table 新表 select 字段 from 旧表;
```
* 特点: **不能复制父表的主键，能够复制父表的数据**
```mysql
create table 新表 like 旧表;
```
* 特点: **只能复制表结构，不能复制表数据**



# 四 SQL语句分类
`DDL（data definition language）`数据库定义语言`CREATE、ALTER、DROP、SHOW`

`DML（data manipulation language）`数据操纵语言`SELECT、UPDATE、INSERT、DELETE`

`DCL（Data Control Language）`数据库控制语言,是用来设置或更改数据库用户或角色权限的语句




