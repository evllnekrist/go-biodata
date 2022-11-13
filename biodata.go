package main

import (
	"fmt"
    "strings"
	"time"
	"math/rand"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PersonAttr struct {
	ID int
	Name string
	Email string
	Addr string
	Occup string
	Reason string
}
type PersonThink interface {
	jobReason() string
}
var people = []interface{}{
	"Thomas Darmawan",
	"Tiara Cassandra Pvee",
	"Raden Putra Andras Tjokro",
	"Tiara Sadien",
	"Ardi Prakoso",
	"Hendra Prakoso",
	"Odin Thomassen",
	"Ruesantri Tiara",
}
var reasons = []interface{}{
	"passion",
	"carry on the family mandate",
	"gifted",
	"just happen",
	"skilled",
	"choose to",
	"nature bless",
}
var people_occup = []PersonAttr{
	{Email:"thomas@gmail.com",Addr:"Jakarta",Occup:"Lecturer"},
	{Email:"pvee@gmail.com",Addr:"Bali",Occup:"Instructor"},
	{Email:"tjokro98@gmail.com",Addr:"Solo",Occup:"Cook Chief"},
	{Email:"tiara@gmail.com",Addr:"Tangerang",Occup:"Dancer"},
	{Email:"ardi@gmail.com",Addr:"Ternate",Occup:"Entrepreneur"},
	{Email:"hendra@gmail.com",Addr:"Padang",Occup:"Merchant"},
	{Email:"odin@gmail.com",Addr:"BauBau",Occup:"Mediator"},
	{Email:"rue@gmail.com",Addr:"Sampit",Occup:"Entrepreneur"},
}

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("*.html")
	// --start------------------------------------------------TEST route
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	router.GET("/myBiodata", getMyBiodata)
	router.GET("/searchBiodata", getSearchedBiodata)
	// --end--------------------------------------------------TEST route
	router.Run("localhost:8080")
}

func getMyBiodata(c *gin.Context) {
	var temp_found PersonAttr
	input := c.Query("email")
	for i, v := range people_occup {
		if(strings.ToLower(v.Email) == strings.ToLower(input)){
			fmt.Println(people_occup[i])
			temp_found = people_occup[i] 
			temp_found.ID = i
			temp_found.Name = people[i].(string)
			temp_found.Reason = temp_found.jobReason()
			break
		}
	}
	c.HTML(http.StatusOK, "myprofile.html", gin.H{
		"content": temp_found,
	})
}

func getSearchedBiodata(c *gin.Context) {
	var people_found []PersonAttr 
	var temp_found PersonAttr
	input := c.Query("name")
	for i, v := range people {
		if(strings.Contains(strings.ToLower(v.(string)), strings.ToLower(input))){
			fmt.Println(people_occup[i])
			temp_found = people_occup[i] 
			temp_found.ID = i
			temp_found.Name = v.(string)
			temp_found.Reason = temp_found.jobReason()
			people_found = append(people_found, temp_found)
		}
	}
	c.HTML(http.StatusOK, "biodata.html", gin.H{
		"search": input,
		"content": people_found,
	})
}

func (p PersonAttr) jobReason() string {
	rand.Seed(time.Now().UnixNano()) 
	return p.Name+" reason is "+reasons[rand.Intn(len(reasons))].(string)
}
