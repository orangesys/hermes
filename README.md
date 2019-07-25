# Janus

janus is orangesys api

```sh
cp .env-sample .env
export $(cat .env|xargs)
```

```sh
http POST localhost:8080//api/v1/user email=foobar1@example.com planid=promeunit \
  'companyname=FooBar Inc.' cardnumber=5555555555554444 expmonth=11 expyear=23 cvc=321 \
  subdomain=thanos
```
