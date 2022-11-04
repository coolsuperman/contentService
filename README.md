
# 练习:文章管理服务搭建
## 设计方案
* 分为API服务与RPC服务（均已开发完毕）
  * API服务只负责处理http请求
    * 使用gin框架处理http请求
    * 通过gin中间件加入xss过滤模块
    * 把各请求通过grpc打到rpc服务上
  * rpc服务负责业务逻辑和数据处理
    * 数据存放在mysql中，对请求较多的接口如列表，详情，加了一层redis缓存，发生修改时删除缓存，等待懒加载重新刷上去（详细逻辑在函数注释里）
    * 通过gorm以及绑定参数的方式规避sql注入
    * 核心业务逻辑已完成单元测试
## 改进
  * 已将rpc服务的mysql/redis manager由懒汉改为饿汉模式，规避初始化时无法监测连接状态
  * rpc与api服务均引入wire进行了DI改造，避免了全局变量的使用
  * 通过cobra对rpc和api进行二级命令整合
  * 配置统一从yaml文件导入，不再部分写死在代码中
  * 清除了未使用的消息题
  * gin绑定时不再直接使用pb struct
## 项目结构
```
.
├── README.md
├── cmd
│   ├── api
│   │   ├── main.go //api服务入口
│   │   ├── wire.go
│   │   └── wire_gen.go
│   └── rpc
│       ├── main.go //rpc服务入口
│       ├── wire.go
│       └── wire_gen.go
├── config
│   └── config.yaml //配置文件
├── docs
│   └── schema
│       └── content.sql //mysql表设计
├── go.mod
├── go.sum
├── internal //业务逻辑
│   ├── api //api服务逻辑
│   │   ├── api.go //api结构定义
│   │   ├── conn
│   │   │   └── rpc_conn.go //rpc链接注入
│   │   ├── content //文章模块
│   │   │   └── content.go  //api服务http处理业务逻辑
│   │   └── provider_set.go //api依赖集合
│   ├── entity
│   │   ├── base_entity.go //基础结构体定义
│   │   ├── req_entity.go //请求结构体定义
│   │   └── resp_entity.go //返回结构体定义
│   └── rpc //rpc服务逻辑
│       ├── content //文章模块
│       │   ├── content.go //核心业务逻辑
│       │   └── content_test.go //核心业务逻辑单元测试（已测试完毕）
│       ├── datamanager //数据库管理模块
│       │   ├── content_operator.go // 文章表Mysql操作
│       │   ├── mysqlmanager.go //mysql数据管理文件
│       │   ├── provider_set.go //datamanager依赖集合
│       │   └── redismanager.go //redis数据管理文件
│       └── rpc.go //rpc结构定义
├── main.go //总入口，只有cobra的逻辑
└── pkg
    ├── config
    │   └── config.go //config解析加载服务
    ├── proto // proto与grpc文件
    │   └── content
    │       ├── content.pb.go
    │       ├── content.proto
    │       └── content_grpc.pb.go
    └── utils //工具
        └── cors_middleware.go //跨域中间件

19 directories, 31 files
```