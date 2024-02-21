# image-temperature-adjustment

## Introduction
this application for adjustment images in API Base
## Package list

Packages which use in this project

#### 👨‍💻 Full list what has been used:
1. [swag](https://github.com/swaggo/swag) Swagger for Go<br/>
2. [Beego](https://github.com/beego/beego) Framework fro Go<br/>
3. [go-playground](https://github.com/go-playground) - Handling Custom Validation, and Translations
4. [resize-image](https://github.com/nfnt/resize) - Resized the image
5. [logging](go.uber.org/zap) - logging
6. [image-coloring] (https://go.dev/blog/image) - image coloring

### Clean Architecture
This project has  4 Domain layer :

 * Models Layer
 * Repository Layer
 * Usecase Layer  
 * Delivery Layer

#### The diagram:

![golang clean architecture](https://github.com/bxcodec/go-clean-arch/raw/master/clean-arch.png)

The explanation about this project's structure  can read from this medium's post : https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047

### How To Run This Project in local use Docker

```bash
# Clone into YOUR directory
git clone https://github.com/radyatamaa/image-temperature-adjustment.git

#move to project
cd image-temperature-adjustment

# if you not installed yet golang can use docker compose 
docker compose -f "docker-compose.yml" up -d --build

# Open at browser this url
http://localhost:8082/swagger/index.html
```

### How To Run This Project in local use Golang

```bash
#move to directory
cd $GOPATH/src/github.com/radyatamaa

# Clone into YOUR $GOPATH/src
git clone https://github.com/radyatamaa/image-temperature-adjustment.git

#move to project
cd image-temperature-adjustment

# Run app 
go run main.go

# if you not installed yet golang can use docker compose 
docker compose -f "docker-compose.yml" up -d --build

# Open at browser this url
http://localhost:8082/swagger/index.html
```


### Swagger UI:
http://localhost:8082/swagger/index.html
![swagger-image](https://github.com/radyatamaa/image-temperature-adjustment/swagger-image.png)

### More about app details test:
![documentation](https://github.com/radyatamaa/image-temperature-adjustment/document-test-cases-result.pdf)