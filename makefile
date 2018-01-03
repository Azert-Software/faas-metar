deploy-local:
	faas-cli build -f metar.yml && \
	faas-cli deploy -f metar.yml

test-local:
	echo -n "Echo Golf Alpha Charlie" | faas-cli invoke metar

deploy-test-local: deploy-local test-local
