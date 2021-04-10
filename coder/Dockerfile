FROM codercom/code-server
USER root
RUN curl -L -o /tmp/golang.tar.gz https://golang.org/dl/go1.16.3.linux-amd64.tar.gz
RUN mkdir -p /usr/local/go
RUN tar -C /usr/local -xzf /tmp/golang.tar.gz
RUN rm -r /tmp/*
ENV PATH="${PATH}:/usr/local/go/bin"
USER coder