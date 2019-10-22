package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"../config"
	"github.com/gin-gonic/gin"
)

func GetDogs(c *gin.Context) {
	files, err := ioutil.ReadDir(config.DOGIMAGES)

	if err != nil {
		log.Fatal(err)
	}

	//generate a random number to serve random image
	s := rand.NewSource(time.Now().Unix())
	// initialize local pseudorandom generator
	r := rand.New(s)
	ranNum := r.Intn(len(files))
	filename := files[ranNum].Name()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf("http://%s%s/dogs/%s", config.HOST, config.PORT, filename))
	}
}
