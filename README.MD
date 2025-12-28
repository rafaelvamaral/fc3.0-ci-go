# SONAR

docker network create sonar-network
docker network connect sonar-network sonarqube
docker run --rm --network="sonar-network" -e SONAR_HOST_URL="http://sonarqube:9000" -v "$(pwd):/usr/src" sonarsource/sonar-scanner-cli -Dsonar.token=<TOKEN> -Dsonar.projectKey=go-project

## SONAR WITH PROPERTIES FILE

docker run --rm   --network="sonar-network"  -v "$(pwd):/usr/src"   sonarsource/sonar-scanner-cli