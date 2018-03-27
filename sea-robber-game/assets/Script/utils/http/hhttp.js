function newHttpRequest(resolve, reject) {
  let xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function () {
    if (xhr.readyState == 4) { // Done
      if (xhr.status >= 200 && xhr.status < 400) {
        resolve(xhr.response, xhr.responseText)
      } else {
        reject(xhr.response, xhr.status)
      }
    }
  };
  return xhr
}


module.exports.get = function (url, options) {
  return new Promise(function (resolve, reject) {
    let xhr = newHttpRequest(resolve, reject)
    xhr.open('GET', url, true);
    if (options.headers) {
       Object.keys(options.headers).forEach( key => {
         xhr.setRequestHeader(key, options.headers[key]);
       })
    }
    xhr.send();
  })
}

module.exports.post = function (url, params, options) {
  return new Promise(function (resolve, reject) {
    let xhr = newHttpRequest(resolve, reject)
    xhr.open('POST', url, true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    if (options.headers) {
      Object.keys(options.headers).forEach( key => {
        xhr.setRequestHeader(key, options.headers[key]);
      })
    }
    xhr.send(JSON.stringify(params));
  })
}