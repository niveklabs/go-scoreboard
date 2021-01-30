# Go Scoreboard

Written using Golang and TinyGo!

## Device 

Wemos D1 Mini

```
tinygo flash -target=d1mini main.go
```

## Console

```
ls /dev/tty.*
screen /dev/tty.usbserial-14110 115200
```