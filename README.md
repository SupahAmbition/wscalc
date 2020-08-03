# wscalc

A calculator web app that utilizes Web Sockects to enable "real time" calculations with all users.
When a user uses this calculator, their calcualtion will be broadcasted to all other users. 

[wscalc.com](http://wscalc.com)


### Build

go build
./wscalc  


### Testing

go test -v


### Timeline: 
Here you can see when progress was made and what was done (by viewing the repo at the time of the commit). 
I time tracked progress on this project so you can see how long I took to do things. 
I tracked time spent coding, testing or debuging. Other time was not recored. 

init  -- 0:00 [49360055c](https://github.com/SupahAmbition/wscalc/commit/49360055c9719cee319f4f33782766e047ae3911)
  
0:58 --  [c2ca51f64](https://github.com/SupahAmbition/wscalc/commit/c2ca51f64f79b19366d8833ce28ad12378a412bf)   
  
3:18 -- [582495bd5](https://github.com/SupahAmbition/wscalc/commit/582495bd57cae9dcd5edb97ac7e5777398ca8e6c) & [e5965eae72](https://github.com/SupahAmbition/wscalc/commit/e5965eae72ceae79b09247d91397c534a017bc16)

3:42 -- [1b3e64439ef](https://github.com/SupahAmbition/wscalc/commit/1b3e64439efad53e1d297834153f6204e3a4c517)

5:21 -- [8e61eadb307](https://github.com/SupahAmbition/wscalc/commit/8e61eadb3070bb31450955f488f905e389851827)

5:53 -- [7580a40a](https://github.com/SupahAmbition/wscalc/commit/7580a40aac567699f7b8a0ebcf0d6de89929d2662)

finished 



#### Final Notes 

In the end I went over the 5 hour mark a bit, but I believe 
the quality of the resulting code was woth it. 
I combed through many (if not all) number formating / calculator bugs, and 
I created a super efficent, thread safe data structure that is sure to help keep AWS costs low. 
  

I lifted a bit of html/css from [here](https://github.com/abarna-codespot/A-simple-Calculator) in order to save time, 
and retrofitted it to allow for decimals, and negitive numbers. 
  
I dont really consider this a bug but to use decimals you do have to have a leading digit such as (0.1).  
