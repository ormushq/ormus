# Swagger

## Generate swagger

For generate swagger file run this command:

```shell
make swagger-generator
```
This command is installs swag if it is not installed before then run `swag format` command then generate swagger files.
At this time swagger is generated for two service:
- Source
- Manager

Each service has separate route to access on swagger service.
- [Manager](http://swagger.ormus.local/manager/index.html)
- [Source](http://swagger.ormus.local/source/index.html)

The main route of [swagger](http://swagger.ormus.local) service include all links of swagger services.

For mor detail of write operations in swag format see the swag [documentation](https://github.com/swaggo/swag).