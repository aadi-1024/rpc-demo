FROM ubuntu:jammy
LABEL authors="aaditya"

COPY mailer .
RUN chmod +x ./mailer

ENTRYPOINT ["./mailer"]