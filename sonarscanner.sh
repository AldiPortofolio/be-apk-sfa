go test ./... -coverprofile=coverage.out --cover
docker run \
    --rm \
    -e SONAR_HOST_URL="http://34.101.147.203:9011" \
    -e SONAR_LOGIN="0aa344097a18447513b8b5b659196392b6ce9373" \
    -v "/Users/putri/go/src/ottosfa-api-apk:/usr/src" \
    sonarsource/sonar-scanner-cli
