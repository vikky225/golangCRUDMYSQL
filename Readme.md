
<p align="center">
  <a href="" rel="noopener">
 <img src="https://djeqr6to3dedg.cloudfront.net/repo-logos/library/golang/live/logo.png" alt="Project logo"></a>
</p]>
<h3 align="center">GOLANG CRUD MOVIE WITH MYSQL DB (DOCKERIZED)</h3>

## modules and package to be installed 
``` bash 
go mod init example.com/go-bookstore
go get "github.com/jinzhu/gorm"
go get "github.com/jinzhu/gorm/dialects/mysql"
go get "github.com/gorilla/mux" 
```

## Please run command from the root package GO-BOOKSTORE
```docker-compose up```  

to make sure mysql is up and running you should have Docker installed in your maching in order to test.

## Please run command from cmd package for build and run
```bash  go build```

```bash  go run .```

## Go ahead and test files inside testing folder testendpoints.http or use postman to configure same endpoints

## Endpoints details

<details>
### Creating the book


POST http://localhost:9010/api/bookstore/books

```json
Content-Type: application/json

{
    "Name": "Zero to One",
    "Author": "post1Authocccr",
    "Publication": "Animal"
}
```

#### Getting all the books
GET http://localhost:9010/api/bookstore/books


#### Getting book by id

GET http://localhost:9010/api/bookstore/books/5




#### Delete book by id 

DELETE http://localhost:9010/api/bookstore/books/


#### Update book by id 

PUT http://localhost:9010/api/bookstore/books/6

Content-Type: application/json

```json
{

  "name": "Zero to One",
  "author": "updating",
  "publication": "Animal"
}
</details>

### Notes:
- Ensure that the URLs (`http://localhost:9010`) match your actual application configuration.
- Replace placeholder values (like book IDs) with real data as per your application's behavior.
- The code blocks for JSON and bash commands are formatted using triple backticks (\`\`\`) followed by the language identifier (`json` or `bash`), ensuring they display correctly with syntax highlighting in markdown previews.
