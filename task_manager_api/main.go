package main 
import "github.com/gin-gonic/gin"
import "fmt"

func main(){
	fmt.Println("Task Manager API")
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context){
		context.JSON(200, gin.H{
			"message": "pongs	",
			"message2": "pongs",

		})
	})
	router.Run()
}