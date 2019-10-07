package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8090", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>SITE</title>

<style>
	body {
		background-image: url('https://tvup.media/wp-content/uploads/2019/07/0_HICLyAdNSIyT0ODU.jpg'); 
		background-repeat: no-repeat;
		background-size: 100% 300%;
	}

</style>
</head>

<body>
<h1>Welcome to Hariharan's page!!!</h1>
<h1>  <a href="/">HOME</a></h1>
<h1>  <a href="/about">ABOUT MYSELF</a></h1>
<h1>  <a href="/contact">CONTACT</a></h1>
</body>
</html>`)
}

func about(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>SITE</title>
<style>
	body{
		background-color: yellow;
	}
</style>

</head>
<body>
<h1>ABOUT MYSELF</h1>
<p>A third year CSE Undergrad studying in Amrita School of Engineering. I am basically an Android developer as well as a Web Developer. My interests include solving various math problems as well as photography. I love reading about the latest inventions and discoveries in the field of technology. Also I am sports lover. My favourite games include cricket,football,tennis and basketball. </p>
<h1>  <a href="/">HOME</a></h1>
<h1>  <a href="/about">ABOUT MYSELF</a></h1>
<h1>  <a href="/contact">CONTACT</a></h1>
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
		background-color: Orange; 
	}
</style>    
</head>
<body>
<h1>CONTACT ME ON!</h1>
<h3>  <a href = "https://github.com/harrycode007">Github</a></h3>
<h3>  <a href = "https://www.linkedin.com/in/hariharan-krishnamoorthy-934a29183/">Linkedin</a></h3> 
<h1>  <a href="/">HOME</a></h1>
<h1>  <a href="/about">ABOUT MYSELF</a></h1>
<h1>  <a href="/contact">CONTACT</a></h1>
</body>
</html>`)
}