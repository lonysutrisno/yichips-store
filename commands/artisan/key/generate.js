var exec = require('child_process').exec;
const crypto = require('crypto');
const fs = require('fs');

const file = '.env';

fs.readFile(file, "utf8", (err, data) => {
	let newKey = "";
	const charSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";

	for (var i = 0; i < 5; i++) {
		newKey += charSet.charAt(Math.floor(Math.random() * charSet.length));
	}

	newKey = crypto.createHash('sha256').update(newKey).digest('base64')
	data = data.replace(/API_KEY(.*)/, "API_KEY=" + newKey)

	fs.writeFile(file, data, function(err) {
	    console.log("NEW API KEY = " + newKey + "\n")
	});
});
