# Use official Node.js image as base
FROM --platform=linux/amd64 node:latest as build

# Set working directory
WORKDIR /app

# Copy package.json and yarn.lock to install dependencies
COPY package.json yarn.lock ./

# Install dependencies
RUN yarn install

# Copy the rest of the application code
COPY . .

# Build the React app
RUN yarn build

# Stage 2 - Serve the React application using Nginx
FROM --platform=linux/amd64 nginx:alpine

# Copy build files from previous stage to nginx directory
COPY --from=build /app/dist /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]