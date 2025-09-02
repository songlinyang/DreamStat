package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// // 创建记录
//
//	type User struct {
//		gorm.Model
//		ID         uint
//		Name       string
//		Age        int
//		Birthday   time.Time
//		CreditCard CreditCard
//		Company    Company
//		Email      string `gorm:"unique"`
//	}
//
//	type Email struct {
//		gorm.Model
//		UserID  uint
//		address string
//	}
//
//	type Company struct {
//		gorm.Model
//		UserID uint
//		Name   string
//	}
//
//	type CreditCard struct {
//		gorm.Model
//		Number string
//		UserID uint
//	}
type Employees struct {
	gorm.Model
	ID         uint
	Name       string
	Department string
	Salary     float64
}
type Books struct {
	gorm.Model
	ID     uint
	Title  string
	Author string
	Price  uint32
}

func main() {
	db, err := gorm.Open(sqlite.Open("identifier.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//创建 表
	db.AutoMigrate(&Employees{})
	db.AutoMigrate(&Books{})
}

//	db.AutoMigrate(&CreditCard{})
//	db.AutoMigrate(&Email{})
//	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
//	//
//	//db.Create(&User{
//	//	Name:       "Jinzhu",
//	//	CreditCard: CreditCard{Number: "411111111111"},
//	//})
//	// skip all associations
//	//db.Omit(clause.Associations).Create(&user)
//	//err = db.Create(&user).Error // 通过数据的指针来创建
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	//
//	//fmt.Println(user.ID)             // 返回插入数据的主键
//	//fmt.Println(result.Error)        // 返回 error
//	//fmt.Println(result.RowsAffected) // 返回插入记录的条数
//	//
//	//users := []*User{
//	//	{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
//	//	{Name: "Jackson", Age: 19, Birthday: time.Now()},
//	//}
//	//
//	//db.Create(users) // pass a slice to insert multiple row
//	//db.Debug().Select("Name", "Age", "CreatedAt").Create(&user)
//	//db.Debug().Omit("Name", "Age", "CreatedAt").Create(&user)
//	//var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
//	//db.Create(&users)
//	//
//	//for _, user := range users {
//	//	fmt.Println(user.ID) // 1,2,3
//	//}
//
//	//	var users = []User{}
//	//	for i := 0; i < 200; i++ {
//	//		users = append(users, User{Name: fmt.Sprint("jinzhu_", i)})
//	//	}
//	//
//	//	// batch size 100
//	//	db.CreateInBatches(users, 100)
//
//	//db.First(user, 10)
//	// SELECT * FROM users WHERE id = 10;
//
//	//db.First(user, "10")
//	// SELECT * FROM users WHERE id = 10;
//
//	//var result User
//
//	//db.Debug().Find(&user, []int{1, 2, 3})
//	//db.Debug().Model(User{ID: 10}).First(&result) ???
//	// SELECT * FROM users WHERE id IN (1,2,3);
//	//users := User{}
//	// Struct
//	//db.Debug().Where(&User{Name: "jinzhu", Age: 20}).First(&users)
//	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;
//
//	// Map
//	//db.Debug().Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
//	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
//
//	// Slice of primary keys
//	//db.Debug().Where([]int64{20, 21, 22}).Find(&users)
//	// SELECT * FROM users WHERE id IN (20, 21, 22);
//
//	//内联条件
//	//user := User{}
//	//// Get by primary key if it were a non-integer type
//	//db.First(&user, "id = ?", "string_primary_key")
//	//// SELECT * FROM users WHERE id = 'string_primary_key';
//	//
//	//// Plain SQL
//	//db.Find(&user, "name = ?", "jinzhu")
//	//// SELECT * FROM users WHERE name = "jinzhu";
//	//
//	//db.Find(&user, "name <> ? AND age > ?", "jinzhu", 20)
//	//// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;
//	//
//	//// Struct
//	//db.Find(&user, User{Age: 20})
//	//// SELECT * FROM users WHERE age = 20;
//	//
//	//// Map
//	//db.Debug().Find(&user, map[string]interface{}{"age": 20})
//	//// SELECT * FROM users WHERE age = 20;
//	// Not条件
//	//user := User{}
//	//res := db.Debug().Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").First(&user)
//	//fmt.Println(res)
//	//db.Debug().Not("name=?", "jinzhu").First(&user)
//	//// Not In
//	//db.Debug().Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&user)
//	//// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");
//
//	//// Struct
//	//db.Debug().Not(User{Name: "jinzhu", Age: 18}).First(&user)
//	//// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;
//	//
//	//// Not In slice of primary keys
//	//db.Debug().Not([]int64{1, 2, 3}).First(&user)
//	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;
//	// GroupBY
//	//res1 := db.Debug().Distinct("name", "age").Order("name, age desc").Find(&user)
//	//fmt.Println(res1)
//	////Having
//	////Distinct
//	//res2 := db.Debug().Distinct("name", "age").Order("name, age desc").Find(&user)
//	//fmt.Println(res2)
//	//type result struct {
//	//	Name  string
//	//	CName string
//	//}
//	//db.Create(&User{
//	//	Name:     "John Doe",
//	//	Age:      25,
//	//	Birthday: time.Now(),
//	//	Email:    "yangsl@qq.com",
//	//})
//	db.AutoMigrate(&Company{})
//	//err = db.Create(&User{
//	//	Name:    "John Doe",
//	//	Age:     42,
//	//	Company: Company{Name: "Company 1"}}).Error
//	//fmt.Println(err)
//	//Left Join
//	//users := User{}
//	//db.Debug().Joins("Company", db.Where(&Company{Name: ""})).Find(&users)
//	//db.Model(&User{}).Debug().Select("users.name, companies.name").Joins("left join companies on companies.user_id = users.id")
//	//fmt.Println(err)
//	//db.Table("users").Select("users.name, companies.name").Joins("left join companies on companies.user_id = users.id")
//	// inner join
//	//db.Debug().InnerJoins("Company").Find(&users)
//	//更新
//
//}
//
//// 创建钩子
//// 开始事务
//func (u *User) BeforeSave(tx *gorm.DB) error {
//	fmt.Println("BeforeSave")
//	//if u.Name != "Jinzhu" {
//	//	return errors.New("invalid role")
//	//}
//	return nil
//}
//
//// BeforeSave, BeforeCreate, AfterSave, AfterCreate
//func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
//	fmt.Println("BeforeCreate")
//	//if u.Name != "Jinzhu" {
//	//	return errors.New("invalid role")
//	//}
//	return nil
//}
//
//func (u *User) AfterCreate(tx *gorm.DB) (err error) {
//	fmt.Println("AfterCreate")
//	//if u.Name != "Jinzhu" {
//	//	return errors.New("invalid role")
//	//}
//	return nil
//}
//
//// 关联前的 save
//// 插入记录至 db
//// 关联后的 save
//
//func (u *User) AfterSave(tx *gorm.DB) (err error) {
//	fmt.Println("AfterSave")
//	//if u.Name != "Jinzhu" {
//	//	return errors.New("invalid role")
//	//}
//	return nil
//}
//
//func (u *CreditCard) BeforeSave(tx *gorm.DB) error {
//	fmt.Println("CreditCard_BeforeSave")
//	return nil
//}
//
//// BeforeSave, BeforeCreate, AfterSave, AfterCreate
//func (u *CreditCard) BeforeCreate(tx *gorm.DB) error {
//	fmt.Println("CreditCard_BeforeCreate")
//	return nil
//}
//
//func (u *CreditCard) AfterCreate(tx *gorm.DB) error {
//	fmt.Println("CreditCard_AfterCreate")
//	return nil
//}
//
//// 关联前的 save
//// 插入记录至 db
//// 关联后的 save
//
//func (u *CreditCard) AfterSave(tx *gorm.DB) error {
//	fmt.Println("CreditCard_AfterSave")
//	return nil
//}
//
//// 提交或回滚事务
