myapp: 
  docker build -t app ./app
  docker run --rm -it -p 80:80 app

clean:
  docker system prune
  
run: myapp
