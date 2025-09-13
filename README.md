# topService

一个用Go语言编写的Web后端微服务，提供用户和产品管理功能。

## 功能特性

- RESTful API设计
- MySQL数据库集成
- GORM ORM框架
- Gin Web框架
- 参数验证
- 错误处理
- 日志记录
- CORS支持
- 环境配置

## 项目结构

```
topService/
├── main.go                 # 程序入口
├── go.mod                  # Go模块文件
├── .env                    # 环境配置文件
├── internal/
│   ├── config/            # 配置相关
│   │   └── config.go
│   ├── database/          # 数据库相关
│   │   └── database.go
│   ├── model/             # 数据模型
│   │   ├── user.go
│   │   └── product.go
│   ├── service/           # 业务逻辑层
│   │   ├── user_service.go
│   │   └── product_service.go
│   ├── handler/           # HTTP处理器层
│   │   ├── user_handler.go
│   │   └── product_handler.go
│   ├── middleware/        # 中间件
│   │   └── middleware.go
│   └── router/            # 路由配置
│       └── router.go
```

## 快速开始

### 1. 环境要求

- Go 1.21+
- MySQL 5.7+

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置数据库

修改 `.env` 文件中的数据库配置:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=topservice_db
```

### 4. 创建数据库

在MySQL中创建数据库:

```sql
CREATE DATABASE topservice_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 5. 运行服务

```bash
go run main.go
```

服务将在 `http://localhost:8080` 启动

## API 接口

### 健康检查
- `GET /health` - 服务健康检查

### 用户管理
- `POST /api/v1/users` - 创建用户
- `GET /api/v1/users` - 获取用户列表
- `GET /api/v1/users/:id` - 获取单个用户
- `PUT /api/v1/users/:id` - 更新用户
- `DELETE /api/v1/users/:id` - 删除用户

### 产品管理
- `POST /api/v1/products` - 创建产品
- `GET /api/v1/products` - 获取产品列表
- `GET /api/v1/products/:id` - 获取单个产品
- `PUT /api/v1/products/:id` - 更新产品
- `DELETE /api/v1/products/:id` - 删除产品

## API 示例

### 创建用户
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "phone": "13800138000"
  }'
```

### 创建产品
```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPhone 15",
    "description": "最新款iPhone",
    "price": 7999.99,
    "stock": 100,
    "category": "电子产品"
  }'
```

## 环境配置

项目支持以下环境变量配置：

| 变量名 | 描述 | 默认值 |
|--------|------|--------|
| DB_HOST | 数据库主机 | localhost |
| DB_PORT | 数据库端口 | 3306 |
| DB_USER | 数据库用户 | root |
| DB_PASSWORD | 数据库密码 | password |
| DB_NAME | 数据库名称 | topservice_db |
| SERVER_HOST | 服务器主机 | 0.0.0.0 |
| SERVER_PORT | 服务器端口 | 8080 |
| APP_ENV | 应用环境 | development |
| APP_DEBUG | 调试模式 | true |