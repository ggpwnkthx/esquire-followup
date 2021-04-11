FROM codercom/code-server
USER root
RUN curl -sL https://deb.nodesource.com/setup_15.x | bash -
RUN apt-get update && apt-get install -y gcc nodejs
RUN curl -L -o /tmp/golang.tar.gz https://golang.org/dl/go1.16.3.linux-amd64.tar.gz
RUN mkdir -p /usr/local/go
RUN tar -C /usr/local -xzf /tmp/golang.tar.gz
RUN rm -r /tmp/*
ENV PATH="${PATH}:/usr/local/go/bin"
USER coder