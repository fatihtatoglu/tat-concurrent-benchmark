const util = require("util");

const delay = util.promisify(setTimeout);

async function runTasks(taskCount) {

    const tasks = [];

    for (let i = 0; i < taskCount; i++) {
        tasks.push(delay(10 * 1000));
    }

    await Promise.all(tasks);

    console.log("All tasks are finished.");
}

const taskCount = parseInt(process.argv[2] || 10000);

runTasks(taskCount).catch(e => {
    console.error("Error:", e);
});