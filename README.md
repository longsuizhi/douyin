# 项目名称

抖音极速版

# 项目构建

打开 /app/config 文件夹下的 config.toml 文件

# 项目文件目录说明

```shell
filetree
├─douyin
│   ├─api:对外提供本服务api接口sdk存放目录；也可以在当前目录实现对自己接口的mock
│   │   └─code:错误码、响应封装目录
│   ├─app:main.go存放目录，一个app代表一个application；与当前application相关依赖的配置文件
│   ├─conf:配置初始化、配置获取逻辑代码存放目录
│   ├─controller:请求参数校验，handler实现等逻辑代码存放目录
│   ├─dao:redis/mysql等存储组件代码逻辑存放目录
│   ├─logger:日志系统代码存放目录
│   ├─model:业务相关 struct、constants 定义代码存放目录
│   ├─middlewares:中间件代码存放目录
│   ├─router:路由注册
│   ├─service:业务主逻辑实现代码存放目录
│   ├─static:静态资源存储目录，存放视频、图片等
│   ├─utils:公共工具代码逻辑存放目录
│   └─douyin.sql:相关表设计
```

# 核心功能

视频推荐

# 项目技术

前端 后端（webSocket,channel/goroutine,gin,temlate,gorm,sql,nosql,mq）

# 系统架构

四层：前端，接入层，逻辑层，持久层

# 需求文档

# 相关表设计

# 功能设计文档
