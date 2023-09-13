import time
from playsound import playsound


print('stone...')

playsound('D:\sankalp stuff\CS SCHOOL\hi.mp3')
time.sleep(1)
print('paper...')
playsound('D:\sankalp stuff\CS SCHOOL\he.mp3')
time.sleep(1)
print('scissor...')
playsound('D:\sankalp stuff\CS SCHOOL\lo.mp3')
time.sleep(1)

c=input('type')

import random
a=('STONE','PAPER','SCISSOR')
b=random.choice(a)


time.sleep(1)
print('COMPUTER : ',b)




if b=='STONE' and c=='paper':
    print('1 point to user')
if b=='STONE' and c=='stone':
    print('draw')
if b=='STONE' and c=='scissor':
    print('1 point to computer')


if b=='SCISSOR' and c=='paper':
    print('1 point to computer')
if b=='SCISSOR' and c=='stone':
    print('1 point to user')
if b=='SCISSOR' and c=='scissor':
    print('draw')
    
if b=='PAPER' and c=='paper':
    print('draw')
if b=='PAPER' and c=='stone':
    print('1 point to computer')
if b=='PAPER' and c=='scissor':
    print('1 point to user')