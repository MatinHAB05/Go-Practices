import sys
ss = open("a.txt", "+w")
sys.stdout = ss

for i in range(33, 127):
    print(chr(i)*1000, end="|")


ss.close()
