FROM node:8.11

# Add the libraries
RUN mkdir -p /usr/src/libraries/javascript
WORKDIR /usr/src/libraries/javascript
COPY ./libraries/javascript .
RUN npm install

# Create app directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# Install app dependencies
RUN npm install -g @vue/cli@3.5.5
COPY ./web.client/.npmrc .
COPY ./web.client/package.json .
RUN npm install

# Expose ports for web access and debugging
EXPOSE 8080 9229

CMD [ "npm", "run", "serve" ]
