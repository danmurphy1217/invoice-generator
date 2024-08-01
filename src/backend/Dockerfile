# Use the official Golang image as the builder
FROM golang:1.22 as builder

# Set the working directory
WORKDIR /app

# Copy the Go application source code
COPY . .
RUN go mod download

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
RUN ["chmod", "775", "main"]

# Use a minimal image for the runtime environment
FROM alpine:latest

# Install ca-certificates, bash, and curl
RUN apk --no-cache add ca-certificates curl bash python3 py3-pip

# Install WeasyPrint dependencies
RUN apk add --no-cache \
    py3-weasyprint \
    py3-cffi \
    py3-pillow \
    py3-cssselect \
    py3-html5lib \
    py3-lxml \
    py3-pygments \
    py3-tinycss2 \
    py3-cairosvg \
    py3-packaging \
    py3-bleach \
    py3-webencodings \
    py3-soupsieve \
    cairo \
    pango \
    gdk-pixbuf \
    libffi-dev \
    libxml2-dev \
    libxslt-dev \
    fontconfig

# Install additional fonts
RUN apk add --no-cache \
    msttcorefonts-installer \
    font-noto \
    font-noto-cjk \
    font-noto-emoji \
    ttf-dejavu \
    && update-ms-fonts \
    && fc-cache -f

# Set the working directory
WORKDIR /root

# Copy the built Go binary from the builder
COPY --from=builder /app .

# Expose the port on which the application runs
EXPOSE 8080

# Command to run the application
CMD "./main"