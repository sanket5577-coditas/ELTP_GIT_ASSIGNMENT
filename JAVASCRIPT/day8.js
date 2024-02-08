const array = [1, 5, 3, "sanket"];

console.log(array["3"]); // it will sonsole sanket

//property of array
console.log(array.length); //Property of array

console.log(array.push("shambhu", "rahul")); //it will return length of array

console.log(array.pop()); //remove last element of array and it will return length of array
console.log(array);
console.log(array.shift()); // remove first element of array and return length of array
console.log(array);
console.log(array.unshift()); //return array length
console.log(array.slice(1, 3)); // last index it exclude from array
console.log(array); //
// console.log(array.splice(1,2,4,5))  // last index is exclude from array
console.log(array.sort()); // it take paramter too to sort the array according to user need..
arrayofnames = ["S", "a"];
console.log(arrayofnames.sort()); // it will give sorted array according to ascii value of element number->Upper->Lower

console.log(array.reverse()); // itwill reverse the array

const number = [10, 20, 30, 40, 50, 60];

console.log("\n" + number.push(70));
console.log(number.splice(2,0,15));
console.log(number.shift(1));
console.log(number.indexOf(40)); // if element of array is not present it will return -1;
console.log(number.indexOf(10,0)); // it will find element from given index if it is indnex orthwise -1
console.log(number.lastIndexOf(30)); // it find from last of array

const ageArray = [21,22,['222','arjun'],[1,2,3]]
console.log(ageArray.flat())
console.log(number.toString())   // it will convert to string and type is string
console.log(10,20===10,20)   //10 false 20
console.log(Array.isArray(number))   // to check wheather it is array or not if it is return TRUE or FALSE
