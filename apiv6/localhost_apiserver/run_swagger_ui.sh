#!/usr/bin/env bash

bash -c "sleep 1; firefox localhost" & # delayed launch

docker run --rm --name cloudexport_swagger -p 80:8080 -e SWAGGER_JSON="/local/cloud_export.openapi.yaml" -v "$(pwd)/../api_spec/openapi_3.0.0/":/local swaggerapi/swagger-ui 