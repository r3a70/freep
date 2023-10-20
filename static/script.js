function uploadFile() {
  const fileInput = document.getElementById('fileInput');
  const file = fileInput.files[0];
  
  const xhr = new XMLHttpRequest();
  const formData = new FormData();
  
  formData.append('file', file);

  xhr.open('POST', '/upload', true);
  
  xhr.upload.onprogress = function(e) {
    if (e.lengthComputable) {
      const percent = (e.loaded / e.total) * 100;
      const progressBar = document.getElementById('progressBar');
      progressBar.style.width = percent + '%';
      progressBar.textContent = percent.toFixed(2) + '%';
    }
  };
  
  xhr.onload = function() {
    // Handle the response after the file is uploaded
    // This function will be called when the upload is complete
    const progressBar = document.getElementById('progressBar');
    progressBar.textContent = xhr.responseText
  };
  
  xhr.send(formData);
}
