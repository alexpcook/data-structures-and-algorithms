package algorithms

//#1 - Sort 10 schools around your house by distance:
// Insertion sort, because it's a small amount of data and simple to code

//#2 - eBay sorts listings by the current Bid amount:
// Radix or counting sort, because we roughly know the range of integers to be sorted ($1-$50k perhaps)

//#3 - Sport scores on ESPN
// Quick sort, to have better memory performance compared to merge sort

//#4 - Massive database (can't fit all into memory) needs to sort through past year's user data
// Merge sort, to avoid worst case scenario time complexity of quick sort

//#5 - Almost sorted Udemy review data needs to update and add 2 new reviews
// Insertion sort, because data are almost sorted

//#6 - Temperature Records for the past 50 years in Canada
// Quick sort, assuming floating point numbers
// Radix or counting sort, assuming no floating point numbers

//#7 - Large user name database needs to be sorted. Data is very random.
// Quick sort, assuming a good fist pivot point can be determined

//#8 - You want to teach sorting for the first time
// Bubble sort or selection sort, because they're simple to understand
