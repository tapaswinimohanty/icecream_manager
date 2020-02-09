# ZALORA ICECREAM API

### App architecture:
- Language: Go
- Database: Postgres
- RESTful API

###  API
[ZALORA ICECREAM API](https://localhost/api/v1/products)

### Setup
- Install Postgres database


```
- Create new postgres database with name: `icecream`.
```
docker exec -it postgres psql -U postgres
CREATE DATABASE icecream;
```



#### Seed Data
Set environment variable `Seed=true` to active seed operation.
This operation will **clean database**, load data from `database/data/sample.json` file and **insert to tables**.
```
docker run --rm --name zalora \
-e DBHost=localhost \
-e DBName=icecream \
-e DBUser=postgres \
-e DBPassword=postgres \
-e DBPort=5432 \
-e Seed=true \
-p 8080:8080 \
--network host \
-it lynam/zalora:1.0
```

### Build from sourcecode with docker
- Clone sourcecode
```
git clone https://bitbucket.com/libertywireless/icecream_manager
cd zalora_icecream
```
- Build with docker (change `lynam/zalora:1.0` to whatever you want)
```
docker build -t lynam/zalora:1.0 .
```
- Run with docker as above

### API Document
#### Login
Default account:
```
username: testUser
password: testPassword
```
Request:
```
curl -X POST \
  http://localhost:8080/api/v1/login \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
	"username": "testUser",
	"password": "testPassword"
}'
```
result:
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDg4NzIyNzYsIm5hbWUiOiJ6YWxvcmEifQ.HT0-oZF5Y4cVdYYOKKnJy_DI2G4uheMmyPh7iDcWylY"
}
```
#### Get product list
```
curl -X GET \
  http://localhost:8080/api/v1/products \
  -H 'Content-Type: application/json'
```
#### Get product by id
Change `2190` to your product id.
```
curl -X GET \
  http://localhost:8080/api/v1/products/2190 \
  -H 'Content-Type: application/json'
```

#### Insert product (need login)
Change the `bearer-token-here` with the login token you received in `login` step.
```
curl -X POST \
  http://localhost:8080/api/v1/products \
  -H 'Authorization: Bearer bearer-token-here' \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F 'name=kem cafe' \
  -F 'image_open=@path_to_image.png' \
  -F 'description=kem cafe' \
  -F 'story=kem cafe ngon lem' \
  -F 'allergy_info=có kem, có cafe' \
  -F 'dietary_certifications=kem, cafe' \
  -F 'sourcing_value_ids=1,2' \
  -F 'ingredient_ids=1,2' \
  -F 'image_closed=@path_to_image.jpg'
```

#### Update product (need login)
Change `2190` to your product id.
Change the `bearer-token-here` with the login token you received in `login` step.
```
curl -X PUT \
  http://localhost:8080/api/v1/products/2190 \
  -H 'Authorization: Bearer bearer-token-here' \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F 'name=kem cafe' \
  -F 'description=kem cafe 2' \
  -F 'story=kem cafe ngon lem 2' \
  -F 'allergy_info=có kem, có cafe 2' \
  -F 'dietary_certifications=kem, cafe 2' \
  -F 'sourcing_value_ids=1,2' \
  -F ingredient_ids=4
```

#### Delete product (need login)
Change `2190` to your product id.
Change the `bearer-token-here` with the login token you received in `login` step.
```
curl -X DELETE \
  http://localhost:8080/api/v1/products/2190 \
  -H 'Authorization: Bearer bearer-token-here' \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache'
```


