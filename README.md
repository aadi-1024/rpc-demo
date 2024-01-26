### A demo project showcasing RPC in Golang 

Consists of three microservices:
- Broker: accepts HTTP JSON requests and forwards the request via RPC to the required microservice
- Mailer: send email
- Cache: stores user id and email in memory for quick access