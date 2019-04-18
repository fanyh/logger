# logger
mini logger library to coverage golang log

## features
* support log setting log level
* time is lmicroseconds
* show logger detail source
* support f and ln output
* auto convert root patch when in bin 
* show log_level in out

# install
go get github.com/fanyh/loggrt

# set log level

```
    [ --log level]
```

# output

```
2018/12/05 18:02:35.980993 [INF @./dh64/main.go:76] **
2018/12/05 18:02:35.981090 [DBG @./main.go:28] **

```
