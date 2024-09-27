# lbry-blobs-downloader

# Usage
```
download blobs or streams from reflector.

Usage:
  blobsdownloader [flags]

Flags:
      --build                       build the file from the blobs
      --claim_id string             claim id of the stream (excludes --hash and forces --stream)
      --concurrent-threads int      Number of concurrent downloads to run (default 16)
      --hash string                 hash of the blob or sdblob
  -h, --help                        help for blobsdownloader
      --http-port string            the port reflector listens to for HTTP connections (default "5569")
      --http3-port string           the port reflector listens to for HTTP3 peer connections (default "5568")
      --mode int                    0: HTTP3, 1: TCP (LBRY), 2: HTTP, 3: use all (default 2)
      --peer-port string            the port reflector listens to for TCP peer connections (default "5567")
      --rename                      attempt renaming the downloaded file to its original name
      --stream                      whether the hash is for a stream or not (download whole file)
      --trace                       print all traces
      --upstream-reflector string   the address of the reflector server (without port) (default "blobcache-eu.odycdn.com")
```

# Example output:

```bash
./blobsdownloader --mode 2 --upstream-reflector blobcache-eu.odycdn.com --claim_id d7f967165c39862d0a66f8db83ac09fc30b6a152 --build --rename --stream
DEBU[0000] tcp server: blobcache-eu.odycdn.com:5567     
DEBU[0000] http3 server: blobcache-eu.odycdn.com:5568   
DEBU[0000] http server: blobcache-eu.odycdn.com:5569    
DEBU[0000] [H] download time: 96 ms     Speed: 0.20 MB/s    
DEBU[0000] 7035ef0d04d24aad341eeebd5a67c0929e38133d992d47d8d17f5e8e04758814f90d2f0f6e273966b27301200a3d6b57
DEBU[0002] [H] download time: 1734 ms   Speed: 1.15 MB/s 
DEBU[0002] 8589360be3fa2dbf69016ae47d9fe8cd7213f0ed3c894d1488225c26c080082f8c6abf728ee602e02d64023a23582658 
...
DEBU[0014] HTTP protocol downloaded at an average of 0.97 MiB/s
```

# As a library
`go get https://github.com/nikooo777/lbry-blobs-downloader`

Then see [example](example/example.go)

# Building
1) Install GO 1.19+ (on ubuntu it's `sudo snap install go --classic`)
2) run `make`
3) ???
4) profit