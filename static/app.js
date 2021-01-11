const Controller = {
  search: (ev) => {
    ev.preventDefault();
    const form = document.getElementById("form");
    const data = Object.fromEntries(new FormData(form));
    const response = fetch(`/search?q=${data.query}`).then((response) => {
      response.json().then((result) => {
        Controller.updateTable(result['results']);
      });
    });
  },

  updateTable: (results) => {
    const table = document.getElementById("table");
    // results can be null because of golang behavior
    if (!results || results.length === 0) {
      table.style.display = "none";
      return;
    }
    table.style.display = "block";
    const tableBody = document.getElementById("table-body");
    const rows = [];
    for (let result of results) {
      const row = `<tr><td>${result.match}</td><td>${result.work}</td></tr>`;
      rows.push(row);
    }
    tableBody.innerHTML = rows.join('');
  },
};

const form = document.getElementById("form");
form.addEventListener("submit", Controller.search);
