package main

import (
	"encoding/gob"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/finkord/building_modern_web_app_with_go/internal/driver"

	"github.com/finkord/building_modern_web_app_with_go/internal/config"
	"github.com/finkord/building_modern_web_app_with_go/internal/handlers"
	"github.com/finkord/building_modern_web_app_with_go/internal/helpers"
	"github.com/finkord/building_modern_web_app_with_go/internal/models"
	"github.com/finkord/building_modern_web_app_with_go/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

var infoLog *log.Logger
var errorLog *log.Logger

// main is the main function
func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	defer close(app.MailChan)

	fmt.Println("Starting mail listener...")
	listenForMail()

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {

	// what am i going to store in my session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(models.RoomRestriction{})
	gob.Register(map[string]int{})

	// read flags
	inProduction := flag.Bool("prod", true, "Application will run in production mode")
	useCache := flag.Bool("cache", false, "Use template cache")
	dbName := flag.String("dbname", "bookings", "Database name")
	dbHost := flag.String("dbhost", "localhost", "Database host")
	dbUser := flag.String("dbuser", "postgres", "Database user")
	dbPass := flag.String("dbpassword", "postgres", "Database password")
	dbPort := flag.String("dbport", "5432", "Database port")
	dbSSL := flag.String("dbssl", "disable", "Database SSL (disable, prefer, require)")

	flag.Parse()

	if *dbUser == "" || *dbPass == "" {
		log.Println("Missing required database credentials (-dbuser and -dbpassword)")
		return nil, errors.New("missing required database credentials")
	}

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	// Change this to true in production
	app.InProduction = *inProduction
	app.UseCache = *useCache

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	log.Println("Connecting to database...")
	// build the database connection string
	// host=localhost port=5432 user=postgres password=postgres dbname=bookings
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", *dbHost, *dbPort, *dbName, *dbUser, *dbPass, *dbSSL)
	db, err := driver.ConnectSQL(connStr)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	log.Println("Connected to database")

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Println("Cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	helpers.NewHelpers(&app)

	render.NewRenderer(&app)

	fmt.Printf("Starting server on port %s\n", portNumber)

	return db, nil
}
