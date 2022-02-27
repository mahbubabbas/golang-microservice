package models

import (
	"context"
	"golang-microsvc/utils"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ClassName string `json:"class_name"`
	RollNo    int    `json:"roll_no"`
}

func (s *Student) Insert(inputStudent *Student) map[string]interface{} {
	coll := mDB.Collection(STUDENT_COLLECTION)

	doc := bson.D{
		{"first_name", inputStudent.FirstName},
		{"last_name", inputStudent.LastName},
		{"class_name", inputStudent.ClassName},
		{"roll_no", inputStudent.RollNo},
	}

	res, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return utils.ErrorResponse("Insert error")
	}

	return utils.SuccessResponse(res.InsertedID)
}

func (s *Student) InsertMany(inputStudent []Student) map[string]interface{} {
	coll := mDB.Collection(STUDENT_COLLECTION)

	docs := []interface{}{}

	for _, stud := range inputStudent {
		doc := bson.D{
			{"first_name", stud.FirstName},
			{"last_name", stud.LastName},
			{"class_name", stud.ClassName},
			{"roll_no", stud.RollNo},
		}

		docs = append(docs, doc)
	}

	res, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		return utils.ErrorResponse("Insert error")
	}

	return utils.SuccessResponse(res.InsertedIDs)
}

func (s *Student) FindStudents() map[string]interface{} {
	coll := mDB.Collection(STUDENT_COLLECTION)

	filter := bson.D{}

	cursor, err := coll.Find(
		context.TODO(),
		filter,
		options.Find().SetProjection(bson.M{"roll_no": 0, "_id": 0}),
	)

	if err != nil {
		return utils.ErrorResponse("Not able to fecth data")
	}

	resp := []bson.M{}
	cursor.All(context.TODO(), &resp)

	return utils.SuccessResponse(resp)
}

func (s *Student) Update(id string) map[string]interface{} {
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", oid}}

	coll := mDB.Collection(STUDENT_COLLECTION)
	doc := bson.M{}

	if s.FirstName != "" {
		doc["first_name"] = s.FirstName
	}

	if s.LastName != "" {
		doc["last_name"] = s.LastName
	}

	res, err := coll.UpdateOne(context.TODO(), filter, bson.D{{"$set", doc}})
	if err != nil {
		return utils.ErrorResponse("Unable to update")
	}

	return utils.SuccessResponse(res.ModifiedCount)
}

func (s *Student) UpdateMany(sArr []Student, ids string) map[string]interface{} {
	idArr := strings.Split(ids, ",")

	myids := []interface{}{}

	for _, id := range idArr {
		oid, _ := primitive.ObjectIDFromHex(id)
		myids = append(myids, oid)
	}

	filter := bson.D{{"_id", bson.D{{"$in", myids}}}}

	coll := mDB.Collection(STUDENT_COLLECTION)

	studentDocs := []interface{}{}

	for _, s := range sArr {
		doc := bson.M{}
		if s.FirstName != "" {
			doc["first_name"] = s.FirstName
		}
		if s.LastName != "" {
			doc["last_name"] = s.LastName
		}

		studentDocs = append(studentDocs, doc)
	}

	res, err := coll.UpdateMany(context.TODO(), filter, bson.D{{"$set", bson.D{{"first_name", "kakakk"}}}})
	if err != nil {
		return utils.ErrorResponse("Unable to update")
	}

	return utils.SuccessResponse(res.ModifiedCount)
}

func (s *Student) Delete(id string) map[string]interface{} {
	oid, _ := primitive.ObjectIDFromHex(id)	
	coll := mDB.Collection(STUDENT_COLLECTION)
	filter := bson.D{{"_id", oid}}

	res, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return utils.ErrorResponse("Delete error")
	}

	return utils.SuccessResponse(res.DeletedCount)
}

