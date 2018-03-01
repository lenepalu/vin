package models

import "github.com/louisevanderlith/db"

type EngineLayout = int

const (
	Inline EngineLayout = iota
	V
	Rotary
	W
	Boxer
)

type FuelType = int

const (
	Petrol FuelType = iota
	Diesel
	Hybrid
	Electric
	LPG
)

type Induction = int

const (
	NA Induction = iota
	Turbo
	Supercharger
)

type Engine struct {
	db.Record
	Family            string
	Series            string
	Code              string
	Displacement      int
	FuelType          FuelType
	Layout            EngineLayout
	Cylinders         int
	Valvetrain        string
	ValvesPerCylinder int
	PowerKW           int
	PowerAt           int
	TorqueNm          int
	TorqueAt          int
	Induction         Induction
	StartYear         int
	EndYear           int
	Models            []*Model
}
