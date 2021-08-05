package gif

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/thedevsaddam/gojsonq"
)

// GetGif will query the Gpihy API for the random GIF and will query the JSON
// response for the original URL of the GIF
func GetGif() {

	const (
		baseURL = "https://api.giphy.com/v1/gifs/random"
		usage   = `Please provide the API key, tag of a GIF, and the optional rating(g, pg, pg-13, r).
	The API key can be stored in a .env file as GIPHY_API_KEY=VALUE
	
	[KEY] [TAG] [RATING]
	`
	)

	var rating, key string
	args := os.Args[1:]

	// loading the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// handle the amount of os arguments being passed to main.go
	switch len(args) {
	case 1:
		rating = "g"
		key = os.Getenv("GIPHY_API_KEY")
		if key == "" {
			log.Fatal(usage)
		}
	case 2:
		rating, key = "g", args[0]
	case 3:
		rating = args[2]
	default:
		log.Fatal(usage)
	}

	tag := args[0]

	// using the GET method to query the gif using the Giphy API
	req, err := http.NewRequest(
		http.MethodGet,
		baseURL+"?api_key="+key+"&tag="+tag+"&rating="+rating,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	// sends the request of the GET above
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// reading the entire body of the returned request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// quering the JSON response from the API to find the original URL of the GIF
	jq := gojsonq.New().JSONString(string(body)).Find("data.image_original_url")

	fmt.Printf("Downloading this random gif: %s\n", jq.(string))

	// downloading the faile and saving to random.gif
	err = DownloadFile("random.gif", jq.(string))
	if err != nil {
		log.Fatal(err)
	}
}
