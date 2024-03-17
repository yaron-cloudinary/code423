# Introduction
This is a POC that demonstrates how an image returning a 423 error, can also show a gracefull screen to the user, and even generate it in the background. 
It consists of a web server that returns 423 for the first two attempts, and an image for the third attempt.
It provides a webpage that shows an error image, as expceted.  However if the image url is tested directly from the webserver URL, it shows a graceful webpage indicating processing in the background. 

# Execution

1. Run the server:
```
go run main.go 
```


2. Open the website in the browser.  Watch the error.  Reload a few times, until you see the image.


3. Open the image url directly in the browser ( "http://localhost:8080/").   Note how it shows a graceful message when it gets a 423 error.  Reload a few times until the image is displayed.

