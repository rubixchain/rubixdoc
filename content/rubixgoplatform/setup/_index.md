---
alwaysopen: false
date: "2022-08-09T10:42:11.812Z"
description: Setup
head: <hr/>
title: Setup
weight: 4
---

The new Rubixgoplatform support command line options to run/configure the Rubix node. To run the application use the follwing format.

```
./rubixgoplatform <cmd> <options>

Use the following commands

                       -v : To get tool version

                       -h : To get help

                      run : To run the rubix core

                     ping : Use the command to ping the peer
```

Run Command
: To run the Rubix node use this command. 
```
./rubixgoplatform run -p node1 -n 0 -s

This following options are used to run the Rubix node
  -c string
        Configuration file for the core (default "api_config.json")
  -k string
        Config file encryption key (default "TestKeyBasic#2022")
  -n uint
        Node number
  -p string
        Working directory path (default "./")
  -s    Start the core
  -testNet
        Run as test net
  -testNetKey string
        Test net key (default "testswarm.key")
```
Ping Command
: To ping any peer in network use this command.
```
./rubixgoplatform ping -peerID 12D3KooWKr8dEQiLXuKacxDCZiHePVEMpgjxk19C3QozuUVQcQHA

This following options are used for this command
  -addr string
        Server/Host Address (default "localhost")
  -peerID string
        Peerd ID
  -port string
        Server/Host port (default "20000")
```
Add Bootstrap Command
: To add bootstrap to node use this command.
```
./rubixgoplatform addbootstrap -peers /ip4/103.60.213.76/tcp/4001/p2p/QmR1VH6SsEN1wf4EmstxXtNMvR35KEetbBetiGWWKWavJ6

This following options are used for this command
  -addr string
        Server/Host Address (default "localhost")
  -peers string
        Bootstrap peers, mutiple peers will be seprated by comma
  -port string
        Server/Host port (default "20000")
```
Remove Bootstrap Command
: To remove bootstrap from node use this command.
```
./rubixgoplatform removebootstrap -peers /ip4/103.60.213.76/tcp/4001/p2p/QmR1VH6SsEN1wf4EmstxXtNMvR35KEetbBetiGWWKWavJ6

This following options are used for this command
  -addr string
        Server/Host Address (default "localhost")
  -peers string
        Bootstrap peers, mutiple peers will be seprated by comma
  -port string
        Server/Host port (default "20000")
```
Remove All Bootstrap Command
: To remove all bootstrap from node use this command.
```
./rubixgoplatform removeallbootstrap

This following options are used for this command
  -addr string
        Server/Host Address (default "localhost")
  -port string
        Server/Host port (default "20000")
```
Enable Explorer Service Command
: To enable explorer service on the node use this command.
```
./rubixgoplatform enableexplorer 

This following options are used for this command
  -addr string
        Server/Host Address (default "localhost")
  -port string
        Server/Host port (default "20000")
  -dbAddress string
        Database address (default "localhost")
  -dbName string
        Explorer database name (default "ExplorerDB")
  -dbPassword string
        Database password (default "password")
  -dbPort string
        Database port number (default "1433")
  -dbType string
        DB Type, supported database are SQLServer, PostgressSQL, MySQL & Sqlite3 (default "SQLServer")
  -dbUsername string
        Database username (default "sa")
```

{{% children description="true"   %}}