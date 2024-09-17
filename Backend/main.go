package main
import (
	"Backend/config"
	// "Backend/routes"
	// "Backend/utils"
	"os"
	// "log"
	"fmt"
)
func main(){
	config.InitDB()

	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
	fmt.Println(port)

    // if err := r.Run(":" + port); err != nil {
    //     log.Fatalf("Failed to run server: %v", err)
    // } else {
    //     log.Printf("Server is running on port %s", port)
    // }
}