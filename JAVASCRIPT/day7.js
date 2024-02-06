// console.log(undefined==false)  //false
// console.log(null==false)   //false
// console.log(NaN==false)    //false



//1) Ternary Operator



// const isLogin = false;
// const isAge  =21;
// isLogin==true?isAge>18?console.log("Done"):console.log("Minor"):console.log("suspended")




// 2) how to use switch and one small assignment 


// const day = 'Monday';

// switch (day) {
//     case 'Monday':
//         console.log("Basic of JS")
//         break;
//         case 'Tuesday':
//         console.log("variable and semantic and how to decalre")
//         break;
//         case 'Wedenday':
//         console.log("datatype in js")
//         break;
//         case 'Thuesday':
//         console.log("method in number")
//         break;
//         case 'Friday':
//         console.log("method in string")
//         break;

//     default:
//         console.log("WeekEnd")
//         break;
// }



// console.log('Sanket'===true)                       //false
// console.log(!!'Sanket'===true)                    //true
// console.log(Boolean('Sanket')===true)            //true
// console.log(Boolean('true'))



// 3) ifElse statment


const age = 21;
const mininmumAgeAllowedInSectionOne = 18;
const mininmumAgeAllowedInSectionTwo = 24;

// if(age >= mininmumAgeAllowedInSectionOne && age <= mininmumAgeAllowedInSectionTwo){
//     console.log("Allowed in section One.....!")
// }
// else if(age > mininmumAgeAllowedInSectionTwo){
//     console.log("Allowed in section Two...!")
// }
// else{
//     console.log("Not Allowed...!")
// }


// Ternary of same problem

// age>=mininmumAgeAllowedInSectionOne && age<=mininmumAgeAllowedInSectionTwo? console.log("Allowed in section One") :age>mininmumAgeAllowedInSectionTwo ? console.log("Allowed in section Two"):console.log("not allowed for any section")

// age>=mininmumAgeAllowedInSectionOne?console.log("Done"):   // else if mandatory in ternary 