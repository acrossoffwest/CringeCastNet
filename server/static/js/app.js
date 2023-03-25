/** global variables **/
var req = "";


function appStart() {
    guiController = new GuiController;
    guiController.attachResponseHandler(handleServerResponse);
}

function reqStrUpdate() {
    req = guiController.getQueryFromQueryBox(queryType);
    guiController.updateQueryDisp(req);
}

function httpGetFromQueryBox(queryType) {
    req = guiController.getQueryFromQueryBox(queryType);
    httpGet(guiController, req, handleServerResponse);
}


function handleServerResponse(responseStatus, responseText) {
    var textToDisplay;

    if (responseStatus == 200) {
        textToDisplay = "Valid request. Response:\n" + responseText;
        drawDeviceStatus("Online");
    }

    /** request error **/
    else if (responseStatus == -1) {
        textToDisplay = "http request error:\n" + responseText;
        drawDeviceStatus("Error");
    }
    /** request timeout **/
    else if (responseStatus == 0) {
        textToDisplay = "http request timeout:\n" + responseText;
        drawDeviceStatus("Offline");
    }
    /** Error: not found **/
    else if (responseStatus == 404) {
        textToDisplay = "Server responded but parameter not found:\n" + responseText;
        drawDeviceStatus("Error");
    }
    /** Error: in teapot mode **/
    else if (responseStatus == 418) {
        textToDisplay = "Welp, this is awkward:\n" + responseText;
        drawDeviceStatus("Teapot");
    }
    /** Error: different error **/
    else {
        textToDisplay = "Server responded with undefined error no.:\n" + responseText;
        drawDeviceStatus("Error");
    }
    updateRespDisp(textToDisplay);
}

function httpGet(controller, param, onLoadFunctionHandle = null) {
    // console.log(controller)
    if (controller == undefined) {
        controller = guiController
    }

    controller.updateQueryDisp(param);
    controller.updateRespDisp("Waiting for response...");
    // console.log("param: " + param);
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", param, true);
    xmlHttp.onload = function (e) {
        if (xmlHttp.readyState === 4) {
            onLoadFunctionHandle(xmlHttp.status, xmlHttp.responseText)
        }
    };
    xmlHttp.timeout = 60000;
    xmlHttp.ontimeout = function (e) {
        onLoadFunctionHandle(0, "timeout");
    };
    xmlHttp.onerror = function (e) {
        onLoadFunctionHandle(-1, "Error");
    };
    // console.log("Starting query...");
    xmlHttp.send(null);
}

/** Rendering HTML section **/

function updateRespDisp(text) {
    guiController.updateRespDisp(text);
}

function updateQueryDisp(text) {
    guiController.updateQueryDisp(text);
}

function drawDeviceStatus(statusText) {
    guiController.drawDeviceStatus(statusText);
}

/** Drag and drop **/

let dropArea = document.getElementById('drop-area');

['dragenter', 'dragover', 'dragleave', 'drop'].forEach(eventName => {
    dropArea.addEventListener(eventName, preventDefaults, false)
});

['dragenter', 'dragover'].forEach(eventName => {
    dropArea.addEventListener(eventName, highlight, false)
});

;['dragleave', 'drop'].forEach(eventName => {
    dropArea.addEventListener(eventName, unhighlight, false)
});

function preventDefaults(e) {
    e.preventDefault()
    e.stopPropagation()
}

function highlight(e) {
    dropArea.classList.add('highlight')
}

function unhighlight(e) {
    dropArea.classList.remove('highlight')
}

dropArea.addEventListener('drop', handleDrop, false);

function handleDrop(e) {
    let dt = e.dataTransfer
    let files = dt.files
    handleFiles(files)
}

function handleFiles(files) {
    ([...files]).forEach(uploadFile)
}

function uploadFile(file) {
    let url = '/uploader'
    let formData = new FormData()

    formData.append('file', file)

    fetch(url, {
        method: 'POST',
        body: formData
    })
        .then(() => { /* Done. Inform the user */
        })
        .catch(() => { /* Error. Inform the user */
        })
}