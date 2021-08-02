package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mission-2/config"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Movie []Movie `json:"Search"`
}

type Movie struct {
	ImdbID string `json:"imdbID"`
	Title  string `json:"Title"`
	Year   int    `json:"Year"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type MovieDetail struct {
	ImdbID     string `json:"imdbID"`
	Title      string `json:"Title"`
	Year       int    `json:"Year"`
	Type       string `json:"Type"`
	Poster     string `json:"Poster"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Country    string `json:"Country"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
}

type ResponseDataMovieList struct {
	Data  []Movie
	Total int
}

func GetMovie(c *gin.Context) {
	search := c.Query("searchword")
	page := c.Query("pagination")
	response := gin.H{
		config.MESSAGE_RESPONSE_INDEX: config.UNKNOWN_ERROR_MESSAGE,
		config.DATA_RESPONSE_INDEX:    []interface{}{},
	}
	if len(search) == 0 || len(page) == 0 {
		log.Printf(config.ERROR_PARSING_REQ_DATA_TEMPLATE)
		response[config.MESSAGE_RESPONSE_INDEX] = config.REQ_PARAM_ERROR_MESSAGE
		c.JSON(http.StatusBadRequest, response)
		return
	}
	api_url := config.API_URL + "?apikey=" + config.API_KEY + "&s=" + search + "&page=" + page

	responseApi, err := http.Get(api_url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(responseApi.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	movielist := ResponseDataMovieList{responseObject.Movie, len(responseObject.Movie)}
	// data, err := json.Marshal(movielist)

	if err != nil {
		log.Printf(config.ERROR_PARSING_REQ_DATA_TEMPLATE)
		response[config.MESSAGE_RESPONSE_INDEX] = config.INTERNAL_SERVER_ERROR_MESSAGE
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response[config.MESSAGE_RESPONSE_INDEX] = config.SUCCESS_MESSAGE
	response[config.DATA_RESPONSE_INDEX] = movielist
	c.JSON(http.StatusOK, response)
}

func GetMovieDetail(c *gin.Context) {
	movieID := c.Param("id")
	println(movieID)

	response := gin.H{
		config.MESSAGE_RESPONSE_INDEX: config.UNKNOWN_ERROR_MESSAGE,
		config.DATA_RESPONSE_INDEX:    []interface{}{},
	}
	if len(movieID) == 0 || movieID == ":" {
		log.Printf(config.ERROR_PARSING_REQ_DATA_TEMPLATE)
		response[config.MESSAGE_RESPONSE_INDEX] = config.REQ_PARAM_ERROR_MESSAGE
		c.JSON(http.StatusBadRequest, response)
		return
	}

	api_url := config.API_URL + "?apikey=" + config.API_KEY + "&i=" + movieID

	responseApi, err := http.Get(api_url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(responseApi.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	movieDetail := ResponseDataMovieList{responseObject.Movie, len(responseObject.Movie)}

	if err != nil {
		log.Printf(config.ERROR_PARSING_REQ_DATA_TEMPLATE)
		response[config.MESSAGE_RESPONSE_INDEX] = config.INTERNAL_SERVER_ERROR_MESSAGE
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response[config.MESSAGE_RESPONSE_INDEX] = config.SUCCESS_MESSAGE
	response[config.DATA_RESPONSE_INDEX] = movieDetail
	c.JSON(http.StatusOK, response)

}
