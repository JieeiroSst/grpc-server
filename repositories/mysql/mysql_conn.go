package mysql

import (
	"log"
	"sync"
	"errors"

	"github.com/JIeeiroSst/go-app/entities"
	"github.com/JIeeiroSst/go-app/inventory"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mutex    sync.Mutex
	instance *Mysqlconn
)

type Mysqlconn struct {
	db *gorm.DB
}

type Config struct {
	DSN string
}

func GetMysqlConnInstance(c *Config) *Mysqlconn {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			dsn := c.DSN
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				panic(err)
			}
			instance = &Mysqlconn{
				db: db,
			}
			db.AutoMigrate(&entities.Item{})
		}
	}
	return instance
}

func NewMysqlConnRepo(c *Config) *Mysqlconn {
	return &Mysqlconn{
		db: GetMysqlConnInstance(c).db,
	}
}

func (db *Mysqlconn) CheckAndOrder(items []*inventory.Item) (bool, error) {
	var itemStore entities.Item
	var compare bool
	for _, item := range items {
		err := db.db.Where("product_id = ?", item.ProductId).Find(&itemStore)
		if err != nil {
			log.Println("server running error", err)
		}
		if itemStore.Quantity == 0 {
			compare = false
		}
		if itemStore.Quantity == item.Quantity {
			compare = true
			err := db.db.Model(&itemStore).Where("product_id = ?", item.ProductId).Update("quantity = ?", 0).Error
			if err != nil {
				log.Println("server running error", err)
			}
		}
		if itemStore.Quantity > item.Quantity {
			compare = true
			remainQuantity := item.Quantity - itemStore.Quantity
			err := db.db.Model(&itemStore).Where("product_id = ?", item.ProductId).Update("quantity = ?", remainQuantity).Error
			if err != nil {
				log.Println("server running error", err)
			}
		}
		if itemStore.Quantity < item.Quantity {
			compare = false
		}
	}
	return compare, nil
}

func (db *Mysqlconn) CreateItem(item entities.Item) error {
	return db.db.Create(item).Error
}

func (db *Mysqlconn) GetById(id string,item entities.Item) (entities.Item , error){
	err:=db.db.Where("product_id = ?",id).Find(&item)
	if err!=nil {
		return entities.Item{},errors.New("get data error")
	}
	return item,nil
}

func (db *Mysqlconn) GetAll() ([]entities.Item,error) {
	var items []entities.Item
	err:=db.db.Find(&items)
	if err!=nil {
		return nil, errors.New("get data error")
	}
	return items,nil 
}