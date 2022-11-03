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
## 项目结构
```
.
├── README.md
├── cmd
│   ├── api
│   │   └── main.go //api服务入口
│   └── rpc
│       └── main.go //rpc服务入口
├── config
│   └── config.yaml //配置文件
├── docs
│   └── schema
│       └── content.sql //mysql表设计
├── go.mod
├── go.sum
├── internal //业务逻辑
│   ├── api //api服务逻辑
│   │   └── content //文章模块
│   │       └── content.go //api服务http处理业务逻辑
│   ├── entity //结构体定义
│   │   ├── base_entity.go //基础结构体定义
│   │   ├── req_entity.go //请求结构体定义
│   │   └── resp_entity.go //返回结构体定义
│   └── rpc //rpc服务逻辑
│       ├── content //文章模块
│       │   ├── content.go //核心业务逻辑
│       │   └── content_test.go //核心业务逻辑单元测试（已测试完毕）
│       └── datamanager //数据库管理模块
│           ├── content_operator.go // 文章表Mysql操作
│           ├── mysqlmanager.go //mysql数据管理文件
│           └── redismanager.go //redis数据管理文件
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

18 directories, 21 files
```