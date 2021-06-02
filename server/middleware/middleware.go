package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func createDBInstance() {
	// DB connection string
	connectionString := os.Getenv("DB_URI")

	// Database Name
	dbName := os.Getenv("DB_NAME")

	// Collection name
	collName := os.Getenv("DB_COLLECTION_NAME")

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTask()

	json.NewEncoder(w).Encode(payload)
	//json.Marshal(&payload) --> another chance to encode the string
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.ToDoList
	//var test1 = models.SubNumber
	fmt.Println("r.Body: ", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&task)
	fmt.Printf("task: %+v\n", task)

	insertOneTask(task)
	json.NewEncoder(w).Encode(task)

}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	//json.Marshal(&params) --> another chance to encode the string
}

func TaskUndo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	//json.Marshal(&params) --> another chance to encode the string
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	//json.Marshal(&params) --> another chance the encode the string
}

func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	count := deleteAllTask()
	json.NewEncoder(w).Encode(count)
	//json.Marchal(params) --> another chance the encode the string

}

func AddSubTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contexxt-Type", "application/x-www-form-urlencoder")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}

	//for '$each'
	addsales := []models.SubNumber{models.SubNumber{
		Numberone:   1,
		Numbertwo:   2,
		Numberthree: 3,
		Substring:   "nono",
	}}

	//not for '$each'
	/*
		addsales := models.SubNumber{
			Numberone: 1,
			Numbertwo: 2,
			Numberthree: 3,
			Substring: "nono"
		}
	*/
	/*
		'$addToSet' is to add the new array if anything is different,
		otherwise; you can not show the same array
		subadd := bson.M{"$addToSet: bson.M{"subadd":addsales}
	*/
	subadd := bson.M{"$addToSet": bson.M{"subtest": bson.M{"$each": addsales}}}

	result, err := collection.UpdateOne(context.Background(), filter, subadd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
	json.NewEncoder(w).Encode(params["id"])

}
func getAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M

	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		//fmt.Printf("show result: %+v", result)
		if e != nil {
			log.Fatal(e)
		}
		//fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// Insert one task in the DB
func insertOneTask(task models.ToDoList) {
	fmt.Println("subtask: ", task)

	insertResult, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

func taskComplete(task string) {
	fmt.Println("Complete: " + task)

	/*
		addsales := []models.SubNumber{models.SubNumber{
			Numberone:   1,
			Numbertwo:   2,
			Numberthree: 3,
			Substring:   "nono",
		}}
		addsales := models.SubNumber{
			Numberone:   1,
			Numbertwo:   2,
			Numberthree: 3,
			Substring:   "hello",
		}
	*/

	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	//update document must contain key beginning with '$' ref:https://www.itread01.com/content/1541224023.html

	update := bson.M{"$set": bson.M{"status": true}}
	//adddd := bson.M{"$push": bson.M{"subtest": bson.M{"$each": addsales, "$slice": 1}}}
	//adddd := bson.M{"$addToSet": bson.M{"subtest": bson.M{"$each": addsales}}}
	//adddd := bson.M{"$addToSet": bson.M{"subtest": addsales}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	//result, err := collection.UpdateOne(context.Background(), filter, adddd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

func undoTask(task string) {
	fmt.Println("Undo: " + task)

	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

func deleteOneTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	//delete the id is actually delete this row
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}

// delete all the tasks from the DB
func deleteAllTask() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
	return d.DeletedCount
}
