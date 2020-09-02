# lbry-blobs-downloader
```
download blobs or streams from reflector.

Usage:
  blobdownloader [flags]

Flags:
      --concurrent-threads int     Number of concurrent downloads to run (default 1)
      --hash string                hash of the blob or sdblob (default "c333e168b1adb5b8971af26ca2c882e60e7a908167fa9582b47a044f896484485df9f5a0ada7ef6dc976489301e8049d")
  -h, --help                       help for blobdownloader
      --mode int                   0: only use QUIC, 1: only use TCP, 2: use both
      --peer-port string           the port reflector listens to for TCP peer connections (default "5567")
      --quic-port string           the port reflector listens to for QUIC peer connections (default "5568")
      --reflector-address string   the address of the reflector server (without port) (default "cdn.reflector.lbry.com")
      --stream                     whether the hash is for a stream or not (download whole file)
```

Example output:

```
./blobdownloader --mode 2 --hash f45c05b04a28eace6c80727db30174a3a234f48752646c885bfb95125a6b47ddaaca57dba8c9a4daa0127d6fc01f8fc0
INFO[0000] tcp server: cdn.reflector.lbry.com:5567      
INFO[0000] quic server: cdn.reflector.lbry.com:5568     
INFO[0000] QUIC protocol:                               
INFO[0000] [Q] download time: 285 ms    Speed: 7.01 MB/s   
INFO[0000] TCP protocol:                                
INFO[0000] [T] download time: 697 ms    Speed: 2.87 MB/s 
```
