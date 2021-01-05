module github.com/namo-io/go-kit

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/antonfisher/nested-logrus-formatter v1.3.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/google/uuid v1.0.0
	github.com/olivere/elastic v6.2.35+incompatible
	github.com/olivere/elastic/v7 v7.0.4
	github.com/prometheus/common v0.15.0
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/viper v1.7.1
	go.mongodb.org/mongo-driver v1.4.4
	gopkg.in/sohlich/elogrus.v7 v7.0.0
	gorm.io/driver/postgres v1.0.6
	gorm.io/gorm v1.20.9
)

replace github.com/namo-io/go-kit => ../../namo-io/go-kit
