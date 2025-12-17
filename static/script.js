function loadConcerts(id) {
    fetch("/api/concerts?id=" + id)
        .then(r => r.json())
        .then(data => {
            let html = ""
            for (let city in data) {
                html += `<p><b>${city}:</b> ${data[city].join(", ")}</p>`
            }
            document.getElementById("concerts").innerHTML = html
        })
}
