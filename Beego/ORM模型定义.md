# 模型定义
默认的表名规则，使用驼峰转蛇形：
```
AuthUser -> auth_user
Auth_User -> auth__user
DB_AuthUser -> d_b__auth_user
```
除了开头的大写字母以外，遇到大写会增加 `_`，原名称中的下划线保留。

## 自定义表名

```go

```


