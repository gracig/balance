# balance
A sample balance line algorithm.

Usage:
balance -master [Master File] -tx [Transaction File]

master: Is the file that contains the actual state of the data. This file should be in an unduplicated and sorted state

transaction: Is the file that will be processed against the master file. This file should also be in an unduplicated and sorted state.

The output will be the union of the lines from the master and transaction files adding at the end of the line one of the possible suffixes (,del ,upt or ,new).

,del: It means that this line is present only in the master file

,upt: It means that this line was present in both files. The outputed line was the one from the transaction file.

,new: It means that this line is present only in the transaction file.


This command is here just to illustrate the implementation of a balance line algorithm in go using goroutines and channels. This can be copied and modified for your own need.
