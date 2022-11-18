package configs

type AppConfig struct {
	MongoConnectionURI string
	DBName             string
	CollectionName     string
}

var Configs = AppConfig{
	MongoConnectionURI: "mongodb://mongo:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false",
	DBName:             "user-api",
	CollectionName:     "Users",
}
