FROM ubuntu:18.04

COPY . .

ENV PATH "$PATH:/usr/local/go/bin:/home/user/go/bin:/home/user/.cargo/bin"

RUN apt-get update -q \
	&& apt-get install -y -q --no-install-recommends curl ca-certificates \
	&& curl -o /usr/local/bin/docker-compose -L \
		"https://github.com/docker/compose/releases/download/${COMPOSE_VERSION}/docker-compose-Linux-x86_64" \
	&& chmod +x /usr/local/bin/docker-compose
    
RUN curl -sL https://deb.nodesource.com/setup_13.x | bash -
RUN apt-get update && apt-get install -y nodejs
RUN apt-get install build-essential -y

RUN bash setup.sh
USER user
