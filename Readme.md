# Gift exchanger


Golang application using Gin. use `Makefile` to interract with project


Requirement:

- Golang 1.22

### Getting started

- Run

```bash
make local.run
```
- Build image 

```bash
make docker.build
```

see `Makefile` for more.



### API

| HTTP Method & Endpoint       | Handler                                                  |
|------------------------------|----------------------------------------------------------|
| GET /v1/health               | github.com/soub4i/giftsxchanger/pkg/api/v1.Healthcheck    |
| GET /v1/members              | github.com/soub4i/giftsxchanger/pkg/api/v1.Get            |
| POST /v1/members             | github.com/soub4i/giftsxchanger/pkg/api/v1.Create         |
| GET /v1/members/:id          | github.com/soub4i/giftsxchanger/pkg/api/v1.Fetch          |
| PUT /v1/members/:id          | github.com/soub4i/giftsxchanger/pkg/api/v1.Update         |
| DELETE /v1/members/:id       | github.com/soub4i/giftsxchanger/pkg/api/v1.Delete         |
| GET /v1/gift-exchange        | github.com/soub4i/giftsxchanger/pkg/api/v1.GetExchange    |
| POST /v1/gift-exchange       | github.com/soub4i/giftsxchanger/pkg/api/v1.Shuffle        |
