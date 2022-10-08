// Package flag di gunakan untuk memparsing command line argument pada package os.args
// Package flag ini mereturn pointer.
package main

import (
	"flag"
	"fmt"
)

var security string

func init() {
	flag.StringVar(&security, "security", "avast", "Put your database security")
}

func main() {
	// flag = .Int | .String | .Bool | .BoolVar
	// flag.String("flagName", "defaultValue", "instruction")

	var host *string = flag.String("host", "localhost", "Put your database host")
	var username = flag.String("username", "mysql", "Put your database user")
	password := flag.String("password", "root", "Put your database password")
	port := flag.Int("port", 8080, "Put your database port")

	flag.Parse() // Mesti di parse dulu

	fmt.Println("DB_HOST :", *host)
	fmt.Println("DB_USER :", *username)
	fmt.Println("DB_PASSWORD :", *password)
	fmt.Println("DB_PORT :", *port)
	security = "smadav"
	fmt.Println("DB_SECURITY :", security)
}
