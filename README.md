# balance
A sample balance line algorithm.

Usage:
balance -master [Master File] -tx [Transaction File]

master: Is the file that contains the actual state of the data. This file should be in a unduplicated an sorted state

transaction: Is the file that will be processed against the master file. 

The output will be the union of the lines from master and transaction with one of the possible suffixes (,del ,upt or ,new).

,del: It means that this line is present only in the master file

,upt: It means that this line was present in both files. Transaction file was written to output.

,new: It means that this line is present only in the transaction file.


This command is here just to illustrate the implementation of a balance line algorith in go using goroutines and channels. This can be copied and modified for your own needs.

Of course I made this tool because I need this functionality and it was fun writing it in go.
