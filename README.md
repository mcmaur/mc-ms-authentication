# ms-authentication
Authentication microservice written in Go #social-login #login #ms #microservice #golang

<br/>

### READY TO USE AUTHENTICATION MICROSERVICE

1. git clone https://github.com/mcmaur/mcm-ms-authentication.git
2. copy env.tom_example to env.toml
3. set social network keys inside env.toml file
4. docker-compose up -d
5. visit localhost:8080

<br/>
<br/>


# Todo
- dockerfile for building image
- write more tests
  - mock with httptest
- add auth protocol: add applications and scope
- refactor to improve readability and maintenance
- map for social key & password in config file?
