const [ , , ...args ] = process.argv;                                                                                                
let totalSeconds = args[0];                                                                                                           
let days = Math.floor(totalSeconds / 86400);                                                                                           
let hours = Math.floor(totalSeconds / 3600) - (days * 24);                                                                             
totalSeconds %= 3600;                                                                                                                 
let minutes = Math.floor(totalSeconds / 60);                                                                                           
let seconds = Math.round(totalSeconds % 60);                                                                                           
console.log(`${days} days, ${hours} hours, ${minutes} minutes and ${seconds} seconds`); 
