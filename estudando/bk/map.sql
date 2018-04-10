
Starting Nmap 7.01 ( https://nmap.org ) at 2018-03-16 13:23 -03
NSE: Loaded 132 scripts for scanning.
NSE: Script Pre-scanning.
Initiating NSE at 13:23
Completed NSE at 13:23, 0.00s elapsed
Initiating NSE at 13:23
Completed NSE at 13:23, 0.00s elapsed
Initiating Ping Scan at 13:23
Scanning 127.0.0.1 [2 ports]
Completed Ping Scan at 13:23, 0.00s elapsed (1 total hosts)
Initiating Connect Scan at 13:23
Scanning localhost (127.0.0.1) [1000 ports]
Discovered open port 3306/tcp on 127.0.0.1
Discovered open port 80/tcp on 127.0.0.1
Discovered open port 631/tcp on 127.0.0.1
Discovered open port 7070/tcp on 127.0.0.1
Discovered open port 3000/tcp on 127.0.0.1
Completed Connect Scan at 13:23, 0.03s elapsed (1000 total ports)
Initiating Service scan at 13:23
Scanning 5 services on localhost (127.0.0.1)
Completed Service scan at 13:24, 76.19s elapsed (5 services on 1 host)
NSE: Script scanning 127.0.0.1.
Initiating NSE at 13:24
Completed NSE at 13:24, 1.87s elapsed
Initiating NSE at 13:24
Completed NSE at 13:24, 0.00s elapsed
Nmap scan report for localhost (127.0.0.1)
Host is up (0.00025s latency).
Not shown: 995 closed ports
PORT     STATE SERVICE         VERSION
80/tcp   open  http            Apache httpd 2.4.18
| http-ls: Volume /
| SIZE  TIME              FILENAME
| -     2018-01-23 16:55  dev.com.br/
| 11K   2018-01-23 13:38  index.html.old
|_
| http-methods: 
|_  Supported Methods: OPTIONS GET HEAD POST
|_http-server-header: Apache/2.4.18 (Ubuntu)
|_http-title: Index of /
631/tcp  open  ipp             CUPS 2.1
| http-methods: 
|   Supported Methods: GET HEAD OPTIONS POST PUT
|_  Potentially risky methods: PUT
| http-robots.txt: 1 disallowed entry 
|_/
|_http-server-header: CUPS/2.1 IPP/2.1
|_http-title: Home - CUPS 2.1.3
3000/tcp open  ppp?
3306/tcp open  mysql           MySQL 5.7.21-0ubuntu0.16.04.1
| mysql-info: 
|   Protocol: 53
|   Version: .7.21-0ubuntu0.16.04.1
|   Thread ID: 21
|   Capabilities flags: 63487
|   Some Capabilities: InteractiveClient, SupportsLoadDataLocal, ConnectWithDatabase, LongColumnFlag, DontAllowDatabaseTableColumn, SupportsCompression, SupportsTransactions, FoundRows, Speaks41ProtocolOld, Support41Auth, Speaks41ProtocolNew, ODBCClient, IgnoreSpaceBeforeParenthesis, LongPassword, IgnoreSigpipes
|   Status: Autocommit
|_  Salt: \x12-8KT-F\x18B%?\x05{g~\x1B\x0E_f_
7070/tcp open  ssl/realserver?
|_ssl-date: TLS randomness does not represent time
1 service unrecognized despite returning data. If you know the service/version, please submit the following fingerprint at https://nmap.org/cgi-bin/submit.cgi?new-service :
SF-Port3000-TCP:V=7.01%I=7%D=3/16%Time=5AABEF8A%P=x86_64-pc-linux-gnu%r(Ge
SF:nericLines,67,"HTTP/1\.1\x20400\x20Bad\x20Request\r\nContent-Type:\x20t
SF:ext/plain;\x20charset=utf-8\r\nConnection:\x20close\r\n\r\n400\x20Bad\x
SF:20Request")%r(GetRequest,C24,"HTTP/1\.0\x20200\x20OK\r\nDate:\x20Fri,\x
SF:2016\x20Mar\x202018\x2016:23:38\x20GMT\r\nContent-Type:\x20text/html;\x
SF:20charset=utf-8\r\n\r\n<!doctype\x20html>\n<html\x20lang=\"en\">\n\x20\
SF:x20\x20\x20<head>\n\x20\x20\x20\x20\x20\x20\x20\x20<title>WS\x20SHORTEN
SF:ER\x20LOGIN</title>\n\x20\x20\x20\x20\x20\x20\x20\x20\n\x20\x20\x20\x20
SF:\x20\x20\x20\x20<meta\x20charset=\"utf-8\">\n\x20\x20\x20\x20\x20\x20\x
SF:20\x20<meta\x20name=\"viewport\"\x20content=\"width=device-width,\x20in
SF:itial-scale=1,\x20shrink-to-fit=no\">\n\n\x20\x20\x20\x20\x20\x20\x20\x
SF:20\n\x20\x20\x20\x20\x20\x20\x20\x20<link\x20rel=\"stylesheet\"\x20href
SF:=\"https://maxcdn\.bootstrapcdn\.com/bootstrap/4\.0\.0-beta\.2/css/boot
SF:strap\.min\.css\"\x20integrity=\"sha384-PsH8R72JQ3SOdhVi3uxftmaW6Vc51MK
SF:b0q5P2rRUpPvrszuE4W1povHYgTpBfshb\"\x20crossorigin=\"anonymous\">\n\x20
SF:\x20\x20\x20\x20\x20\x20\x20<link\x20rel=\"stylesheet\"\x20href=\"/stat
SF:ic/css/style\.css\"\x20>\n\x20\x20\x20\x20\x20\x20\x20\x20<link\x20rel=
SF:\"stylesheet\"\x20href=\"https://cdnjs\.cloudflare\.com/ajax/libs/font-
SF:awesome/4\.7\.0/css/font-awesome\.css\">\n\x20\x20\x20\x20</head>\n\x20
SF:\x20\x20\x20<body>\n\x20\x20\x20\x20\x20\x20\x20\x20<div\x20class=\"con
SF:tainer\">\n\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20<div\x20clas
SF:s=\"row\">\n\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x2
SF:0\x20<div\x20class=\"col-12\x20text-center\x20mt-5\">\n\x20\x20\x20\x20
SF:\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20<p><h1\x20\x20c
SF:las")%r(Help,67,"HTTP/1\.1\x20400\x20Bad\x20Request\r\nContent-Type:\x2
SF:0text/plain;\x20charset=utf-8\r\nConnection:\x20close\r\n\r\n400\x20Bad
SF:\x20Request")%r(HTTPOptions,84,"HTTP/1\.0\x20405\x20Method\x20Not\x20Al
SF:lowed\r\nDate:\x20Fri,\x2016\x20Mar\x202018\x2016:23:43\x20GMT\r\nConte
SF:nt-Length:\x200\r\nContent-Type:\x20text/plain;\x20charset=utf-8\r\n\r\
SF:n")%r(RTSPRequest,67,"HTTP/1\.1\x20400\x20Bad\x20Request\r\nContent-Typ
SF:e:\x20text/plain;\x20charset=utf-8\r\nConnection:\x20close\r\n\r\n400\x
SF:20Bad\x20Request");
Service Info: Host: 127.0.1.1

NSE: Script Post-scanning.
Initiating NSE at 13:24
Completed NSE at 13:24, 0.00s elapsed
Initiating NSE at 13:24
Completed NSE at 13:24, 0.00s elapsed
Read data files from: /usr/bin/../share/nmap
Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
Nmap done: 1 IP address (1 host up) scanned in 79.10 seconds
