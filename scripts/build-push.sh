
######
## !! Run this script in root of the project
#######

docker build -t nofrostoo/todo-app:dev ./backend
docker push nofrostoo/todo-app:dev

docker build -t nofrostoo/todo-app-frontend:dev ./frontend
docker push nofrostoo/todo-app-frontend:dev
