module blobdownloader

go 1.14

require (
	github.com/lbryio/lbry.go/v2 v2.7.2-0.20210416195322-6516df1418e3
	github.com/lbryio/reflector.go v1.1.3-0.20210521160537-bc889001bbe9
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.3
)

//replace github.com/lbryio/reflector.go => /home/niko/go/src/github.com/lbryio/reflector.go
replace github.com/btcsuite/btcd => github.com/lbryio/lbrycrd.go v0.0.0-20200203050410-e1076f12bf19
