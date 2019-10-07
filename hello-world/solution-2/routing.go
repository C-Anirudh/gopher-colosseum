package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
	<title>Chandu's Page</title>
	<style>
	body {
		background-image: url('https://timesofindia.indiatimes.com/thumb/msid-67151704,width-400,resizemode-4/67151704.jpg?imglength=215618'); 
		background-repeat: no-repeat;
		background-size: 100% 600%;
	}
</style>
</head>
<body>
<h1>Welcome</h1>
<button onclick="window.location.href = '/';">Home</button>
<button onclick="window.location.href = '/about';">About Me</button>
<button onclick="window.location.href = '/contact';">Contact</button>
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
<p>A Third year CSE Undergrad studying in Amrita School of Engineering. I am a Web Developer. My interests are  solving maths problems and reading about history . I update myself with the latest inventions and discoveries in the field of technology. I am sports lover. Cricket is my favourite game. Die hard fan of virat kohli and R.C.B  </p>
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
		background-image: url('https://timesofindia.indiatimes.com/thumb/msid-67151704,width-400,resizemode-4/67151704.jpg?imglength=215618'); 
		background-repeat: no-repeat;
		background-size: 100% 600%;
	}
</style>
</head>
<body>
<h1>To Contact Me</h1>
<button onclick="window.location.href = 'https://github.com/chanduch1999';">Github</button>
<button onclick="window.location.href = 'https://www.linkedin.com/in/chandu-ch-0a9b5218b/';">Linkedin</button>
<button onclick="window.location.href = '/';">Home</button>
</body>
</html>`)
}
