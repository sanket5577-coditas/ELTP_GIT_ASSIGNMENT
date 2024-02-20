// const data = {
//     user:[{
//         name:"Sanket",
//         lastName:"Kumar"
//     },{
//         name:"Shambu",
//         lastName:"Deshmukh"
//     }]
// }

// const {user:[person1]}=data
// console.log(person1);
// const{name} = person1;
// // console.log(name);


// setInterval(() => {
//     console.log("Hello");
//     setTimeout(() => {
//         console.log("sdan");
//     }, 100);
    
// }, 5000);


const getUserPromise = (onSucess,onFailer)=>{
    const user=[{name:"Sanket"},{lastName:"Kumar"}]
    const duration = Math.random()*5000;
    setTimeout(() => {
        const random = Math.random();
        if (random>0.5) {
            return onSucess({data:user,statusCode:200,error:null})
        }
       else{
        return onFailer({data:null,statusCode:500,error:("something went wrong")})
       }
    }, duration);

}


getUserPromise((Response)=>console.log(Response),
(error)=>console.log(error))