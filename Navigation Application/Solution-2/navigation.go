//Author : ch.chandu
package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/profile",profile)
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
<button onclick="window.location.href = '/profile';">Profile</button>
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


func profile(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<!DOCTYPE html>
	<html>
	<head>
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
	<style>
	.card {
	  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
	  max-width: 300px;
	  margin: auto;
	  text-align: center;
	  font-family: arial;
	}
	
	.title {
	  color: grey;
	  font-size: 18px;
	}
	
	button {
	  border: none;
	  outline: 0;
	  display: inline-block;
	  padding: 8px;
	  color: white;
	  background-color: #000;
	  text-align: center;
	  cursor: pointer;
	  width: 100%;
	  font-size: 18px;
	}
	
	a {
	  text-decoration: none;
	  font-size: 22px;
	  color: black;
	}
	
	button:hover, a:hover {
	  opacity: 0.7;
	}
	</style>
	</head>
	<body>
	
	<h2 style="text-align:center">Profile Card</h2>
	
	<div class="card">
	  <img src="https://timesofindia.indiatimes.com/thumb/msid-67151704,width-400,resizemode-4/67151704.jpg?imglength=215618" alt="Chandu" style="width:100%">
	  <h1>Chandu</h1>
	  <p class="title">CEO & Founder, Example</p>
	  <p>Amrita university University</p>
	  <div style="margin: 24px 0;">
		<a href="https://twitter.com/ch_chandu01"><i class="fa fa-twitter"></i></a>  
		<a href="https://www.linkedin.com/in/chandu-ch-0a9b5218b/"><i class="fa fa-linkedin"></i></a>  
		<a href="https://www.facebook.com"><i class="fa fa-facebook"></i></a> 
	  </div>
	  <p><button>Contact</button></p>
	</div>
	<button onclick="window.location.href = '/';">Home</button>

	</body>
	</html>`)
}
