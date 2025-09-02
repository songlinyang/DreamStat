package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint
	UserName   string
	Password   string
	Email      string
	PostsCount int
	Posts      []Post
}
type Post struct {
	gorm.Model
	ID            uint
	Title         string
	Body          string
	UserID        uint
	CommentStatue string
	Comments      []Comments
}
type Comments struct {
	gorm.Model
	ID             uint
	Content        string
	PostID         uint
	GoodStateCount uint `gorm:"default:0"`
	BadStateCount  uint `gorm:"default:0"`
}
type PostResult struct {
	ID       uint
	Title    string
	Body     string
	Comments []Comments `gorm:"-"`
}
type CommentResult struct {
	ID             uint
	Content        string
	GoodStateCount uint
	BadStateCount  uint
}
type UserResult struct {
	UserName string
	Password string
	Email    string
	Posts    []Post
}
type maxPostsResult struct {
	PostID      uint
	PostsCounts int
}

func main() {
	db, err := gorm.Open(sqlite.Open("identifier.sqlite"), &gorm.Config{})
	//题目1：模型定义
	//db.AutoMigrate(&User{}, &Post{}, &Comments{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//测试数据插入
	//db.Create(&User{
	//	UserName: "John Doe",
	//	Password: "123456",
	//	Email:    "john@126.com",
	//	Post:     []Post{{Title: "John's Post", Body: "I'm John hello world", Comments: []Comments{{Content: "John123"}, {Content: "John345"}}}, {Title: "John good state", Body: "John good state", Comments: []Comments{{Content: "John678"}, {Content: "John910"}}}},
	//})
	//db.Create(&User{
	//	UserName: "Jack",
	//	Password: "123456",
	//	Email:    "jack@126.com",
	//	Post:     []Post{{Title: "Jack's Post", Body: "I'm Jack Hello world", Comments: []Comments{{Content: "jack123"}, {Content: "jack345"}}}, {Title: "jack good state", Body: "jack good state", Comments: []Comments{{Content: "jack 678"}, {Content: "jack 910"}}}},
	//})

	// 题目2、1 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	var u User
	var UserName = User{UserName: "John Doe"}
	if err := db.Where(UserName).First(&u).Error; err != nil {
		panic(err)
	}

	var posts []Post
	if err := db.Model(&u).Association("Posts").Find(&posts); err != nil {
		panic(err)
	}
	fmt.Println(posts)
	var Cs []Comments
	for _, post := range posts {
		var postResult PostResult
		if err := db.Model(&post).Association("Comments").Find(&Cs); err != nil {
			panic(err)
		}
		db.Model(&post).Scan(&postResult)
		fmt.Println(postResult)
		for _, comment := range Cs {
			var commentResult CommentResult
			db.Model(&comment).Scan(&commentResult)
			fmt.Println(commentResult)
		}
	}

	// 题目2、2 编写Go代码，使用Gorm查询评论数量最多的文章信息。
	//select post_id,COUNT(1) as postsCount from comments GROUP BY post_id order by postsCount desc
	var c maxPostsResult
	if err := db.Model(&Comments{}).Select("post_id,count(post_id) as postCounts").Group("post_id").Order("postCounts desc").Limit(1).Scan(&c).Error; err != nil {
		panic(err)
	}
	fmt.Println("评论数量最多的文章ID是：", c.PostID)

	//题目3：钩子函数
	//3.1为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	if err := db.Select("user_id", "title", "body").Create(&Post{UserID: 1, Title: "你好", Body: "你好"}).Error; err != nil {
		panic(err)
	}
	//3.2为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
}

// 创建钩子
// 开始事务
func (p *Post) BeforeSave(tx *gorm.DB) error {
	fmt.Println("BeforeSave")
	//if u.Name != "Jinzhu" {
	//	return errors.New("invalid role")
	//}
	return nil
}

// BeforeSave, BeforeCreate, AfterSave, AfterCreate
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate")
	return nil
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterCreate")
	//if u.Name != "Jinzhu" {
	//	return errors.New("invalid role")
	//}
	return nil
}

// 关联前的 save
// 插入记录至 db
// 关联后的 save
type PostsCountResult struct {
	PostsCount int `gorm:"column:posts_count"`
}

func (p *Post) AfterSave(tx *gorm.DB) (err error) {
	fmt.Println("AfterSave")
	fmt.Println("BeforeCreate")
	//先查出有多少次
	var user User
	var postsCountResult PostsCountResult
	if err := tx.Select("posts_count").Where("id=?", p.UserID).First(&user).Scan(&postsCountResult).Error; err != nil {
		panic(err)
	}
	fmt.Println("当前用户文章数：", postsCountResult.PostsCount)
	postsCountResult.PostsCount = postsCountResult.PostsCount + 1
	if err := tx.Debug().Model(&user).Where("id=?", p.UserID).Update("posts_count", postsCountResult.PostsCount).Error; err != nil {
		panic(err)
	}
	fmt.Println("user.PostsCount累计：", postsCountResult.PostsCount)
	fmt.Println("hhhh=>", user)
	//if u.Name != "Jinzhu" {
	//	return errors.New("invalid role")
	//}
	return err
}
