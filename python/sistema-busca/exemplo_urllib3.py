import urllib3

http = urllib3.PoolManager()

pagina = http.request('GET', 'http://www.iaexpert.com.br')
st = pagina.status
print(st)
dt = pagina.data[0:50]
print(dt)

