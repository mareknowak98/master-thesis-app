FROM node:16.13.2-alpine

# make the 'app' folder the current working directory
WORKDIR /app

# copy both 'package.json' and 'package-lock.json' (if available)
COPY package*.json ./
RUN npm install

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY . .

EXPOSE 8080

RUN npm run build

# install gserve
RUN npm install -g serve

EXPOSE 8080

CMD [ "serve", "-s", "dist", "-l", "8080" ]


#commands:
# docker build -t mylearn .          - build
# docker run -p 8080:8080 mylearn         - run
# docker rm -vf $(docker ps -aq)         -delete all containers
# docker rmi -f $(docker images -aq)       -delete all images