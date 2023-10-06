package store

import (
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
)

type Template struct {
    ID           interface{} `json:"id"`
    Name         string      `json:"name" binding:"required"`
    Value        string      `json:"value" binding:"required"`
    NpmName      string      `json:"npmName" binding:"required"`
    Version      string      `json:"version" binding:"required"`
    Ignore       []string    `json:"ignore" binding:"required"`
    ForceInstall bool        `json:"forceInstall"`
}

func CreateTemplate(ctx *gin.Context, t *Template) (interface{}, error) {
    collection := db.mongo.Collection("template")
    result, err := collection.InsertOne(ctx, bson.D{
        {"name", t.Name},
        {"value", t.Value},
        {"npmName", t.NpmName},
        {"version", t.Version},
        {"ignore", t.Ignore},
        {"forceInstall", t.ForceInstall},
    })
    if err != nil {
        return "", err
    }
    return result.InsertedID, nil
}

func TemplateList(ctx *gin.Context) ([]*Template, int64, error) {
    collection := db.mongo.Collection("template")
    var list []*Template
    count, err := collection.CountDocuments(ctx, bson.D{})
    if err != nil {
        return nil, 0, err
    }
    cur, err := collection.Find(ctx, bson.D{})
    if err != nil {
        return nil, 0, err
    }
    defer cur.Close(ctx)
    for cur.Next(ctx) {
        var tmp Template
        if err := cur.Decode(&tmp); err != nil {
            return nil, 0, err
        }
        list = append(list, &tmp)
    }
    if err := cur.Err(); err != nil {
        return nil, 0, err
    }
    return list, count, nil
}

func UpdateTemplate(c *gin.Context, t *Template) (interface{}, error) {
    collection := db.mongo.Collection("template")
    id, err := collection.UpdateOne(c, bson.D{{"value", t.Value}}, bson.D{{"$set", bson.D{
        {"name", t.Name},
        {"value", t.Value},
        {"npmName", t.NpmName},
        {"version", t.Version},
        {"ignore", t.Ignore},
        {"forceInstall", t.ForceInstall},
    }}})
    if err != nil {
        return "", err
    }
    return id, nil
}

func DeleteTemplate(c *gin.Context, value string) (int64, error) {
    collection := db.mongo.Collection("template")
    one, err := collection.DeleteOne(c, bson.D{{"value", value}})
    if err != nil {
        return 0, err
    }
    return one.DeletedCount, nil
}
