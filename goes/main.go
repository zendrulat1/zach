package main

import (
	"context"
	"fmt"
	//"github.com/davecgh/go-spew/spew"
  	//"html/template"
	. "github.com/logrusorgru/aurora"
	"github.com/rs/cors"
	
	Home "github.com/zendrulat1/goes/Handlers/Home"
	
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)
type ViewData struct {
  Name string
}
var err error

type Contexter struct {
	M string
	U *url.URL
	P string
	B    io.ReadCloser
	//Gb   func() (io.ReadCloser, error)
	Host string
	Form url.Values
	Cancel <-chan struct{}
	R      *http.Response
	//H      http.Header
  //D Duration
}
 var CC Contexter
  

func AddContext(ctx context.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
     Start := time.Now()
		Duration := time.Now().Sub(Start)
    fmt.Println(Duration)
   // spew.Dump(Duration)
		CC = Contexter{
			M:      r.WithContext(ctx).Method,
			U:      r.WithContext(ctx).URL,
			P:      r.WithContext(ctx).Proto,
			B:      r.WithContext(ctx).Body,
			Host:   r.WithContext(ctx).Host,
			Form:   r.WithContext(ctx).Form,
			Cancel: r.WithContext(ctx).Cancel,
			R:      r.WithContext(ctx).Response,
     // D: Duration,
			//H:      r.WithContext(ctx).Header,
		}
    
    

	
		fmt.Println(Blue("/ʕ◔ϖ◔ʔ/````````````````````````````````````````````"))
		fmt.Printf("Method:%s\n - URL:%s\n - Proto:%s\n - Body:%v\n - Host:%s\n - Form:%v\n - Cancel:%d\n - Response:%d\n - Dur:%02d-00:00 \n - DBConn:%v \r\n  - Header:%s\n ",
			Cyan(CC.M),
			Red(CC.U),
			Brown(CC.P),
			Blue(CC.B),
			Yellow(CC.Host),
			BgRed(CC.Form),
			BgGreen(CC.Cancel),
			BgBrown(CC.R),
			//BgMagenta(CC.D),
			//Red(DBs.DBS().Stats()),
			//Green(CC.H),
		)

		cookie, _ := r.Cookie("username")

		if cookie != nil {
			//Add data to context
			ctx := context.WithValue(r.Context(), "Username", cookie.Value)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {

			if err != nil {
				// Error occurred while parsing request body
				w.WriteHeader(http.StatusBadRequest)

				return
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func StatusPage(w http.ResponseWriter, r *http.Request) {
	//Get data from context
	if username := r.Context().Value("Username"); username != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello " + username.(string) + "\n"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Logged in"))
	}
}

//starts the server
func main() {

	mux := http.NewServeMux()
	//handlers

	
	mux.HandleFunc("/", StatusPage)
	mux.HandleFunc("/home", Home.Home)
	
	handler := cors.Default().Handler(mux)
	c := context.Background()
	log.Fatal(http.ListenAndServe(":8082", AddContext(c, handler)))

}