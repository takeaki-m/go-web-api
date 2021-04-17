package main
import (
	// import mapping moculed
	"go-gin-api/mappings"
)

func main() {
	mappings.CreateUrlMappings()
	mappings.Router.Run(":8080")
}