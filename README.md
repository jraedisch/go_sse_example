About

This is a simple example of proxying Server-Sent Events via Caddy proxy to a GopherJS front end.

Setup

- $ go get -u github.com/jraedisch/go_sse_example
- $ cd $GOPATH/src/github.com/jraedisch/go_sse_example
- $ openssl req -x509 -nodes -days 365000 -newkey rsa:2048 -keyout selfsigned.key -out selfsigned.crt
- add streaming.example.com to /etc/hosts
- http://www.robpeck.com/2010/10/google-chrome-mac-os-x-and-self-signed-ssl-certificates
- $ go get -u github.com/mholt/caddy/caddy
- $ go get -u github.com/gopherjs/gopherjs
- $ go get -u honnef.co/go/js/dom
- $ gopherjs build -mwv client/client.go -o static/client.min.js
- $ go run server.go
- $ sudo caddy
- open https://streaming.example.com
- done

Misc

- There is no hot reloading, but GopherJS will watch for file changes.
  Changes to webapp.go require a restart.
- Do not vendor GopherJS, since it will result in hard to decipher bugs.
  See https://github.com/gopherjs/gopherjs/issues/415.
- For live deployment you need to provide a separate Caddyfile, that could be much simpler.
  See https://caddyserver.com/docs/getting-started.
- Production ready projects should probably add some specs :)
