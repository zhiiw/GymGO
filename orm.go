package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	gorm.Model
	ID       uint `gorm:"primaryKey; autoIncrement"`
	UserName string
	Password string
}

type Trainer struct {
	gorm.Model
	ID            uint `gorm:"primaryKey; autoIncrement"`
	AvailableTime string
	Speciality    string
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
	Employee uint `gorm:"foreignKey:ID"`
}

type Equipment struct {
	gorm.Model
	ID   uint `gorm:"primaryKey; autoIncrement"`
	Name string
	Pos  string
	Img  string
}

type Customer struct {
	gorm.Model
	ID                    uint `gorm:"primaryKey; autoIncrement"`
	UserName              string
	PassWord              string
	RegisterTime          time.Time
	ManagerOfCustomer     uint `gorm:"foreignKey:ID"`
	Weight                float64
	Bodyfat               float64
	CalculatedDailyIntake string
}

type Curriculum struct {
	gorm.Model
	ID   uint `gorm:"primaryKey; autoIncrement"`
	Name string
	Type string
	Size string
	Time time.Time
}

type VIPCard struct {
	gorm.Model
	ID         uint `gorm:"primaryKey; autoIncrement"`
	Balance    float64
	Due        time.Time
	CustomerId uint `gorm:"foreignKey:ID"`
}

type CurriculumCustomer struct {
	gorm.Model
	ID         uint       `gorm:"primaryKey; autoIncrement"`
	Customer   uint   `gorm:"foreignKey: ID"`
	Curriculum uint `gorm:"foreignKey: ID"`
	Attendance uint
}

type CurriculumTrainer struct {
	gorm.Model
	ID         uint       `gorm:"primaryKey; autoIncrement"`
	Trainer    uint    `gorm:"foreignKey: ID"`
	Curriculum uint `gorm:"foreignKey: ID"`
	LastClass  time.Time
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

type MaintainerEquipment struct {
	gorm.Model
	ID                uint       `gorm:"primaryKey; autoIncrement"`
	MaintainerOfEquip uint `gorm:"foreignKey: ID"`
	Equipment         uint  `gorm:"foreignKey: ID"`
	LastMaintain      time.Time
}

type FoodRank struct {
	gorm.Model
	ID         uint `gorm:"primaryKey; autoIncrement"`
	Food       uint `gorm:"foreignKey: ID"`
	Rank       uint
	Desciption string
}

type DailyIntakeFood struct {
	gorm.Model
	ID          uint        `gorm:"primaryKey; autoIncrement"`
	Food        uint        `gorm:"foreignKey: ID"`
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
	db.AutoMigrate(&Trainer{})
	db.AutoMigrate(&Manager{})
	db.AutoMigrate(&Maintainer{})
	db.AutoMigrate(&Equipment{})
	db.AutoMigrate(&Customer{})
	db.AutoMigrate(&Curriculum{})
	db.AutoMigrate(&VIPCard{})
	db.AutoMigrate(&CurriculumCustomer{})
	db.AutoMigrate(&CurriculumTrainer{})
	db.AutoMigrate(&Food{})
	db.AutoMigrate(&Nutrition{})
	db.AutoMigrate(&DailyIntake{})
	db.AutoMigrate(&MaintainerEquipment{})
	db.AutoMigrate(&FoodRank{})
	db.AutoMigrate(&DailyIntakeFood{})


	// Create

	// Read

	// Update - update multiple fields

}
