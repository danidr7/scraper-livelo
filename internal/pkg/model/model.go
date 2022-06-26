package model

type PartnerTreated struct {
	Code string `json:"code" bson:"code"`
	Name string `json:"name" bson:"name"`
	Link string `json:"link" bson:"link"`
}