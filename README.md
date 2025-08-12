# Gin Template

一个基于 Gin 框架的现代化 Go Web 应用模板，采用清洁架构设计，集成了常用的中间件和工具，帮助您快速构建高质量的 Web 应用程序。

## ✨ 功能特性

- 🚀 **高性能**：基于 Gin 框架，提供出色的性能表现
- 🏗️ **清洁架构**：采用分层架构设计，代码结构清晰，易于维护
- 🔌 **依赖注入**：使用 Google Wire 进行依赖注入管理
- 🗄️ **数据库集成**：支持 MySQL 数据库，使用 GORM ORM 框架
- 📦 **缓存支持**：集成 Redis 缓存，提升应用性能
- 📝 **日志系统**：使用 Logrus 提供结构化日志记录
- 🔐 **身份认证**：内置身份认证中间件
- ⚙️ **配置管理**：使用 Viper 进行配置文件管理
- 🛠️ **代码生成**：支持 GORM 模型自动生成
- 📋 **统一响应**：标准化的 API 响应格式

## 🛠️ 技术栈

| 技术 | 版本 | 用途 |
|------|------|------|
| [Gin](https://gin-gonic.com/) | v1.10.1 | Web 框架 |
| [GORM](https://gorm.io/) | v1.30.0 | ORM 框架 |
| [Go-Redis](https://redis.uptrace.dev/) | v9.11.0 | Redis 客户端 |
| [Logrus](https://github.com/sirupsen/logrus) | v1.9.3 | 日志框架 |
| [Wire](https://github.com/google/wire) | v0.6.0 | 依赖注入 |
| [Viper](https://github.com/spf13/viper) | v1.20.1 | 配置管理 |

## 📁 项目结构

```
gin-template/
├── build/                      # 构建输出目录
├── cmd/                        # 应用程序入口
│   ├── main.go                 # 主程序入口
│   └── wire/                   # Wire 依赖注入配置
│       ├── wire.go             # Wire 配置文件
│       └── wire_gen.go         # Wire 生成的代码
├── config/                     # 配置文件目录
│   └── application.yml         # 应用配置文件
├── internal/                   # 内部应用代码
│   ├── config/                 # 配置相关
│   │   ├── config.go           # 配置结构体定义
│   │   └── time.go             # 自定义时间类型
│   ├── constant/               # 常量定义
│   │   └── dict/               # 字典常量
│   ├── dao/                    # 数据访问层
│   ├── database/               # 数据库连接
│   │   ├── mysql.go            # MySQL 连接配置
│   │   ├── redis.go            # Redis 连接配置
│   │   └── provider_set.go     # Wire 提供者集合
│   ├── dto/                    # 数据传输对象
│   │   └── response.go         # 统一响应结构
│   ├── engine/                 # Gin 引擎配置
│   │   ├── engine.go           # 引擎初始化
│   │   ├── middleware/         # 中间件
│   │   │   ├── auth.go         # 身份认证中间件
│   │   │   ├── error.go        # 错误处理中间件
│   │   │   └── provider_set.go # Wire 提供者集合
│   │   └── v1/                 # V1 版本路由
│   │       └── routes.go       # 路由定义
│   ├── errs/                   # 错误定义
│   ├── handler/                # 控制器层
│   ├── model/                  # 数据模型
│   ├── service/                # 业务逻辑层
│   └── utils/                  # 工具函数
├── tools/                      # 工具脚本
│   └── gorm_gen/               # GORM 代码生成工具
│       └── gorm_gen.go         # 模型生成脚本
├── web/                        # 静态资源目录
├── Makefile                    # 构建脚本
├── go.mod                      # Go 模块文件
└── go.sum                      # Go 模块校验文件
```

## 🚀 快速开始

### 环境要求

- Go 1.23.9 或更高版本
- MySQL 5.7 或更高版本
- Redis 3.0 或更高版本

### 安装依赖

```bash
# 克隆项目
git clone https://github.com/gradyyoung/gin-template.git
cd gin-template

# 下载依赖
go mod download
```

### 配置数据库

1. 创建 MySQL 数据库
```sql
CREATE DATABASE test CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. 修改配置文件 `config/application.yml`
```yaml
mysql:
  dsn: root:your_password@tcp(127.0.0.1:3306)/your_database?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai

redis:
  host: localhost
  port: 6379
  password: ""
```

### 生成数据库模型

```bash
# 使用 Makefile 生成模型
make gen

# 或直接运行生成脚本
go run tools/gorm_gen/gorm_gen.go
```

### 生成依赖注入代码

```bash
# 使用 Makefile 生成 Wire 代码
make wire

# 或手动生成
cd cmd/wire && wire
```

### 运行应用

```bash
# 使用 Makefile 运行（推荐）
make run

# 或直接运行
go run cmd/main.go
```

应用将在 `http://localhost:8080` 启动

## 🔧 Makefile 命令

| 命令 | 描述 |
|------|------|
| `make build` | 构建应用程序 |
| `make run` | 运行应用程序 |
| `make clean` | 清理构建产物 |
| `make fmt` | 格式化代码 |
| `make tidy` | 整理 Go 模块 |
| `make wire` | 生成 Wire 依赖注入代码 |
| `make gen` | 生成 GORM 模型 |
| `make help` | 显示帮助信息 |

## 📖 使用指南

### 添加新的 API 路由

1. 在 `internal/handler/` 目录下创建处理器
2. 在 `internal/service/` 目录下创建业务逻辑
3. 在 `internal/engine/v1/routes.go` 中注册路由
4. 更新 Wire 配置文件

### 数据库模型管理

1. 在数据库中创建表结构
2. 修改 `tools/gorm_gen/gorm_gen.go` 中的配置
3. 运行 `make gen` 生成模型代码

### 中间件使用

项目内置了以下中间件：

- **身份认证中间件**：`internal/engine/middleware/auth.go`
- **错误处理中间件**：`internal/engine/middleware/error.go`

### 配置管理

所有配置都在 `config/application.yml` 文件中管理，支持以下配置项：

- 服务器配置（端口、认证等）
- 数据库配置（MySQL、Redis）
- 日志配置
- 时间格式配置

## 🚀 部署指南

### Docker 部署

```bash
# 构建镜像
docker build -t gin-template .

# 运行容器
docker run -p 8080:8080 gin-template
```

### 二进制部署

```bash
# 构建二进制文件
make build

# 部署到服务器
./build/gin-template
```

## 🤝 贡献指南

1. Fork 本仓库
2. 创建您的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交您的修改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开一个 Pull Request

## 📄 许可证

本项目基于 MIT 许可证开源 - 查看 [LICENSE](LICENSE) 文件了解详情

## 🔗 相关链接

- **Gitee**：https://gitee.com/gradyyoung/gin-template
- **GitHub**：https://github.com/gradyyoung/gin-template
- **文档**：[项目文档](https://github.com/gradyyoung/gin-template/wiki)
- **问题反馈**：[Issues](https://github.com/gradyyoung/gin-template/issues)

## 🙏 致谢

感谢以下开源项目为本模板提供的支持：

- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [Google Wire](https://github.com/google/wire)
- [Logrus](https://github.com/sirupsen/logrus)

---

如果这个项目对您有帮助，请给个 ⭐️ 支持一下！
