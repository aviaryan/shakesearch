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
    const count = document.getElementById("count-value");
    // results can be null because of golang behavior
    if (!results || results.length === 0) {
      table.style.display = "none";
      count.innerHTML = "0";
      return;
    }
    table.style.display = "table";
    const tableBody = document.getElementById("table-body");
    const rows = [];
    for (let result of results) {
      const formattedMatch = result.match.replace(/\r\n/g, '<br/>')
      const row = `<tr><td>${formattedMatch}</td><td>${result.work}</td></tr>`;
      rows.push(row);
    }
    tableBody.innerHTML = rows.join('');
    count.innerHTML = results.length;
  },
};

const form = document.getElementById("form");
form.addEventListener("submit", Controller.search);
