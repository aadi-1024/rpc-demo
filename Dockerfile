FROM ubuntu:jammy
LABEL authors="aaditya"

COPY MailHog .

EXPOSE 8025
EXPOSE 1025

RUN chmod +x ./MailHog

ENTRYPOINT "./MailHog"