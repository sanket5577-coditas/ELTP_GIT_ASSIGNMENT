// const array = [[0,1],[2,3]]

// for (const [key,value] of Object.entries(array)) {
//     console.log(`${key} and values are ${value}` )
    
// }

const myObjectInfo ={
    name:"sanket",
    age:21
}

const copyOfMyObjectInfo = Object.assign(myObjectInfo)
copyOfMyObjectInfo["surname"] = "chauhdary"
console.log(copyOfMyObjectInfo)
console.log(myObjectInfo)