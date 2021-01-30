# A-chat-backend
A-chat-backend

# Nginx部署
```
server {   
  listen 80;   
  server_name chat.example.com;      

  location ~ /ws {      
    proxy_pass http://127.0.0.1:2022;      
    proxy_http_version 1.1;      
    proxy_set_header Upgrade $http_upgrade;      
    proxy_set_header Connection "upgrade";   
  }      

  location / {      
     proxy_pass http://127.0.0.1:2022;   
  }
}
```

# 推荐
## 项目banner制作
- [network-science](http://www.network-science.de/ascii/)
- [patorjk](http://patorjk.com/software/taag/#p=display&f=Graffiti&t=Type%20Something%20)
