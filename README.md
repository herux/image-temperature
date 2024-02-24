# image-temperature
Adjust the color of the image by temperature
it was written using GO

## run as script
```
go run cmd/main.go -input path/image-input-here.jpg -output path/image-output-here.jpg -temperature 10
```

## run after build
```
go build -o imgadjuster cmd/main.go
```
```
./imgadjuster -input path/image-input-here.jpg -output path/image-output-here.jpg -temperature 10
```
