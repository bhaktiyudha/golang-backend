package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestGetMovie(t *testing.T) {
	makeNewRequest := func(page string, search string) *httptest.ResponseRecorder {
		router := gin.Default()
		router.GET("/movie", func(c *gin.Context) {
			if c.Params != nil {
				search = c.Param("searchword")
				page = c.Param("pagination")
			}
		}, GetMovie)

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/movie", nil)
		req.Header.Set("Content-Type", "application/json")
		q := req.URL.Query()
		q.Add("searchword", search)
		q.Add("pagination", page)
		req.URL.RawQuery = q.Encode()

		fmt.Println(req.URL.String())
		req.URL.RawQuery = q.Encode()
		router.ServeHTTP(w, req)

		return w
	}

	res := makeNewRequest("", "")
	fmt.Println(res.Body)

	assert.Equal(t, res.Code, 400)

	res = makeNewRequest("1", "batman")
	fmt.Println(res.Body)

	assert.Equal(t, res.Code, 200)

	res.Flush()
}

func TestGetMovieDetail(t *testing.T) {
	makeNewRequest := func(id string) *httptest.ResponseRecorder {
		router := gin.Default()
		router.GET("/movie/:id", func(c *gin.Context) {
			if c.Params != nil {
				id = c.Param("id")
			}
		}, GetMovieDetail)

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/movie/"+id, nil)
		req.Header.Set("Content-Type", "application/json")
		q := req.URL.Query()
		req.URL.RawQuery = q.Encode()

		fmt.Println(req.URL.String())
		req.URL.RawQuery = q.Encode()
		fmt.Println(req)
		router.ServeHTTP(w, req)

		return w
	}

	res := makeNewRequest("")
	fmt.Println(res.Body)

	assert.Equal(t, res.Code, 404)

	res = makeNewRequest(":")
	fmt.Println(res.Body)

	assert.Equal(t, res.Code, 400)

	res = makeNewRequest(":tt0111161")
	fmt.Println(res.Body)

	assert.Equal(t, res.Code, 200)

	res.Flush()
}
