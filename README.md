# WebCheck

WebCheck is a Go library for validating common resources found on the internet such as IP Addresses, URLs, and Email Addresses.

Documentation is [here](https://godoc.org/github.com/kochie/webcheck)

## Getting Started

To use this library you can install it from a terminal using Go.
```bash
go get github.com/kochie/webcheck
```
Or if your project is using `dep`
```bash
dep ensure --add github.com/kochie/webcheck
```

## Currently Supported Validators
At this moment there are several validators available. Including:
- URLs
- Email Addresses
- IPv4 Addresses
- IPv6 Addresses

## Contributing

If there is a resource you would like to add feel free to make an issue, or better yet, fork the repo and make a PR. Just remember to write a few tests to make sure the validators work as expected.
