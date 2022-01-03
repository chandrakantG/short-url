// run following commands
go mod tidy
sudo docker build --tag short-url .
sudo docker run -d -p 8082:8082 short-url

encode Url:
Request:
curl --location --request POST 'localhost:8082/urlEncoder' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url" :"https://play.golang.com/"
}'

Response:
{
    "status": 200,
    "success": true,
    "data": "http://localhost:8082/N0Q9H0NdYuk"
}

decode Url :
Request:
curl --location --request GET 'http://localhost:8082/N0Q9H0NdYuk'

Response:
{
    "status": 200,
    "success": true,
    "data": "https://play.golang.com/"
}

Test cmd:
/usr/bin/go test -timeout 30s -v -run ^TestEncode$ short-url/urlshortner

/usr/bin/go test -timeout 30s -v -run ^TestDecode$ short-url/urlshortner