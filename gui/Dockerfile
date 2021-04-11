FROM node:alpine3.10
RUN mkdir -p /app/node_modules
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
ARG NODE_ENV=production
ENV NODE_ENV $NODE_ENV
EXPOSE 3000
COPY ./entrypoint.sh /bin/entrypoint.sh
RUN chmod +x /bin/entrypoint.sh
RUN npm install -g npm@7.9.0
CMD [ "/bin/entrypoint.sh" ]