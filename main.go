package main

import (
	"html/template"
	"log"
	"net/http"
)

type page struct {
	Title, Kicker, Heading, Intro, Photo string
	Cards                                []card
}

type card struct{ Title, Text string }

var pages = map[string]page{
	"/about/": {
		Title: "About | Nelson Alapa", Kicker: "ABOUT", Heading: "A life shaped by service and community", Photo: "about-nelson-alapa.jpg",
		Intro: "Hon. Chief (Dr.) Nelson G.O. Alapa is a public servant, entrepreneur and community advocate from Benue South.",
		Cards: []card{
			{"Biography", "Born in the mid-1960s to Chief Onyilokwu Alapa in Otukpo, Benue State, Nelson Alapa’s background connects the Idoma and Igede communities of Benue South. His public life has been guided by a belief that leadership must be accessible, practical and rooted in the everyday concerns of the people."},
			{"Educational background", "His published academic record includes the First School Leaving Certificate from Army Children School, Owerri; WAEC from Community Secondary School, Makurdi; an ND in Criminology and Security Management from the University of Jos; an Advanced Diploma and BSc in Business Management from Abubakar Tafawa Balewa University; a PGD in Peace and Conflict Resolution from the National Open University of Nigeria; an MSc in Legislative Studies from the University of Benin; and a PhD in Legislative Studies from the Federal University, Lokoja."},
			{"Public service and recognition", "During his time in the House of Representatives, the public record lists committee service touching Agriculture, Petroleum Downstream, electoral matters and the Nigeria Army. The existing site also records honours including Icon of the Niger, an Excellence Award from the National Association of Idoma Students, and recognition for rural development and youth empowerment."},
		},
	},
	"/politics/": {
		Title: "Politics | Nelson Alapa", Kicker: "POLITICAL JOURNEY", Heading: "Representation must deliver visible change", Photo: "politics-community.jpg",
		Intro: "A political journey defined by grassroots connection, public accountability and a commitment to expand the reach of service.",
		Cards: []card{
			{"A mandate for service", "The public account of Nelson Alapa’s journey describes a move into politics after years of community-facing philanthropic work. He was elected to represent the Otukpo/Ohimini Federal Constituency and inaugurated into the National Assembly in 2007."},
			{"Work in office", "The original website attributes a range of constituency-focused interventions to this period: school and health-centre development, water initiatives, community electrification, mobility support, assistance for widows and donations of wheelchairs to people with disabilities. Each item should be supported with dates, locations and project photographs before publication."},
			{"A standard for leadership", "This page presents politics as a responsibility to listen, report back and convert public trust into visible outcomes—not simply a contest for office. It is the right place for a verified timeline, legislation, constituency projects and official statements."},
		},
	},
	"/business/": {
		Title: "Business | Nelson Alapa", Kicker: "BUSINESS", Heading: "Building legacies that create opportunity", Photo: "business-enterprise.jpg",
		Intro: "Enterprise, innovation and long-term investment are presented as complementary paths to growth and self-reliance.",
		Cards: []card{
			{"Maritime", "The business portfolio identifies maritime activity as an area of interest, reflecting a focus on connections, trade and new routes to economic opportunity."},
			{"Real estate", "Real-estate development is framed around lasting foundations—projects that can support growth, improve places and create value over time."},
			{"Farming and human investment", "Agriculture and human-capital development are positioned together: practical investment in the land, skills, young people and resilient local livelihoods."},
		},
	},
	"/philanthropy/": {
		Title: "Philanthropy | Nelson Alapa", Kicker: "PHILANTHROPY", Heading: "Empowering lives with practical support", Photo: "philanthropy-education.jpg",
		Intro: "The philanthropic agenda focuses on access to education, care, dignity and pathways to independence.",
		Cards: []card{
			{"Education and infrastructure", "The existing public site highlights scholarship support, mentoring and investment in educational opportunity. This section should document each school intervention with the project name, community, completion year, verified beneficiary data and approved photography."},
			{"Health, water and wellbeing", "Healthcare support, clean-water initiatives and assistance for vulnerable households are recurring themes in the published record. Presenting these with clear locations and partners will make the impact easier for communities to follow."},
			{"Dignity through empowerment", "Support for widows, feeding initiatives, skills and mobility assistance reflects a simple principle: assistance should preserve dignity and widen opportunity. This page is designed to hold beneficiary stories only where consent has been explicitly granted."},
		},
	},
	"/news/": {
		Title: "News | Nelson Alapa", Kicker: "NEWSROOM", Heading: "Updates from the field", Photo: "news-field.jpg",
		Intro: "Verified announcements, community stories and public statements—presented with context and clear dates.",
		Cards: []card{
			{"Published update", "The existing site published a news item titled ‘Nelson Alapa attends Meeting In Ibadan’ on 7 April 2025. Future posts should preserve this standard of clear title, author, publication date and supporting evidence."},
			{"A better newsroom", "Every story should have an accurate headline, a short summary, full date, credited photographer and a relevant image. This makes the platform useful to citizens, journalists and researchers—not merely a collection of announcements."},
		},
	},
	"/events/": {
		Title: "Events | Nelson Alapa", Kicker: "EVENTS", Heading: "Showing up where community happens", Photo: "events-community.jpg",
		Intro: "A calendar for service, conversation and meaningful community connection.",
		Cards: []card{
			{"Community programmes", "The current public website identifies Children’s Prayer, Christmas Carol, Widows Outreach, political campaigns and birthday events as recurring event categories."},
			{"What every event needs", "Before an event is published, include its date, time, location, host, accessibility information, RSVP details and a named contact. Afterward, add a short report and a carefully selected photo gallery."},
		},
	},
	"/contacts/": {
		Title: "Contact | Nelson Alapa", Kicker: "CONTACT", Heading: "Start a conversation", Photo: "contact-office.jpg",
		Intro: "Questions, invitations, media enquiries and community feedback are welcome.",
		Cards: []card{
			{"Office contact", "Published contact details: +234-7039009991 and info@nelsonalapa.com."},
			{"Office address", "10 Enenche Onobu Close, GRA Otukpo, Benue South."},
			{"Secure contact form", "Before launch, this route should connect to a secure Go endpoint with spam protection, privacy consent and a defined response workflow for the office team."},
		},
	},
	"/video-gallery/": {
		Title: "Video Gallery | Nelson Alapa", Kicker: "VIDEO GALLERY", Heading: "Watch the story unfold", Photo: "video-gallery.jpg",
		Intro: "A focused library for speeches, field visits, interviews and community moments.",
		Cards: []card{
			{"Featured stories", "Use this space for official video releases with concise titles, publishing dates, descriptions, captions and a transcript for every speech or interview."},
			{"Publishing standard", "Use approved footage only, optimize the thumbnail and provide captions. A consistent gallery makes important messages easy to find and accessible to more people."},
		},
	},
}

var view = template.Must(template.New("page").Parse(`<!doctype html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><title>{{.Title}}</title><link rel="preconnect" href="https://fonts.googleapis.com"><link rel="preconnect" href="https://fonts.gstatic.com" crossorigin><link href="https://fonts.googleapis.com/css2?family=DM+Sans:wght@400;500;600;700&family=Playfair+Display:wght@600;700&display=swap" rel="stylesheet"><link rel="stylesheet" href="/styles.css"><link rel="stylesheet" href="/inner.css"></head><body><div class="topbar"><div class="container topbar-inner"><span>Leadership. Service. Progress.</span><span><a href="/contacts/">Connect with the team</a></span></div></div><header class="site-header"><div class="container nav-wrap"><a class="brand" href="/"><span class="seal">NA</span><span><b>HON. CHIEF (DR.)</b><strong>Nelson Alapa</strong></span></a><button class="menu-button" aria-label="Open navigation" aria-expanded="false">☰</button><nav class="main-nav"><a href="/about/">About</a><a href="/politics/">Politics</a><a href="/business/">Business</a><a href="/philanthropy/">Philanthropy</a><a href="/news/">News</a><a href="/events/">Events</a><a href="/contacts/">Contact</a><a class="nav-action" href="/video-gallery/">Video gallery</a></nav></div></header><main><section class="inner-hero"><div class="container"><p class="eyebrow light">{{.Kicker}}</p><h1>{{.Heading}}</h1><p>{{.Intro}}</p></div></section><section class="route-photo" style="background-image:linear-gradient(90deg,rgba(16,44,62,.25),rgba(16,44,62,.25)),url('/assets/images/{{.Photo}}')"><div>Approved image: {{.Heading}}</div></section><section class="section"><div class="container content-layout">{{range .Cards}}<article class="content-card"><p class="eyebrow">NELSON ALAPA</p><h2>{{.Title}}</h2><p>{{.Text}}</p></article>{{end}}</div></section></main><footer><div class="container footer-bottom"><span>© 2026 Nelson Alapa. All rights reserved.</span><span><a href="/contacts/">Contact</a> &nbsp; <a href="/video-gallery/">Video gallery</a></span></div></footer><script src="/script.js"></script></body></html>`))

func main() {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.Handle("/styles.css", http.FileServer(http.Dir(".")))
	mux.Handle("/business.css", http.FileServer(http.Dir(".")))
	mux.Handle("/inner.css", http.FileServer(http.Dir(".")))
	mux.Handle("/script.js", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/business/" {
			http.ServeFile(w, r, "business.html")
			return
		}
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "index.html")
			return
		}
		if data, ok := pages[r.URL.Path]; ok {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_ = view.Execute(w, data)
			return
		}
		http.NotFound(w, r)
	})
	log.Println("Serving at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
