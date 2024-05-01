ARG GO_IMAGE_NAME
ARG GO_IMAGE_VERSION

FROM ${GO_IMAGE_NAME}:${GO_IMAGE_VERSION}

# Set working directory
WORKDIR /home/www

RUN go install github.com/cosmtrek/air@latest

# Expose port 8000
#EXPOSE 8000