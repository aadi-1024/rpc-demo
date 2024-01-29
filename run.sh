CGO_BUILD=0 go build -C broker/ -o broker
CGO_BUILD=0 go build -C cache/ -o cache
CGO_BUILD=0 go build -C mailer/ -o mailer

docker compose up -d --build
#~/go/bin/MailHog