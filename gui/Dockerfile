FROM node:alpine3.10
RUN mkdir -p /app/node_modules
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
EXPOSE 3000
COPY ./entrypoint.sh /bin/entrypoint.sh
RUN chmod +x /bin/entrypoint.sh
CMD [ "/bin/entrypoint.sh" ]