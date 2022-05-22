package configs

import "os"

var (
	URL string
)

func init() {
	URL = os.Getenv("URL")
}
