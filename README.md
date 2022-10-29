# hypercheck
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

### Example

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

If you have databases and api credentials you your environment variables, you can simply run:
```
~ hypercheck --auto
```

### Development

```
docker-compose up -d
go test -v -coverprofile=coverage.out ./...
```