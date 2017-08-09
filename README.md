# LoyaltyOne Homework Project

## Outline

This project was essentially to create a bare bones platform for threaded conversations. It consists of two main components: a RESTful API back end written in Go and a front end UI/API client written with ES6.

The back end functions as a static file server and exposes two API endpoints. The first endpoint, located at `/echo`, responds to `POST` requests and simply returns the contents of the `POST` body as the response body.

The second endpoint is at `/text` and is responsible for management of text objects.

A `GET` request to this endpoint returns all existing text objects in their order of insertion as a JSON encoded array.

A `POST` request to this endpoint accepts a single JSON encoded text object, commits it to the database, and returns it after populating additional fields.

#### POST Request Object Format

```JSON
{
  "city": "Barcelona",
  "content": "This is a test message.",
  "username": "User1"
}
```

To support threaded conversations, this endpoint supports an optional `parent_id` field in the request which specifies the `id` of another text object to which the new one will belong.

#### POST Response Object Format

```JSON
{
  "city": "Barcelona",
  "content": "This is a test message.",
  "created_at": "2017-08-08T22:35:45.438185Z",
  "id": 32,
  "latitude": 41.3850639,
  "longitude": 2.1734035,
  "parent_id": null,
  "temperature": 20.82,
  "username": "User1"
}
```

## Building

First, install the Go tool chain and add `$GOPATH/bin` to your `$PATH`

Then execute the following to install dependencies and utilities

```
go get github.com/jteeuwen/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs/...
go get -v -t -d ./...
```

Front end code is located in `www/` and changes must be registered for compilation by running
```
go-bindata-assetfs www/...
```

Migrations are located in `db/migrations` and changes must be registered for compilation by running
```
cd db/migrations && go-bindata -pkg migrations .
```

## Deploying

All deployments are handled through CircleCI. See `.circleci/config.yml` for details.

### Port Forwarding Configuration

```
iptables -A INPUT -i eth0 -p tcp --dport 80 -j ACCEPT
iptables -A INPUT -i eth0 -p tcp --dport 8080 -j ACCEPT
iptables -A PREROUTING -t nat -i eth0 -p tcp --dport 80 -j REDIRECT --to-port 8080

iptables -A INPUT -i eth0 -p tcp --dport 443 -j ACCEPT
iptables -A INPUT -i eth0 -p tcp --dport 8443 -j ACCEPT
iptables -A PREROUTING -t nat -i eth0 -p tcp --dport 443 -j REDIRECT --to-port 8443
```
