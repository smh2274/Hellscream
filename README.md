# Hellscream
## file server

![image](https://tse2-mm.cn.bing.net/th/id/OIP.UJkwAzrKMmv5TBxF_Ro37wAAAA?pid=Api&rs=1)

### this project used [gin](https://github.com/gin-gonic/gin)

##### step 1
> clone this repository && cd this project
 
##### step 2
> docker network -d bridge $(your network)

##### step 3
> docker build . -t hellscream

##### step 4
> docker  run --name hellscream -v $(your file path):/Azeroth/Hellscream/file -v $(your config path):/Azeroth/Hellscream/config --network $(your network) --network-alias hellscream -it -d hellscream

--- 
       
#### configure 

* this file server divided into two parts, you can configure hellscream_conf.yaml
  * need auth
    *  you add map struct into file.protect, the key is url, value is file path
    
    * for example:
       ```yaml
        file:
          protect
            avatar:
              /Azeroth/Hellscream/file/avatar
       ```
       you can access 127.0.0.1:8088/hellscream/protect/avatar/xxx.txt with jwt
  
  * open access
    *  you add map struct into file.public, the key is url, value is file path
        
     * for example:
        ```yaml
            file:
              public
                img:
                  /Azeroth/Hellscream/file/img
        ```
        you can access 127.0.0.1:8088/hellscream/public/img/xxx.jpg without jwt
      

* if you want change access network gateway, you can do some change on envoy.yaml
