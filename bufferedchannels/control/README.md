# Shows memory usage of a channel

Comparing the difference of a buffered channel with structs vs pointers. Also seeing how much space is allocated for a channel

To run `bash anaylze.sh` . It will output the memory usuage of main()

```
SmallSize: 25
MediumSize:100
LargeSize:500

    2.39MB     2.39MB     40: sa := make(chan SmallStruct, 100000)
         .   784.09kB     41: sb := make(chan *SmallStruct, 100000)
      128B       128B     42: sc := make(chan SmallStruct, 1)
         .       104B     43: sd := make(chan *SmallStruct, 1)
         .          .     44:
    9.54MB     9.54MB     45: ma := make(chan MediumStruct, 100000)
         .   784.09kB     46: mb := make(chan *MediumStruct, 100000)
      208B       208B     47: mc := make(chan MediumStruct, 1)
         .       104B     48: md := make(chan *MediumStruct, 1)
         .          .     49:
   47.69MB    47.69MB     50: la := make(chan LargeStruct, 100000)
         .   784.09kB     51: lb := make(chan *LargeStruct, 100000)
      640B       640B     52: lc := make(chan LargeStruct, 1)
         .       104B     53: ld := make(chan *LargeStruct, 1)
```