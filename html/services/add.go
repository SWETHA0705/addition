package services

import (
	"add/interfaces"
	"add/models"
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
)

type Add struct{
	Collection *mongo.Collection
	ctx context.Context
}

func InitService(collection *mongo.Collection,ctx context.Context)interfaces.Add{
	return &Add{collection,ctx}
}

func (a * Add)Add(m *models.Add)(string){
	num1,_:= strconv.Atoi(m.FirstNumber)
	num2,_:= strconv.Atoi(m.SecondNumber)

	Sum := num1 +num2
	m.Sum = strconv.Itoa(Sum)
	_,err:=a.Collection.InsertOne(a.ctx,m)
	if err != nil {
		return "nil"
	}
	return m.Sum
}