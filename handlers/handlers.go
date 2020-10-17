package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"../models"
)

var meetings *mongo.Collection
var participants *mongo.Collection

// Database Name
const dbName = "appointy"

// Collection name
const meetingsCollection = "meetings"
const participantsCollection = "participants"

func connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/confusion")
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

	meetings = client.Database(dbName).Collection(meetingsCollection)
	participants = client.Database(dbName).Collection(participantsCollection)
}

//ScheduleMeeting handler schedules a meeting for the user
func ScheduleMeeting(w http.ResponseWriter, req *http.Request) {
	connect()
	if req.Method == "POST" {

		var meeting models.Meeting

		err := json.NewDecoder(req.Body).Decode(&meeting)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := meetings.InsertOne(context.Background(), meeting)
		if err != nil {
			log.Fatal(err)
		}

		/*
			the below code finds the document just created using the insertedId and returns back the document
			in JSON format.
		*/
		err1 := meetings.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&meeting)
		if err1 != nil {
			log.Fatal(err1)
		}
		meetingJSON, _ := json.Marshal(meeting)
		w.Header().Set("Content-Type", "application/json")
		w.Write(meetingJSON)

		// now if we check the mongo shell for the document, we'll see that the document has been created with the timestamp

	} else {
		err := errors.New("Invalid operation")
		fmt.Println(w, "%s", err)
	}
}

// function to extract the params of the url
func getCode(r *http.Request, defaultCode int) (int, string) {
	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {
		return defaultCode, p[0]
	} else if len(p) > 1 {
		code, err := strconv.Atoi(p[0])
		if err == nil {
			return code, p[2]
		} else {
			return defaultCode, p[2]
		}
	} else {
		return defaultCode, ""
	}
}

//GetMeetingByID returns the meeting which has the same id as that present in the params of url
func GetMeetingByID(w http.ResponseWriter, req *http.Request) {
	connect()
	var meeting models.Meeting
	if req.Method == "GET" {
		_, param := getCode(req, 200)
		oid, _ := primitive.ObjectIDFromHex(param)

		err := meetings.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&meeting)
		meetingJSON, _ := json.Marshal(meeting)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(meetingJSON)

	} else {
		err := errors.New("Invalid operation")
		fmt.Println(w, "%s", err)
	}
}

func GetMeetingsWithinTimeRange(w http.ResponseWriter, req *http.Request) {}

func GetMeetingsOfParticipant(w http.ResponseWriter, req *http.Request) {

	/*
		I didn't get time to implement this endpoint however I would just like to talk about my approach if in case
		you think of considering it. I think this approach will surely work.
		- I will extract the query params from the url string
		- I will then find all the documents which has the given email as a participant
		- I'll then return all such documents
	*/
}
