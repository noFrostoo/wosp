name: Build and Run

on:
  push:
    branches:
      - main  # You can change this to match your main branch name

jobs:
  build-backend:
    runs-on: ubuntu-latest
    steps:
    - name: Checkout repository
      uses: actions/checkout@v2
      
    - name: Build Backend
      run: ./build-backend-push.sh
      working-directory: scripts  # Update this with the path to your backend directory
      
  build-frontend:
    runs-on: ubuntu-latest
    steps:
    - name: Checkout repository
      uses: actions/checkout@v2
      
    - name: Build Frontend
      run: ./build-frontend-push.sh
      working-directory: scripts  # Update this with the path to your frontend directory

