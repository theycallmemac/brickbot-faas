const http = require('http');
let type = process.argv[2];
let stop = process.argv[3];
http.get("http://localhost:8000/"+type+"/stop/"+stop, (resp) => {
  let data = "";
  resp.on('data', (chunk) => {
    data += chunk;
  });
  resp.on('end', () => {
    console.log(JSON.parse(data));
  });
}).on("error", (err) => {
  console.log("Error: " + err.message);
});
