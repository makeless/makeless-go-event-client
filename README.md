# Makeless - SaaS Ecosystem - Golang Event Client

[![Build Status](https://ci.loeffel.io/api/badges/makeless/makeless-go-event-client/status.svg)](https://ci.loeffel.io/makeless/makeless-go-event-client)

Event Client to test and debug Makeless Events (Server-Sent Events)

<img src="https://i.imgur.com/DTOS73n.gif" alt="command" width="600">

## Example

```bash
make build
./makeless-go-event-client \
  -url https://localhost:3000/api/auth/event \
  -token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Imx1Y2FzQGxvZWZmZWwuaW8iLCJlbWFpbFZlcmlmaWNhdGlvbiI6dHJ1ZSwiZXhwIjoxNjE0MTg5MDU2LCJvaG5nb040ZzdmQUd6WjlieXJkS01pTmtNYnFRek5tQyI6MSwib3JpZ19pYXQiOjE2MTQxODU0NTZ9.SvvUIvHOW_B7HLwdqt5zsYzrpxks6PLsj2_tVxp21bo
```