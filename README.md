# Janus

janus is orangesys api

```sh
cp .env-sample .env
export $(cat .env|xargs)
```

```sh
http POST localhost:8080//api/v1/user email=hogehoge3@example.com planid=promeunit \
  'companyname=Orangesys Inc.' cardnumber=4242424242424242 expmonth=11 expyear=23 cvc=123
```
