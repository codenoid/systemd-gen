# systemd-gen

![cool systemd logo](https://raw.githubusercontent.com/codenoid/systemd-gen/master/systemd-logo.png)

## Download

1. Download binary (Built on Ubuntu 18.04) from [release page](https://github.com/codenoid/systemd-gen/releases/download/0.0.1/systemd-gen)
2. `go get -u github.com/codenoid/systemd-gen`

```bash
sudo wget https://github.com/codenoid/systemd-gen/releases/download/0.0.1/systemd-gen
sudo mv systemd-gen /usr/bin
# alright
```

## How to use it

```bash
# username is logged username
$ sudo systemd-gen service-name username '/path/to/project/dir' '/usr/bin/python3 something.py'
Service File Created.
$ sudo service service-name status
● service-name.service - service-name Service                                                 
   Loaded: loaded (/etc/systemd/system/service-name.service; enabled; vendor preset: enabled)  
   Active: active (running) since Fri 2020-01-10 16:43:29 WIB; 14min ago                        
 Main PID: 3667 (python3)                                                                       
    Tasks: 2 (limit: 4915)                                                                      
   CGroup: /system.slice/service-name.service                                                  
           ├─3667 /usr/bin/python3 something.py                                                     
           └─3683 /usr/bin/python3 something.py                                                     
                                                                                                
Jan 10 16:43:29 host-name systemd[1]: Started service-name Service.                        
```