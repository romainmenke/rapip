package reader

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTransform(t *testing.T) {
	{
		const templateIndicatorB = `TEMPLATE
`
		t.Log(len(templateIndicatorB))
	}

	{
		b := []byte(`TEMPLATE
HTTP/1.1 200 OK
Connection: close
Cache-Control: max-age=0, private, must-revalidate
Content-Type: text/html; charset=utf-8
Date: Fri, 29 Sep 2017 20:29:44 GMT

<!DOCTYPE HTML>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1, minimal-ui"> <title>Cranky &amp; The Law</title> <link rel="stylesheet" type="text/css" href="/assets/css/bundle.min.css"> <script type="text/javascript" src="/assets/js/bundle.min.js"></script> <meta name="Description" content="Cranky &amp; The Law"> <link rel="canonical" href="https://crankyandthelaw.com"/> <meta property="og:url" content="https://crankyandthelaw.com"/> <meta property="og:type" content="website"/> <meta property="og:title" content="Cranky &amp; The Law"/> <meta property="og:description" content="Cranky &amp; The Law are two nieces from Antwerp, Belgium. Armed with a ukulele, a guitar or a keyboard (The Law) and a bass guitar (Cranky) they play original indie-pop songs."/> <meta property="og:image" content="https://crankyandthelaw.com/assets/img/cranky_and_the_law_share.jpg"/> <script type="application/ld+json"> { "@context": "http://schema.org", "@type": "Organization", "name": "Cranky & The Law", "url": "https://crankyandthelaw.com", "sameAs": [ "https://www.facebook.com/CrankyAndTheLaw" ], "logo": "https://crankyandthelaw.com/assets/img/cranky-and-the-law_logo_round.png" } </script> </head> <body class="home dark"> <div class="bg-wrapper"> <div class="bg-color"> </div> <div class="bg-sizer"> <img class="bg-img" src="/assets/img/cranky_and_the_law_bg.jpg" alt="Cranky &amp; The Law"> </div> </div> <nav> <div class="bg-color"></div> <a class="logo" href="/"></a> <div class="nav-buttom" onclick="toggleMenu()"></div> <div id="menu"> <script type="text/javascript">toggleMenu();</script> <a id="news" href="/news">News</a> <a id="about" href="/about">About</a> <a id="media" href="/media">Media</a> <a id="gigs" href="/gigs">Gigs</a> <a id="booking" href="/booking">Booking</a> <a id="press" href="/press">Press</a> </div> </nav> <div class="color-wrapper__inner"> <article class="content"> </article> </div> </body> <footer> <script> (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){ (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o), m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m) })(window,document,'script','https://www.google-analytics.com/analytics.js','ga'); ga('create', 'UA-93180050-1', 'auto'); ga('set', 'anonymizeIp', true); ga('send', 'pageview'); </script> </footer> </html>
`)

		t.Log(fmt.Sprintf("@%s@", string(b[9:])))
	}

	{
		f := "foo"
		b, err := json.Marshal(f)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(b))
	}

}
