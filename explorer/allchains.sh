#!/bin/bash

for ((i=1; i<=227; i++)); do
   echo $i
   OFF=$((i*15))
   echo $OFF
   echo "https://apiplus-api-prod-mainnet.factom.com/v2/chains?limit=15&offset=${OFF}&stages=factom%2Cbitcoin"
	curl "https://apiplus-api-prod-mainnet.factom.com/v2/chains?limit=15&offset=${OFF}&stages=factom%2Cbitcoin" -H 'Origin: https://explorer.factom.com' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-US,en;q=0.9' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36' -H 'content-type: application/json' -H 'accept: application/json' -H "Referer: https://explorer.factom.com/chains?page=$i" -H "factom-provider-token: $1" -H 'Connection: keep-alive' --compressed >> out
done
