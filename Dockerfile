FROM golang

# Create app directory
WORKDIR /go/src/github.com/jacobrs/Shopify2019-Intern-Challenge

# Install app
COPY . .

# Install vendor manager
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Install dependencies
RUN dep ensure

RUN go build -o shop

EXPOSE 3000
CMD ["./shop"]
