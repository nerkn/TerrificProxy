# TerrificProxy
HTTP reverse proxy that uses one basic config.txt file. Learning curve composed of copy paste.



# Config File
Format of the config.txt is World url, space, Local/world url to redirect
FromUrl ToUrl
mydomain.com 192.168.8.2:3303
alex.mydomain.com 192.168.8.2:3300
google.mydomain.com google.com





# Overview
Imagine that you have deployed a bunch of microservices or websites in one server/cluster. You want to access/redirect/rename from one location. Terrefic Proxy handles all from one location. Because it does one thing its super simple and super fast. 

Traditional reverse-proxies require so much knowlage, config info and space and cpu power to make basic redirects. With Terrefic Proxy you just forget its there and when you want to publish new project, you just add a line ( or many, I like to buy similar domain names)
 
