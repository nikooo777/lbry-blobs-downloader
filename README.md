# lbry-blobs-downloader

# Usage
```
download blobs or streams from reflector.

Usage:
  blobsdownloader [flags]

Flags:
      --build                       build the file from the blobs
      --concurrent-threads int      Number of concurrent downloads to run (default 1)
      --hash string                 hash of the blob or sdblob (default "c333e168b1adb5b8971af26ca2c882e60e7a908167fa9582b47a044f896484485df9f5a0ada7ef6dc976489301e8049d")
  -h, --help                        help for blobsdownloader
      --http-port string            the port reflector listens to for HTTP connections (default "5569")
      --http3-port string           the port reflector listens to for HTTP3 peer connections (default "5568")
      --mode int                    0: HTTP3, 1: TCP (LBRY), 2: HTTP, 3: use all
      --peer-port string            the port reflector listens to for TCP peer connections (default "5567")
      --stream                      whether the hash is for a stream or not (download whole file)
      --trace                       print all traces
      --upstream-reflector string   the address of the reflector server (without port) (default "reflector.lbry.com")
```

# Example output:

```
./blobsdownloader --mode 2 --upstream-reflector blobcache-eu.lbry.com --stream --trace
INFO[0000] tcp server: blobcache-eu.lbry.com:5567       
INFO[0000] http3 server: blobcache-eu.lbry.com:5568     
INFO[0000] http server: blobcache-eu.lbry.com:5569      
INFO[0000] [0](blobcache-eu) origin: disk - timing: 318.112µs - delta: 0s
[1](blobcache-eu) origin: db-backed - timing: 1.733389ms - delta: 1.415277ms
[2](blobcache-eu) origin: sf_db-backed - timing: 1.740339ms - delta: 6.95µs
[3](blobcache-eu) origin: caching - timing: 1.758839ms - delta: 18.5µs
[4](nikubuntu) origin: http - timing: 152.753668ms - delta: 150.994829ms 
INFO[0000] [H] download time: 152 ms    Speed: 0.02 MB/s   
INFO[0000] 098140c687145e2a5dcfbbc25f0143994c884c7e59aec61313c8c703628165e7566db256b3960ba51a745d291fae3029 
INFO[0002] [0](blobcache-eu) origin: disk - timing: 4.457131ms - delta: 0s
[1](blobcache-eu) origin: db-backed - timing: 5.542876ms - delta: 1.085745ms
[2](blobcache-eu) origin: sf_db-backed - timing: 5.546686ms - delta: 3.81µs
[3](blobcache-eu) origin: caching - timing: 5.556186ms - delta: 9.5µs
[4](nikubuntu) origin: http - timing: 2.244775968s - delta: 2.239219782s 
INFO[0002] [H] download time: 2244 ms   Speed: 0.89 MB/s  
INFO[0002] ce6757b67a03be058ef94bb8df5d4a799eeb87b829f22974878e5c7122e16f3443ab8c2e9d2eddbd2c7feb09ffe23085 
INFO[0004] [0](blobcache-eu) origin: disk - timing: 9.062184ms - delta: 0s
[1](blobcache-eu) origin: db-backed - timing: 10.206859ms - delta: 1.144675ms
[2](blobcache-eu) origin: sf_db-backed - timing: 10.212759ms - delta: 5.9µs
[3](blobcache-eu) origin: caching - timing: 10.230229ms - delta: 17.47µs
[4](nikubuntu) origin: http - timing: 1.797285994s - delta: 1.787055765s 
INFO[0004] [H] download time: 1797 ms   Speed: 1.11 MB/s
...
INFO[0023] HTTP protocol downloaded at an average of 1.16 MiB/s
```

# As a library
`go get https://github.com/nikooo777/lbry-blobs-downloader.git`

Then see [example](example/example.go)

# Building
1) Install GO 1.16+ (on ubuntu it's `sudo snap install go --classic`)
2) run `make`
3) ???
4) profit