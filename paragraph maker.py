

import string
import sys
import time
'''

print ('|' + 'sam'+  95* ' ' + '|')

print(100 * '-')

'''

def tameez():
    
    print(100*'-')
    y='y'
    while y=='y':
        a=input('type : ')
        sys.stdout.write('\x1b[1A')
        sys.stdout.write('\x1b[2K')
        print('|' + a + (98-len(a))* ' ' + '|')
        y=input('for next line type y or else press enter :')
        if y=='k':
            print(100*'-')
        sys.stdout.write('\x1b[1A')
        sys.stdout.write('\x1b[2K')
    print(100*'-')
    
tameez()