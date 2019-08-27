# Janus

janus is orangesys api

```sh
cp .env-sample .env
export $(cat .env|xargs)
```

クレジットカード番号 |カードの種類
---|---
4111111111111111|Visa
4242424242424242|Visa
4012888888881881|Visa
5555555555554444|MasterCard
5105105105105100|MasterCard
378282246310005|American Express
371449635398431|American Express
30569309025904|Diner's Club
38520000023237|Diner's Club
3530111333300000|JCB
3566002020360505|JCB

## Create stripe consumer

```sh
http POST localhost:8080//api/v1/user email=foobar1@example.com planid=promeunit \
  'companyname=FooBar Inc.' cardnumber=5555555555554444 expmonth=11 expyear=23 cvc=321 \
  subdomain=thanos
```
