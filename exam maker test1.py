
a=input('WELCOME TO QUESTION PAPER MAKER ,if you want to exit press small letter n or else press enter button!')
if a=='n':
    print('QUITING THE PROGRAM')
    exit()
else:
    print('Lets continue !')
print('Type only dichotomous question.')    
ques1=input('Type your first question ?')
ans1=input('Type the answer in yes or no . If answer is yes for the above question type  letter Y and if answer is no for the above questyion type letter N')
if ans1=='Y':
    print('Ok , the correct answer for ques1 is registered as yes')
elif ans1=='y':
    print('Ok , the correct answer for ques2 is registered as yes')
elif ans1=='N':
    print('Ok , the correct answer for ques1 is registered as no')
elif ans1=='n':
    print('Ok , the correct answer for ques2 is registered as no')
else:
    print('invalid answer format for ques1')
ques2=input('Type your second question ?')
ans2=input('Type the answer in yes or no . If answer is yes for the above question type letter Y and if answer is no for the above questyion type letter N')
if ans2=='Y':
    print('Ok , the correct answer for ques2 is registered as yes')
elif ans2=='y':
    print('Ok , the correct answer for ques2 is registered as yes')
elif ans2=='N':
    print('Ok , the correct answer for ques2 is registered as no')
elif ans2=='n':
    print('Ok , the correct answer for ques2 is registered as no')
else:
    print('invalid answer format for ques2')
ask=input('if you want to add one more question press Y or else press enter button')
if ask=='y':
    print('ok')
    ques3=input('Type your third question ?')
    ans3=input('Type the answer in yes or no . If answer is yes for the above question type letter Y and if answer is no for the above questyion type letter N')
    if ans3=='Y':
         print('Ok , the correct answer for ques3 is registered as yes')
    elif ans3=='y':
         print('Ok , the correct answer for ques3 is registered as yes')
    elif ans3=='N':
         print('Ok , the correct answer for ques3 is registered as no')
    elif ans2=='n':
         print('Ok , the correct answer for ques3 is registered as no')
    else:
         print('invalid answer format for ques3')
elif ask=='Y':
    print('ok')
    ques3=input('Type your third question ?')
    ans3=input('Type the answer in yes or no . If answer is yes for the above question type letter Y and if answer is no for the above questyion type letter N')
    if ans3=='Y':
         print('Ok , the correct answer for ques3 is registered as yes')
    elif ans3=='y':
         print('Ok , the correct answer for ques3 is registered as yes')
    elif ans3=='N':
         print('Ok , the correct answer for ques3 is registered as no')
    elif ans2=='n':
         print('Ok , the correct answer for ques3 is registered as no')
    else:
         print('invalid answer format for ques3')
else:
    ques3=None
    print('lets start assignment')
import sys
import time
sys.stdout.write('\x1b[1A')
sys.stdout.write('\x1b[2K')
sys.stdout.write('\x1b[1A')  
sys.stdout.write('\x1b[2K')
sys.stdout.write('\x1b[1A')
sys.stdout.write('\x1b[2K')
sys.stdout.write('\x1b[1A')
sys.stdout.write('\x1b[2K')
sys.stdout.write('\x1b[1A')
sys.stdout.write('\x1b[2K')
sys.stdout.write('\x1b[1A')
sys.stdout.write('\x1b[2K')
sys.stdout.write('\x1b[1A')
sys.stdout.write('\x1b[2K')
sys.stdout.write('\x1b[1A')
sys.stdout.write('\x1b[2K')
sys.stdout.write('\x1b[1A')
sys.stdout.write('\x1b[2K')
sys.stdout.write('\x1b[1A')
sys.stdout.write('\x1b[2K')
sys.stdout.write('\x1b[1A')
sys.stdout.write('\x1b[2K')
sys.stdout.write('\x1b[1A')
sys.stdout.write('\x1b[2K')
sys.stdout.write('\x1b[1A')
sys.stdout.write('\x1b[2K')
sys.stdout.write('\x1b[1A')
sys.stdout.write('\x1b[2K')
sys.stdout.write('\x1b[1A')
sys.stdout.write('\x1b[2K')






print('Assignment Started')
# test chhapna
marks=0
prashan1=input(ques1)
if prashan1==ans1:
    print('Correct Answer')
    marks=marks+1
else:
    print('wrong answer')
prashan2=input(ques2)
if prashan2==ans2:
    print('Correct Answer')
    marks=marks+1
else:
    print('wrong answer')
prashan3=input(ques3)
if ques3!=None:
    if prashan3==ans3:
        print('Correct Answer')
        marks=marks+1
        print('user has obtained ',marks,' marks out of 3 marks')
    else:
        print('wrong answer')
        print('user has obtained ',marks,' marks out of 3 marks')
else:
    print('user has obtained ',marks,' marks out of 3 marks')















