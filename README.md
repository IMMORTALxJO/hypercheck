# hypercheck
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![DeepSource](https://deepsource.io/gh/IMMORTALxJO/hypercheck.svg/?label=active+issues&show_trend=true&token=SaCVhzg7Sci39dpzTEGdpLsS)](https://deepsource.io/gh/IMMORTALxJO/hypercheck/?ref=repository-badge)

Single binary to test everything you need. Perfectly fits for kubernetes pod healthchecks and init-containers.

Supported checks:
- **Auto** - generate checks automagicly based on environment variables 🧙
- **HTTP** - status codes, content body, headers
- **TCP** - address reachability, latency
- **File System** - existence, type ( dir or regular ), size, count, owners, etc.
- **DNS** - records content and count ( A, CNAME, MX, TXT, etc. )
- **Redis** - ping-pong check
- Databases connectivity:
  - **mysql**
  - **postgres**

### Examples

If you have some connection descriptions in your environment variables, you can simply run:
```
~ env
...
API_ENDPOINT=https://postman-echo.com/basic-auth
API_USER=postman
API_PASS=password
...
~ hypercheck --auto
✅ DNS postman-echo.com 'A:count>0' 3 > 0
✅ HTTP https://postman:password@postman-echo.com:443/basic-auth 'code==200' 200 == 200
```
> detection magic powered by [scheme-detector](https://github.com/IMMORTALxJO/scheme-detector)

Manual checks are also possible:
```
~ hypercheck \
  --db online 'mysql://user:password@localhost:3306/database' \
  --db online 'postgres://user:password@localhost:5432/postgres?sslmode=disable' \
  --http 'code==200,headers:count>1' https://postman-echo.com/status/200 \
  --dns 'online,A:count>1' google.com \
  --fs regular ./README.md \
  --redis online localhost:6379 \
  --tcp 'online,latency<1s' 1.1.1.1:53
✅ Database mysql://user:password@localhost:3306/database 'online' is true
✅ Database postgres://user:password@localhost:5432/postgres?sslmode=disable 'online' is true
✅ HTTP https://postman-echo.com/status/200 'code==200' 200 == 200
✅ HTTP https://postman-echo.com/status/200 'headers:count>1' 6 > 1
✅ DNS google.com 'online' is true
✅ DNS google.com 'A:count>1' 2 > 1
✅ FileSystem ./README.md
✅ Redis localhost:6379 'online' is true
✅ TCP 1.1.1.1:53 'online' is true
✅ TCP 1.1.1.1:53 'latency<1s' 0 < 1
```

### Usage
```
~ hypercheck --help
--auto
 Generate probes automaticaly based on current environment variables ( List )
--fs
 Filesystem files check, attributes:
        exists - at least one file found ( Bool )
        dir - is directory ( List[Bool] )
        regular - is regular file ( List[Bool] )
        uid - files UID ( List[Number] )
        user - files username ( List[String] )
        count - files count ( Number )
        size - files size ( List[Number] )
        gid - files GID ( List[Number] )
        group - files groupname ( List[String] )
--http
 Check http resource, attributes:
        code - response status code ( Number )
        content - response content ( String )
        online - status code 200 ( Bool )
        offline - status code is not 200 ( Bool )
        headers - headers content ( List[String] )
--tcp
 Check tcp port, attributes:
        online - is reachable ( Bool )
        offline - is unreachable ( Bool )
        latency - duration ( Number )
--dns
 Check dns query response, attributes:
        online - A record is not empty ( Bool )
        offline - A record is empty ( Bool )
        A - A record content ( List[string] )
        NS - NS record content ( List[string] )
        TXT - TXT record content ( List[string] )
        MX - MX record content ( List[string] )
        CNAME - CNAME record content ( String )
--tcp
 Check tcp port, attributes:
        online - is reachable ( Bool )
        offline - is unreachable ( Bool )
        latency - duration ( Number )
--redis
 Test redis kv database, attributes:
        online - PING-PONG success ( Bool )
        offline - PING-PONG failed ( Bool )
--db
 Check database ( pgsql, mysql ), attributes:
        offline - has connection errors ( Bool )
        online - no connection errors ( Bool )
```

### Development
```
docker-compose up -d
go test -v -coverprofile=coverage.out ./...
```
