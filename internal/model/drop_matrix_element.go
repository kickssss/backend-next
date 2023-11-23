package model

import (
	"time"

	"github.com/uptrace/bun"
)

type DropMatrixElement struct {
	bun.BaseModel `bun:"drop_matrix_elements,alias:dme"`

	ElementID       int         `bun:",pk,autoincrement" json:"id"`
	StageID         int         `json:"stageId"`
	ItemID          int         `json:"itemId"`
	StartTime       *time.Time  `json:"startTime"`
	EndTime         *time.Time  `json:"endTime"`
	DayNum          int         `json:"dayNum"`
	Quantity        int         `json:"quantity"`
	Times           int         `json:"times"`
	QuantityBuckets map[int]int `bun:"type:jsonb" json:"quantityBuckets"`
	Server          string      `json:"server"`
	SourceCategory  string      `json:"sourceCategory"` // sourceCategory can be: "automated", "manual", "all"

	RangeID int `bun:"-" json:"-"`
	// TimeRange field is for those elements whose time range is not saved in DB, but a customized one
	TimeRange *TimeRange `bun:"-" json:"-"`
}
