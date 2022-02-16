module github.com/collins/crabada

go 1.16

require (
	decred.org/dcrwallet/v2 v2.0.0-rc3
	github.com/c-ollins/crabada v0.0.0-00010101000000-000000000000
	github.com/decred/dcrwallet/wallet v1.3.0
	github.com/decred/slog v1.2.0
	github.com/ethereum/go-ethereum v1.10.11
	github.com/go-echarts/go-echarts/v2 v2.2.4
	github.com/jrick/logrotate v1.0.0
	github.com/lib/pq v1.0.0
	github.com/shopspring/decimal v1.3.1
	go.etcd.io/bbolt v1.3.6
	gopkg.in/tucnak/telebot.v2 v2.4.0
)

replace github.com/c-ollins/crabada => ./
