# hypercheck
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![DeepSource](https://deepsource.io/gh/IMMORTALxJO/hypercheck.svg/?label=active+issues&show_trend=true&token=SaCVhzg7Sci39dpzTEGdpLsS)](https://deepsource.io/gh/IMMORTALxJO/hypercheck/?ref=repository-badge)

Single binary to test everything you need.

Supported checks:
- auto
- http
- tcp
- file system
- dns
- redis
- db
  - mysql
  - postgres

### Usage
```
--fs
 Filesystem files check, attributes:
        size - files size ( List[Number] )
        regular - is regular file ( List[Bool] )
        user - files username ( List[String] )
        exists - at least one file found ( Bool )
        dir - is directory ( List[Bool] )
        uid - files UID ( List[Number] )
        gid - files GID ( List[Number] )
        group - files groupname ( List[String] )
        count - files count ( Number )
--http
 Check http resource, attributes:
        code - response status code ( Number )
        content - response content ( String )
        online - is online ( Bool )
        offline - is offline ( Bool )
        headers - headers content ( List[String] )
--tcp
 Check tcp port, attributes:
        latency - duration ( Number )
        online - is reachable ( Bool )
        offline - is unreachable ( Bool )
--dns
 Check dns query response, attributes:
        CNAME - CNAME record content ( String )
        online - A record is not empty ( Bool )
        offline - A record is empty ( Bool )
        A - A record content ( List[string] )
        NS - NS record content ( List[string] )
        TXT - TXT record content ( List[string] )
        MX - MX record content ( List[string] )
--tcp
 Check tcp port, attributes:
        online - is reachable ( Bool )
        offline - is unreachable ( Bool )
        latency - duration ( Number )
--redis
 Test redis kv database, attributes:
        online - is reachable ( Bool )
        offline - is unreachable ( Bool )
--db
 Check database ( pgsql, mysql ), attributes:
        online - is reachable ( Bool )
        offline - is unreachable ( Bool )
--auto
 Generate probes automaticaly based on current environment variables ( List )
```
### Examples

If you have connections in your environment variables, you can simply run:
```
~ env
...
API_ENDPOINT=https://postman-echo.com/basic-auth
API_USER=postman
API_PASS=passowrd
...
~ hypercheck --auto

```
> automatic dependency detection powered by [scheme-detector](https://github.com/IMMORTALxJO/scheme-detector)

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

Checking 'online' mysql://user:password@localhost:3306/database ...
        ✅  DB online mysql://user:password@localhost:3306/database
Checking 'online' postgres://user:password@localhost:5432/postgres?sslmode=disable ...
        ✅  DB online postgres://user:password@localhost:5432/postgres?sslmode=disable
Checking 'code==200,headers:count>1' https://postman-echo.com/status/200 ...
        ✅  HTTP code==200 https://postman-echo.com/status/200
        ✅  HTTP headers:count>1 https://postman-echo.com/status/200
Checking 'online,A:count>1' google.com ...
        ✅  DNS online google.com
        ✅  DNS A:count>1 google.com
Checking 'regular' ./README.md ...
        ✅  FS regular ./README.md
Checking 'online' localhost:6379 ...
        ✅  redis online localhost:6379
Checking 'online,latency<1s' 1.1.1.1:53 ...
        ✅  TCP online 1.1.1.1:53
        ✅  TCP latency<1s 1.1.1.1:53
```

### Development
```
docker-compose up -d
go test -v -coverprofile=coverage.out ./...
```
