FROM ubuntu:jammy
LABEL authors="aaditya"

COPY broker .
RUN chmod +x ./broker

ENTRYPOINT ["./broker"]