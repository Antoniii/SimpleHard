def f3nplus1(n, m=0):
	if n == 1:
		print(m)
	elif n%2 == 0:
		n = n/2
		m += 1
		f3nplus1(n,m)
	else:
		n = n*3+1
		m+=1
		f3nplus1(n,m)

f3nplus1(27)
f3nplus1(7)
f3nplus1(26)
f3nplus1(94)
f3nplus1(97)