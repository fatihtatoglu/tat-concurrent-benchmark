# Alpine Linux tabanlı Node.js image
FROM node:alpine

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json files to the container's workspace
#COPY package*.json ./

# Install dependencies
#RUN npm install

# Copy the local package files to the container's workspace
COPY . .

# Command to run the application
CMD ["node", "index.js"]
