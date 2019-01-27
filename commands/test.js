var exec = require('child_process').exec;
const fs = require('fs');

const servicesFolder = './services/';
var services = [];

fs.readdirSync(servicesFolder).forEach(folder => {
	services.push(folder);
});

runTest();

async function runTest() {
	for (x in services) {
		await exec("go test ./services/" + services[x] + "/test", function (error, stdout, stderr) {
			console.log(stdout);
		});
	}
}