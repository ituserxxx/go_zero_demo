curl --location --request POST 'http://172.16.9.103:9880/api/http/hello/list' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--header 'Accept: */*' \
--header 'Host: 172.16.9.103:9880' \
--header 'Connection: keep-alive' \
--data-raw '{
"page":1,
"pageSize":2
}'

