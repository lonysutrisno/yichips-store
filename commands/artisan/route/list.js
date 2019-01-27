var exec = require('child_process').exec;
const fs = require('fs');

const servicesFolder = './services/';
var services = [];

console.log("\n[ ROUTE LIST ]");

let adminBaseUrl = "";
let mobileBaseUrl = "";

fs.readFile("./config/app.go", "utf8", (err, data) => {
	routeBases = data.match(/(admin_base_url|mobile_base_url)\"(.*)\"/g);

	routeBases.forEach((base, index) => {
		const url = base.match(/ (.*)/g)[0].replace(/(\"| )/g, "");

		if (base.indexOf("admin_base_url") != -1) {
			adminBaseUrl = url;
		} else {
			mobileBaseUrl = url;
		}
	});
});

fs.readdirSync(servicesFolder).forEach(folder => {
	const route = "./services/" + folder + "/connector/rest/route.go";

	fs.readFile(route, "utf8", (err, data) => {
		console.log("\nService: " + folder);
		console.log("===================================\n");

		const routes = data.match(/(e.(GET|POST|PUT|DELETE)\((adminBaseUrl|mobileBaseUrl)\+\"(.*)\")(.*)/g);

		routes.forEach((route, index) => {
			const method = route.match(/(GET|POST|PUT|DELETE)/)[0];
			const baseUrl = route.indexOf("adminBaseUrl") ? adminBaseUrl : mobileBaseUrl;
			const path = route.match(/\"(.*)\"/)[0].replace(/\"/g, "");
			const use = route.match(/\/\/(.*)/g)[0].replace("// ", "");

			console.log("\t" + method + "\t:  " + baseUrl + path + "\t\t[ " + use + " ]")
		})

		console.log("\n");
	});

});