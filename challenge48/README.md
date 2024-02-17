# [Challenge 48](https://codingchallenges.substack.com/p/coding-challenge-48-data-privacy)

## Pre-requisite
[Install Docker](https://www.docker.com/products/docker-desktop/)

## To run the service

`git clone https://github.com/Prajwalgn-07/JOHN_CRICKETT_CODING_CHALLENGES.git`

`cd challenge48`

`make run`

### To get the token
```
curl --location 'https://localhost:8080/token' \
--header 'Content-Type: application/json' \
--data '{
    "Username":"prajwal",
    "Password":"1234"
}'
```
### To encode the data
```
curl --location 'https://localhost:8080/tokenize' \
--header 'Authorization: <Token>' \
--header 'Content-Type: application/json' \
--data '{
    "ID":"test23",
	"data": {
		"field1": "value1wee",
		"field2": "value2",
		"fieldn": "valuen"
	}
}'
```
### To decode the data
```
curl --location 'https://localhost:8080/detokenize' \
--header 'Authorization: <Token>' \
--header 'Content-Type: application/json' \
--data '{
    "id": "test23",
    "data": {
        "field1": "XXf1uYhIyEFRd1YQx4z+2w==",
        "field2": "Z56OcZAcOj0tEojLpjxe8w==",
        "fieldn": "i0acSrv/cNtqcvQ7Cuulbw=="
    }
}'
```
### To clean restart
`make restart`

### To clean
`make clean`