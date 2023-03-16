FROM golang:1.20.2-alpine

# Install air for live-reloading
RUN go install github.com/cosmtrek/air@latest

# Set the working directory to /usr/src/app
WORKDIR /usr/src/app

# Install packages
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the code into the container
COPY . .

# Set the environment variable for air
ENV AIR_WD /usr/src/app

# Start the air tool with live-reloading
CMD ["air"]