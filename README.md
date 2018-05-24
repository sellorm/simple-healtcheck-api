# simple-healtcheck-api

A really simple API, written in Go, to provide status information about the host system. Useful for monitoring purposes.

## Starting the service

In the simple-healthcheck-api directory, run the following:

```
./simple-healtcheck-api
```

## Dashboard

Once started, the service presents a web based dashboard to the user. This is created using the 'www/dashboard.html'.

The dashboard is available by visiting `http://localhost:1234/`.

## API

The dashboard gets its data from the API, which is available at `http://localhost:1234/__healthcheck__`.

The API returns JSON is the following format:

```
{cpu: 5, load: 1.97, mem: 61.24, version: "0.1"}
```
