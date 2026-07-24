package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

var routes = map[string]string{
	"/":               "index.html",
	"/about/":         "about.html",
	"/politics/":      "politics.html",
	"/business/":      "business.html",
	"/philanthropy/":  "philanthropy.html",
	"/impact/":        "impact.html",
	"/news/":          "news.html",
	"/events/":        "events.html",
	"/contacts/":      "contacts.html",
	"/video-gallery/": "video-gallery.html",
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.Handle("/styles.css", http.FileServer(http.Dir(".")))
	mux.Handle("/inner.css", http.FileServer(http.Dir(".")))
	mux.Handle("/business.css", http.FileServer(http.Dir(".")))
	mux.Handle("/brand.css", http.FileServer(http.Dir(".")))
	mux.Handle("/images.css", http.FileServer(http.Dir(".")))
	mux.Handle("/route-gallery.css", http.FileServer(http.Dir(".")))
	mux.Handle("/script.js", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if file, ok := routes[r.URL.Path]; ok {
			servePage(w, file)
			return
		}
		http.NotFound(w, r)
	})
	log.Println("Serving at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func servePage(w http.ResponseWriter, file string) {
	page, err := os.ReadFile(file)
	if err != nil {
		http.NotFound(w, nil)
		return
	}
	page = []byte(strings.Replace(string(page), "</head>", `<link rel="icon" type="image/jpeg" href="/assets/images/logo.jpeg?v=1"><link rel="stylesheet" href="/brand.css"><link rel="stylesheet" href="/images.css"><link rel="stylesheet" href="/route-gallery.css"></head>`, 1))
	page = []byte(strings.Replace(string(page), "</main>", routeGallery(file)+`</main>`, 1))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(page)
}

func routeGallery(file string) string {
	galleries := map[string]string{
		"about.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>Leadership in focus</h2><p>Moments of public engagement, service and connection with the people of Benue South.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img2.jpeg" alt="Dr. Nelson Alapa at a public engagement" loading="eager"><figcaption>Public engagement</figcaption></figure><figure><img src="/assets/images/img11.jpeg" alt="Dr. Nelson Alapa meeting community leaders" loading="lazy"><figcaption>Community connection</figcaption></figure><figure><img src="/assets/images/img20.jpeg" alt="A public event in Benue South" loading="lazy"><figcaption>Service in action</figcaption></figure></div></div></section>`,
		"politics.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>Listening. Leading. Delivering.</h2><p>Leadership is strongest when it stays close to the people it serves.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img1.jpeg" alt="Dr. Nelson Alapa at a political gathering" loading="eager"><figcaption>Grassroots engagement</figcaption></figure><figure><img src="/assets/images/img20.jpeg" alt="Community event" loading="lazy"><figcaption>Community dialogue</figcaption></figure><figure><img src="/assets/images/img21.jpeg" alt="Public gathering" loading="lazy"><figcaption>Shared purpose</figcaption></figure></div></div></section>`,
		"business.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>Relationships that create opportunity</h2><p>Enterprise grows through trust, collaboration and a long-term commitment to people.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img6.jpeg" alt="Dr. Nelson Alapa at a public engagement" loading="eager"><figcaption>Leadership and enterprise</figcaption></figure><figure><img src="/assets/images/img11.jpeg" alt="Dr. Nelson Alapa in conversation" loading="lazy"><figcaption>Building partnerships</figcaption></figure><figure><img src="/assets/images/img2.jpeg" alt="Community event" loading="lazy"><figcaption>Connecting people</figcaption></figure></div></div></section>`,
		"philanthropy.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>Hope expressed through action</h2><p>Community support begins with presence, compassion and a commitment to restore dignity.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img3.jpeg" alt="Dr. Nelson Alapa at a community gathering" loading="eager"><figcaption>Compassion in action</figcaption></figure><figure><img src="/assets/images/img20.jpeg" alt="Community event" loading="lazy"><figcaption>Standing with communities</figcaption></figure><figure><img src="/assets/images/img21.jpeg" alt="Public engagement" loading="lazy"><figcaption>Hope for tomorrow</figcaption></figure></div></div></section>`,
		"impact.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>Progress people can see</h2><p>Programmes gain meaning when communities are present, heard and strengthened.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img6.jpeg" alt="Dr. Nelson Alapa at a public event" loading="eager"><figcaption>Community impact</figcaption></figure><figure><img src="/assets/images/img11.jpeg" alt="Community leadership engagement" loading="lazy"><figcaption>Shared progress</figcaption></figure><figure><img src="/assets/images/img22.jpeg" alt="Public community event" loading="lazy"><figcaption>People first</figcaption></figure></div></div></section>`,
		"news.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>From the field</h2><p>Official moments, public engagements and leadership stories in pictures.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img6.jpeg" alt="Dr. Nelson Alapa at an official gathering" loading="eager"><figcaption>Official engagement</figcaption></figure><figure><img src="/assets/images/img11.jpeg" alt="Dr. Nelson Alapa with leaders" loading="lazy"><figcaption>Leadership update</figcaption></figure><figure><img src="/assets/images/img20.jpeg" alt="Community event" loading="lazy"><figcaption>In the community</figcaption></figure></div></div></section>`,
		"events.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>Moments that bring people together</h2><p>Community gatherings, public conversations and events that strengthen shared purpose.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img20.jpeg" alt="Community celebration" loading="eager"><figcaption>Community gathering</figcaption></figure><figure><img src="/assets/images/img21.jpeg" alt="Public event" loading="lazy"><figcaption>Public occasion</figcaption></figure><figure><img src="/assets/images/img22.jpeg" alt="Community event" loading="lazy"><figcaption>Shared moments</figcaption></figure></div></div></section>`,
		"contacts.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>A leadership team that stays connected</h2><p>For invitations, media enquiries, partnership ideas and community feedback, the office is ready to listen.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img2.jpeg" alt="Dr. Nelson Alapa at a public event" loading="eager"><figcaption>Connect with the office</figcaption></figure><figure><img src="/assets/images/img11.jpeg" alt="Dr. Nelson Alapa in conversation" loading="lazy"><figcaption>Open dialogue</figcaption></figure><figure><img src="/assets/images/img6.jpeg" alt="Community gathering" loading="lazy"><figcaption>Community first</figcaption></figure></div></div></section>`,
	}
	return galleries[file]
}
