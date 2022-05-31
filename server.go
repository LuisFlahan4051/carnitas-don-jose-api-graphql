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
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

const (
	DEFAULTHOST     = "localhost"
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

// --------------------- GRAPH SERVER
func addGraphqlServer(mux *mux.Router, uriApp string, uriApi string) *mux.Router {
	graph := generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{},
	})

	graphServer := handler.NewDefaultServer(graph)
	graphServer.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "http://"+uriApp
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", graphServer)

	log.Println(color.Ize(color.Green, ">>> Connect to http://"+uriApi+"/query for GraphQL playground"))
	return mux
}

// ------------------------ SERVE STATICS FILES
//Need the react app built, the address is defined here
func index(writer http.ResponseWriter, request *http.Request) {
	indexTemplate := template.Must(template.ParseFiles("build/ui/index.html"))
	indexTemplate.Execute(writer, nil)
}

func addUIHandler(mux *mux.Router) *mux.Router {
	staticFiles := http.FileServer(http.Dir("build/ui/static/"))

	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", staticFiles))
	mux.HandleFunc("/", index)

	return mux
}

// ----------------------- MIX AND PREPARE ALL PORTS IN ONE SERVER
func newMux(portApp string, hostApp string, portApi string, hostApi string) *mux.Router {
	mux := mux.NewRouter()

	//Use this for enable all origins of requests
	mux.Use(cors.AllowAll().Handler)

	//Use this for enable specific origins
	/* mux.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{
			"http://"+ uirApp,
			"http://localhost:8080",
		},
		AllowCredentials: true,
		Debug:            true,
	}).Handler) */

	// NOTE: Make some func for iterating the mux and add eny server.

	mux = addUIHandler(mux)
	mux = addGraphqlServer(mux, portApi, hostApi)
	return mux
}

func runServer(mux *mux.Router, port string) {
	fmt.Println("Server working fine!")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

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

	prepareMux := newMux(portApp, hostApp, portApi, hostApi)

	runServer(prepareMux, portApi)

}
