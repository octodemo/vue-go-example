FROM ubuntu:18.04

COPY . .

ENV PATH "$PATH:/usr/local/go/bin:/home/user/go/bin:/home/user/.cargo/bin"
    
RUN curl -sL https://deb.nodesource.com/setup_13.x | bash -
RUN apt-get update && apt-get install -y nodejs
RUN apt-get install build-essential -y

RUN bash setup.sh
USER user
