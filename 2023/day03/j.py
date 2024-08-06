def readfile(nome_arq):
	with open(nome_arq, "r") as f:
		linhas = f.readlines()
	return linhas

arr = []
def monntaNro(i, j):	
	jini = j
	str_nro = ''	
	while j<n_colunas and mat[i][j].isnumeric():				
		str_nro += str(mat[i][j])		
		mat[i][j] = '.'
		j+=1
	j = jini-1
	while mat[i][j].isnumeric() and j>=0:
		str_nro = str(mat[i][j]) + str_nro
		mat[i][j] = '.'
		j-=1	
	
	# print("str_nro = ", str_nro)
	arr.append(int(str_nro))
	return int(str_nro)


def day_3_part1():

	symbols = ""	
	i=0
	j=0
	tot=0
	for linha in linhas:		
		j=0
		for c in linha:
			if not (c.isnumeric()) and c !="." and c not in symbols:
				symbols +=c
			if j<n_colunas and c !='\n':						
				mat[i][j] = c				
				j+=1
		i+=1

	print(symbols)



	for i in range(0,n_linhas):
		for j in range(0,n_colunas):
			if mat[i][j] in symbols:				
				if mat[i-1][j].isnumeric():					
					tot+= monntaNro(i-1,j)
				if mat[i-1][j+1].isnumeric():					
					tot+= monntaNro(i-1,j+1)
				if mat[i-1][j-1].isnumeric():											
					tot+= monntaNro(i-1,j-1)									
				if mat[i+1][j].isnumeric():						
					tot+= monntaNro(i+1,j)
				if mat[i+1][j+1].isnumeric():									
					tot+= monntaNro(i+1,j+1)
				if mat[i+1][j-1].isnumeric():					
					tot+= monntaNro(i+1,j-1)						
				if j-1 >= 0 and mat[i][j-1].isnumeric():					
					tot+= monntaNro(i,j-1)				
				if j+1 < n_colunas and mat[i][j+1].isnumeric():				
					tot+= monntaNro(i,j+1)
		# print(tot, i)

	print(tot)

def day_3_part2():

	symbols = ""	
	i=0
	j=0
	tot=0
	for linha in linhas:		
		j=0
		for c in linha:
			if not (c.isnumeric()) and c !="." and c not in symbols:
				symbols +=c
			if j<n_colunas and c !='\n':						
				mat[i][j] = c				
				j+=1
		i+=1

	for i in range(0,n_linhas):
		for j in range(0,n_colunas):
			if mat[i][j] == '*':				
				nros = []				
				if mat[i-1][j].isnumeric():					
					nros.append(monntaNro(i-1,j))					
				if mat[i-1][j+1].isnumeric():					
					nros.append(monntaNro(i-1,j+1))					
				if mat[i-1][j-1].isnumeric():		
					nros.append(monntaNro(i-1,j-1))														
				if mat[i+1][j].isnumeric():						
					nros.append(monntaNro(i+1,j))					
				if mat[i+1][j+1].isnumeric():	
					nros.append(monntaNro(i+1,j+1))					
				if mat[i+1][j-1].isnumeric():		
					nros.append(monntaNro(i+1,j-1))					
				if j-1 >= 0 and mat[i][j-1].isnumeric():					
					nros.append(monntaNro(i,j-1))									
				if j+1 < n_colunas and mat[i][j+1].isnumeric():	
					nros.append(monntaNro(i,j+1))
				if len(nros) == 2:
					tot += nros[0] * nros[1]
	print(tot)

linhas = readfile("input.txt")			
# n_colunas = len(linhas[0])
n_colunas = len(linhas)
print(n_colunas)
n_linhas = len(linhas)
mat = [[""]*n_colunas for _ in range(n_linhas)]
day_3_part1()

print(sorted(arr))
# linhas = readfile("input.txt")
# n_colunas = len(linhas)
# n_linhas = len(linhas)
# mat = [[""]*n_colunas for _ in range(n_linhas)]
# day_3_part2()