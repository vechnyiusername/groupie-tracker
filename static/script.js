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

function toggleSection(header) {
    const content = header.nextElementSibling;
    
    if (content.classList.contains('hidden')) {
        // Expand
        content.style.maxHeight = content.scrollHeight + 'px';
        content.classList.remove('hidden');
        header.classList.remove('collapsed');
        // Reset max-height after transition completes
        setTimeout(() => {
            content.style.maxHeight = '';
        }, 300);
    } else {
        // Collapse
        content.style.maxHeight = content.scrollHeight + 'px';
        // Force reflow
        content.offsetHeight;
        content.style.maxHeight = '0';
        content.classList.add('hidden');
        header.classList.add('collapsed');
    }
}

// Initialize: hide all sections by default
document.addEventListener('DOMContentLoaded', function() {
    const sections = document.querySelectorAll('.artist-section');
    sections.forEach(section => {
        const header = section.querySelector('h3');
        const content = section.querySelector('.artist-section-content');
        if (header && content) {
            header.classList.add('collapsed');
            content.classList.add('hidden');
            content.style.maxHeight = '0';
        }
    });
});