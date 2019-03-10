# shoot

**echo 实现的 go 语言版本的 restful API 的 app**

> export GOPROXY=https://goproxy.io

## 使用环境变量改变配置

```bash
export JERRY_ADDR=:5000 && ./shoot
```

## 统一数据返回格式

如果请求成功会走200字段，前端会在resolve中拿到数据

如果失败走其它字段，前端会在reject中拿到异常信息，因此没必要通过其它字段来判断是否正确

**异常返回信息**

```json
{
  "err_code": 0,
  "message": "ok!",
  "url": ""
}
```

| 字段名   |                       作用                       |   类型 |
| -------- | :----------------------------------------------: | -----: |
| err_code | 判断当前的请求是否异常，0 表示成功，其它均为异常         |    int |
| message  |        返回的信息，成功和异常均会返回信息              | string |
| url      |       如果出现异常，则返回请求的 url 路径             | string |

**正常返回的信息**

```json
{
  "field": "",
  "some": {}
}
```

## TODO LIST

- [x] 全局异常处理
- [x] 参数检验
- [x] JWT 支持
- [x] ORM 框架集成
- [x] 配置文件驱动
- [x] cors 跨域
- [x] 测试驱动
- [x] 统一数据返回格式
- [ ] 优化，美观

