# image-temperature
Adjust the color of the image by temperature 

it was written using GO

## run this before any others
```
go mod tidy
```

## run as script
```
go run cmd/main.go -input path/image-input-here.jpg -output path/image-output-here.jpg -temperature 10
```

## run after build
```
go build -o imgadjuster cmd/main.go
```
then
```
./imgadjuster -input path/image-input-here.jpg -output path/image-output-here.jpg -temperature 10
```
