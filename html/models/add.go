package models

type Add struct {
	FirstNumber  string `json:"firstnumber,omitempty" bson:"firstnumber"`
	SecondNumber string `json:"secondnumber,omitempty" bson:"secondnumber"`
	Sum          string `json:"sum,omitempty" bson:"sum"`
}
