# Guess-it-1

The program tests given number as standard input, prints out a range in which the next number provided should be.

The data received by the program, as always, will be presented as the following example

189
113
121
114
145
110
...

This data represents a graph in which the values of the x axis are the number of the lines (0, 1, 2, 3, 4, 5 ...) and the values of the y axis are the actual numbers (189, 113, 121, 114, 145, 110...).

Each of the numbers will be your standard input and the purpose of your program is for you to find the range in which the next number will be in. This range should have a space separating the lower limit from the upper one like in the example:

# Testing

1. Download this zip file https://assets.01-edu.org/guess-it/guess-it.zip  
2. Clone this repository to your computer
3. Rename guess-it-1 to student and upload to the folder
4. Give all the permissions to all the files 
    chmod +x *
5. Install the node.js package manager
    npm install
6. node server.js

    TESTING


big-range http://localhost:3000/?guesser=big-range

average   http://localhost:3000/?guesser=average

median    http://localhost:3000/?guesser=median

# Functional

Try running the student program against the big-range program, using Data 2 and then Data 3 3 times each.
Did the student won against big-range most of the times (at least 2 out of 3 times for each)?

Try running the student program against the average program, using Data 1, Data 2 and Data 3 3 times each.
Did the student won against average most of the times (at least 2 out of 3 times for each)?

Try running the student program against the median program, using Data 1, Data 2 and Data 3 3 times each.
Did the student won against median most of the times (at least 2 out of 3 times for each)?

Bonus

Try running the student program against the nic program, using Data 1, Data 2 and Data 3 3 times each.
+Did the student won against nic most of the times (at least 2 out of 3 times for each)?