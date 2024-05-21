# B2CMallSystem

# B2C Mall System

这是一个基于 Beego 框架构建的  电商系统，后端使用 MySQL 和 Redis 数据库，前端使用 HTML5、CSS 和 jQuery。

## 项目结构
.
├── certfiles
├── common
│   ├── functions.go
│   └── utils.go
├── conf
│   └── app.conf
├── controllers
│   ├── api
│   │   ├── base.go
│   │   ├── product.go
│   │   ├── user.go
│   ├── backend
│   │   ├── administrator.go
│   │   ├── auth.go
│   │   ├── banner.go
│   │   ├── login.go
│   │   ├── main.go
│   │   ├── menu.go
│   │   ├── order.go
│   │   ├── productCate.go
│   │   ├── product.go
│   │   ├── productTypeAttr.go
│   │   ├── productType.go
│   │   ├── role.go
│   │   ├── setting.go
│   ├── frontend
│   │   ├── auth.go
│   │   ├── buy.go
│   │   ├── cart.go
│   │   ├── elasticsearch.go
│   │   ├── index.go
│   │   ├── product.go
│   │   ├── public.go
│   │   ├── user.go
├── models
│   ├── administrator.go
│   ├── auth.go
│   ├── banner.go
│   ├── menu.go
│   ├── order.go
│   ├── productCate.go
│   ├── product.go
│   ├── productTypeAttr.go
│   ├── productType.go
│   ├── role.go
│   ├── setting.go
│   └── user.go
├── routers
│   ├── api.go
│   ├── backend.go
│   ├── frontend.go
│   └── router.go
├── static
│   ├── css
│   ├── img
│   ├── js
│   └── lib
├── tests
│   ├── default_test.go
├── views
│   ├── backend
│   │   ├── administrator
│   │   │   ├── add.html
│   │   │   ├── edit.html
│   │   │   └── index.html
│   │   ├── auth
│   │   │   ├── add.html
│   │   │   ├── edit.html
│   │   │   └── index.html
│   │   ├── banner
│   │   │   ├── add.html
│   │   │   ├── edit.html
│   │   │   └── index.html
│   │   ├── login
│   │   │   └── login.html
│   │   ├── main
│   │   │   ├── index.html
│   │   │   └── welcome.html
│   │   ├── menu
│   │   │   ├── add.html
│   │   │   ├── edit.html
│   │   │   └── index.html
│   │   ├── order
│   │   │   ├── edit.html
│   │   │   └── order.html
│   │   ├── product
│   │   │   ├── add.html
│   │   │   ├── edit.html
│   │   │   └── index.html
│   │   ├── productCate
│   │   │   ├── add.html
│   │   │   ├── edit.html
│   │   │   └── index.html
│   │   ├── productType
│   │   │   ├── add.html
│   │   │   ├── edit.html
│   │   │   └── index.html
│   │   ├── productTypeAttribute
│   │   │   ├── add.html
│   │   │   ├── edit.html
│   │   │   └── index.html
│   │   ├── public
│   │   │   ├── error.html
│   │   │   ├── page_aside.html
│   │   │   ├── page_header.html
│   │   │   ├── page_menu.html
│   │   │   └── success.html
│   │   ├── role
│   │   │   ├── add.html
│   │   │   ├── auth.html
│   │   │   ├── edit.html
│   │   │   └── index.html
│   │   ├── setting
│   │       └── index.html
│   └── frontend
│       ├── auth
│       │   ├── login.html
│       │   ├── register_step1.html
│       │   ├── register_step2.html
│       │   └── register_step3.html
│       ├── buy
│       │   ├── checkout.html
│       │   └── confirm.html
│       ├── cart
│       │   ├── addcart_success.html
│       │   └── cart.html
│       ├── elasticsearch
│       │   └── list.html
│       ├── index
│       │   └── index.html
│       ├── product
│       │   ├── item.html
│       │   ├── list.html
│       │   ├── secitem.html
│       │   ├── sec-kill-list1.html
│       │   └── sec-kill-list.html
│       ├── public
│       │   ├── banner.html
│       │   ├── page_footer.html
│       │   ├── page_header.html
│       │   └── user_left.html
│       └── user
│           ├── order.html
│           ├── order_info.html
│           └── welcome.html
└── go.mod
└── go.sum
└── main.go
```

## 功能

- 管理员管理: 添加、编辑、删除管理员。
- 授权管理: 添加、编辑、删除授权。
- 横幅管理: 添加、编辑、删除横幅。
- 登录: 管理员和用户的安全登录功能。
- 主仪表盘: 系统概览。
- 菜单管理: 添加、编辑、删除菜单。
- 订单管理: 添加、编辑、删除订单，查看订单详情。
- 产品管理: 添加、编辑、删除产品，管理产品图片。
- 产品类别管理: 添加、编辑、删除产品类别。
- 产品类型管理: 添加、编辑、删除产品类型。
- 产品类型属性管理: 添加、编辑、删除产品类型属性。
- 角色管理: 添加、编辑、删除角色，分配权限。
- 设置管理: 查看和编辑系统设置。
- 用户管理: 用户注册、登录和个人资料管理。
- 购物车管理: 将产品添加到购物车，查看购物车，进行结算。
- Elasticsearch 集成: 产品搜索功能。
- 前端页面: 用户认证、产品列表、购物车、订单详情等。

## 设置与安装

### 先决条件

- Go 1.18 或更高版本
- MySQL
- Redis

### 步骤

1. **克隆仓库**:
   
   git clone https://github.com/yourusername/b2c-ecommerce-system.git
   cd b2c-ecommerce-system
 

2. 设置 MySQL:
   - 创建一个名为 `shop` 的 MySQL 数据库。
   - 更新 `app.conf` 文件中的 MySQL 凭证。

3. 设置 Redis:
   - 确保 Redis 已安装并在系统上运行。

4. 安装依赖项:
   
   go mod tidy
 

5. 运行应用程序:
   
   go run main.go
   

6. 访问应用程序:
   - 后台: `http://localhost:8080/backend`
   - 前台: `http://localhost:8080`

## 注意
为你的实际 GitHub 用户名。
2. 确保在你的仓库中包含 `go.mod` 和 `go.sum` 文件，以便进行依赖管理。
3. 确保所有必需的数据库和缓存服务（如 MySQL 和 Redis）已正确配置和运行。
