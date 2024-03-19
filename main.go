package main

import (
	"fmt"
	"io"
	"net/http"
)

var counter int = 0
var imageURL string = "https://res-nightly.cloudinary.com/itai/image/fetch/w_2000/x_0/https://hips.hearstapps.com/hmg-prod/images/2023-mclaren-artura-101-1655218102.jpg"

func main() {

	http.HandleFunc("/myimage", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	counter++
	if counter%3 == 0 {
		streamImage(w, r)
		return
	}
	w.Header().Set("x-code-error", "async call")
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(423)

	// webdav := `<?xml version="1.0" encoding="utf-8" ?>
	// <D:error xmlns:D="DAV:">
	//   <D:lock-token-submitted>
	// 	<D:href>Generating asset in the background... Retry/reload later.<br>(HTTP 423)</D:href>
	//   </D:lock-token-submitted>
	// </D:error>`
	// fmt.Fprint(w, webdav)

	// fmt.Fprint(w, "<html><body>Generating...</body><script>setTimeout(function() {\n  location.reload();\n}, 2000);\n</script></html>")
	fmt.Fprint(w, "<html><body>Cloudinary is generating the asset in the background. Retry later.<br>(HTTP 423)</body></html>")

}

func streamImage(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get(imageURL)
	if err != nil {
		http.Error(w, "Failed to fetch image", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Image URL returned non-200 status code", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	// Copy the image stream to the response writer
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Failed to write image stream", http.StatusInternalServerError)
		return
	}
}
