FROM ubuntu:18.04

COPY . .

ENV PATH "$PATH:/usr/local/go/bin:/home/user/go/bin:/home/user/.cargo/bin"

RUN apk add --no-cache py-pip python-dev libffi-dev openssl-dev gcc libc-dev make && pip install docker-compose
    
RUN curl -sL https://deb.nodesource.com/setup_13.x | bash -
RUN apt-get update && apt-get install -y nodejs

RUN bash setup.sh
USER user
