module github.com/nikooo777/lbry-blobs-downloader

go 1.17

require (
	github.com/lbryio/lbry.go/v2 v2.7.2-0.20220208210038-a0391bec7915
	github.com/lbryio/reflector.go v1.1.3-0.20220209235713-2f7d67794f93
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.4.0
)

require (
	github.com/aws/aws-sdk-go v1.27.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bluele/gcache v0.0.2 // indirect
	github.com/brk0v/directio v0.0.0-20190225130936-69406e757cf7 // indirect
	github.com/btcsuite/btcd v0.0.0-20190824003749-130ea5bddde3 // indirect
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d // indirect
	github.com/btcsuite/go-socks v0.0.0-20170105172521-4720035b7bfd // indirect
	github.com/btcsuite/websocket v0.0.0-20150119174127-31079b680792 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-errors/errors v1.1.1 // indirect
	github.com/go-ini/ini v1.48.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af // indirect
	github.com/karrick/godirwalk v1.16.1 // indirect
	github.com/lbryio/chainquery v1.9.0 // indirect
	github.com/lbryio/lbry.go v1.1.2 // indirect
	github.com/lbryio/types v0.0.0-20201019032447-f0b4476ef386 // indirect
	github.com/lucas-clemente/quic-go v0.24.0 // indirect
	github.com/lyoshenka/bencode v0.0.0-20180323155644-b7abd7672df5 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/marten-seemann/qpack v0.2.1 // indirect
	github.com/marten-seemann/qtls-go1-16 v0.1.4 // indirect
	github.com/marten-seemann/qtls-go1-17 v0.1.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/pelletier/go-toml v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.10.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.18.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/shopspring/decimal v0.0.0-20191009025716-f1972eb1d1f5 // indirect
	github.com/spf13/afero v1.4.1 // indirect
	github.com/spf13/cast v1.3.0 // indirect
	github.com/spf13/jwalterweatherman v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/volatiletech/inflect v0.0.0-20170731032912-e7201282ae8d // indirect
	github.com/volatiletech/null v8.0.0+incompatible // indirect
	github.com/volatiletech/sqlboiler v3.4.0+incompatible // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/net v0.0.0-20210428140749-89ef3d95e781 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/tools v0.1.1 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/ini.v1 v1.51.0 // indirect
	gopkg.in/nullbio/null.v6 v6.0.0-20161116030900-40264a2e6b79 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//replace github.com/lbryio/reflector.go => /home/niko/go/src/github.com/lbryio/reflector.go
replace github.com/btcsuite/btcd => github.com/lbryio/lbrycrd.go v0.0.0-20200203050410-e1076f12bf19
