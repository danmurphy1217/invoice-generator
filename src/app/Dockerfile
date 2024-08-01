# Use an official Node.js image as the base
FROM node:22 AS build

# Set working directory
WORKDIR /app

# Copy pnpm lockfile and package.json
COPY pnpm-lock.yaml ./
COPY package.json ./

# Install pnpm
RUN npm install -g pnpm

# Install dependencies
RUN pnpm install

# Copy the rest of the application
COPY . .

EXPOSE 5173

# Run the application
CMD ["pnpm", "run", "dev","--","--host","0.0.0.0"]

