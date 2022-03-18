module github.com/nikooo777/lbry-blobs-downloader

go 1.16

require (
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/lbryio/lbry.go v1.1.2 // indirect
	github.com/lbryio/lbry.go/v2 v2.7.2-0.20210416195322-6516df1418e3
	github.com/lbryio/reflector.go v1.1.3-0.20211214213601-4d8e7739d704
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/afero v1.4.1 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
)

//replace github.com/lbryio/reflector.go => /home/niko/go/src/github.com/lbryio/reflector.go
replace github.com/btcsuite/btcd => github.com/lbryio/lbrycrd.go v0.0.0-20200203050410-e1076f12bf19
