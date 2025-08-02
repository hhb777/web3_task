# 任务4项目要求：
使用 Go 语言结合 Gin 框架和 GORM 库开发一个个人博客系统的后端，实现博客文章的基本管理功能，包括文章的创建、读取、更新和删除（CRUD）操作，同时支持用户认证和简单的评论功能。 

1.项目初始化
创建一个新的 Go 项目，使用 go mod init 初始化项目依赖管理。
安装必要的库，如 Gin 框架、GORM 以及数据库驱动（如 MySQL 或 SQLite）。

2.数据库设计与模型定义
设计数据库表结构，至少包含以下几个表：
users 表：存储用户信息，包括 id 、 username 、 password 、 email 等字段。
posts 表：存储博客文章信息，包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、 created_at 、 updated_at 等字段。
comments 表：存储文章评论信息，包括 id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段。
使用 GORM 定义对应的 Go 模型结构体。

3.用户认证与授权
实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码。
使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。

4.文章管理功能
实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
实现文章的更新功能，只有文章的作者才能更新自己的文章。
实现文章的删除功能，只有文章的作者才能删除自己的文章。

5.评论功能
实现评论的创建功能，已认证的用户可以对文章发表评论。
实现评论的读取功能，支持获取某篇文章的所有评论列表。

6.错误处理与日志记录
对可能出现的错误进行统一处理，如数据库连接错误、用户认证失败、文章或评论不存在等，返回合适的 HTTP 状态码和错误信息。
使用日志库记录系统的运行信息和错误信息，方便后续的调试和维护。

# 代码示例参考
数据库连接与模型定义
package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Email    string `gorm:"unique;not null"`
}

type Post struct {
    gorm.Model
    Title   string `gorm:"not null"`
    Content string `gorm:"not null"`
    UserID  uint
    User    User
}

type Comment struct {
    gorm.Model
    Content string `gorm:"not null"`
    UserID  uint
    User    User
    PostID  uint
    Post    Post
}

func main() {
    db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // 自动迁移模型
    db.AutoMigrate(&User{}, &Post{}, &Comment{})
}
 
用户注册与登录示例
package main

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "net/http"
    "time"
)

func Register(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // 加密密码
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    user.Password = string(hashedPassword)

    if err := db.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var storedUser User
    if err := db.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // 验证密码
    if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // 生成 JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":       storedUser.ID,
        "username": storedUser.Username,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte("your_secret_key"))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }
    // 剩下的逻辑...
}
 
# 提交要求
提交完整的项目代码，包括必要的配置文件和依赖管理文件。
提供项目的 README 文件，说明项目的运行环境、依赖安装步骤和启动方式。
    1.项目运行环境：windows10 x64系统
    2.依赖安装步骤：
        go mod init myblok
        go get -u gorm.io/gorm
        go get -u gorm.io/driver/sqlite
        go get -u gorm.io/driver/mysql
        go get -u github.com/gin-gonic/gin
        go get github.com/golang-jwt/jwt

    3.启动方式：go run ./main.go
可以使用 Postman 或其他工具对接口进行测试，并提供测试用例和测试结果。
测试用例OpenAPI：curl 'https://workspace.apipost.net/proxy/v2/runner' -X POST -H "Accept: text/event-stream" -H "Content-Type: application/json" --data-binary '{"option":{"scene":"auto_test","lang":"zh-cn","globals":{},"project":{"request":{"header":{"parameter":[{"key":"Accept","value":"*/*","is_checked":1,"field_type":"String","is_system":1},{"key":"Accept-Encoding","value":"gzip, deflate, br","is_checked":1,"field_type":"String","is_system":1},{"key":"User-Agent","value":"PostmanRuntime-ApipostRuntime/1.1.0","is_checked":1,"field_type":"String","is_system":1},{"key":"Connection","value":"keep-alive","is_checked":1,"field_type":"String","is_system":1}]},"query":{"parameter":[]},"body":{"parameter":[]},"cookie":{"parameter":[]},"auth":{"type":"noauth"},"pre_tasks":[],"post_tasks":[]}},"env":{"env_id":"1","env_name":"默认环境","env_pre_url":"","env_pre_urls":{"1":{"server_id":"1","name":"默认服务","sort":1000,"uri":""},"default":{"server_id":"1","name":"默认服务","sort":1000,"uri":""}},"environment":{}},"cookies":{"switch":1,"data":[]},"system_configs":{"send_timeout":0,"auto_redirect":-1,"max_redirect_time":5,"auto_gen_mock_url":-1,"request_param_auto_json":-1,"proxy":{"type":-1,"envfirst":1,"bypass":[],"protocols":["http"],"auth":{"authenticate":-1,"host":"","username":"","password":""}},"ca_cert":{"open":-1,"path":""},"client_cert":{}},"custom_functions":{},"collection":[{"target_id":"db82b2839b278","target_type":"api","parent_id":"db80edbf9b277","name":"registerapi","request":{"auth":{"type":"inherit"},"body":{"mode":"json","parameter":[],"raw":"{\r\n    \"username\": \"hhb\",\r\n    \"password\": \"123456\",\r\n    \"email\": \"123@.com\"\r\n}","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"query_add_equal":1,"parameter":[]},"cookie":{"cookie_encode":1,"parameter":[]},"restful":{"parameter":[]}},"parents":[{"target_id":"db80edbf9b277","target_type":"folder"}],"method":"PUT","protocol":"http/1.1","url":"127.0.0.1:8080/register","pre_url":""},{"target_id":"db80edbf9b277","target_type":"folder","parent_id":"0","name":"myblog","request":{"auth":{"type":"inherit"},"pre_tasks":[],"post_tasks":[],"body":{"parameter":[]},"header":{"parameter":[]},"query":{"parameter":[]},"cookie":{"parameter":[]}},"parents":[],"server_id":"0","pre_url":""},{"target_id":"db90225b9b2d5","target_type":"api","parent_id":"db80edbf9b277","name":"loginapi","request":{"auth":{"type":"inherit"},"body":{"mode":"json","parameter":[{"param_id":"db92e6bb71302","description":"","field_type":"string","is_checked":-1,"key":"username","not_null":1,"value":"hhb","content_type":"","file_name":"","file_base64":"","schema":{"type":"string"}},{"param_id":"db931e8f71336","description":"","field_type":"string","is_checked":-1,"key":"password","not_null":1,"value":"123456","content_type":"","file_name":"","file_base64":"","schema":{"type":"string"}}],"raw":"{\r\n    \"username\": \"hhb\",\r\n    \"password\": \"123456\"\r\n}","raw_parameter":[],"raw_schema":{"type":"object","required":["username","password"],"properties":{"password":{"type":"string","example":"123456"},"username":{"type":"string","example":"hhb"}},"x-schema-orders":["username","password"]},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"query_add_equal":1,"parameter":[]},"cookie":{"cookie_encode":1,"parameter":[]},"restful":{"parameter":[]}},"parents":[{"target_id":"db80edbf9b277","target_type":"folder"}],"method":"PUT","protocol":"http/1.1","url":"127.0.0.1:8080/login","pre_url":""},{"target_id":"eb71b6079b3bc","target_type":"api","parent_id":"db80edbf9b277","name":"createpost","request":{"auth":{"type":"inherit"},"body":{"mode":"json","parameter":[],"raw":"{\r\n    \"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQyMTUzOTAsImlkIjoxLCJ1c2VybmFtZSI6ImhoYiJ9.nQ-dJBADeosRJrFFJuT6Vcuo9o3lieqnaLlDysF1I-I\",\r\n    \"title\": \"hhbweb3test\",\r\n    \"user_id\":1,\r\n    \"content\": \"About web3 blog test content\"\r\n\r\n}","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"query_add_equal":1,"parameter":[]},"cookie":{"cookie_encode":1,"parameter":[]},"restful":{"parameter":[]}},"parents":[{"target_id":"db80edbf9b277","target_type":"folder"}],"method":"PUT","protocol":"http/1.1","url":"127.0.0.1:8080/article/create","pre_url":""},{"target_id":"ebe1c7839b3f2","target_type":"api","parent_id":"db80edbf9b277","name":"listpost","request":{"auth":{"type":"inherit"},"body":{"mode":"none","parameter":[],"raw":"","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"query_add_equal":1,"parameter":[]},"cookie":{"cookie_encode":1,"parameter":[]},"restful":{"parameter":[]}},"parents":[{"target_id":"db80edbf9b277","target_type":"folder"}],"method":"GET","protocol":"http/1.1","url":"127.0.0.1:8080/article/list","pre_url":""},{"target_id":"eef6df779b007","target_type":"api","parent_id":"db80edbf9b277","name":"readpost","request":{"auth":{"type":"inherit"},"body":{"mode":"none","parameter":[],"raw":"","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"query_add_equal":1,"parameter":[{"param_id":"eef849379b00b","field_type":"String","is_checked":1,"key":"title","not_null":1,"value":"hhbweb3test","description":""},{"description":"","field_type":"string","is_checked":1,"key":"userid","value":"1","not_null":1,"schema":{"type":"string"},"param_id":"eef8a7c79b037"}]},"cookie":{"cookie_encode":1,"parameter":[]},"restful":{"parameter":[]}},"parents":[{"target_id":"db80edbf9b277","target_type":"folder"}],"method":"GET","protocol":"http/1.1","url":"127.0.0.1:8080/article?title=hhbweb3test&userid=1","pre_url":""},{"target_id":"ef1e43479b10c","target_type":"api","parent_id":"db80edbf9b277","name":"updatepost","request":{"auth":{"type":"inherit"},"body":{"mode":"json","parameter":[],"raw":"{\r\n    \"title\": \"hhbweb3test\",\r\n    \"user_id\":1,\r\n    \"content\": \"About web3 blog test content,add and update content for test!\"\r\n\r\n}","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"parameter":[],"query_add_equal":1},"cookie":{"parameter":[],"cookie_encode":1},"restful":{"parameter":[]},"tabs_default_active_key":"query"},"parents":[{"target_id":"db80edbf9b277","target_type":"folder"}],"method":"PUT","protocol":"http/1.1","url":"127.0.0.1:8080/article/update","pre_url":""},{"target_id":"ef2ad4839b13f","target_type":"api","parent_id":"db80edbf9b277","name":"deletepost","request":{"auth":{"type":"inherit"},"body":{"mode":"json","parameter":[],"raw":"{\r\n    \"title\":\"hhbweb3testupdate\",\r\n    \"userid\":\"1\"\r\n}","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"parameter":[{"description":"","field_type":"string","is_checked":-1,"key":"title","value":"hhbweb3testupdate","not_null":1,"schema":{"type":"string"},"param_id":"ef2addfb9b140"},{"description":"","field_type":"string","is_checked":-1,"key":"userid","value":"1","not_null":1,"schema":{"type":"string"},"param_id":"ef2d52439b186"}],"query_add_equal":1},"cookie":{"parameter":[],"cookie_encode":1},"restful":{"parameter":[]},"tabs_default_active_key":"query"},"parents":[{"target_id":"db80edbf9b277","target_type":"folder"}],"method":"DELETE","protocol":"http/1.1","url":"127.0.0.1:8080/article","pre_url":""},{"target_id":"efc5bb139b229","target_type":"api","parent_id":"db80edbf9b277","name":"createcomment","request":{"auth":{"type":"inherit"},"body":{"mode":"json","parameter":[],"raw":"{\r\n    \"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQyMTUzOTAsImlkIjoxLCJ1c2VybmFtZSI6ImhoYiJ9.nQ-dJBADeosRJrFFJuT6Vcuo9o3lieqnaLlDysF1I-I\",\r\n    \"content\": \"web3 is chance and ...!\",\r\n    \"user_id\":1,\r\n    \"post_id\": 1\r\n\r\n}","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"query_add_equal":1,"parameter":[]},"cookie":{"cookie_encode":1,"parameter":[]},"restful":{"parameter":[]}},"parents":[{"target_id":"db80edbf9b277","target_type":"folder"}],"method":"PUT","protocol":"http/1.1","url":"127.0.0.1:8080/comment","pre_url":""},{"target_id":"efdc253b9b232","target_type":"api","parent_id":"db80edbf9b277","name":"listcomment","request":{"auth":{"type":"inherit"},"body":{"mode":"none","parameter":[],"raw":"","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"query_add_equal":1,"parameter":[{"description":"","field_type":"number","is_checked":1,"key":"post_id","value":"1","not_null":1,"schema":{"type":"number"},"param_id":"efde06fb9b234"},{"description":"","field_type":"number","is_checked":1,"key":"user_id","value":"1","not_null":1,"schema":{"type":"number"},"param_id":"efde42179b279"}]},"cookie":{"cookie_encode":1,"parameter":[]},"restful":{"parameter":[]}},"parents":[{"target_id":"db80edbf9b277","target_type":"folder"}],"method":"GET","protocol":"http/1.1","url":"127.0.0.1:8080/comment?post_id=1&user_id=1","pre_url":""}],"database_configs":{},"name":"myblog","ignore_error":-1,"enable_sandbox":-1,"iterationCount":1,"sleep":0,"testing_id":"f02664ef9b322","iterates_data_id":"0","iterationData":[]},"test_events":[{"type":"api","auto_sync":false,"test_id":"f02664ef9b322","event_id":"f02a275f9b323","enabled":1,"data":{"target_id":"db82b2839b278","project_id":"4cebef7a9821000","parent_id":"db80edbf9b277","target_type":"api","apiData":{"name":"registerapi","method":"PUT","protocol":"http/1.1","url":"127.0.0.1:8080/register","request":{"auth":{"type":"inherit"},"body":{"mode":"json","parameter":[],"raw":"{\r\n    \"username\": \"hhb\",\r\n    \"password\": \"123456\",\r\n    \"email\": \"123@.com\"\r\n}","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"query_add_equal":1,"parameter":[]},"cookie":{"cookie_encode":1,"parameter":[]},"restful":{"parameter":[]}}}}},{"type":"api","auto_sync":false,"test_id":"f02664ef9b322","event_id":"f02a275f9b324","enabled":1,"data":{"target_id":"db90225b9b2d5","project_id":"4cebef7a9821000","parent_id":"db80edbf9b277","target_type":"api","apiData":{"name":"loginapi","method":"PUT","protocol":"http/1.1","url":"127.0.0.1:8080/login","request":{"auth":{"type":"inherit"},"body":{"mode":"json","parameter":[{"param_id":"db92e6bb71302","description":"","field_type":"string","is_checked":-1,"key":"username","not_null":1,"value":"hhb","content_type":"","file_name":"","file_base64":"","schema":{"type":"string"}},{"param_id":"db931e8f71336","description":"","field_type":"string","is_checked":-1,"key":"password","not_null":1,"value":"123456","content_type":"","file_name":"","file_base64":"","schema":{"type":"string"}}],"raw":"{\r\n    \"username\": \"hhb\",\r\n    \"password\": \"123456\"\r\n}","raw_parameter":[],"raw_schema":{"type":"object","required":["username","password"],"properties":{"password":{"type":"string","example":"123456"},"username":{"type":"string","example":"hhb"}},"x-schema-orders":["username","password"]},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"query_add_equal":1,"parameter":[]},"cookie":{"cookie_encode":1,"parameter":[]},"restful":{"parameter":[]}}}}},{"type":"api","auto_sync":false,"test_id":"f02664ef9b322","event_id":"f02a27639b325","enabled":1,"data":{"target_id":"eb71b6079b3bc","project_id":"4cebef7a9821000","parent_id":"db80edbf9b277","target_type":"api","apiData":{"name":"createpost","method":"PUT","protocol":"http/1.1","url":"127.0.0.1:8080/article/create","request":{"auth":{"type":"inherit"},"body":{"mode":"json","parameter":[],"raw":"{\r\n    \"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQyMTUzOTAsImlkIjoxLCJ1c2VybmFtZSI6ImhoYiJ9.nQ-dJBADeosRJrFFJuT6Vcuo9o3lieqnaLlDysF1I-I\",\r\n    \"title\": \"hhbweb3test\",\r\n    \"user_id\":1,\r\n    \"content\": \"About web3 blog test content\"\r\n\r\n}","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"query_add_equal":1,"parameter":[]},"cookie":{"cookie_encode":1,"parameter":[]},"restful":{"parameter":[]}}}}},{"type":"api","auto_sync":false,"test_id":"f02664ef9b322","event_id":"f02a27639b326","enabled":1,"data":{"target_id":"ebe1c7839b3f2","project_id":"4cebef7a9821000","parent_id":"db80edbf9b277","target_type":"api","apiData":{"name":"listpost","method":"GET","protocol":"http/1.1","url":"127.0.0.1:8080/article/list","request":{"auth":{"type":"inherit"},"body":{"mode":"none","parameter":[],"raw":"","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"query_add_equal":1,"parameter":[]},"cookie":{"cookie_encode":1,"parameter":[]},"restful":{"parameter":[]}}}}},{"type":"api","auto_sync":false,"test_id":"f02664ef9b322","event_id":"f02a27639b327","enabled":1,"data":{"target_id":"eef6df779b007","project_id":"4cebef7a9821000","parent_id":"db80edbf9b277","target_type":"api","apiData":{"name":"readpost","method":"GET","protocol":"http/1.1","url":"127.0.0.1:8080/article?title=hhbweb3test&userid=1","request":{"auth":{"type":"inherit"},"body":{"mode":"none","parameter":[],"raw":"","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"query_add_equal":1,"parameter":[{"param_id":"eef849379b00b","field_type":"String","is_checked":1,"key":"title","not_null":1,"value":"hhbweb3test","description":""},{"description":"","field_type":"string","is_checked":1,"key":"userid","value":"1","not_null":1,"schema":{"type":"string"},"param_id":"eef8a7c79b037"}]},"cookie":{"cookie_encode":1,"parameter":[]},"restful":{"parameter":[]}}}}},{"type":"api","auto_sync":false,"test_id":"f02664ef9b322","event_id":"f02a27639b328","enabled":1,"data":{"target_id":"ef1e43479b10c","project_id":"4cebef7a9821000","parent_id":"db80edbf9b277","target_type":"api","apiData":{"name":"updatepost","method":"PUT","protocol":"http/1.1","url":"127.0.0.1:8080/article/update","request":{"auth":{"type":"inherit"},"body":{"mode":"json","parameter":[],"raw":"{\r\n    \"title\": \"hhbweb3test\",\r\n    \"user_id\":1,\r\n    \"content\": \"About web3 blog test content,add and update content for test!\"\r\n\r\n}","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"parameter":[],"query_add_equal":1},"cookie":{"parameter":[],"cookie_encode":1},"restful":{"parameter":[]},"tabs_default_active_key":"query"}}}},{"type":"api","auto_sync":false,"test_id":"f02664ef9b322","event_id":"f02a27639b329","enabled":1,"data":{"target_id":"ef2ad4839b13f","project_id":"4cebef7a9821000","parent_id":"db80edbf9b277","target_type":"api","apiData":{"name":"deletepost","method":"DELETE","protocol":"http/1.1","url":"127.0.0.1:8080/article","request":{"auth":{"type":"inherit"},"body":{"mode":"json","parameter":[],"raw":"{\r\n    \"title\":\"hhbweb3testupdate\",\r\n    \"userid\":\"1\"\r\n}","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"parameter":[{"description":"","field_type":"string","is_checked":-1,"key":"title","value":"hhbweb3testupdate","not_null":1,"schema":{"type":"string"},"param_id":"ef2addfb9b140"},{"description":"","field_type":"string","is_checked":-1,"key":"userid","value":"1","not_null":1,"schema":{"type":"string"},"param_id":"ef2d52439b186"}],"query_add_equal":1},"cookie":{"parameter":[],"cookie_encode":1},"restful":{"parameter":[]},"tabs_default_active_key":"query"}}}},{"type":"api","auto_sync":false,"test_id":"f02664ef9b322","event_id":"f02a27639b32a","enabled":1,"data":{"target_id":"efc5bb139b229","project_id":"4cebef7a9821000","parent_id":"db80edbf9b277","target_type":"api","apiData":{"name":"createcomment","method":"PUT","protocol":"http/1.1","url":"127.0.0.1:8080/comment","request":{"auth":{"type":"inherit"},"body":{"mode":"json","parameter":[],"raw":"{\r\n    \"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQyMTUzOTAsImlkIjoxLCJ1c2VybmFtZSI6ImhoYiJ9.nQ-dJBADeosRJrFFJuT6Vcuo9o3lieqnaLlDysF1I-I\",\r\n    \"content\": \"web3 is chance and ...!\",\r\n    \"user_id\":1,\r\n    \"post_id\": 1\r\n\r\n}","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"query_add_equal":1,"parameter":[]},"cookie":{"cookie_encode":1,"parameter":[]},"restful":{"parameter":[]}}}}},{"type":"api","auto_sync":false,"test_id":"f02664ef9b322","event_id":"f02a27639b32b","enabled":1,"data":{"target_id":"efdc253b9b232","project_id":"4cebef7a9821000","parent_id":"db80edbf9b277","target_type":"api","apiData":{"name":"listcomment","method":"GET","protocol":"http/1.1","url":"127.0.0.1:8080/comment?post_id=1&user_id=1","request":{"auth":{"type":"inherit"},"body":{"mode":"none","parameter":[],"raw":"","raw_parameter":[],"raw_schema":{"type":"object"},"binary":null},"pre_tasks":[],"post_tasks":[],"header":{"parameter":[]},"query":{"query_add_equal":1,"parameter":[{"description":"","field_type":"number","is_checked":1,"key":"post_id","value":"1","not_null":1,"schema":{"type":"number"},"param_id":"efde06fb9b234"},{"description":"","field_type":"number","is_checked":1,"key":"user_id","value":"1","not_null":1,"schema":{"type":"number"},"param_id":"efde42179b279"}]},"cookie":{"cookie_encode":1,"parameter":[]},"restful":{"parameter":[]}}}}}]}'
测试结果在myblog.json文件中记录

