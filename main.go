package main

// https://godocs.io/go.mau.fi/whatsmeow#example-package

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// yuppieseaticetruck09

func NewReservation(c *gin.Context) {

	var newReservation Reservation

	// Call BindJson to bind recieved JSON to new Album

	err := c.BindJSON(&newReservation)

	if err != nil || newReservation.MainFirstName == "" {

		c.IndentedJSON(http.StatusBadRequest, "BAD DATA")
		// this return won't break the program
		return
	}

	c.IndentedJSON(http.StatusCreated, "OK")
}

func SendReservationToWhatsApp() {

}

func main() {
	fmt.Println("Running API on port 3020")

	router := gin.Default()

	router.POST("/new-reservation", NewReservation)

	router.Run("localhost:3020")
}
