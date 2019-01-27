var fs = require("fs");

const servicesFolder = "./services/";

fs.readdirSync(servicesFolder).forEach(folder => {
	fs.createReadStream(".env.test").pipe(fs.createWriteStream("services/" + folder + "/test/.env"));
	console.log(".env.test has been copied to service " + folder);
});