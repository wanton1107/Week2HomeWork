package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"homework3/config"
	"homework3/model"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	dataSource := config.AppConfig.Datasource
	DB, err = dbConnect(dataSource.Host, dataSource.Port, dataSource.Username, dataSource.Password, dataSource.Database)
	if err != nil {
		return err
	}
	return nil
}

func dbConnect(host string, port int, username string, password string, database string) (*gorm.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_",
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitData() error {
	DB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	if err := DB.Exec("TRUNCATE TABLE tb_comment").Error; err != nil {
		return err
	}
	if err := DB.Exec("TRUNCATE TABLE tb_post").Error; err != nil {
		return err
	}
	if err := DB.Exec("TRUNCATE TABLE tb_user").Error; err != nil {
		return err
	}
	DB.Exec("SET FOREIGN_KEY_CHECKS = 1")

	// 准备完整关联数据
	users := []model.User{
		{
			Username: "zhangsan",
			Password: "$argon2id$v=19$m=65536,t=3,p=4$xxxx$hash1",
			Email:    "zhangsan@example.com",
			Posts: []model.Post{
				{
					Title:   "Go语言并发编程详解",
					Content: "Go语言的并发模型基于CSP理论，通过goroutine和channel实现高效并发。goroutine是轻量级线程，初始栈仅2KB，可动态增长。channel用于goroutine间通信，遵循\"不要通过共享内存来通信，而要通过通信来共享内存\"的理念。select语句可同时监听多个channel，实现非阻塞通信。",
					Status:  1,
					Comments: []model.Comment{
						{UserID: 2, Content: "写得很清楚，受教了！"},
						{UserID: 3, Content: "channel的缓冲大小怎么设置合适？"},
						{UserID: 2, Content: "示例代码能再详细点吗"},
					},
				},
				{
					Title:   "GORM关联查询技巧",
					Content: "GORM的Preload方法可预加载关联数据，避免N+1查询问题。Joins适用于需要筛选关联数据的场景。Association方法可处理多对多关系的增删改。注意外键字段命名规范，建议使用表名+ID格式，便于自动识别。",
					Status:  1,
					Comments: []model.Comment{
						{UserID: 3, Content: "Preload和Joins的区别讲得很透"},
					},
				},
				{
					Title:   "微服务架构设计经验",
					Content: "微服务拆分粒度是关键，过细导致运维复杂，过粗失去微服务优势。建议按业务边界拆分，每个服务独立部署、独立扩展。服务间通信优先使用gRPC，性能优于HTTP JSON。配置中心、服务发现、熔断降级是必备基础设施。",
					Status:  1,
					Comments: []model.Comment{
						{UserID: 2, Content: "拆分粒度确实很难把握"},
						{UserID: 3, Content: "gRPC的学习成本怎么样"},
						{UserID: 2, Content: "有具体的项目案例吗"},
						{UserID: 3, Content: "期待更新"},
					},
				},
				{
					Title:   "Redis缓存策略总结",
					Content: "缓存穿透使用布隆过滤器或缓存空值解决。缓存击穿通过互斥锁或逻辑过期应对。缓存雪崩需设置随机过期时间加集群高可用。缓存与数据库一致性采用Cache-Aside模式，更新时先更新数据库再删缓存。",
					Status:  0,
					Comments: []model.Comment{
						{UserID: 2, Content: "布隆过滤器有推荐实现吗"},
					},
				},
			},
		},
		{
			Username: "lisi",
			Password: "$argon2id$v=19$m=65536,t=3,p=4$xxxx$hash2",
			Email:    "lisi@test.com",
			Posts: []model.Post{
				{
					Title:   "Docker容器化部署实践",
					Content: "Docker镜像构建要遵循最小化原则，使用多阶段构建减小体积。Docker Compose适合开发环境，生产环境推荐Kubernetes。容器健康检查配置liveness和readiness探针，确保服务可用性。资源限制防止单个容器耗尽宿主机资源。",
					Status:  1,
					Comments: []model.Comment{
						{UserID: 1, Content: "多阶段构建能减小多少体积？"},
						{UserID: 3, Content: "K8s的学习曲线确实陡"},
					},
				},
				{
					Title:   "Kubernetes入门指南",
					Content: "K8s核心概念包括Pod、Deployment、Service、Ingress。Pod是最小调度单元，Deployment管理Pod副本和滚动更新。Service提供集群内服务发现和负载均衡。Ingress管理外部访问，支持HTTPS配置。",
					Status:  1,
					Comments: []model.Comment{
						{UserID: 1, Content: "Deployment和StatefulSet区别是啥"},
						{UserID: 3, Content: "Ingress控制器用哪个好"},
						{UserID: 1, Content: "Nginx Ingress比较成熟"},
					},
				},
				{
					Title:   "MySQL性能优化案例",
					Content: "慢查询日志定位问题SQL，explain分析执行计划。索引优化遵循最左前缀原则，避免回表查询。大表分页使用覆盖索引加延迟关联。读写分离使用中间件或应用层实现，注意主从延迟问题。",
					Status:  1,
					Comments: []model.Comment{
						{UserID: 3, Content: "覆盖索引具体怎么设计"},
						{UserID: 1, Content: "主从延迟有解决方案吗"},
						{UserID: 3, Content: "强制走主库查询"},
					},
				},
				{
					Title:   "架构设计原则分享",
					Content: "SOLID原则是面向对象设计的基础，单一职责要求类只负责一件事。开闭原则对扩展开放对修改关闭。依赖倒置依赖抽象而非具体实现。微服务设计还需考虑CAP权衡，根据业务场景选择一致性或可用性优先。",
					Status:  1,
					Comments: []model.Comment{
						{UserID: 1, Content: "实际项目中很难完全遵循"},
					},
				},
				{
					Title:   "代码重构心得",
					Content: "重构前必须有完善的单元测试保障。小步快跑，每次只改一个地方。代码坏味道包括过长函数、过大类、重复代码等。设计模式是重构的有力工具，但不要过度设计。重构后代码应更易读、易测试、易扩展。",
					Status:  0,
					Comments: []model.Comment{
						{UserID: 3, Content: "测试覆盖率要到多少"},
						{UserID: 1, Content: "核心业务建议80%以上"},
						{UserID: 3, Content: "收到，感谢分享"},
					},
				},
			},
		},
		{
			Username: "wangwu",
			Password: "$argon2id$v=19$m=65536,t=3,p=4$xxxx$hash3",
			Email:    "wangwu@demo.cn",
			Posts: []model.Post{
				{
					Title:   "接口设计规范",
					Content: "RESTful API设计使用HTTP动词表示操作，URL使用名词复数。响应格式统一包装，包含code、message、data字段。错误码按模块划分，便于定位问题。版本控制通过URL路径或Header实现，推荐URL路径方式。",
					Status:  1,
					Comments: []model.Comment{
						{UserID: 1, Content: "GraphQL会不会替代REST"},
						{UserID: 2, Content: "各有适用场景吧"},
					},
				},
				{
					Title:   "日志收集方案对比",
					Content: "ELK Stack功能强大但资源消耗大，适合大规模日志。Loki轻量级与Grafana集成好，适合云原生环境。Fluentd和Filebeat是常用采集器，支持多种输出插件。日志分级很重要，生产环境避免输出debug日志。",
					Status:  1,
					Comments: []model.Comment{
						{UserID: 1, Content: "Loki的查询语法需要适应"},
						{UserID: 2, Content: "比ES的DSL简单多了"},
						{UserID: 1, Content: "确实，上手更快"},
					},
				},
				{
					Title:   "Git工作流总结",
					Content: "Git Flow适合版本发布周期明确的项目，分支管理规范但复杂。GitHub Flow简化流程，适合持续部署。分支命名要有意义，feature/功能描述、bugfix/问题描述。提交信息遵循规范，方便生成changelog。",
					Status:  1,
					Comments: []model.Comment{
						{UserID: 2, Content: "我们团队用的GitLab Flow"},
						{UserID: 1, Content: "和GitHub Flow差不多吧"},
					},
				},
			},
		},
	}

	// 使用FullSaveAssociations批量插入所有关联数据
	if err := DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&users).Error; err != nil {
		return err
	}

	if err := DB.Table("tb_post").Where("1=1").Updates(map[string]interface{}{"comment_status": 1}).Error; err != nil {
		return err
	}

	return nil

}
