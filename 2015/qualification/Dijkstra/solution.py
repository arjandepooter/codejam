def quaternions(a,b):
    if len(a+b)%2==0:
        sign=""
    else:
        sign="-"
    a=a[-1]
    b=b[-1]
    table = {("1","1"):"1",("1","i"):"i",("1","j"):"j",("1","k"):"k",("i","1"):"i",("i","i"):"-1",("i","j"):"k",("i","k"):"-j",("j","1"):"j",("j","i"):"-k",("j","j"):"-1",("j","k"):"i",("k","1"):"k",("k","i"):"j",("k","j"):"-i",("k","k"):"-1"}
    output = sign+table[(a,b)]
    if len(output)>1:
        if output[1]=="-":
            output = output[2:]
    return output

def evaluate(string):
    current="1"
    for i in range(len(string)):
        current = quaternions(current, string[i])
    return current

def get_ijk(string):
    if evaluate(string)!= "-1":
        return "NO"
    else:
        foundi = False
        foundk= False
        current="1"
        i=0
        while not foundi and i < len(string):
            current = quaternions(current, string[i])
            if current == "i":
                position_i = i
                foundi = True
            i += 1
        if foundi:
            current = "1"
            k=len(string)-1
            while not foundk and k>position_i:
                current = quaternions(string[k], current)
                if current == "k":
                    position_k = k
                    foundk = True
                k -= 1
            if foundk:
                return "YES"
        return "NO"

def get_output(instance):
    inputdata = open(instance + ".in", 'r')
    output = open(instance+ ".out", 'w')
    T = int(inputdata.readline())
    for t in range(T):
        L, X = inputdata.readline().split()
        X = int(X)
        string = inputdata.readline()[:-1]
        string *= X
        output.write('Case #' + str(t+1) +': ' + get_ijk(string) +  "\n")
    return None
