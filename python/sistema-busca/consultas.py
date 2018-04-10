import pymysql
import nltk

def getIdPalavra(palavra):
    retorno = -1
    stemmer = nltk.stem.RSLPStemmer()
    conexao = pymysql.connect(host='localhost', user='root', passwd='1234', db='indice')
    cursor = conexao.cursor()
    
    cursor.execute('select idpalavra from palavras where palavra = %s', stemmer.stem(palavra))
    if cursor.rowcount > 0:
        retorno = cursor.fetchone()[0]
    cursor.close()
    conexao.close()
    return retorno

def buscaMaisPalavras(consulta):
    listacampos =  'p1.idurl'
    listatabelas = ''
    listaclausulas = ''
    palavrasid = []

    palavras = consulta.split(' ')
    numeroTabela = 1
    for palavra in palavras:
        idpalavra = getIdPalavra(palavra)
        if idpalavra > 0:
            palavrasid.append(idpalavra)
            if numeroTabela > 1:
                listatabelas += ', '
                listaclausulas += ' and '
                listaclausulas += ' p%d.idurl = p%d.idurl and ' % (numeroTabela - 1, numeroTabela)
            listacampos += ', p%d.localizacao ' % numeroTabela
            listatabelas += ' palavra_localizacao p%d' % numeroTabela
            listaclausulas += 'p%d.idpalavra = %d' % (numeroTabela, idpalavra)
            numeroTabela += 1
    consultacompleta = 'select %s from %s where %s' % (listacampos, listatabelas, listaclausulas)

    conexao = pymysql.connect(host='localhost', user='root', passwd='1234', db='indice')
    cursor = conexao.cursor()
    cursor.execute(consultacompleta)
    linhas = [linha for linha in cursor]

    cursor.close()
    conexao.close()
    return linhas, palavrasid

def getUrl(idurl):
    retorno = -1
    stemmer = nltk.stem.RSLPStemmer()
    conexao = pymysql.connect(host='localhost', user='root', passwd='1234', db='indice')
    cursor = conexao.cursor()
    
    cursor.execute('select url from urls where idurl = %s', idurl)
    if cursor.rowcount > 0:
        retorno = cursor.fetchone()[0]
    cursor.close()
    conexao.close()
    return retorno

def frequenciaScore(linhas):
    contagem = dict([linha[0], 0] for linha in linhas)
    for linha in linhas:
        contagem[linha[0]] += 1
    return contagem

def pesquisa(consulta):
    linhas, palavrasid = buscaMaisPalavras(consulta)
    #scores = dict([linha[0], 0] for linha in linhas)
    #scores = frequenciaScore(linhas)
    #scores = localizacaoScore(linhas)
    scores = distanciaScore(linhas)
    #for url, score in scores.items():
    #    print(str(url) + ' - ' + str(score))
    scoresordenado = sorted([(score, url) for (url, score) in scores.items()], reverse=0)
    for (score, idurl) in scoresordenado[0:10]:
        print('%f\t%s' % (score, getUrl(idurl)))

def localizacaoScore(linhas):
    localizacoes = dict([linha[0], 10000] for linha in linhas)
    for linha in linhas:
        soma = sum(linha[1:])
        if soma < localizacoes[linha[0]]:
            localizacoes[linha[0]] = soma
    return localizacoes

def distanciaScore(linhas):
    if len(linhas[0]) <= 2:
        return dict([(linha[0], 1.0) for linha in linhas])

    distancias = dict([(linha[0], 10000) for linha in linhas])
    for linha in linhas:
        dist = sum([abs(linha[i] - linha[i - 1]) for i in range(2, len(linha))])
        if dist < distancias[linha[0]]:
            distancias[linha[0]] = dist
    return distancias

def buscaUmaPalavra(palavra):

    idpalavra = getIdPalavra(palavra)
    conexao = pymysql.connect(host='localhost', user='root', passwd='1234', db='indice')
    cursor = conexao.cursor()
    cursor.execute('select urls.url from palavra_localizacao plc inner join urls on plc.idurl = urls.idurl where plc.idpalavra = %s', idpalavra)
    paginas = set()
    for url in cursor:
        #print(url[0])
        paginas.add(url[0])
    print('Páginas encontradas: ' + str(len(paginas)))
    for url in paginas:
        print(url)
    cursor.close()
    conexao.close()

#buscaUmaPalavra('programação')
pesquisa('python programação')
#linhas, score = buscaMaisPalavras('python php programação')
#print('Frequencia: ' + frequenciaScore(linhas) )