package db
import (
    "fmt"
    "context"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson/primitive"
)
type User struct{
    Name string `json:"name"`
    Password string `json:"password"`
    Role string `json:"role"`
    DB string `json:"db"`
}
func Mongo_Connect(uri string) (*mongo.Client, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, err
    }
    return client, nil
}
func List_Collections(db string, connection *mongo.Client) ([]string, error) {
    database := connection.Database(db)
    collections, err := database.ListCollectionNames(context.TODO(), bson.D{})
    if err != nil {
        return nil, fmt.Errorf("error listing collections: %v", err)
    }
    return collections, nil
}
func List_Databases(connection *mongo.Client) ([]string, error){
    databases, err := connection.ListDatabaseNames(context.TODO(), bson.D{})
    if err != nil {
        return nil, fmt.Errorf("error listing databases: %v", err)
    }
    return databases, nil
}
func Get_Document_by_id(id string, db string , col string, connection *mongo.Client)(*primitive.M ,error){
    database := connection.Database(db)
    collection := database.Collection(col)
    // convert the string obj id into mongo id
    obj_id , err := primitive.ObjectIDFromHex(id)
    if err != nil {
        fmt.Println(err)
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    // filter
    filter := bson.M{"_id": obj_id}
    //var output bson.M
    var result primitive.M
    err = collection.FindOne(ctx, filter).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, fmt.Errorf("error fetching document: %v", err)
        }
        return nil , fmt.Errorf("Error fetching document: %v", err)
    }
    return &result, nil
}
// have to create a function for check the user name is there or not
func Get_data_by_username(user_name  string, connection *mongo.Client)(User, error){
    database := connection.Database("user-data")
    collection := database.Collection("records")
    filter := bson.M{"name": user_name}
    var output User
    err := collection.FindOne(context.TODO(), filter).Decode(&output)
    if err != nil {
        return User{}, fmt.Errorf("error fetching document: %v", err)
    }
    return output, nil
}
