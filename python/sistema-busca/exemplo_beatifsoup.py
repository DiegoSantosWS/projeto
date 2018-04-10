from bs4 import BeautifulSoup
import urllib3

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
http = urllib3.PoolManager()
#pag = http.request('GET', 'http://www.wsitebrasil.com.br')
pag = http.request('GET', 'https://pt.wikipedia.org/wiki/Linguagem_de_programa%C3%A7%C3%A3o')
st = pag.status


sopa = BeautifulSoup(pag.data, "lxml")
print(sopa)
#print(sopa.title)
#print(sopa.title.string)
links = sopa.find_all('a')
print(len(links))
#print(links)
for link in links:
    print(link.get('href'))
    print(link.contents)