### RUN DEVELOPMENT:
##### - web service
``go run web/cmd/main.go``

##### - balance service
``go run balance/cmd/main.go``

##### - threshold service
``go run threshold/cmd/main.go``

### API
``POST : http://localhost:8080/deposit``

``GET : http://localhost:8080/detail?wallet_id=<id>``