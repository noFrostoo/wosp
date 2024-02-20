
######
## !! Run this script in root of the project
#######

docker build -t nofrostoo/todo-app:dev ./backend
docker push nofrostoo/todo-app:dev
