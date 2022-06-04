package config

import "log"

// create struct for containing dependences all web app
type Application struct {
	ErrorLog *log.Logger
	InfoLog *log.Logger
}
