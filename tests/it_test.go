package integration_test

import (
	"github.com/josuebrunel/clausify"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/url"
	"testing"
	"time"
)

var today = time.Now()

type Product struct {
	ID       int64 `gorm:"primary_key, AUTO_INCREMENT"`
	Name     string
	Category string
	Price    float64
	Quantity int64
	Status   bool
	Created  time.Time
}

var products = []Product{
	{Name: "mobile", Category: "tech", Price: 500, Quantity: 50, Created: today},
	{Name: "apple", Category: "fruit", Price: 2, Quantity: 100, Created: today},
	{Name: "tesla", Category: "car", Price: 40000, Quantity: 10, Created: today},
	{Name: "mango", Category: "fruits", Price: 5, Quantity: 70, Created: today},
	{Name: "ball", Category: "tech", Price: 15, Quantity: 100, Created: today},
}

var usecases = []struct {
	Input    string
	Expected int64
}{
	{Input: "name=apple", Expected: 1},
	{Input: "category__like=fruit", Expected: 2},
	{Input: "category__nlike=fruit", Expected: 3},
	{Input: "price__gte=10000", Expected: 1},
	{Input: "quantity__lt=20", Expected: 1},
	{Input: "price__between=0,20", Expected: 3},
	{Input: "price__nbetween=0,20", Expected: 2},
	{Input: "quantity__in=0,10,20", Expected: 1},
	{Input: "status__neq=false", Expected: 0},
}

func qs2cond(qs string) (c clausify.Clause) {
	u, _ := url.Parse("https://httpbin.org/?" + qs)
	c, _ = clausify.Clausify(u.Query())
	log.Printf("[CLAUSE] => %s, %s", c.Conditions, c.Variables)
	return
}

func setup() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost port=5445 user=test password=test dbname=test"), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	db = db.Debug()
	db.AutoMigrate(&Product{})
	db.Create(&products)
	return db
}

func teardown(db *gorm.DB) {
	db.Migrator().DropTable(&Product{})
}

func TestClausifyGORMIntegration(t *testing.T) {
	db := setup()
	for _, usecase := range usecases {
		c := qs2cond(usecase.Input)
		pp := []Product{}
		r := db.Where(c.Conditions, c.Variables...).Find(&pp)
		if r.RowsAffected != usecase.Expected {
			t.Errorf("Expected: %d \t Got: %d\n", usecase.Expected, r.RowsAffected)
		}
	}
	teardown(db)
}
