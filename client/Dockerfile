FROM ghcr.io/cirruslabs/flutter:3.29.3 AS build-web

WORKDIR /client

COPY ./pubspec.yaml .
COPY ./analysis_options.yaml .

RUN flutter pub get

COPY ./lib/ ./lib/
COPY ./web/ ./web/

RUN flutter build web --release