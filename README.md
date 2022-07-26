# TerrificProxy
HTTP reverse proxy that uses one basic config.txt file. Learning curve composed of copy paste.



# Config File
Format of the config.txt is World url, space, Local/world url to redirect
 *   FromUrl ToUrl
 *   mydomain.com 192.168.8.2:3303
 *   alex.mydomain.com 192.168.8.2:3300
 *   google.mydomain.com google.com





# Overview
Imagine that you have deployed a bunch of microservices or websites in one server/cluster. You want to access/redirect/rename from one location. Terrefic Proxy handles all from one location. Because it does one thing its super simple and super fast. 
![Terrific Proxy architecture](/docs/TerrificProxy-architecture-big.webp)


Traditional reverse-proxies require so much knowlage, config info and space and cpu power to make basic redirects. With Terrefic Proxy you just forget its there and when you want to publish new project, you just add a line ( or many, I like to buy similar domain names)
 
 
# Installation 
```
  [install go](https://go.dev/doc/install)
  git clone https://github.com/nerkn/TerrificProxy
  cd TerrificProxy

  edit config file 
  nano config.txt
  sudo go run terrific.go
```

  
# Todo 
- [ ] https://github.com/nerkn/TerrificProxy/issues/1  auto https support :drooling_face:
- [ ] auto config loading :robot:
- [ ] add docker, so users not bother :running_woman:
- [ ] benchmark, find techluencer and pomp them up :rooster:
- [x] make github repository  :cupcake:
- [ ] find supporters :hugs:
