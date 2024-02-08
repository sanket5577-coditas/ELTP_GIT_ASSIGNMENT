//Object

const numbers = [1, 2, 4, 5, 6, 7];


for (let number in numbers) {
  console.log(number);
}
for (const number of numbers) {
  console.log(number);
}
// console.log(numbers.slice());
const arrayNumbers = numbers.slice(); // Shallow clone
// arrayNumbers.push(10);
// console.log(arrayNumbers); //[1, 2, 4, 5, 6, 7, 10]
// console.log(numbers); //[1, 2, 4, 5, 6, 7]
const arrayOfNumbersForShallowClone = [...arrayNumbers]

// console.log(arrayOfNumbersForShallowClone)

const internshipStudentsObject ={
    name:"sanket",
    age:21,
    employee_Id:5577,
    address:
    {
        pinCode:413006
    }
}

const {address:{pinCode}} = internshipStudentsObject

const{age} = internshipStudentsObject;
console.log(age)
const{name, ...internshipInfo} = internshipStudentsObject;
console.log(internshipInfo)

// const internshipInfoSecondStudent ={...internshipStudentsObject};
const internshipInfoSecondStudent =JSON.parse(JSON.stringify(internshipStudentsObject));
console.log(internshipInfoSecondStudent)


internshipInfoSecondStudent.name="Rahul"; 
console.log(internshipStudentsObject)  //{ name: 'sanket', age: 21, employee_Id: 5577 }
console.log(internshipInfoSecondStudent)  //{ name: 'Rahul', age: 21, employee_Id: 5577 }

