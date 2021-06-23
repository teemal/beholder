docker-compose -f ../persist/docker-compose-dynamodb-local.yaml down 
docker stop $(docker ps -a -q --filter="name=wonderful_black")