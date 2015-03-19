#http://stackoverflow.com/questions/699866/python-int-to-binary
def binary(i):
    if i == 0:
        return [0]
    s = []
    while i:
        if i & 1 == 1:
            s = [1] + s
        else:
            s = [0] + s
        i >>= 1
    return s

def decimal(binlist):
    return sum(binlist[-j]*2**(j-1) for j in range(len(binlist),0,-1))

def waystomakezero(Abit,Bbit):
    if decimal(Abit)<decimal(Bbit):
        Abit, Bbit = Bbit, Abit
    nA = len(Abit)
    nB = len(Bbit)
    if sum(Bbit) == 0:
        return decimal(Abit)+1
    if nA == nB ==1:
        if Abit[0]+Bbit[0]==2:
            return 3
        elif Abit[0] + Bbit[0] ==1:
            return 2
        else:
            return 1
    if nA > nB:
        #if A has more bits than B than the first can be arranged arbitrarily, except for the maximum, which has a restriction on the last bits to make the whole <A.
        return decimal(Abit[:(nA-nB)])*waystomakezero(nB*[1],Bbit)+waystomakezero(Abit[-nB:],Bbit)
    else:
        #from now on we have nA = nB
        n = nB
        #three cases: first bit is 0 for both,
        #or 1 for one and 0 for the other.
        #case 1 A=0, B=0
        count = 3**(n-1)
        #case 2 A=1, B=0
        count += waystomakezero(Abit[1:],(n-1)*[1])
        #case 3 A=0, B=1
        count += waystomakezero((n-1)*[1],Bbit[1:])
        return count

def update(bitlist, place):
    updated = False
    output = bitlist[:]
    while place >= 0:
        if bitlist[place]==1:
            output[place]=0
            output[(place+1):]= (len(bitlist)-place-1)*[1]
            return output
        else:
            place-=1
    return "Something went wrong, perhaps A<K?"

def timeszero(Astring, Bstring):
    #number of ways to make zero keeping Bstring fixed
    n=len(Astring)
    if n== 1:
        if Bstring[0]==1 or Astring[0]==0:
            return 1
        else:
            return 2
    if Astring[0] ==0 and Bstring ==0:
        return timeszero(Astring[1:],Bstring[1:])
    elif Astring[0] ==1 and Bstring ==0:
        return timeszero(Astring[1:],Bstring[1:])+timeszero([1]*(n-1),Bstring[1:])
    elif Astring[0] ==0 and Bstring ==1:
        return timeszero(Astring[1:],Bstring[1:])
    elif Astring[0] ==1 and Bstring ==1:
        return timeszero([1]*(n-1),Bstring[1:])
    return "Something is wrong"

def numberofwinners(A,B,K):
    A=A-1
    B=B-1
    K=K-1
    Kbit = binary(K)
    Bbit = binary(B)
    Abit = binary(A)
    nK = len(Kbit)
    nA = len(Abit)
    nB = len(Bbit)
    if nA > nB:
        Bbit = [0]*(nA-nB) + Bbit
    else:
        Abit = [0]*(nB-nA) + Abit
    if nK==1:
        if Kbit[0]==0:
            return waystomakezero(Abit,Bbit)
        else:
            count = waystomakezero(Abit,Bbit)
            if A%2==0 and A>0:
                Abit = binary(A-1)
            if B%2==0 and B>0:
                Bbit = binary(B-1)
            return waystomakezero(Abit[:-1],Bbit[:-1])+count
    else:
        #k=0
        #a[:-nK]=A[:-nk] and b[:-nk]=B[:-nk]: 1 time
        count = numberofwinners(decimal(Abit[-(nK-1):])+1, decimal(Bbit[-(nK-1):])+1, decimal([1]*(nK-1))+1)
        #a[:-nK]<A[:-nk] and b[:-nk]=B[:-nk]: ?times
        timeszero(Abit[:-(nK-1)],Bbit[:-(nK-1)]) *numberofwinners(decimal([1]*(nK-1))+1, decimal(Bbit[-(nK-1):])+1, decimal([1]*(nK-1))+1)
        #a[:-nK]=A[:-nk] and b[:-nk]<B[:-nk]: ?times
        timeszero(Bbit[:-(nK-1)],Abit[:-(nK-1)]) *numberofwinners(decimal(Abit[-(nK-1):])+1, decimal([1]*(nK-1))+1, decimal([1]*(nK-1))+1)
        #a[:-nK]<A[:-nk] and b[:-nk]<B[:-nk]: ?times
        count += waystomakezero(binary(decimal(Abit[:-(nK-1)])-1),binary(decimal(Bbit[:-(nK-1)])-1))*4**(nK-1)
        #k=1
        if A>K:
            if Abit[-nK]==0:
                Abit=update(Abit, len(Abit)-nK)
            Abit.pop(-nK)
        if B>K:
            if Bbit[-nK]==0:
                Bbit=update(Bbit, len(Bbit)-nK)
            Bbit.pop(-nK)
        count+=numberofwinners(decimal(Abit)+1, decimal(Bbit)+1, decimal(Kbit[1:])+1)
        return count
