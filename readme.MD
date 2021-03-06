<p align="center">
<h1 align="center">Shortened service</h1>
</p>

## This service can do:
 - Create,read and delete shortened URLs.
 - Redirect the shortned URL to original
 - Counts the redirects made
 - Show the total counter per redirect

## Tecnical Details
 - Programming language: Go 1.15
 - Web Framework: Gearbox :gear:
 - ORM Library: Gorm
 - Database: PostgreSQL
 - Log: ZeroLog


## Pattern to the service

![image](https://user-images.githubusercontent.com/9199380/176738028-bdd6ebfe-4108-448c-89a0-962f3235be90.png)

## General Diagram

![image](https://user-images.githubusercontent.com/9199380/176750568-670371cf-cf4b-440f-855c-8126e4e04f7b.png)


## Api Documentation

You can access to the documentation for the next url: http://localhost:8080/ce99870d

## API Consideration

### Postman Collection
In the folder postman there is a postman collection that you could use.

### Detail
for the get and delete link you have to replace ed2fc7c5 for your custom code, ed2fc7c5 is just a example.

## How use the redirection
At your web browser set the next url:
http://localhost:8080/{code} for example this code f2a13dbb (http://localhost:8080/f2a13dbb) send to you https://www.theguardian.com/international 

## Steps to running the service

- clone this repo

### Build the service's container image:
docker build -t yourbuildname .

### Start an app container
docker run -dp 8080:8080 yourbuildname

Open your postman or another favorite client, read the documentation and enjoy.

## Next Features:
 - Cache with redis
 - Observability and traceability with Datadog (LOG y APM)
 - Unit Testing
 - A new interface with gRPC