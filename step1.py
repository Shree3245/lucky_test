userPoints = {}
while True:
    spendings =input("Please input users spendings in a list seperated by a comma (','). If you wish to stop inputting, please type 'stop':\t").replace(" ", "").split(",")
    if "stop" in spendings:
        break
    for i in range(len(spendings)):
        try:
            out = int(float("{:.2f}".format(float(spendings[i]))))
            
            if str(i) in userPoints:
                userPoints[str(i)] += out
            else:
                userPoints[str(i)] = out
        except ValueError:
            print("Please input a number")
    
print(userPoints)