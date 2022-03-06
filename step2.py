#This initial set of code is taken from step 1.py
from pprint import pprint
from collections import defaultdict

userPoints = defaultdict(dict)
while True:
    spendings =input("Please input users spendings in a list seperated by a comma (','). If you wish to stop inputting, please type 'stop':\t").replace(" ", "").split(",")
    if "stop" in spendings:
        break
    spendings = [int(float("{:.2f}".format(float(i)))) for i in spendings]
    counter = sum([int(i) for i in spendings])
    pointTally = []

    for i in range(len(spendings)):
        out = spendings[i]
        if str(i) in userPoints: 
            out = int(out) + int(userPoints[str(i)]['history'][-1]) + int(counter) *(int(userPoints[str(i)]['history'][-1])/sum([int(float(i)) for i in userPoints['previous']]))
            out = (float("{:.2f}".format(float(out))))
            userPoints[str(i)]['spending'] += out
            userPoints[str(i)]['history'].append(out)
            userPoints[str(i)]['todays_total_points']=out
            pointTally.append(out)
        else:
            userPoints[str(i)]['spending'] = out
            userPoints[str(i)]['history'] =[out]
            userPoints[str(i)]['todays_total_points']=out

            pointTally.append(out)

        
    userPoints["previous"] = pointTally
    pprint(userPoints)
