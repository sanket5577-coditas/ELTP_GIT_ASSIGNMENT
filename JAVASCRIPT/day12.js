// Loops

// For Loop

const person={
    name :"sanket",
    age:12,
    skill: false

}
// const keys =Object.keys(person)
// for (let index = 0; index < keys.length; index++) {
//     const key = keys[index]
//     console.log(person[key])
// }



// const numebrsArray= [1,2,4,5,6];

// for (const key in numebrsArray) {
//     console.log(key)
// }

// for (const iterator of numebrsArray) {
//     console.log(iterator)
// }
// for (const key in person) {
//     console.log(key)
// }

// // for (const iterator of person) {
// //     // console.log(iterator)
// //     console.log(iterator[person])
// // }



// // function

// printName("sanket")
// function printName(Name){       // Best Hosting
//     // console.log(this)
//     console.log(`My name is ${Name}`)
//     // return "n"
// }
//    //My name is sanket
// const some = printName("sanket");
// console.log(typeof(some))

// const sum = (firstName,lastName)=>{
//     // console.log(this)
//     console.log(`${firstName} , ${lastName}`)
// }

// sum("sanket","Kumar")


function squareOfNumber(number){
    return number**2
}
console.log(squareOfNumber(2))

const squareOfNumbers = (number)=>{
    return number**2
} 

console.log(squareOfNumbers(2))