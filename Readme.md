
# How secure is AES-ECB?

## Here is our input picture!
![RC In](rc.png)

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
![RC out](enc.png)


## Notes
1. _RC logo_ belongs to [Recurse Center](https://recurse.com)
2. Readme has PNG versions of ppm images created with `convert name.ppm name.png`
