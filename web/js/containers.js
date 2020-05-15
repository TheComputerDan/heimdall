const app = document.getElementById('root');

// const logo = document.createElement('img');
// logo.src = 'images/logo.png';
var hostname = window.location.hostname;
const endpoint = "http://" + hostname + ':8095/containers'

const container = document.createElement('div');
container.setAttribute('class', 'uk-container uk-container-large');
container.setAttribute('uk-sortable',true);
// container.classList.add('uk-child-width-1-2@s', 'uk-grid-match');

// app.appendChild(logo);
app.appendChild(container);

var request = new XMLHttpRequest();
request.open('GET', endpoint, true);
request.onload = function () {

  // Begin accessing JSON data here
var data = JSON.parse(this.response);
if (request.status >= 200 && request.status < 400) {
  data.forEach(docker_container => {
      const card = document.createElement('div');
      card.setAttribute('class', 'card');

      card.classList.add("uk-card" ,
        "uk-card-hover",
        "uk-card-body", 
        "uk-card-default",
        "uk-margin",
        "uk-background-muted");
      // card.classList.add("uk-card-hover");

      var stat = document.createElement('span');
      stat.setAttribute('class','uk-label uk-card-badge');

      const h1 = document.createElement('h1');
      h1.setAttribute('class','uk-card-title');
      h1.textContent = docker_container.Names;
      h1.textContent = h1.textContent.split('/')[1]

      const p = document.createElement('p');

      p.setAttribute('style', 'white-space: pre;');

      p.textContent = `State: ${docker_container.State}\r\n`;
      p.textContent += `Image: ${docker_container.Image}\r\n`;
      p.textContent += `ID: ${docker_container.Id.slice(0,11)}`;

      if (docker_container.State == "exited"){
        stat.classList.add("uk-label-danger");
        stat.textContent = `${docker_container.State}`;
        // stat.style.backgroundColor = "#F14E44";
      } else if (docker_container.State == "running") {
        stat.classList.add("uk-label-success");
        stat.textContent = `${docker_container.State}`;
        // dot.style.backgroundColor = "#84fab0";
      } else {
        stat.classList.add("uk-label-warning");
        stat.textContent = `${docker_container.State}`;
        // dot.style.backgroundColor = "#fcff37";
      }

      container.appendChild(card);
      card.appendChild(h1);
      card.appendChild(stat);
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