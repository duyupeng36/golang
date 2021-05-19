# 一、MySQL简介
**数据库（Database）是按照数据结构来组织、存储和管理数据的仓库**
MySQL就是一款基于网络通信的数据库管理软件。
> 任何基于网络通信的软件底层都是使用socket进行开发的
> 1. 服务端
     >     * 基于网络通信
>     * 收发消息
> 2. 客户端
     >     * 基于网络通信
>     * 收发消息

**支持使用不同的语言编写的MySQL客户端**：存在服务端与客户端语言的兼容性问题
> **解决方案**
> 1. 服务端支持全部语言(不合理的)
> 2. 采用同一的语言(使用服务端自己确定的数据交互语言)

**为解决SQL数据库与不同的编程语言的交流，确定了结构性查询语言(SQL)**

## 1.1 相关的重要概念
* 库: 类似于文件夹
* 表: 类似于文件夹中的文件
* 记录: 文件中的数据
* 表头: 表格的第一行
* 字段: 每一列对应的名字

# 二 安装

## 2.1 Windows

1. 下载压缩包: https://dev.mysql.com/downloads/mysql/
2. 解压到`c`盘根目录，并重名为`mysql`
3. 将`mysql/bin`添加到`Path`环境变量
4. 打开`cmd`以管理员方式运行
5. 初始化数据库: `mysqld --initialize --console`，需要记录默认生成的密码
6. 注册为Windows系统服务: `mysqld -install`
8. 启动服务: `net start mysql`或`sc start mysql`
9. 登录数据库: `mysql -u root -p`
10. 修改密码: `alter user 'root'@'localhost' identified by '想要设置的密码';`
11. 退出
12. 新增`mysql/my.ini`配置文件,键入如下内容
     ```
     [mysqld] 
     # 设置mysql的安装目录，也就是刚才我们解压的目录
     basedir=C:/mysql
     # 设置mysql数据库的数据的存放目录
     datadir=C:/mysql/data
     # 设置默认使用的端口
     port=3306
     # 允许最大连接数
     max_connections=200
     # 允许连接失败的次数。这是为了防止有人试图攻击数据库
     max_connect_errors=10
     # 服务端使用的字符集
     character-set-server=utf8mb4
     # 数据库字符集对应一些排序等规则使用的字符集
     collation-server=utf8mb4_general_ci
     # 创建新表时将使用的默认存储引擎
     default-storage-engine=INNODB
     # 默认使用“mysql_native_password”插件作为认证加密方式
     # MySQL8.0默认认证加密方式为caching_sha2_password
     # default_authentication_plugin=mysql_native_password
      
     [mysql]
     # 设置mysql客户端默认字符集
     default-character-set=utf8mb4
      
     [client]
     # 其他连接方式
     default-character-set=utf8mb4
     ```
13. 重启服务室配置文件生效
     * 关闭服务: `net stop mysql`或`sc stop mysql`
     * 启动服务: `net start mysql`或`sc start mysql`

**删除服务: `sc delete mysql`或`mysqld -remove`（需先停止服务）**

## 2.2 Linux

### 2.2.1 安装编译好的二进制压缩包
1. 下载编译好的二进制包: `wget https://cdn.mysql.com//Downloads/MySQL-5.7/mysql-5.7.34-linux-glibc2.12-x86_64.tar.gz`

<br>

2. 创建一个mysql目录: `mkdir /usr/local/myql`

<br>

3. 将二进制文件包解压到该目录: `tar -zvxf mysql-5.7.34-linux-glibc2.12-x86_64.tar.gz -C /usr/local/mysql`

<br>

4. 创建用户组: `groupadd mysql`

<br>

5. 创建一个不可登录用户并添加目录: `useradd -r -s /sbin/nologin -g mysql mysql -d /usr/local/mysql/`

<br>

6. 修改目录属主与属组: `chown -R mysql:mysql /usr/local/mysql`

<br>

7. 修改配置文件: `vim /etc/my.cnf`
    ```
    [mysqld]
          
    bind-address=0.0.0.0
    port=3306 
    user=mysql
    basedir=/usr/local/mysql 
    datadir=/usr/local/mysql/data
    socket=/tmp/mysql.sock
    
    log-error=/usr/local/mysql/data/mysql.err
    pid-file=/usr/local/mysql/data/mysql.pid
    #character config
    character_set_server=utf8mb4 
    
    symbolic-links=0
    
    explicit_defaults_for_timestamp=true
    ```
   
<br>

8. 初始化: `/usr/local/mysql/bin/mysqld  --defaults-file=/etc/my.cnf --user=mysql --basedir=/usr/local/mysql --datadir=/usr/local/mysql/data  --initialize`
   记录初始化的秘密
     * 出现错误执行: `yum install -y libaio`

<br>

9. 开启`ssl`: `/usr/local/mysql/bin/mysql_ssl_rsa_setup --datadir=/usr/local/mysql/data`

<br>

10. 修改启动配置文件: 
     * `cp /usr/local/mysql/support-files/mysql.server /etc/init.d/mysql`
     * `vim /etc/init.d/mysql`
      ```
      basedir=/usr/local/mysql
      datadir=/usr/local/mysql/data
      mysqld_pid_file_path=/usr/local/mysql/data
      ```

<br>

11. 建立软连接: `ln -s /usr/local/mysql/bin/mysql /usr/bin`

<br>

12. 登录mysql-修改密码
    * `mysql -uroot -p`
    * `set password = password('123456');` 设置密码
    
<br>

13. 设置开机启动
    * `chmod 775 /etc/init.d/mysql`

    * `chkconfig --add mysql`

    * `chkconfig --level 345 mysql on`
    
### 2.2.2 命令安装

#### 下载
基于`centos7.5`

```SHELL
wget https://mirrors.sohu.com/mysql/MySQL-5.7/mysql-5.7.27-1.el7.x86_64.rpm-bundle.tar
```

#### 解压

```SHELL
tar xf mysql-5.7.27-1.el7.x86_64.rpm-bundle.tar
```
#### 安装

```shell
yum install -y *.rpm
```

默认安装位置：`/var/lib/mysql`

如果开启了selinux则会出现如下错误
```shell
2019-08-30T11:18:22.976635Z 0 [Warning] Can't create test file /mydata/mysql/localhost.lower-test
2019-08-30T11:18:22.976687Z 0 [Note] /usr/sbin/mysqld (mysqld 5.7.27) starting as process 2788 ...
2019-08-30T11:18:22.980289Z 0 [Warning] Can't create test file /mydata/mysql/localhost.lower-test
2019-08-30T11:18:22.980338Z 0 [Warning] Can't create test file /mydata/mysql/localhost.lower-test
```
只需要关闭selinux
```shell
setenforce 0
```

#### 重置密码
```
# 查看默认密码
grep 'pass' /var/log/mysqld.log

# 重置密码
mysql_secure_installation
输入root密码
是否要修改密码
是否要修改root密码 大小写 数字 特殊字符 
是否要删除匿名用户
是否禁止root远程登录
是否要删除test数据库
是否要刷新表的权限
```
#### 密码校验规则

```mysql
set global validate_password_policy=0;
-- 0 校验级别最低，只校验密码的长度，长度可以设定
-- 1 必须包括大写字母、小写字母、数字、特殊字符
-- 2 必须满足上面两条，并追加，对于密码中任意连续的4个（或者4个以上） 字符不能是字典中的单词
set global validate_password_length=3; -- 修改密码的最短长度
```

#### 创建用户
```mysql
create user 'username'@'ip' identified by 'password';
-- 全部ip的话则是%
```

#### 查看权限和授权用户
```mysql
show grants;  -- 查看权限

grant all on *.* to 'username'@'ip' identified by 'password';  -- 给权限给用户

flush privileges;  -- 刷新权限表
```


