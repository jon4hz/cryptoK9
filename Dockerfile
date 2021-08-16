FROM golang:latest as builder
# add build dependencies
RUN apt-get update -qq
RUN apt-get install -y -qq libtesseract-dev libleptonica-dev
# build the binary
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN make build


# runtime image
FROM golang:latest

RUN apt-get update -qq
RUN apt-get install -y -qq libtesseract-dev libleptonica-dev
RUN apt-get install -y -qq \
  tesseract-ocr-eng

COPY --from=builder /app/cryptoK9 /app/
WORKDIR /app 
CMD [ "/app/cryptoK9" ]