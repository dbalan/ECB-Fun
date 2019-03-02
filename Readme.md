
# How secure is AES-ECB?

## Here is our input picture!
![RC In](rc.ppm)

## Run encrypt tool
```
go run encrypt.go
```
## Trick unix into believing this is an image file

Write the same magic number into the file
```
dd conv=notrunc if=rc.ppm of=enc.ppm bs=1 count=16
```

## There we go
![RC out](enc.ppm)
