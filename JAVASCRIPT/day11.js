// const numberArray = [1,2,3,4[4,5]];
// const copyOfnumberArray = numberArray.slice();
// copyOfnumberArray[3] = 10;
// console.log(numberArray);
// console.log(copyOfnumberArray);
const person = [
  { name: "Akhil" },
  { name: "Arjun" },
  { name: "Sanket" },
  { name: "Yash" },
  { name: "Shambhu" },
];

for (const values of person) {
  console.log(`My Name is ${values.name}`);
}
console.log(">>>>>>>>>>>>>");
for (const { name } of person) {
  console.log(`My Name is ${name}`);
}
// for (const [keys,{person:value}] of Object.entries(person)) {
//     console.log(`My Name is ${keys} and ${{value}}`)

// }

// map

const numberArray = [1, 2, 3, 4, 5];
const squareOfNumberArray = numberArray.map((numberArray) => {
  console.log(numberArray ** 2);
});

// console.log(numberArray);

//filter

const oddNumberArray = numberArray.filter((numberArray) => {
  return (numberArray % 2===0);
});
console.log(oddNumberArray)


// reduce

const sum = numberArray.reduce((acc, currentIndex) => {
    return acc+=currentIndex
},0)
console.log(sum)

const maxNumberOfArray = numberArray.reduce((acc, currentIndex) => {
    if(acc<currentIndex){
        acc=currentIndex
    }
    return acc;
},numberArray[0])
console.log(maxNumberOfArray)

const something = numberArray.every((numberArray)=>{
    if(numberArray%2===0){
        return true;
    }
})
// console.log(something)
// const sumFucntion = numberArray.sum((numberArray)=>{
//     if(numberArray%2===0){
//         return true;
//     }
// })
console.log(sumFucntion)

numberArray.reduce((acc,currentValue)=>{
    console.log(currentValue)
}
)

