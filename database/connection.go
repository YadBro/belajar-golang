package database

var connection string

// Init function adalah sebuah function yang akan di eksekusi pertama kali ketika ada yang mengimportnya,
// sama seperti dalam OOP yaitu constructor.
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
