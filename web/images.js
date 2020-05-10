const app = document.getElementById('root');

const logo = document.createElement('img');
logo.src = 'images/logo.png';

const container = document.createElement('div');
container.setAttribute('class', 'container');

app.appendChild(logo);
app.appendChild(container);

var request = new XMLHttpRequest();
request.open('GET', 'http://localhost:8080/images', true);
request.onload = function () {

  // Begin accessing JSON data here
var data = JSON.parse(this.response);
if (request.status >= 200 && request.status < 400) {
  data.forEach(docker_images => {
      const card = document.createElement('div');
      card.setAttribute('class', 'card');

      if (docker_images.RepoTags == null) {
        return;
      }

      const repoTagsLength = docker_images.RepoTags.length;
      
      const h1 = document.createElement('h1');
      h1.textContent = docker_images.RepoTags[0];
      h1.textContent = h1.textContent.split(':')[0]

      const p = document.createElement('p');

      p.setAttribute('style', 'white-space: pre;');

      p.textContent = `ID: ${docker_images.Id}\r\n`;
      p.textContent += `Size: ${docker_images.Size}\r\n`;
      p.textContent += `Created: ${docker_images.Created}\r\n`;
      p.textContent += `Tags: ${repoTagsLength}\r\n`;

      container.appendChild(card);
      card.appendChild(h1);
      card.appendChild(p);
      
      // Testing making cards into buttons
      card.onclick = function () {
        location.href = "images.html";
      };

    });
  } else {
    const errorMessage = document.createElement('marquee');
    errorMessage.textContent = `Gah, it's not working!`;
    app.appendChild(errorMessage);
  }
}

request.send();