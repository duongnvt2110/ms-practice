package config

type App struct {
	Host string `envconfig:"APP_HOST" default:"localhost"`
	Port string `envconfig:"APP_PORT" default:"3000"`
	Google
	JWT
	Mysql
}

type JWT struct {
	Secret          string `envconfig:"JWT_SECRET" default:"abcdxxx"`
	AccessTokenExp  int    `envconfig:"JWT_ACCESS_TOKEN_EXP" default:"300"`
	RefreshTokenExp int    `envconfig:"JWT_RERESH_TOKEN_EXP" default:"1800"`
}

type Google struct {
	OauthClientID     string   `envconfig:"GOOGLE_OAUTH_CLIENT_ID"`
	OauthClientSecret string   `envconfig:"GOOGLE_OAUTH_CLIENT_SECRET"`
	OauthScopes       []string `envconfig:"GOOGLE_OAUTH_SCOPES" default:"https://www.googleapis.com/auth/userinfo.email"`
	OauthGoogleUrlAPI string   `envconfig:"GOOGLE_OAUTH_URL_API" default:"https://www.googleapis.com/oauth2/v2/userinfo?access_token="`
}

type Mysql struct {
	PrimaryHosts []string `envconfig:"MYSQL_PRIMARY_HOSTS" default:"localhost"`
	ReadHosts    []string `envconfig:"MYSQL_READ_HOSTS" default:"localhost"`
	Port         string   `envconfig:"MYSQL_PORT" default:"3306"`
	User         string   `envconfig:"MYSQL_USER" default:"admin"`
	Password     string   `envconfig:"MYSQL_PASSWORD" default:"admin"`
	DBName       string   `envconfig:"MYSQL_DB_NAME" default:"test"`
}
