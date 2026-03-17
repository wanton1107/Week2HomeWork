package main

import (
	"encoding/json"
	"homework3/config"
	"homework3/model"
	"homework3/repository"
	"log"
)

func main() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	if err := repository.InitDB(); err != nil {
		panic(err)
	}

	if err := repository.DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}); err != nil {
		panic(err)
	}

	if err := repository.InitData(); err != nil {
		panic(err)
	}
	// 查询指定用户已发布的文章和评论
	var user model.User
	repository.DB.Preload("Posts", "status=?", 1).Preload("Posts.Comments").Where(&model.User{ID: 1}).First(&user)
	data, _ := json.MarshalIndent(user, "", "\t")
	log.Println("----------------题目2-1-------------------")
	log.Println(string(data))

	// 查询评论数最多的文章
	statistics := make([]map[string]interface{}, 0)
	err := repository.DB.Model(&model.Post{}).
		Joins("left join tb_comment on tb_post.id = tb_comment.post_id").
		Select("tb_post.*,count(*) as comment_count").
		Group("tb_post.id").
		Order("count(*) desc").
		First(&statistics).Error
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("------------------------题目2-2---------------------------")
	data, _ = json.MarshalIndent(statistics, "", "\t")
	log.Println(string(data))

	// 测试新增文章钩子函数
	log.Println("-----------------------题目3-1--------------------------")
	var testUser model.User
	repository.DB.Where(&model.User{ID: 1}).First(&testUser)
	log.Printf("新增前文章数量：%d\n", testUser.PostCount)
	repository.DB.Model(&model.Post{}).Create(&model.Post{
		Title:   "GitAAAA",
		Content: "Git Flowdfdffffffffffff，分支管理规范但复杂。GitHub Flow简化流程，适合持续部署。分支命名要有意义，feature/功能描述、bugfix/问题描述。提交信息遵循规范，方便生成changelog。",
		Status:  1,
		UserID:  1,
		Comments: []model.Comment{
			{UserID: 1, Content: "和GitHub Flow差不多吧"},
		},
	})
	repository.DB.Where(&model.User{ID: 1}).First(&testUser)
	log.Printf("新增后文章数量：%d\n", testUser.PostCount)

	// 测试删除评论钩子函数
	log.Println("------------------------题目3-2---------------------------")
	var testPost model.Post
	repository.DB.Where(&model.Post{ID: 2}).First(&testPost)
	log.Printf("删除前文章评论状态：%d\n", testPost.CommentStatus)
	var testComment model.Comment
	repository.DB.Where("post_id=?", 2).First(&testComment)
	if err := repository.DB.Delete(&testComment).Error; err != nil {
		log.Println(err.Error())
	}
	repository.DB.Where(&model.Post{ID: 2}).First(&testPost)
	log.Printf("删除后文章评论状态：%d\n", testPost.CommentStatus)

}
