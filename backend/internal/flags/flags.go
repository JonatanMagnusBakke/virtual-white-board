package flags

import (
	"os"
	"strconv"
)

//Flags env flags struct
type Flags struct {
	APIPort          string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDbName   string
	AllowCORS        bool
}

func lookupEnvOrDefault(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

//ParseFlags parses the environment variables
func ParseFlags() (*Flags, error) {

	apiPort := lookupEnvOrDefault("API_PORT", "3000")

	postgresHost := lookupEnvOrDefault("POSTGRES_HOST", "localhost")

	postgresPort := lookupEnvOrDefault("POSTGRES_PORT", "5432")

	postgresUser := lookupEnvOrDefault("POSTGRES_USER", "postgres")

	postgresPassword := lookupEnvOrDefault("POSTGRES_PASSWORD", "postgres")

	postgresDbName := lookupEnvOrDefault("POSTGRES_DBNAME", "docker")

	corsString := lookupEnvOrDefault("CORS", "false")
	allowCors, err := strconv.ParseBool(corsString)
	if err != nil {
		allowCors = false
	}

	flags := &Flags{
		APIPort:          apiPort,
		PostgresHost:     postgresHost,
		PostgresPort:     postgresPort,
		PostgresUser:     postgresUser,
		PostgresPassword: postgresPassword,
		PostgresDbName:   postgresDbName,
		AllowCORS:        allowCors,
	}
	return flags, nil
}
