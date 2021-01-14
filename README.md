# lbry-blobs-downloader

# Usage
```
download blobs or streams from reflector.

Usage:
  blobdownloader [flags]

Flags:
      --build                      build the file from the blobs
      --concurrent-threads int     Number of concurrent downloads to run (default 1)
      --hash string                hash of the blob or sdblob (default "c333e168b1adb5b8971af26ca2c882e60e7a908167fa9582b47a044f896484485df9f5a0ada7ef6dc976489301e8049d")
  -h, --help                       help for blobdownloader
      --mode int                   0: only use QUIC, 1: only use TCP, 2: use both
      --peer-port string           the port reflector listens to for TCP peer connections (default "5567")
      --quic-port string           the port reflector listens to for QUIC peer connections (default "5568")
      --reflector-address string   the address of the reflector server (without port) (default "cdn.reflector.lbry.com")
      --stream                     whether the hash is for a stream or not (download whole file)
      --trace                      print all traces
```

# Example output:

```
./blobdownloader --stream --mode 0 --reflector-address 192.168.0.103 --trace --build
INFO[0000] tcp server: 192.168.0.103:5567                                                                                                                                                                                                                                                                                    
INFO[0000] quic server: 192.168.0.103:5568                                                                                                                                                                                                                                                                                   
INFO[0000] [0](blobcache3) origin: disk - timing: 128.422µs - delta: 0s                                                                                                                                                                                                                                                      
[1](blobcache3) origin: db-backed - timing: 338.481µs - delta: 210.059µs                                                                                                                                                                                                                                                     
[2](blobcache3) origin: caching - timing: 342.686µs - delta: 4.205µs                                                                                                                                                                                                                                                         
[3](player4) origin: http3 - timing: 7.797947ms - delta: 7.455261ms                                                                                                                                                                                                                                                          
INFO[0000] [Q] download time: 39 ms     Speed: 0.08 MB/s                                                                                                                                                                                                                                                                     
INFO[0000] 098140c687145e2a5dcfbbc25f0143994c884c7e59aec61313c8c703628165e7566db256b3960ba51a745d291fae3029                                                                                                                                                                                                                  
INFO[0000] [0](blobcache3) origin: disk - timing: 1.33516ms - delta: 0s                                                                                                                                                                                                                                                      
[1](blobcache3) origin: db-backed - timing: 1.547805ms - delta: 212.645µs                                                                                                                                                                                                                                                    
[2](blobcache3) origin: caching - timing: 1.553207ms - delta: 5.402µs                                                                                                                                                                                                                                                        
[3](player4) origin: http3 - timing: 59.54254ms - delta: 57.989333ms                                                                                                                                                                                                                                                         
INFO[0000] [Q] download time: 59 ms     Speed: 33.47 MB/s                                                                                                                                                                                                                                                                    
INFO[0000] ce6757b67a03be058ef94bb8df5d4a799eeb87b829f22974878e5c7122e16f3443ab8c2e9d2eddbd2c7feb09ffe23085                                                                                                                                                                                                                  
INFO[0000] [0](blobcache3) origin: disk - timing: 1.172721ms - delta: 0s                                                                                                                                                                                                                                                     
[1](blobcache3) origin: db-backed - timing: 1.376992ms - delta: 204.271µs                                                                                                                                                                                                                                                    
[2](blobcache3) origin: caching - timing: 1.383173ms - delta: 6.181µs                                                                                                                                                                                                                                                        
[3](player4) origin: http3 - timing: 106.754469ms - delta: 105.371296ms                                                                                                                                                                                                                                                      
INFO[0000] [Q] download time: 106 ms    Speed: 18.69 MB/s                                                                                                                                                                                                                                                                    
[...]
INFO[0000] QUIC protocol downloaded at an average of 36.13 MiB/s
```

# As a library
`go get https://github.com/nikooo777/lbry-blobs-downloader.git`

Then see [example](example/example.go)