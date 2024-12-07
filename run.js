//Starting the server and running the correct satarting commands/files based on the operation system.
const { spawn } = require('child_process');
const os = require('os');

const isWindows = os.platform() === 'win32';

function runCommand(command, cwd) {
    const options = { cwd, shell: true };
    const process = spawn(command, [], options);

    process.stdout.on('data', (data) => console.log(data.toString()));
    process.stderr.on('data', (data) => console.error(data.toString()));
    process.on('close', (code) => console.log(`Process exited with code ${code}`));
}

// Start the server
console.log("Starting server...");
runCommand(isWindows ? 'start_server.bat' : 'bash start_server.sh', './server');

// Start the client
console.log("Starting client...");
runCommand(isWindows ? 'start_client.bat' : 'bash start_client.sh', './client');
