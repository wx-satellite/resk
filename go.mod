module resk

go 1.12

require (
	github.com/go-ini/ini v1.60.0 // indirect
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.3.0+incompatible
	github.com/lestrrat-go/strftime v1.0.3 // indirect
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/segmentio/ksuid v1.0.2
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/sirupsen/logrus v1.6.0
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a
	github.com/tietang/dbx v1.0.0
	github.com/tietang/go-utils v0.1.3 // indirect
	github.com/tietang/props v2.2.0+incompatible
	github.com/valyala/fasttemplate v1.2.1 // indirect
	github.com/x-cray/logrus-prefixed-formatter v0.5.2
	golang.org/x/sys v0.0.0-20200817155316-9781c653f443 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/ini.v1 v1.60.0 // indirect
)

replace (
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => github.com/golang/crypto v0.0.0-20190403202508-8e1b8d32e692
	golang.org/x/net v0.0.0-20190311183353-d8887717615a => github.com/golang/net v0.0.0-20190403144856-b630fd6fe46b
	golang.org/x/sys v0.0.0-20190403152447-81d4e9dc473e => github.com/golang/sys v0.0.0-20190403152447-81d4e9dc473e
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	golang.org/x/tools v0.0.0-20190328211700-ab21143f2384 => github.com/golang/tools v0.0.0-20190404132500-923d25813098
)
