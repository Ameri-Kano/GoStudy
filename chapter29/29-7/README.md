## HTTPS 인증서 간단하게 만드는 방법

### 인증서 생성 방법
`openssl req -new -newkey rsa:2048 -nodes -keyout localhost.key -out localhost.csr`

### 인증서 인증 방법(실제 환경에서 사용하면 안됨)
`openssl x509 -req -days 365 -in localhost.csr -signkey localhost.key -out localhost.crt`