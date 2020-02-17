package main

import (
	_ "app/docs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"utils"

	"repository"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// GetBeer godoc
// @Summary Get Beer
// @Description Get Beer for by beer_id
// @Tags beer
// @Param  id path string true "Beer Id"
// @Produce  json
// @Success 200
// @Router /api/getbeer/{id} [get]
func GetBeer(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := params["id"]
	err, beer := repository.GetBeer(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(beer)
}

// AddBeer godoc
// @Summary Add a Beer
// @Description add by json beer
// @Tags beers
// @Accept  json
// @Produce  json
// @Param beer body Beer true "Add Beer"
// @Success 200 {object} Beer
// @Router /api/addbeer [post]
func AddBeer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var beer repository.Beer
	_ = json.NewDecoder(request.Body).Decode(&beer)
	err, beer := repository.AddBeer(beer)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(beer)
}

// @title Sample API
// @version 1.0
// @description This is a sample serice for managing brewery
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	bucket := utils.Bucket()
	fmt.Println(bucket.Name())
	router := mux.NewRouter()

	//router.Path("/api/bucketinfo").Methods("GET").HandlerFunc(BucketInfo)
	router.Path("/api/getbeer/{id}").Methods("GET").HandlerFunc(GetBeer)
	router.Path("/api/addbeer").Methods("POST").HandlerFunc(AddBeer)

	//swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	fmt.Println("starting server...")
	// Listen on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))

	//config := utils.GetInstance()
	//fmt.Println(config.Connection)

}
