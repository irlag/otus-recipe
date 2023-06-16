ARG GOLANG_VERSION

FROM golang:${GOLANG_VERSION}-alpine AS build

WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ./

ARG VERSION
ENV CGO_ENABLED=0

RUN go build -ldflags "-s -X otus-recipe/cmd.Version=$VERSION" -v -a -o /bin/otus-recipe main.go

FROM scratch
COPY --from=build /bin/otus-recipe /bin/otus-recipe
COPY --from=build /app/migrations /migrations

ENTRYPOINT ["/bin/otus-recipe"]