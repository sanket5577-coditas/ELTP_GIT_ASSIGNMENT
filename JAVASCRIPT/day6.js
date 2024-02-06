// Boolean 
// logical_operator 
// and , or ,not 

// falsy - NaN , FALSE,0,Undefined,null,''
// truthy - all are true other than falsy



//  comparison

// console.log(5<'123') //true
// console.log(5<'') //false
// console.log(5<'a') //false
// console.log(5<5) //false
// console.log(5=='5')
// console.log(5==='5')
// console.log(5!=='5')

// console.log(NaN==NaN) //asignment 
// console.log(NaN===NaN)// assignment

// //how to compare array
// console.log(!!'true' == !!'false') //false
// console.log('false' == 'false') // true


console.log("simple authentication system")
const logInId = true;
const age = 21;
if(logInId===true)
{
    if(age>=18)
    {
        console.log("Wlcome in the authentication system..!")
    }
    else
    {
        console.log("ur Age is not 18")
    } 
}
else
{
    console.log("Account is suspended")
}