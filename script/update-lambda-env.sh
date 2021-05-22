#!/bin/bash


token=$(jq -r '.ZOOM_AUTH_TOKEN' < ../secret/secret.json)

aws lambda get-function --function-name "${FN_NAME}" | jq -c '.Configuration.Environment.Variables' > "/tmp/${FN_NAME}-env-old"
aws lambda update-function-configuration --function-name "${FN_NAME}" --environment="{\"Variables\":$(jq -c ".ZOOM_AUTH_TOKEN=\"${token}\"" < "/tmp/${FN_NAME}-env-old")}" > /dev/null
