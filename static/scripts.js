var articles;

function submitArticle(){
  fetch('article', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Accept':       'application/json'
    },
    body: '{ \"Id\": \"\", \"Title\": \"' + document.getElementById("Title").value 
    +'\", \"Description\": \"' + document.getElementById("Description").value
    +'\", \"Content\": \"' + document.getElementById("Content").value +"\"}"
  })
  .then((res) => res.json())
  .then((data) => console.log(data))
  .catch((error) => console.log(error))
}

function getFormData() {
  var formData = '{ \"Id\": \"\", \"Title\": \"' + document.getElementById("Title").value 
  +'\", \"Description\": \"' + document.getElementById("Description").value
  +'\", \"Content\": \"' + document.getElementById("Content").value +' \"}'
  console.log(formData);
  return formData;
}

function getArticles(){
  fetch('all', {
    method: 'GET',
  })
  .then((res) => res.json())
  .then((data) => {
    data.forEach( function(cur) {
      document.getElementById("artikler").innerHTML 
        += "<tr><th>"+cur.Id+"</th><th>"+cur.Title+"</th><th>"
        + cur.Description + "</th><th>"+ cur.Content + "</th></tr>"
    })
    console.log(data)
  })
  .catch((error) => console.log(error))
}
