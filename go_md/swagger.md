接口文档生成工具swagger
======

# 1.下载swag
```shell script
go get -u github.com/swaggo/swag/cmd/swag
```

# 2.在项目跟路径(包含main.go文件)执行swag init命令。
> Run swag init in the project's root folder which contains the main.go file. 
  This will parse your comments and generate the required files (docs folder and docs/docs.go).
```shell script
swag init
```

# 3.初始化swagger
```go
func initSwagger() {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "OpenAPI"
	docs.SwaggerInfo.Description = "API接口文档"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "example.api.org"
	docs.SwaggerInfo.BasePath = "/tools"
	docs.SwaggerInfo.Schemes = []string{"http","https"}
}
```

# 4.在相应的接口上添加接口操作注释,例如:
```go
// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.DefaultError
// @Router /accounts/{id} [get]
func (c *Controller) ShowAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account, err := model.AccountOne(aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, account)
}
```

# 5.添加路由
```go
import (
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// use ginSwagger middleware to serve the API docs
r.GET("/tools/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

# 6.执行swag init命令，更新接口文件

### Usage
https://github.com/swaggo/swag