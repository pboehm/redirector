package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
)

func redirect(g *gin.Context, url string) {
	http.Redirect(g.Writer, g.Request, url, http.StatusFound)
	g.Abort(-1)
}

func RunRedirector(conn *RedisConnection) {
	r := gin.Default()
	r.SetHTMLTemplate(BuildTemplate())

	r.GET("/", func(g *gin.Context) {
		fqdn := g.Request.Host
		if strings.Index(fqdn, ":") != -1 {
			fqdn, _, _ = net.SplitHostPort(fqdn)
		}

		log.Printf("FQDN: %s", fqdn)

		redirection := conn.GetAndIncrementCount(fqdn)

		if redirection.To != "" {
			redirect(g, redirection.To)
		} else {
			g.String(404, "No redirection found for this host")
		}
	})

	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"user": "password",
	}))

	authorized.GET("/index", func(g *gin.Context) {
		g.HTML(200, "index.html", gin.H{"fqdn": RedirectorFQDN})
	})

	authorized.GET("/available/:fqdn", func(g *gin.Context) {
		fqdn := g.Params.ByName("fqdn")

		g.JSON(200, gin.H{
			"available": !conn.Exist(fqdn),
		})
	})

	authorized.GET("/create", func(g *gin.Context) {
		q := g.Request.URL.Query()

		fqdn := q["fqdn"][0]
		target := q["target"][0]

		conn.Create(fqdn, target)

		redirect(g, "/admin/index/")
	})

	r.Run(WebListenSocket)
}

// Get index template from bindata
func BuildTemplate() *template.Template {
	html, err := template.New("index.html").Parse(string(indexTemplate))
	HandleErr(err)

	return html
}

func ValidHostname(host string) (string, bool) {
	valid, _ := regexp.Match("^[a-z0-9]{1,32}$", []byte(host))

	return host, valid
}
