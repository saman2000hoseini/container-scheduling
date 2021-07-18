FROM alpine:3.14.0

ENV TZ=Asia/Tehran \
    PATH="/app:${PATH}"

RUN apk add --update --no-cache \
      gcc \
      musl-dev \
      build-base

WORKDIR /app

COPY . .

RUN g++ -o min min.cpp
RUN g++ -o max max.cpp
RUN g++ -o average average.cpp
RUN g++ -o sort sort.cpp
RUN g++ -o word_count word_count.cpp

RUN rm *.cpp