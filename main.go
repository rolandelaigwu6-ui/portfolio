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
		"about.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>Leadership in focus</h2><p>A journey shaped by service, learning and an enduring connection to Benue South.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img2.jpeg" alt="Dr. Nelson Alapa at a public engagement" loading="eager"><figcaption>A leader rooted in community</figcaption></figure><figure><img src="/assets/images/img3.jpeg" alt="Dr. Nelson Alapa with fellow leaders" loading="lazy"><figcaption>Service through relationships</figcaption></figure><figure><img src="/assets/images/img4.jpeg" alt="Dr. Nelson Alapa at an official occasion" loading="lazy"><figcaption>Experience with purpose</figcaption></figure></div></div></section>`,
		"politics.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>Listening. Leading. Delivering.</h2><p>Leadership is strongest when it stays close to the people it serves.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img1.jpeg" alt="Dr. Nelson Alapa at a political gathering" loading="eager"><figcaption>Grassroots engagement</figcaption></figure><figure><img src="/assets/images/img5.jpeg" alt="Dr. Nelson Alapa at a public gathering" loading="lazy"><figcaption>Conversations that matter</figcaption></figure><figure><img src="/assets/images/img6.jpeg" alt="Dr. Nelson Alapa at an official event" loading="lazy"><figcaption>United by shared purpose</figcaption></figure></div></div></section>`,
		"business.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>Relationships that create opportunity</h2><p>Enterprise grows through trust, collaboration and a long-term commitment to people.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img7.jpeg" alt="Dr. Nelson Alapa at a public engagement" loading="eager"><figcaption>Leadership that builds trust</figcaption></figure><figure><img src="/assets/images/img8.jpeg" alt="Dr. Nelson Alapa meeting people" loading="lazy"><figcaption>Creating meaningful connections</figcaption></figure><figure><img src="/assets/images/img9.jpeg" alt="Dr. Nelson Alapa at an official occasion" loading="lazy"><figcaption>Opportunity through partnership</figcaption></figure></div></div></section>`,
		"philanthropy.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>Hope expressed through action</h2><p>Community support begins with presence, compassion and a commitment to restore dignity.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img10.jpeg" alt="Dr. Nelson Alapa at a community gathering" loading="eager"><figcaption>Compassion in action</figcaption></figure><figure><img src="/assets/images/img11.jpeg" alt="Dr. Nelson Alapa with community leaders" loading="lazy"><figcaption>Standing with communities</figcaption></figure><figure><img src="/assets/images/img12.jpeg" alt="Dr. Nelson Alapa at an event" loading="lazy"><figcaption>Hope for tomorrow</figcaption></figure></div></div></section>`,
		"impact.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>Progress people can see</h2><p>Programmes gain meaning when communities are present, heard and strengthened.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img13.jpeg" alt="Dr. Nelson Alapa at a public event" loading="eager"><figcaption>Community impact</figcaption></figure><figure><img src="/assets/images/img14.jpeg" alt="Leadership engagement" loading="lazy"><figcaption>Shared progress</figcaption></figure><figure><img src="/assets/images/img15.jpeg" alt="Community gathering" loading="lazy"><figcaption>People first</figcaption></figure></div></div></section>`,
		"news.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>From the field</h2><p>Official moments, public engagements and leadership stories in pictures.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img16.jpeg" alt="Dr. Nelson Alapa at an official gathering" loading="eager"><figcaption>Official engagement</figcaption></figure><figure><img src="/assets/images/img17.jpeg" alt="Dr. Nelson Alapa with leaders" loading="lazy"><figcaption>Leadership update</figcaption></figure><figure><img src="/assets/images/img18.jpeg" alt="Community event" loading="lazy"><figcaption>In the community</figcaption></figure></div></div></section>`,
		"events.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>Moments that bring people together</h2><p>Community gatherings, public conversations and events that strengthen shared purpose.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img19.jpeg" alt="Community celebration" loading="eager"><figcaption>A community gathering</figcaption></figure><figure><img src="/assets/images/img20.jpeg" alt="Public event" loading="lazy"><figcaption>Public occasion</figcaption></figure><figure><img src="/assets/images/img21.jpeg" alt="Community event" loading="lazy"><figcaption>Shared moments</figcaption></figure></div></div></section>`,
		"contacts.html": `<section class="route-gallery"><div class="container"><div class="route-gallery-heading"><h2>A leadership team that stays connected</h2><p>For invitations, media enquiries, partnership ideas and community feedback, the office is ready to listen.</p></div><div class="route-gallery-grid"><figure><img src="/assets/images/img22.jpeg" alt="Dr. Nelson Alapa at a public event" loading="eager"><figcaption>Connect with the office</figcaption></figure><figure><img src="/assets/images/img23.jpeg" alt="Dr. Nelson Alapa in conversation" loading="lazy"><figcaption>Open dialogue</figcaption></figure><figure><img src="/assets/images/img24.jpeg" alt="Community gathering" loading="lazy"><figcaption>Community first</figcaption></figure></div></div></section>`,
	}
	return galleries[file]
}
