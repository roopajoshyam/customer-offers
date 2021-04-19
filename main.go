package main

import (
	"flag"
	"fmt"
	"github.com/golangcollege/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/jroopa/offers/routes"

	//"github.com/jroopa/offers/routes"
	"log"
	"net/http"
	"os"
	"time"
)

var db *gorm.DB
var err error

type application struct {
	session *sessions.Session
}

func main() {
	e := godotenv.Load()

	if e != nil {
		fmt.Print("error loading env")
	}
	fmt.Println(e)
	port := os.Getenv("PORT")

	defer db.Close()
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	//app := &application{
	//	session: session,
	//}

	//srv := &http.Server{
	//	Handler: app.Handlers(),
	//}
	// Handle routes
	http.Handle("/", routes.Handlers())

	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
