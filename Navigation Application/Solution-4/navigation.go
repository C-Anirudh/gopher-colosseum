// Author: teknas
package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/info", info)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
	<title>teknas's Page</title>
	<style>
	body {
		background-image: url('https://images.pexels.com/photos/1591447/pexels-photo-1591447.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500'); 
		background-repeat: no-repeat;
		background-size: 100% 600%;
	}
	h1{
		color:white;
	}
	.container{
		text-align:center;
		padding-top:20px;
		font-size:25px;
	}
</style>
</head>
<body>
<div class="container">
<h1>Welcome</h1>
<button onclick="window.location.href = '/';">Home</button>
<button onclick="window.location.href = '/about';">About Me</button>
<button onclick="window.location.href = '/contact';">Contact</button>
<button onclick="window.location.href = '/info';">Info</button>
</div>
</body>
</html>`)
}

func about(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
	<title>SITE</title>
</head>
<body>
<h1>ABOUT MYSELF</h1>
<p>A Third year CSE Undergrad studying in Amrita School of Engineering. I am a Web Developer.</p>
<button onclick="window.location.href = '/';">Home</button>
</body>
</html>`)
}

func contact(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
	<title>SITE</title>
	<style>
	body {
		background-image: url('https://images.pexels.com/photos/1591447/pexels-photo-1591447.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500'); 
		background-repeat: no-repeat;
		background-size: 100% 600%;
	}
	h1{
		color:white;
	}
</style>
</head>
<body>
<h1>To Contact Me</h1>
<button onclick="window.location.href = 'https://github.com/teknas07';">Github</button>
<button onclick="window.location.href = '/';">Home</button>
</body>
</html>`)
}

func info(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<!DOCTYPE html>
	<html>
	<head>
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
	<style>
	.title {
	  color: black;
	  font-size: 18px;
	}
	</style>
	</head>
	<body>
	
	<h2 style="text-align:center">Profile Card</h2>
	
	<div class="container">
	  <h1>teknas</h1>
	  <p class="title">WEB DEVELOPER</p>
	  <p>Amrita university University</p>
	  <p><button>Contact</button></p>
	</div>
	<button onclick="window.location.href = '/';">Home</button>

	</body>
	</html>`)
}
