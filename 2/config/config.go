package config

import "os"

var (
	API_KEY     = "faf7e5bb&s"
	API_URL     = "http://www.omdbapi.com/"
	APIPort     = ":5000"
	db_database = "log"
	db_username = "root"
	db_password = "root"
	db_host     = "127.0.0.1"
	db_port     = "3306"
)

const (
	UNKNOWN_ERROR_MESSAGE           string = "unknown error"
	SUCESS_RESPONSE_INDEX           string = "success"
	MESSAGE_RESPONSE_INDEX          string = "message"
	DATA_RESPONSE_INDEX             string = "data"
	ERROR_PARSING_REQ_DATA_TEMPLATE string = "Error parsing request data with request"
	REQ_PARAM_ERROR_MESSAGE         string = "Please check your parameters"
	INTERNAL_SERVER_ERROR_MESSAGE   string = "There has been an internal server error, please try again later"
	SUCCESS_MESSAGE                 string = "Success"
)

func init() {
	if os.Getenv("API_Port") != "" {
		APIPort = ":" + os.Getenv("API_Port")
	}

	if os.Getenv("DB_Database") != "" {
		db_database = os.Getenv("DB_Database")
	}

	if os.Getenv("DB_Username") != "" {
		db_username = os.Getenv("DB_Username")
	}

	if os.Getenv("DB_Password") != "" {
		db_password = os.Getenv("DB_Password")
	}

	if os.Getenv("DB_Host") != "" {
		db_host = os.Getenv("DB_Host")
	}

	if os.Getenv("DB_Port") != "" {
		db_port = os.Getenv("DB_Port")
	}
	if os.Getenv("API_KEY") != "" {
		API_KEY = os.Getenv("API_KEY")
	}
	if os.Getenv("API_URL") != "" {
		API_URL = os.Getenv("API_URL")
	}

}
