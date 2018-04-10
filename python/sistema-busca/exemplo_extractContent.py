import urllib3
from bs4 import BeautifulSoup

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
http = urllib3.PoolManager()
pag  = http.request('GET', 'https://pt.wikipedia.org/wiki/Linguagem_de_programa%C3%A7%C3%A3o')

sopa = BeautifulSoup(pag.data, "lxml")
for tags in sopa(['script','style']):
    tags.decompose()

conteudo = ' '.join(sopa.stripped_strings)
print(conteudo)