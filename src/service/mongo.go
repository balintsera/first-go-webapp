package service

import mgo "gopkg.in/mgo.v2"

var mongoSession *mgo.Session
var collections map[string]*mgo.Collection
var initialized bool

// Create a connection to mongoDB
func connectToMongo() {
	var err error
	mongoSession, err = mgo.Dial("localhost:27017")
	if err != nil {
		panic("Can't connect to database")
	}
	mongoSession.SetMode(mgo.Monotonic, true)
}

// GetMongoCollection returns a collection specified in the name param
func GetMongoCollection(collectionName string) (collection *mgo.Collection) {
	// Initialize the connection if its not initialized
	if !initialized {
		connectToMongo()
	}

	// Make packagewise collections map
	if len(collections) < 1 {
		collections = make(map[string]*mgo.Collection)
	}

	// Check if the collection is already registered and register if not
	if _, ok := collections[collectionName]; !ok {
		collection = mongoSession.DB("testgo").C(collectionName)
		collections[collectionName] = collection
	}
	return collections[collectionName]
}
