# GolangTesseract

GolangTesseract is an easy-to-use web API on Golang for OCR with Tesseract. The main idea is to put all OCR stuff inside the API and not to worry about reusing it in different projects.
Developed as a part of [project](https://github.com/Robertg47/ConserseAIEnhanced) [Robertg47](https://github.com/Robertg47).

## Prerequisites
- Golang
- Docker
- Git

## Installation


To install and run the project, follow these steps:

- Clone the repository: 
```
git clone https://github.com/egorpechoring/tesseractgolang
```
- Navigate to root directory 
- Run throught go run or the script with dockerization: 

```
go run .
```
or
```
sh sh_scripts/dockerize_and_run.sh
```


## Usage

### /ocr
To use the API, make a POST request to http://localhost:8080/ocr with a form-data payload containing an image file.

* method POST
* link http://localhost:8080/ocr
* body is form-data
* image as a file

#### Request example
```
```
#### Response example
```
```

## Contributing

Pull requests are welcome.

## Authors 
* [jegor_petsorin](https://github.com/egorpechoring)


## License

This project is licensed under the [MIT](https://choosealicense.com/licenses/mit/)  License. 