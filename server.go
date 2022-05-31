package main

// GENERATE THE CODE USING > go run github.com/99designs/gqlgen init
import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	//"github.com/LuisFlahan4051/maximonet/api/database" Example adding local files
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-graphql/graph"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-graphql/graph/generated"
	"github.com/TwiN/go-color"
	"github.com/go-chi/chi"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

const (
	DEFAULTHOST     = "http://localhost"
	DEFAULTPORT_API = "8080"
	DEFAULTHOST_API = DEFAULTHOST
	DEFAULTPORT_DB  = "27017"
	DEFAULTHOST_DB  = DEFAULTHOST
	DEFAULTPORT_APP = "3000"
	DEFAULTHOST_APP = DEFAULTHOST
)

func catch(err error) {
	if err != nil {
		log.Fatal(color.Ize(color.Red, err.Error()))
	}
}

//Need the react app built, the address is defined here
func index(writer http.ResponseWriter, request *http.Request) {
	indexTemplate := template.Must(template.ParseFiles("ui/build/index.html"))
	indexTemplate.Execute(writer, nil)
}

func addUIHandler(mux *mux.Router, port string, host string) *mux.Router {
	staticFiles := http.FileServer(http.Dir("ui/build/static/"))

	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", staticFiles))
	mux.HandleFunc("/", index)

	fmt.Println("Port " + port + "Added Successfully!")
	fmt.Println("Can open " + host + ":" + port + " in your web browser\n")
	return mux
}

// func newMux() *mux.Router {
// 	mux := mux.NewRouter()

// 	//Use this for enable all origins of requests
// 	mux.Use(cors.AllowAll().Handler)

// 	//Use this for enable specific origins
// 	/* mux.Use(cors.New(cors.Options{
// 		AllowedOrigins:   []string{
// 			"http://localhost:8080",
// 			"http://localhost:"+port,
// 		},
// 		AllowCredentials: true,
// 		Debug:            true,
// 	}).Handler) */
// 	mux = addUIHandler(mux)
// 	mux = api.AddGraphqlServer(port, graphDoor, mux)
// 	return mux
// }

// func runServer(mux *mux.Router) {
// 	fmt.Println("Server working fine!")
// 	log.Fatal(http.ListenAndServe(":"+port, mux))
// }

// func main() {
// 	//FOR BUILD > go build -ldflags "-H windowsgui" -o main.exe

// 	prepareMux := newMux()

// 	go runServer(prepareMux)

// 	database.TestConnection()

// 	runElectron()
// }

func main() {

	// READ ARGS OF THE CONSOLE WHEN RUN >go run main.go -port=27017 -host=127.0.0.1
	portApiFlag := flag.String("portApi", "", "a string")
	hostApiFlag := flag.String("hostApi", "", "a string")
	portAppFlag := flag.String("portApp", "", "a string")
	hostAppFlag := flag.String("hostApp", "", "a string")
	portDBFlag := flag.String("portDB", "", "a string")
	hostDBFlag := flag.String("hostDB", "", "a string")

	flag.Parse()

	portApi := *portApiFlag
	if portApi == "" {
		portApi = DEFAULTPORT_API
	}
	hostApi := *hostApiFlag
	if hostApi == "" {
		hostApi = DEFAULTHOST_API
	}
	portApp := *portAppFlag
	if portApp == "" {
		portApp = DEFAULTPORT_APP
	}
	hostApp := *hostAppFlag
	if hostApp == "" {
		hostApp = DEFAULTHOST_APP
	}
	portDB := *portDBFlag
	if portDB == "" {
		portDB = DEFAULTPORT_DB
	}
	hostDB := *hostDBFlag
	if hostDB == "" {
		hostDB = DEFAULTHOST_DB
	}
	uriApi := hostApi + ":" + portApi
	uriApp := hostApp + ":" + portApp
	uriDB := hostDB + ":" + portDB
	log.Println(color.Ize(color.Blue, "\nURI-API: "+uriApi+"\nURI-DB: "+uriDB+"\nURI-APP: "+uriApp))

	//--------------------------------------------------------------------------

	router := chi.NewRouter()
	// Add CORS middleware around every request. More inf  https://github.com/rs/cors
	router.Use(cors.AllowAll().Handler)
	/* router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler) */

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == uriApp
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/", playground.Handler("KrisstalNet", "/query"))
	router.Handle("/query", srv)

	log.Println(color.Ize(color.Green, ">>> Connect to "+uriApi+"for GraphQL playground"))
	err := http.ListenAndServe(":"+portApi, router) //:8080
	catch(err)
}
