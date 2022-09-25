#!/bin/bash

if [[ -z "${DIREKTIV_TEST_URL}" ]]; then
	echo "Test URL is not set, setting it to http://localhost:9191"
	DIREKTIV_TEST_URL="http://localhost:9191"
fi

if [[ -z "${DIREKTIV_SECRET_apikey}" ]]; then
	echo "Secret apikey is required, set it with DIREKTIV_SECRET_apikey"
	exit 1
fi

if [[ -z "${DIREKTIV_SECRET_apisecret}" ]]; then
	echo "Secret apisecret is required, set it with DIREKTIV_SECRET_apisecret"
	exit 1
fi

if [[ -z "${DIREKTIV_SECRET_accesstoken}" ]]; then
	echo "Secret accesstoken is required, set it with DIREKTIV_SECRET_accesstoken"
	exit 1
fi

if [[ -z "${DIREKTIV_SECRET_accesstokensecret}" ]]; then
	echo "Secret accesstokensecret is required, set it with DIREKTIV_SECRET_accesstokensecret"
	exit 1
fi

docker run --network=host -v `pwd`/tests/:/tests direktiv/karate java -DtestURL=${DIREKTIV_TEST_URL} -Dlogback.configurationFile=/logging.xml -Dapikey="${DIREKTIV_SECRET_apikey}" -Dapisecret="${DIREKTIV_SECRET_apisecret}" -Daccesstoken="${DIREKTIV_SECRET_accesstoken}" -Daccesstokensecret="${DIREKTIV_SECRET_accesstokensecret}"  -jar /karate.jar /tests/v1.0/karate.yaml.test.feature ${*:1}