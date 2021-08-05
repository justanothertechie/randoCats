package gif

import (
	"io"
	"net/http"
	"os"
)

// WebHandler will handel the request and serve the random.gif
func WebHandler(w http.ResponseWriter, r *http.Request) {
	cat, err := os.Open("random.gif")
	if err != nil {
		panic(err)
	}

	defer cat.Close()
	w.Header().Set("Content-Type", "image/gif")
	io.Copy(w, cat)
}
