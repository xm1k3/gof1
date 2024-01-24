package models

import (
	"time"

	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model
	DriverID    int       `gorm:"column:id" gorm:"primary_key" csv:"driverId"`
	DriverRef   string    `gorm:"column:driverRef" csv:"driverRef"`
	Number      string    `gorm:"column:number" csv:"number"`
	Code        string    `gorm:"column:code" csv:"code"`
	Forename    string    `gorm:"column:forename" csv:"forename"`
	Surname     string    `gorm:"column:surname" csv:"surname"`
	DOB         time.Time `gorm:"column:dob" csv:"dob"`
	Nationality string    `gorm:"column:nationality" csv:"nationality"`
	URL         string    `gorm:"column:url" csv:"url"`
}

type DriverStanding struct {
	DriverID int     `json:"driverId"`
	Forename string  `json:"forename"`
	Surname  string  `json:"surname"`
	Points   float64 `json:"points"`
}

type Constructor struct {
	gorm.Model
	ConstructorID  int    `gorm:"column:id" gorm:"primary_key" csv:"constructorId"`
	ConstructorRef string `gorm:"column:constructorRef" csv:"constructorRef"`
	Name           string `gorm:"column:name" csv:"name"`
	Nationality    string `gorm:"column:nationality" csv:"nationality"`
	URL            string `gorm:"column:url" csv:"url"`
}

type ConstructorStanding struct {
	ConstructorID int     `json:"constructorId"`
	Name          string  `json:"name"`
	Points        float64 `json:"points"`
}

type Circuit struct {
	gorm.Model
	CircuitID  int    `gorm:"column:id" gorm:"primary_key" csv:"circuitId"`
	CircuitRef string `gorm:"column:circuitRef" csv:"circuitRef"`
	Name       string `gorm:"column:name" csv:"name"`
	Location   string `gorm:"column:location" csv:"location"`
	Country    string `gorm:"column:country" csv:"country"`
	URL        string `gorm:"column:url" csv:"url"`
}

type Race struct {
	gorm.Model
	RaceID     int        `gorm:"column:id" gorm:"primary_key" csv:"raceId"`
	Year       int        `gorm:"column:year" csv:"year"`
	Round      int        `gorm:"column:round" csv:"round"`
	CircuitID  int        `gorm:"column:circuitId" csv:"circuitId"`
	Name       string     `gorm:"column:name" csv:"name"`
	Date       time.Time  `gorm:"column:date" csv:"date"`
	Time       string     `gorm:"column:time" csv:"time"`
	URL        string     `gorm:"column:url" csv:"url"`
	QualiDate  *time.Time `gorm:"column:quali_date" csv:"quali_date"`
	QualiTime  string     `gorm:"column:quali_time" csv:"quali_time"`
	SprintDate *time.Time `gorm:"column:sprint_date" csv:"sprint_date"`
	SprintTime string     `gorm:"column:sprint_time" csv:"sprint_time"`
}

type Result struct {
	gorm.Model
	ResultID        int     `gorm:"column:id;primary_key" csv:"resultId"`
	RaceID          int     `gorm:"column:raceId" csv:"raceId"`
	DriverID        int     `gorm:"column:driverId" csv:"driverId"`
	ConstructorID   int     `gorm:"column:constructorId" csv:"constructorId"`
	Number          int     `gorm:"column:number" csv:"number"`
	Grid            int     `gorm:"column:grid" csv:"grid"`
	Position        int     `gorm:"column:position" csv:"position"`
	PositionText    string  `gorm:"column:positionText" csv:"positionText"`
	PositionOrder   int     `gorm:"column:positionOrder" csv:"positionOrder"`
	Points          float64 `gorm:"column:points" csv:"points"`
	Laps            int     `gorm:"column:laps" csv:"laps"`
	Time            string  `gorm:"column:time" csv:"time"`
	Milliseconds    int     `gorm:"column:milliseconds" csv:"milliseconds"`
	FastestLap      int     `gorm:"column:fastestLap" csv:"fastestLap"`
	Rank            int     `gorm:"column:rank" csv:"rank"`
	FastestLapTime  string  `gorm:"column:fastestLapTime" csv:"fastestLapTime"`
	FastestLapSpeed string  `gorm:"column:fastestLapSpeed" csv:"fastestLapSpeed"`
	StatusID        int     `gorm:"column:statusId" csv:"statusId"`
}
