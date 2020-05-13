const app = document.getElementById('root');

// const logo = document.createElement('img');
// logo.src = 'images/logo.png';

const container = document.createElement('div');
container.setAttribute('class', 'container');

// app.appendChild(logo);
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

      card.classList.add("uk-card" ,"uk-card-hover");
      // card.classList.add("uk-card-hover");
      
      const dot = document.createElement('uk-icon');
      dot.classList.add('heart');

      const h1 = document.createElement('h1');
      h1.textContent = docker_container.Names;
      h1.textContent = h1.textContent.split('/')[1]

      const p = document.createElement('p');

      p.setAttribute('style', 'white-space: pre;');

      p.textContent = `State: ${docker_container.State}\r\n`;
      p.textContent += `Image: ${docker_container.Image}\r\n`;
      p.textContent += `ID: ${docker_container.Id.slice(0,11)}`;

      if (docker_container.State == "exited"){
        dot.style.backgroundColor = "#F14E44";
      } else if (docker_container.State == "running") {
        dot.style.backgroundColor = "#84fab0";
      } else {
        dot.style.backgroundColor = "#fcff37";
      }

      container.appendChild(card);
      card.appendChild(h1);
      card.appendChild(dot);
      card.appendChild(p);

      // Testing making cards into buttons
      card.onclick = function () {
        location.href = "containers.html";
      };

    });
  } else {
    const errorMessage = document.createElement('marquee');
    errorMessage.textContent = `Gah, it's not working!`;
    app.appendChild(errorMessage);
  }
}

request.send();