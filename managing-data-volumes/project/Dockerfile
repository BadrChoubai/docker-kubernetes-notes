# Build your image up from another base image
FROM node

# Subsequent commands will be run relative to WORKDIR
WORKDIR /app

COPY package.json /app

# Executes with the creation of our image
RUN npm install

# Which files from our machine should go into our image/container
COPY . /app

EXPOSE 80

# Mount an anonoymous volume for storage of application data
# VOLUME [ "/app/feedback" ]

# Executes with the starting of our container
CMD [ "npm", "start" ]
