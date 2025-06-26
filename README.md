# 技术选型

- Gin：Web 框架
- Gorm：ORM 框架
- Go-Redis：Redis 框架
- Logrus：日志框架
- Wire：依赖注入管理

```shell
# Gin
go get -u github.com/gin-gonic/gin

# Go-Redis
go get -u github.com/redis/go-redis/v9

# Gorm
go get -u gorm.io/gorm
# 数据库驱动
go get -u gorm.io/driver/mysql
# Gorm-Gen
go get -u gorm.io/gen

# Logrus
go get -u github.com/sirupsen/logrus
# logrus-formatter
go get -u github.com/antonfisher/nested-logrus-formatter

# Wire-Cli
go install github.com/google/wire/cmd/wire@latest
# Wire
go get -u github.com/google/wire
```

# 项目结构

```txt
├── cmd                         # cmd
│   ├── gorm_gen.go             # gorm代码生成
│   ├── main.go                 # main.go
│   └── wire                    # wire
│       ├── wire.go             # 依赖注入管理
│       └── wire_gen.go         # wire生成
├── config                      # 配置目录
│   └── application.yml         # 应用配置文件
├── go.mod                      # go.mod
├── go.sum                      # go.sum
├── internal                    # 内部包
│   ├── config                  # 应用配置
│   │   ├── config.go           # 配置文件读取
│   │   └── time.go             # 自定义时间序列化
│   ├── dao                     # dao
│   │   ├── gen.go              # gorm生成
│   │   └── sys_user.gen.go     # gorm生成
│   ├── database                # 持久化层
│   │   ├── mysql.go            # MySQL
│   │   ├── provider_set.go     # Wire ProviderSets 分组管理
│   │   └── redis.go            # Redis
│   ├── dto                     # DTO
│   │   └── response.go         # 控制层统一响应结构
│   ├── engine                  # 管理GIN引擎及路由
│   │   ├── api_v1              # V1版本路由
│   │   │   └── routes.go       # 统一管理路由
│   │   ├── engine.go           # GIN引擎管理
│   │   └── middleware          # GIN中间件
│   │       └── error.go        # 全局异常处理中间件
│   ├── handler                 # 控制层处理器
│   │   ├── provider_set.go     # Wire ProviderSets 分组管理
│   │   └── sys_user.go         # sys_user 处理器
│   ├── model                   # 模型
│   │   └── sys_user.gen.go     # sys_user 模型
│   └── service                 # 业务层
│       ├── provider_set.go     # Wire ProviderSets 分组管理
│       └── sys_user.go         # sys_user 业务层
```

# 使用方式

## 管理路由

所有的路由均在`engine/api_v1`下管理

## 生成数据库model以及dao

在`cmd/gorm_gen.go`中配置数据库连接信息，并运行

## Wire生成依赖注入

在`cmd/wire`目录中执行`wire`命令

# 模板代码仓库地址

- Gitee：https://gitee.com/gradyyoung/gin-template
- Github：https://github.com/gradyyoung/gin-template
