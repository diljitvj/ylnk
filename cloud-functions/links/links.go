package links

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go"
)

// GetLists : This is some function
func GetLists(w http.ResponseWriter, r *http.Request) {
	config := &firebase.Config{ProjectID: "ylnk-7b451"}

	app, err := firebase.NewApp(context.Background(), config)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Printf("Error in initializing firesbase app: %v", err)
		return
	}

	fireStoreClient, err := app.Firestore(context.Background())
	defer fireStoreClient.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Printf("Error in initializing firestore client: %v", err)
		return
	}

	result, err := fireStoreClient.Collection("links").Doc("slug1").Get(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Printf("Error in fetching data: %v", err)
		return
	}

	data := result.Data()

	linkJSON, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Printf("Error in parsing json: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(linkJSON)
	fmt.Printf("Document data: %#v\n", data)
}
