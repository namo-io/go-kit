module github.com/namo-io/go-kit

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/antonfisher/nested-logrus-formatter v1.3.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/google/uuid v1.0.0
	github.com/olivere/elastic/v7 v7.0.4
	github.com/ryanuber/columnize v0.0.0-20160712163229-9b3edd62028f
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	go.mongodb.org/mongo-driver v1.4.4
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	gopkg.in/sohlich/elogrus.v7 v7.0.0
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gorm.io/driver/postgres v1.0.6
	gorm.io/gorm v1.20.9
)

replace github.com/namo-io/go-kit => ../../namo-io/go-kit
