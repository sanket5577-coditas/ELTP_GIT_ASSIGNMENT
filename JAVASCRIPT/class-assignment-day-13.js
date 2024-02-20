// Explain the difference between forEach, map, filter, and reduce.
const numbersArray = [1, 2, 3, 4, 5];

const printedNumbers = numbersArray.forEach((number) => {
  //   console.log(number); // it will give ouput but u will get undefined with o/p
  return number; // The forEach will always return undefined and it will return every
});

// console.log(printedNumbers);    //1 2 3 4 5

const incrementedArray = numbersArray.map((number) => {
  return number + 1; // it will return an array and it will return undefined if return is missing and it return an modified array
});
// console.log(incrementedArray);    [ 2, 3, 4, 5, 6 ]

const evenNumbersArray = numbersArray.filter((number) => {
  return !(number % 2); // it will return an filter array which will have only that element which will fullfill the given condition
});

// console.log(evenNumbersArray);//[2,4]

const sumOfNumbersArray = numbersArray.reduce((accumulator, initalValue) => {
  return accumulator + initalValue; // it will return single value and take two parameter which is accumulator ,initalValue
});

// console.log(sumOfNumbersArray)  //15

// Define a callback function before you use it in forEach, map, filter or reduce.

function isEven(element) {
  return !(element % 2);
}

const evenNumbers = numbersArray.filter(isEven);
// console.log(evenNumbers); //[2,4]

const countries = ["Finland", "Sweden", "Denmark", "Norway", "IceLand"];

// Use forEach to console.log each country in the countries array.
const loggedCountries = countries.forEach((country) => console.log(country));

// console.log(loggedCountries)     //Finland Sweden Denmark Norway IceLand

// Use forEach to console.log each name in the names array.
const names = ["Asabeneh", "Mathias", "Elias", "Brook"];

const loggedNames = names.forEach((name) => console.log(name));

console.log(loggedNames); // Asabeneh Mathias Elias Brook

// Use forEach to console.log each number in the numbers array.
const numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

const loggedNumbers = numbers.forEach((number) => console.log(number));
// console.log(loggedNumbers);    // 1 2 3 4 5 6 7 8 9 10

// Use map to create a new array by changing each country to uppercase in the countries array.

const capitalizedCountryNames = countries.map((country) => {
  return country.toUpperCase();
});

// console.log(capitalizedCountryNames);

// Use map to create an array of countries length from countries array.

const countryNameLengths = countries.map((countryName) => {
  return countryName.length;
});

// console.log(countryNameLengths);

// Use map to create a new array by changing each number to square in the numbers array

const squareOfNumbersArray = numbersArray.map((number) => {
  return number ** 2;
});

// console.log(squareOfNumbersArray);

// Use map to change to each name to uppercase in the names array
const uppercaseNames = names.map((name) => {
  return name.toUpperCase();
});

// console.log(uppercaseNames);

// Use map to map the products array to its corresponding prices.

const products = [
  { product: "banana", price: 3 },
  { product: "mango", price: 6 },
  { product: "potato", price: " " },
  { product: "avocado", price: 8 },
  { product: "coffee", price: 10 },
  { product: "tea", price: "" },
];

const prices = [];
for (const { price, ...product } of products) {
  prices.push(price);
}
// console.log(prices);   //[ 3, 6, ' ', 8, 10, '' ]

// Use filter to filter out countries containing land.
const filteredCountries = countries.filter((country) => {
  return country.includes("land");
});

// console.log(filteredCountries);  //[ 'Finland' ]

// Use filter to filter out countries having six character.
const filterSixCharacterCountry = [];
const filteredCountry = countries.filter((country) => {
  if (country.length === 6) {
    filterSixCharacterCountry.push(country);
  }
  return filterSixCharacterCountry;
});

// console.log(filterSixCharacterCountry);   //[ 'Sweden', 'Norway' ]

// Use filter to filter out countries containing six letters and more in the country array.
const longCountriesName = countries.filter((country) => {
  return country.length > 6;
});

console.log(longCountriesName); //[ 'Finland', 'Denmark', 'IceLand' ]

// Use filter to filter out country start with 'E';
const countryNameStartWithE = countries.filter((countryName) => {
  return countryName.startsWith("E");
});

console.log(countryNameStartWithE); //[]

// Use filter to filter out only prices with values.

// const pricesWithValues = products.filter(({ price }) => price);

// console.log(pricesWithValues);

// Declare a function called getStringLists which takes an array as a parameter and then returns an array only with string items.
let stringArray = [];
function getStringLists(inputArray) {
  for (let index = 0; index < inputArray.length; index++) {
    if (typeof inputArray[index] === "string") {
      stringArray.push(inputArray[index]);
    }
  }
  return stringArray;
}

console.log(getStringLists([1, 2, "sanket", "omkar", 134]));

//Use reduce to sum all the numbers in the numbers array.

const sum = numbersArray.reduce((accumulator, initalValue) => {
  return accumulator + initalValue;
});

console.log(sum); //15

// Use reduce to concatenate all the countries and to produce this
//sentence: Estonia, Finland, Sweden, Denmark, Norway, and IceLand are north European countries
// const countries = ["Finland", "Sweden", "Denmark", "Norway", "IceLand"];
// const sanket =""
// const sank = countries.reduce((accumulator,initalValue)=>{
//     if (initalValue > countries.length-1) {
//         accumulator+=", "+initalValue;
//     } else {
//         accumulator+=initalValue+" and "
//     }
//     return accumulator;
// })

// console.log(`${sank} are north European countries`)

// Explain the difference between some and every

// const allNumbersGreaterThanTwo = numbersArray.some((number)=> {
//   return number>2;   //it will check only one condition and if it found it true it return true and leve the rest of number in an array
// })
// console.log(allNumbersGreaterThanTwo);  //True

const allNumbersGreaterThanTwo = numbersArray.every((number) => {
  return number > 2; // the every check each and every condition one by one if any one condition fails it will return false otherwise it will return true
});
console.log(allNumbersGreaterThanTwo); //false

// Use some to check if some names' length greater than seven in names array
const hasNamesGreaterThanSevenCharacters = names.some((name) => {
  return name.length > 7;
});

console.log(hasNamesGreaterThanSevenCharacters); //true

// Use every to check if all the countries contain the word land
const allCountriesContainLand = countries.every((country) => {
  return country.toLowerCase().includes("land");
});

console.log(allCountriesContainLand); //false

// Explain the difference between find and findIndex.

// const foundNumber = numbers.find((num) => {
//   return num > 2  // it will return the first number which fullFill the given condition else undefined
// });
// console.log(foundNumber); // Output: 3

const foundNumber = numbers.findIndex((num) => {
  return num > 2; // it return the index of that number which fullFill the given condition else -1
});
console.log(foundNumber); // Output: 2

// Use find to find the first country containing only six letters in the countries array
const firstCountryWithSixLetters = countries.find((country) => {
  return country.length >= 6;
});

console.log(firstCountryWithSixLetters); //Finland

// Use findIndex to find the position of the first country containing only six letters in the countries array

const IndexOfFirstCountryWithSixLetters = countries.findIndex((country) => {
  return country.length >= 6;
});

console.log(IndexOfFirstCountryWithSixLetters); //0

// Use findIndex to find the position of Norway if it doesn't exist in the array you will get -1.
const indexOfNorway = countries.findIndex((country) => {
  return country === "Norway";
});

console.log(indexOfNorway); //3

// Use findIndex to find the position of Russia if it doesn't exist in the array you will get -1.
const indexOfRussia = countries.findIndex((country) => {
  return country === "Russia";
});

console.log(indexOfRussia); //-1

// Find the total price of products by chaining two or more array iterators(eg. arr.map(callback).filter(callback).reduce(callback))

const totalPriceOfProducts = products
  .map((product) => {
    return product.price;
  })
  .reduce((accumulator, initalValue) => {
    return parseInt((accumulator += initalValue));
  });
console.log(`The total price is :${totalPriceOfProducts}`);

// Find the sum of price of products using only reduce reduce(callback))
const totalPrice = products.reduce((acc, product) => {
  const productPrice = parseFloat(product.price);

  if (!isNaN(productPrice)) {
    return acc + productPrice;
  } else {
    return acc;
  }
}, 0);

console.log(totalPrice);


//   
