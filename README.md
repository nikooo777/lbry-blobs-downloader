# lbry-blobs-downloader
```
download blobs or streams from reflector.

Usage:
  blobdownloader [flags]

Flags:
      --concurrent-threads int     Number of concurrent downloads to run (default 1)
      --hash string                hash of the blob or sdblob (default "58742ec8f86abbaadf11ad45e22a78c01e3f89ac3d9f3f1c0d1b77198d34b52672aad8f908a68c763d6767858761c247")
  -h, --help                       help for blobdownloader
      --mode int                   0: only use QUIC, 1: only use TCP, 2: use both
      --peer-port string           the port reflector listens to for TCP peer connections (default "5567")
      --quic-port string           the port reflector listens to for QUIC peer connections (default "5568")
      --reflector-address string   the address of the reflector server (without port) (default "reflector.lbry.com")
      --stream                     whether the hash is for a stream or not (download whole file)
```

Example output:

```
./blobdownloader --mode 2 --hash f45c05b04a28eace6c80727db30174a3a234f48752646c885bfb95125a6b47ddaaca57dba8c9a4daa0127d6fc01f8fc0
INFO[0000] QUIC protocol:                               
INFO[0006] download time: 6169 ms       Speed: 3.32 MB/s      
INFO[0006] TCP protocol:                                
INFO[0007] download time: 1241 ms       Speed: 1.61 MB/s  
```
