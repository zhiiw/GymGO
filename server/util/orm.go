package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	gorm.Model
	ID        uint `gorm:"primaryKey; autoIncrement"`
	Name      string
	UserName  string
	Password  string
	Salary    float64
	Address   string
	Gender    string
	Email     string
	Photo     string
	EntryTime string
	Img       string
	Birthday  time.Time
}

type Coach struct {
	gorm.Model
	ID            uint `gorm:"primaryKey; autoIncrement"`
	AvailableTime string
	Speciality    string
	NickName      string
	Employee      uint `gorm:"foreignKey:ID"`
}

type Manager struct {
	gorm.Model
	ID          uint `gorm:"primaryKey; autoIncrement"`
	RegisterCnt uint
	TotalIncome float64
	Employee    uint `gorm:"foreignKey:ID"`
}

type Maintainer struct {
	gorm.Model
	ID       uint `gorm:"primaryKey; autoIncrement"`
	Type     string
	Employee uint `gorm:"foreignKey:ID"`
}

type Equipment struct {
	gorm.Model
	ID   uint `gorm:"primaryKey; autoIncrement"`
	Name string
	Pos  string
	Img  string
}

type MaintainerEquipment struct {
	gorm.Model
	ID           uint `gorm:"primaryKey; autoIncrement"`
	MaintainerId uint `gorm:"foreignKey: ID"`
	EquipmentId  uint `gorm:"foreignKey: ID"`
	MaintainTime time.Time
	Description  string
}

type Customer struct {
	gorm.Model
	ID                    uint `gorm:"primaryKey; autoIncrement"`
	UserName              string
	PassWord              string
	RegisterTime          time.Time
	ManagerOfCustomer     uint `gorm:"foreignKey:ID"`
	Weight                float64
	BodyFat               float64
	CalculatedDailyIntake string
}

type Curriculum struct {
	gorm.Model
	ID      uint `gorm:"primaryKey; autoIncrement"`
	Name    string
	Type    string
	CoachId uint `gorm:"foreignKey:ID"`
}

type VIPCard struct {
	gorm.Model
	ID         uint `gorm:"primaryKey; autoIncrement"`
	Balance    float64
	Due        time.Time
	Level      uint
	TotalSpend float64
	CustomerId uint `gorm:"foreignKey:ID"`
}

type CurriculumCustomer struct {
	gorm.Model
	ID           uint `gorm:"primaryKey; autoIncrement"`
	CustomerID   uint `gorm:"foreignKey: ID"`
	CurriculumID uint `gorm:"foreignKey: ID"`
	Attendance   uint
}

type CurriculumRecord struct {
	gorm.Model
	ID             uint `gorm:"primaryKey; autoIncrement"`
	CustomerID     uint `gorm:"foreignKey: ID"`
	CurriculumID   uint `gorm:"foreignKey: ID"`
	AttendanceTime time.Time
	Rating         uint
	Description    string
}

type CalculatedIntake struct {
	gorm.Model
	ID               uint `gorm:"primaryKey; autoIncrement"`
	Customer         uint
	BreakfastEnergy  float64
	BreakfastFat     float64
	BreakfastProtein float64
	BreakfastCarbo   float64
	BreakfastSurge   float64
	LunchEnergy      float64
	LunchFat         float64
	LunchProtein     float64
	LunchCarbo       float64
	LunchSurge       float64
	SupperEnergy     float64
	SupperFat        float64
	SupperProtein    float64
	SupperCarbo      float64
	SupperSurge      float64
}

type Food struct {
	gorm.Model
	ID       uint `gorm:"primaryKey; autoIncrement"`
	Name     string
	Unit     float64
	Calories float64
	Brand    string
}

type Nutrition struct {
	gorm.Model
	ID      uint `gorm:"primaryKey; autoIncrement"`
	Type    string
	Amount  float64
	Unit    float64
	FoodRef uint `gorm:"foreignKey: ID"`
}

type DailyIntake struct {
	gorm.Model
	ID       uint `gorm:"primaryKey; autoIncrement"`
	Date     time.Time
	Customer uint `gorm:"foreignKey: ID"`
	MealType string
}

type FoodRank struct {
	gorm.Model
	ID          uint `gorm:"primaryKey; autoIncrement"`
	Food        uint `gorm:"foreignKey: ID"`
	Rank        uint
	Description string
}

type DailyIntakeFood struct {
	gorm.Model
	ID          uint `gorm:"primaryKey; autoIncrement"`
	Food        uint `gorm:"foreignKey: ID"`
	DailyIntake uint `gorm:"foreignKey: ID"`
	Amount      uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Employee{})
	db.AutoMigrate(&Coach{})
	db.AutoMigrate(&Manager{})
	db.AutoMigrate(&Maintainer{})
	db.AutoMigrate(&Equipment{})
	db.AutoMigrate(&Customer{})
	db.AutoMigrate(&Curriculum{})
	db.AutoMigrate(&VIPCard{})
	db.AutoMigrate(&CurriculumCustomer{})
	db.AutoMigrate(&Food{})
	db.AutoMigrate(&Nutrition{})
	db.AutoMigrate(&DailyIntake{})
	db.AutoMigrate(&MaintainerEquipment{})
	db.AutoMigrate(&FoodRank{})
	db.AutoMigrate(&CalculatedIntake{})
	db.AutoMigrate(&CurriculumRecord{})
	db.AutoMigrate(&DailyIntakeFood{})

	// Create

	// Read

	// Update - update multiple fields

}
