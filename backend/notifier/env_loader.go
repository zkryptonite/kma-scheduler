package notifier

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var (
	domain        = os.Getenv("DOMAIN")
	privateApiKey = os.Getenv("PRIVATE_API_KEY")
	templateName  = os.Getenv("TEMPLATE_NAME")
	dbName        = os.Getenv("DB_NAME")
)
