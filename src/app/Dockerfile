# Use an official node image as a base
FROM node:18

# Install pnpm globally
RUN npm install -g pnpm

# Set the working directory
WORKDIR /app

COPY package.json pnpm-lock.yaml vite.config.ts ./

# Install dependencies using pnpm
RUN pnpm install

# Copy the rest of the application code
COPY . .

# Expose the port your app runs on
EXPOSE 5173

# Start the application using pnpm
CMD ["pnpm", "run", "dev"]
