## Go Homework 01

## 使用说明

1. 编辑 `homework.go`，完成各个函数的实现。
2. 在本地运行 `go test -v` 验证代码。
3. 提交代码到你的分支。GitHub Actions 会自动运行测试。

## 题目列表

1. Single Number (只出现一次的数字)
2. Is Palindrome (回文数)
3. Is Valid Parentheses (有效的括号)
4. Longest Common Prefix (最长公共前缀)
5. Plus One (加一)
6. Remove Duplicates (删除有序数组中的重复项)
7. Merge Intervals (合并区间)
8. Two Sum (两数之和)

## Go Homework 02
在homework2路径下

## Go Homework 03
在homework3路径下
### 已完成功能
* 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
* 编写Go代码，使用Gorm创建这些模型对应的数据库表。
* 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
* 编写Go代码，使用Gorm查询评论数量最多的文章信息。
* 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
* 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

## Go Homework 04
在homework4路径下，使用Gin+Gorm完成博客系统开发
### 项目结构
* cmd 项目启动文件
* config 配置处理相关，包含数据库配置、jwt配置、log配置
* docs swagger文档
* internal 业务相关代码
  * dto 请求结构体
  * handler 接口处理函数
  * middleware 中间件，包含鉴权中间件、通用错误处理中间件和log处理
  * model 数据库模型
  * repository 持久层代码
  * service 业务层代码
* pkg 
  * apperror 错误处理
  * jwt 鉴权
  * logger 日志
* postman PostMan 测试用例
* router 路由
* application.yml 配置文件
* application-dev.yml