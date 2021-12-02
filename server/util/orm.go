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
	Employee      Employee `gorm:"FOREIGNKEY:ID"`
}

type Manager struct {
	gorm.Model
	ID          uint `gorm:"primaryKey; autoIncrement"`
	RegisterCnt uint
	TotalIncome float64
	Employee    Employee `gorm:"FOREIGNKEY:ID"`
}

type Maintainer struct {
	gorm.Model
	ID       uint `gorm:"primaryKey; autoIncrement"`
	Type     string
	Employee Employee `gorm:"FOREIGNKEY:ID"`
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
	ID           uint       `gorm:"primaryKey; autoIncrement"`
	Maintainer   Maintainer `gorm:"FOREIGNKEY:ID"`
	Equipment    Equipment  `gorm:"FOREIGNKEY:ID"`
	MaintainTime time.Time
	Description  string
}

type Customer struct {
	gorm.Model
	ID                    uint `gorm:"primaryKey; autoIncrement"`
	UserName              string
	PassWord              string
	RegisterTime          time.Time
	ManagerOfCustomer     Manager `gorm:"FOREIGNKEY:ID"`
	Weight                float64
	BodyFat               float64
	CalculatedDailyIntake string
}

type Curriculum struct {
	gorm.Model
	ID    uint `gorm:"primaryKey; autoIncrement"`
	Name  string
	Type  string
	Coach Coach `gorm:"FOREIGNKEY:ID"`
}

type VIPCard struct {
	gorm.Model
	ID         uint `gorm:"primaryKey; autoIncrement"`
	Balance    float64
	Due        time.Time
	Level      uint
	TotalSpend float64
	Customer   Customer `gorm:"FOREIGNKEY:ID"`
}

type CurriculumCustomer struct {
	gorm.Model
	ID         uint       `gorm:"primaryKey; autoIncrement"`
	Customer   Customer   `gorm:"FOREIGNKEY:ID"`
	Curriculum Curriculum `gorm:"FOREIGNKEY:ID"`
	Attendance uint
}

type CurriculumRecord struct {
	gorm.Model
	ID             uint       `gorm:"primaryKey; autoIncrement"`
	Customer       Customer   `gorm:"FOREIGNKEY:ID"`
	Curriculum     Curriculum `gorm:"FOREIGNKEY:ID"`
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
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Name     string
	Unit     float64
	Calories float64
	Brand    string
}

type Nutrition struct {
	gorm.Model
	ID     uint `gorm:"primaryKey; autoIncrement"`
	Type   string
	Amount float64
	Unit   float64
	Food   Food `gorm:"FOREIGNKEY:ID"`
}

type DailyIntake struct {
	gorm.Model
	ID       uint `gorm:"primaryKey; autoIncrement"`
	Date     time.Time
	Customer Customer `gorm:"FOREIGNKEY:ID"`
	MealType string
}

type FoodRank struct {
	gorm.Model
	ID          uint `gorm:"primaryKey; autoIncrement"`
	Food        Food `gorm:"foreignKey:ID"`
	Rank        uint
	Description string
}

type DailyIntakeFood struct {
	gorm.Model
	ID          uint        `gorm:"primaryKey; autoIncrement"`
	Food        Food        `gorm:"foreignKey:ID"`
	DailyIntake DailyIntake `gorm:"foreignKey:ID"`
	Amount      uint
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Employee{}, &Coach{}, &Manager{}, &Maintainer{}, &Equipment{}, &MaintainerEquipment{}, &Customer{}, &Curriculum{}, &VIPCard{}, &CurriculumCustomer{}, &CurriculumRecord{}, &CalculatedIntake{}, &Food{}, &Nutrition{}, &DailyIntake{}, &FoodRank{}, &DailyIntakeFood{}, &CalculatedIntake{}, &CurriculumRecord{})
	if err != nil {
		return
	}

	// Create

	// Read

	// Update - update multiple fields

}
