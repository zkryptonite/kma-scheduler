package repository

import (
	"backend/entity"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClassRepository struct {
	Db *mongo.Database
}

func NewClassRepository(db *mongo.Database) IClassRepository {
	return &ClassRepository{Db: db}
}

func (mongo *ClassRepository) FindClass(className string) (entity.Class, error) {
	class := entity.Class{}
	singleRes := mongo.Db.Collection("Class").FindOne(context.Background(), bson.M{"class_name": className})
	err := singleRes.Decode(&class)
	if err != nil {
		return class, err
	}

	return class, nil
}

func (mongo *ClassRepository) GetScheduleOfClassesByDay(classNames []string, date string) ([]entity.Class, error) {
	var (
		match = bson.M{
			"$match": bson.M{
				"$and": []bson.M{
					{
						"class_name": bson.M{"$in": classNames},
					},
					{
						fmt.Sprintf("%s.%s", "schedule", date): bson.M{"$exists": true},
					},
				},
			},
		}
		project = bson.M{
			"$project": bson.M{
				"class_name":                           1,
				"subject_name":                         1,
				fmt.Sprintf("%s.%s", "schedule", date): 1,
				"lecturer":                             1,
				"room":                                 1,
			},
		}
		query = []bson.M{match, project}
		cls   []entity.Class
	)

	cursor, err := mongo.Db.Collection("Class").Aggregate(context.Background(), query)
	if err != nil {
		return cls, err
	}

	if err = cursor.All(context.Background(), &cls); err != nil {
		return cls, err
	}
	return cls, nil
}

func (mongo *ClassRepository) Insert(cls ...entity.Class) error {
	var documents []interface{}

	for _, cl := range cls {
		if _, err := mongo.FindClass(cl.ClassName); err != nil {
			documents = append(documents, cl)
		}
	}

	_, err := mongo.Db.Collection("Class").InsertMany(context.Background(), documents)
	if err != nil {
		return err
	}
	return nil
}
