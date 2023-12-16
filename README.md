# shortly

A simple URL shortener.

- Accepts `POST /` requests with a URL in the body.
- Returns an alias for the given URL on `POST /` requests.
- Redirects `GET /<alias>` requests to the URL mapped by the `<alias>`.
