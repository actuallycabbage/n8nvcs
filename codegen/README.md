# Codegen

## n8n-api
We use `oapi-codegen` to do the heavy lifting when it comes to writing the n8n client. The package that it produces is not very clean but it's certainly usable and saves us quite some time.

The tricky part is fetching the `openapi.yml` spec from n8n. It makes heavy use of the JSON Schema `$ref` directive (for the uninitiated, it's an extension that allows nested referencing to other files). `oapi-codegen` isn't really sure what to do about those as those nested references sit inside the n8n repo.

The root of their `openapi.yml` sits [here](https://github.com/n8n-io/n8n/master/packages/cli/src/PublicApi/v1). I've found a handy little tool that'll resolve all those pesky `$ref`s and produce a single file ready for use with `oapi-codegen` [swagger-cli](https://github.com/APIDevTools/swagger-cli). Navigate to that `openapi.yml` and run `npx swagger-cli bundle --type yaml --outfile openapi-bundle.yml openapi.yml`. Once you've figured that out, plop down that newly created file into `./codegen/n8nclient` and run `make generated`. Should be good to go.

