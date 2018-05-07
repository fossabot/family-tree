package types

import (
	"time"
)

type Group struct {
	ID          uint64    `json:"id" bson:"_id,omitempty"`
	GroupName   string    `bson:"groupName" json:"groupName"`
	StartYear   int       `bson:"startYear" json:"startYear"`
	EndYear     int       `bson:"endYear" json:"endYear"`
	CreatedTime time.Time `bson:"createdTime" json:"createdTime"`
	Users       []User    `bson:"users" json:"users"`
	UserIDs     []uint64  `bson:"userIDs" json:"userIDs"`
}