package orm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

type BaseModel struct {
	ID        int64          `json:"id,string" gorm:"primaryKey;comment:ID"`                       //主键ID
	CreatedBy int64          `json:"createdBy,string" gorm:"comment:创建者"`                          //创建者
	UpdatedBy int64          `json:"updatedBy,string" gorm:"comment:修改者"`                          //修改者
	CreatedAt time.Time      `json:"createdAt" gorm:"type:datetime;default null;comment:创建时间"`     //新增时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"type:datetime;default null;comment:更新时间"`     //修改时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"type:datetime;default null;comment:删除时间" sql:"index"` //删除时间
}

func (baseModel *BaseModel) GetPrimaryKey() string {
	return "id"
}

func (baseModel *BaseModel) GetPrimaryValue() int64 {
	return baseModel.ID
}

var ormDb *gorm.DB

type Transaction struct {
	Context *gorm.DB
}

type Model interface {
	GetPrimaryKey() string
	GetPrimaryValue() int64
}

func InitGorm() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("mysql.user"),
		os.Getenv("mysql.password"),
		os.Getenv("mysql.host"),
		os.Getenv("mysql.port"),
		os.Getenv("mysql.database"),
	)
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("connect mysql error:", err)
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(20)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(10 * time.Minute)
	ormDb = db
}

func AutoMigrate(models ...interface{}) {
	_ = Default().Context.AutoMigrate(models...)
}

func Default() *Transaction {
	if ormDb == nil {
		InitGorm()
	}
	return &Transaction{
		Context: ormDb,
	}
}
