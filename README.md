# honestbee-test

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/mrmoneyc/honestbee-test/master/LICENSE)

honestbee take home test

## Demo

[![TCP server demo](https://img.youtube.com/vi/pwg2YZaAmwM/0.jpg)](https://www.youtube.com/watch?v=pwg2YZaAmwM)

## Build

### Server
```
make build_server
```

### Client
```
make build_client
```


## Usage

### Server
```
./bin/tcpsrv
```

TCP server default timeout: 10 sec

### Client
```
./bin/tcpcli --help
```

### Statistics API
```
curl -X GET localhost:9528/stat
```

License
---------------

MIT license
