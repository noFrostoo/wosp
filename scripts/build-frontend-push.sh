
######
## !! Run this script in root of the project
#######

docker build -t nofrostoo/todo-app-frontend:dev ./frontend
docker push nofrostoo/todo-app-frontend:dev
