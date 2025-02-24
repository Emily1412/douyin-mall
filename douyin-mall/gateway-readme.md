Q:后续如果要添加其他服务路由，需要如何添加？

除了payment还有user、product、cart、order、checkout这五个服务

按照现有的 payment 服务结构，添加其他服务路由需要：
1.在 gateway/handler/ 下创建对应的处理器:

```go
package user
import (
    "github.com/gin-gonic/gin"
    pb "douyin-mall/user-service/api"  // 导入对应服务的 proto
)

func Register(c *gin.Context) {
    // 实现注册逻辑
}

func Login(c *gin.Context) {
    // 实现登录逻辑
}

func Profile(c *gin.Context) {
    // 实现获取用户信息逻辑
}
```

2.在 gateway/router/routes/ 下添加路由文件:

```go
package routes

import (
    "douyin-mall/gateway/handler/user"
    "douyin-mall/gateway/middleware"
    "github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup) {
    users := r.Group("/users")
    {
        users.POST("/register", user.Register)
        users.POST("/login", user.Login)
         // 需要认证的路由
    	auth := users.Use(middleware.AuthMiddleware())
    	{
        	auth.GET("/profile", user.Profile)
    	}
	}
}
```



3. 在 router.go 中注册新路由:

  ```go
  func SetupRouter() *gin.Engine {
   // ... 现有代码 ...
  
   api := r.Group("/api/v1")
   {
       routes.RegisterPaymentRoutes(api)
       routes.RegisterUserRoutes(api)
       routes.RegisterProductRoutes(api)
       routes.RegisterCartRoutes(api)
       routes.RegisterOrderRoutes(api)
       routes.RegisterCheckoutRoutes(api)
   }
  
   return r
  }
  ```


其他服务也是类似的结构，只是具体的处理逻辑和路由不同。需要我详细说明某个服务的具体路由吗？

# **主要看这里！**

前面当补充看

Q: 请把全部服务添加服务路由的步骤详细说明

我来详细说明添加所有服务路由的步骤：
1.创建各服务的 Proto 文件:

```protobuf
user-service/api/user.proto
product-service/api/product.proto
cart-service/api/cart.proto
order-service/api/order.proto
checkout-service/api/checkout.proto
```

2.在 gateway/handler/ 下创建对应的处理器:

```go
package user

import (
    "github.com/gin-gonic/gin"
    pb "douyin-mall/user-service/api" // 导入对应服务的 proto
)

func Register(c *gin.Context) {
    // 实现注册逻辑
}
func Login(c *gin.Context) {
    // 实现登录逻辑
}
func Profile(c *gin.Context) {
    // 实现获取用户信息逻辑
}
```

其他服务类似：

```
gateway/handler/product/product.go
gateway/handler/cart/cart.go
gateway/handler/order/order.go
gateway/handler/checkout/checkout.go
```



3.在 gateway/router/routes/ 下创建路由文件:

```go
package routes

import (
    "douyin-mall/gateway/handler/user"
    "douyin-mall/gateway/middleware"
    "github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup) {
    users := r.Group("/users")
    {
        users.POST("/register", user.Register)
        users.POST("/login", user.Login)
        //这里和文章开头的第二步代码有些不同，如果这里报错的话可以试一下上面的
        users.GET("/profile", middleware.AuthMiddleware(), user.Profile)
    }
}
```

类似地创建：

```
routes/product.go (商品列表、详情、搜索)
routes/cart.go (购物车增删改查)
routes/order.go (订单创建、查询、取消)
routes/checkout.go (结算相关)
```



4.更新 router.go 注册所有路由:

```go
func SetupRouter() *gin.Engine {
    r := gin.Default()
    // ... 现有中间件代码 ...

	api := r.Group("/api/v1")
	{
    	routes.RegisterUserRoutes(api)
    	routes.RegisterProductRoutes(api)
    	routes.RegisterCartRoutes(api)
    	routes.RegisterOrderRoutes(api)
    	routes.RegisterCheckoutRoutes(api)
    	routes.RegisterPaymentRoutes(api)
	}

	return r
}
```



5.更新健康检查:

在 **gateway/handler/health/health.go** 中取消注释并实现所有服务的健康检查：

```go
var serviceChecks = map[string]healthCheckFunc{
	"payment": checkPaymentService,
	// "user":     checkUserService,
	// "order":    checkOrderService,
	// "product":  checkProductService,
	// "cart":     checkCartService,
	// "checkout": checkCheckoutService,
}
```



6.更新配置文件:

```yaml
services:
  user: "localhost:50051"
  product: "localhost:50052"
  cart: "localhost:50053"
  order: "localhost:50054"
  checkout: "localhost:50055"
  payment: "localhost:50056"
```

