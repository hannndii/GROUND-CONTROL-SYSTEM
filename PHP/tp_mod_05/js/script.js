document.addEventListener("DOMContentLoaded", function () {
    const LoadButton = document.getElementById("show-tasks");
    if (LoadButton) {
        LoadButton.addEventListener("click", loadTasks);
    }
});

async function loadTasks() {
    try {
        const response = await fetch("js/tasks.json");

        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const data = await response.json();
        const tasks = data.tasks; 

        if (tasks && tasks.length > 0) {
            displayTasks(tasks);
        } else {
            const taskListBody = document.getElementById("task-list");
            taskListBody.innerHTML = '<tr><td colspan="2">Tidak ada tugas yang ditemukan.</td></tr>';
        }

    } catch (error) {
        const taskListBody = document.getElementById("task-list");
        taskListBody.innerHTML = `<tr><td colspan="2" style="color: red;">Error: ${error.message}.</td></tr>`;
    }
}

function displayTasks(tasks) {
    const taskListBody = document.getElementById("task-list");
    taskListBody.innerHTML = "";

    tasks.forEach(function (task) {
        const row = document.createElement("tr");

        const isCompleted = task.completed;
        const statusText = isCompleted ? "Completed" : "Not Completed";
        const statusClass = isCompleted ? "completed" : "not-completed"; 
        
        row.classList.add(statusClass);

        let cellTask = document.createElement("td");
        cellTask.textContent = task.text;
        row.appendChild(cellTask);

        let cellStatus = document.createElement("td");
        cellStatus.textContent = statusText;
        cellStatus.classList.add(statusClass); 
        row.appendChild(cellStatus);

        taskListBody.appendChild(row);
    });
}