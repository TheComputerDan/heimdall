const app = document.getElementById('root');

const logo = document.createElement('img');
logo.src = 'logo.png';

const container = document.createElement('div');
container.setAttribute('class', 'container');

app.appendChild(logo);
app.appendChild(container);

var request = new XMLHttpRequest();
request.open('GET', 'http://localhost:8080/containers', true);
request.onload = function () {

  // Begin accessing JSON data here
  var data = JSON.parse(this.response);
  if (request.status >= 200 && request.status < 400) {
    data.forEach(docker_container => {
        const card = document.createElement('div');
        card.setAttribute('class', 'card');

        const h1 = document.createElement('h1');
        h1.textContent = docker_container.Names;
        h1.textContent = h1.textContent.split('/')[1]

        const p = document.createElement('p');

        p.setAttribute('style', 'white-space: pre;');

        p.textContent = `State: ${docker_container.State}\r\n`;
        p.textContent += `Image: ${docker_container.Image}\r\n`;
        p.textContent += `ID: ${docker_container.Id}`;

        container.appendChild(card);
        card.appendChild(h1);
        card.appendChild(p);
    });
  } else {
    const errorMessage = document.createElement('marquee');
    errorMessage.textContent = `Gah, it's not working!`;
    app.appendChild(errorMessage);
  }
}

request.send();