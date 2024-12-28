package config

// Database configuration key value
type database struct {
	// Host name where the Database is hosted
	Host string `json:"host"`

	// Port number of Database connection
	Port int `json:"port"`

	// User name of Database connection
	User string `json:"user"`

	// Password of Database conenction
	Password string `json:"password"`

	// Name of Database that want to connect
	Name string `json:"name"`

	// Dialect is varian or type of database query language
	Dialect string `json:"dialect"`
}
